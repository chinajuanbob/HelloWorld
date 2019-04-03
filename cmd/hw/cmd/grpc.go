package cmd

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	todo "github.com/chinajuanbob/helloworld/pb"
	"github.com/chinajuanbob/helloworld/pkg/service"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
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
	serveCmd.AddCommand(grpcCmd)
	grpcCmd.PersistentFlags().String("address", "0.0.0.0:6666", "address to serve")
	viper.BindPFlag("address", grpcCmd.PersistentFlags().Lookup("address"))
}

func run(cmd *cobra.Command, args []string) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error(err)
		}
	}()

	s := grpc.NewService(
		micro.Name("todo"),
		micro.Address(viper.GetString("address")),
	)
	s.Init()

	todo.RegisterTodoServiceHandler(s.Server(), new(service.TodoService))

	if err := s.Run(); err != nil {
		glog.Fatal(err)
	}
}
