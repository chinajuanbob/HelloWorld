module github.com/chinajuanbob/helloworld

go 1.12

require (
	github.com/DeanThompson/ginpprof v0.0.0-20170218162546-8c0e31bfeaa8
	github.com/cpuguy83/go-md2man v1.0.10 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/gzip v0.0.1
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.1
	github.com/itsjamie/gin-cors v0.0.0-20160420130702-97b4a9da7933
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/micro/go-grpc v1.0.0
	github.com/micro/go-micro v1.0.0
	github.com/micro/go-plugins v0.24.1
	github.com/micro/go-web v0.6.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.1
	github.com/savaki/swag v0.0.0-20170722173931-3a75479e44a3
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.2
)

replace(
     golang.org/x/sync => github.com/golang/sync latest
     k8s.io/api => github.com/kubernetes/api latest
     k8s.io/apimachinery => github.com/kubernetes/apimachinery latest
     k8s.io/utils => github.com/kubernetes/utils latest
)
