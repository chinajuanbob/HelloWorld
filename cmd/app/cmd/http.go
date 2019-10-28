package cmd

import (
	"github.com/chinajuanbob/helloworld/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/micro/go-micro/web"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Run:   runHttp,
}

func init() {
	serveCmd.AddCommand(httpCmd)
	httpCmd.PersistentFlags().String("web", "0.0.0.0:9999", "address to serve")
	viper.BindPFlag("web", httpCmd.PersistentFlags().Lookup("web"))
	httpCmd.PersistentFlags().String("target", "0.0.0.0:6666", "address of grpc service")
	viper.BindPFlag("target", httpCmd.PersistentFlags().Lookup("target"))

}

func runHttp(cmd *cobra.Command, args []string) {
	// Create service
	s := web.NewService(
		web.Name("go.micro.api"),
		web.Address(viper.GetString("web")),
		web.BeforeStart(func() error {
			glog.Info("BeforeStart")
			return nil
		}),
		web.AfterStart(func() error {
			glog.Info("AfterStart")
			return nil
		}),
		web.BeforeStop(func() error {
			glog.Info("BeforeStop")
			return nil
		}),
		web.AfterStop(func() error {
			glog.Info("AfterStop")
			return nil
		}),
	)
	s.Init()

	// Create RESTful handler (using Gin)
	gin.SetMode(gin.DebugMode)
	httpservice := service.NewTodoHttpService(viper.GetString("target"))

	// Register Handler
	s.Handle("/", httpservice.GetRouter())

	// Run server
	if err := s.Run(); err != nil {
		glog.Fatal(err)
	}
}
