package cmd

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	todo "github.com/chinajuanbob/helloworld/pb"
	"github.com/chinajuanbob/helloworld/pkg/service"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
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
	Run: runGrpc,
}

func init() {
	serveCmd.AddCommand(grpcCmd)
	grpcCmd.PersistentFlags().String("address", "0.0.0.0:6666", "address to serve")
	viper.BindPFlag("address", grpcCmd.PersistentFlags().Lookup("address"))
}

func runGrpc(cmd *cobra.Command, args []string) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error(err)
		}
	}()

	s := grpc.NewService(
		micro.Name("todo"),
		micro.Address(viper.GetString("address")),
		micro.BeforeStart(func() error {
			glog.Info("BeforeStart")
			return nil
		}),
		micro.AfterStart(func() error {
			glog.Info("AfterStart")
			return nil
		}),
		micro.BeforeStop(func() error {
			glog.Info("BeforeStop")
			return nil
		}),
		micro.AfterStop(func() error {
			glog.Info("AfterStop")
			return nil
		}),
	)
	s.Init()

	server := service.TodoService{}
	server.Init()
	todo.RegisterTodoServiceHandler(s.Server(), &server)

	if err := s.Run(); err != nil {
		glog.Fatal(err)
	}
}
