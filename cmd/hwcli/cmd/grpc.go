package cmd

import (
	"context"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	todo "github.com/chinajuanbob/helloworld/pb"
	"github.com/davecgh/go-spew/spew"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/client/grpc"
	"github.com/spf13/viper"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: run,
}

func init() {
	clientCmd.AddCommand(grpcCmd)
	grpcCmd.PersistentFlags().String("address", "0.0.0.0:6666", "address of service")
	viper.BindPFlag("address", grpcCmd.PersistentFlags().Lookup("address"))

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run(cmd *cobra.Command, args []string) {
	glog.Info("grpc called")
	service := micro.NewService(
		micro.Name("todo"),
		micro.Client(grpc.NewClient()),
	)
	service.Init()

	c := todo.NewTodoService("todo", service.Client())

	rsp, err := c.List(context.Background(), &todo.ListTodoRequest{}, client.WithAddress(viper.GetString("address")))
	if err != nil {
		glog.Error(err)
	}
	spew.Dump(rsp)
}
