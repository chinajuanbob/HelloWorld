package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	todo "github.com/chinajuanbob/helloworld/pb"
	_ "github.com/davecgh/go-spew/spew"
	"github.com/golang/glog"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/client/grpc"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var todoClient todo.TodoService

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:              "grpc",
	Short:            "A brief description of your command",
	TraverseChildren: true,
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			glog.Fatal("need connent")
		}
		rsp, err := todoClient.Add(context.Background(), &todo.AddTodoRequest{
			Content: args[0],
		}, client.WithAddress(viper.GetString("address")))
		if err != nil {
			glog.Error(err)
		} else {
			// spew.Dump(rsp)
			output([]*todo.Todo{
				rsp.GetTodo(),
			})
		}
	},
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		rsp, err := todoClient.List(context.Background(), &todo.ListTodoRequest{}, client.WithAddress(viper.GetString("address")))
		if err != nil {
			glog.Error(err)
		} else {
			// spew.Dump(rsp)
			output(rsp.GetTodos())
		}
	},
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update id status",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires id and status")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("bad id")
		}
		statusStr := strings.ToUpper(args[1])
		statusInt, ok := todo.TodoStatus_value[statusStr]
		if !ok {
			return errors.New("bad status, available status are 'NEW', 'INPROGRESS' and 'DONE'.")
		}
		status := todo.TodoStatus_NEW
		if statusInt == int32(todo.TodoStatus_INPROGRESS) {
			status = todo.TodoStatus_INPROGRESS
		} else if statusInt == int32(todo.TodoStatus_DONE) {
			status = todo.TodoStatus_DONE
		}
		rsp, err := todoClient.Update(context.Background(), &todo.UpdateTodoRequest{
			Id:     int32(id),
			Status: status,
		}, client.WithAddress(viper.GetString("address")))
		if err != nil {
			glog.Error(err)
		} else {
			// spew.Dump(rsp)
			output([]*todo.Todo{
				rsp.GetTodo(),
			})
		}
		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("bad id")
		}
		rsp, err := todoClient.Delete(context.Background(), &todo.DeleteTodoRequest{
			Id: int32(id),
		}, client.WithAddress(viper.GetString("address")))
		if err != nil {
			glog.Error(err)
		} else {
			glog.Info(rsp.Success)
		}
		return nil
	},
}

func init() {
	clientCmd.AddCommand(grpcCmd)
	grpcCmd.PersistentFlags().String("address", "0.0.0.0:6666", "address of service")
	viper.BindPFlag("address", grpcCmd.PersistentFlags().Lookup("address"))

	service := micro.NewService(
		micro.Name("todo"),
		micro.Client(grpc.NewClient()),
	)
	service.Init()
	todoClient = todo.NewTodoService("todo", service.Client())

	grpcCmd.AddCommand(addCmd)
	grpcCmd.AddCommand(listCmd)
	grpcCmd.AddCommand(updateCmd)
	grpcCmd.AddCommand(deleteCmd)
}

func output(todos []*todo.Todo) {
	// spew.Dump(todos)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status", "Last Update", "Content"})
	for _, v := range todos {
		table.Append([]string{fmt.Sprintf("%d", v.Id), v.Status.String(), v.LastUpdated.String(), v.Content})
	}
	table.Render()
}
