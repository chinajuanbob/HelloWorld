package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	pbTodo "github.com/chinajuanbob/helloworld/pb"
	"github.com/chinajuanbob/helloworld/pkg/gen/client"
	"github.com/chinajuanbob/helloworld/pkg/gen/client/todo"
	"github.com/chinajuanbob/helloworld/pkg/gen/modules"
	"github.com/golang/glog"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var todoHttpClient *client.TodoClient

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:              "http",
	Short:            "A brief description of your command",
	TraverseChildren: true,
}

// addHttpCmd represents the add command
var addHttpCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			glog.Fatal("need connent")
		}
		param := todo.NewAddTodoParams()
		param.Body = &modules.PbAddTodoRequest{
			Content: args[0],
		}
		rsp, err := todoHttpClient.Todo.AddTodo(param)
		if err != nil {
			glog.Error(err)
		} else {
			// spew.Dump(rsp)
			outputHttp([]*modules.PbTodo{
				rsp.Payload.Data,
			})
		}
	},
}

// listHttpCmd represents the list command
var listHttpCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		rsp, err := todoHttpClient.Todo.ListTodos(todo.NewListTodosParams())
		if err != nil {
			glog.Error(err)
		} else {
			// spew.Dump(rsp)
			outputHttp(rsp.Payload.Data)
		}
	},
}

// updateHttpCmd represents the update command
var updateHttpCmd = &cobra.Command{
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
		statusInt, ok := pbTodo.TodoStatus_value[statusStr]
		if !ok {
			return errors.New("bad status, available status are 'NEW', 'INPROGRESS' and 'DONE'.")
		}
		param := todo.NewUpdateTodoParams()
		param.Body = &modules.PbUpdateTodoRequest{
			ID:     int32(id),
			Status: statusInt,
		}
		rsp, err := todoHttpClient.Todo.UpdateTodo(param)
		if err != nil {
			glog.Error(err)
		} else {
			// spew.Dump(rsp)
			outputHttp([]*modules.PbTodo{
				rsp.Payload.Data,
			})
		}
		return nil
	},
}

var deleteHttpCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("bad id")
		}
		param := todo.NewDeleteTodoParams()
		param.TodoID = int64(id)
		rsp, err := todoHttpClient.Todo.DeleteTodo(param)
		if err != nil {
			glog.Error(err)
		} else {
			glog.Info(rsp.Payload.Data)
		}
		return nil
	},
}

func init() {
	clientCmd.AddCommand(httpCmd)
	httpCmd.PersistentFlags().String("web", "0.0.0.0:9999", "address of service")
	viper.BindPFlag("web", httpCmd.PersistentFlags().Lookup("web"))

	todoHttpClient = client.NewHTTPClientWithConfig(nil, client.DefaultTransportConfig().WithHost(viper.GetString("web")))

	httpCmd.AddCommand(addHttpCmd)
	httpCmd.AddCommand(listHttpCmd)
	httpCmd.AddCommand(updateHttpCmd)
	httpCmd.AddCommand(deleteHttpCmd)
}

func outputHttp(todos []*modules.PbTodo) {
	// spew.Dump(todos)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status", "Last Update", "Content"})
	for _, v := range todos {
		status := pbTodo.TodoStatus_NEW
		if v.Status == int32(pbTodo.TodoStatus_INPROGRESS) {
			status = pbTodo.TodoStatus_INPROGRESS
		} else if v.Status == int32(pbTodo.TodoStatus_DONE) {
			status = pbTodo.TodoStatus_DONE
		}
		table.Append([]string{fmt.Sprintf("%d", v.ID), status.String(),
			fmt.Sprintf("seconds:%d            \nnanos:%d", v.LastUpdated.Seconds, v.LastUpdated.Nanos), v.Content})
	}
	table.Render()
}
