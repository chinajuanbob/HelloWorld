package cmd

import (
	"github.com/chinajuanbob/helloworld/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/micro/go-web"
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
	httpCmd.PersistentFlags().String("address", "0.0.0.0:9999", "address to serve")
	viper.BindPFlag("address", httpCmd.PersistentFlags().Lookup("address"))
}

func runHttp(cmd *cobra.Command, args []string) {
	// Create service
	s := web.NewService(
		web.Name("go.micro.api"),
		web.Address(viper.GetString("address")),
	)
	s.Init()

	// Create RESTful handler (using Gin)
	gin.SetMode(gin.DebugMode)
	httpservice := service.NewHealthService()

	// Register Handler
	s.Handle("/", httpservice.GetRouter())

	// Run server
	if err := s.Run(); err != nil {
		glog.Fatal(err)
	}
}
