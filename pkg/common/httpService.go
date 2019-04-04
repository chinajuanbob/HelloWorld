package common

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/DeanThompson/ginpprof"
	"github.com/chinajuanbob/helloworld/pkg/constant"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	cors "github.com/itsjamie/gin-cors"
	"github.com/savaki/swag"
	"github.com/savaki/swag/swagger"
)

type HttpRange struct {
	Start, Length int64
}

func (r *HttpRange) ContentRange(size int64) string {
	return fmt.Sprintf("bytes %d-%d/%d", r.Start, r.Start+r.Length-1, size)
}

type CommonResultError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CommonResult struct {
	Status string             `json:"status"`
	Error  *CommonResultError `json:"error,omitempty"`
	Data   gin.H              `json:"data,omitempty"`
}

func ReturnError(c *gin.Context, err error) {
	glog.Errorln(c.Request.URL, err.Error())
	c.JSON(http.StatusOK, CommonResult{
		Status: constant.HTTPResultStatusError,
		Error: &CommonResultError{
			Message: err.Error(),
		},
		Data: nil,
	})
}

type HttpService struct {
	secret      string
	authorizeFn func(c *gin.Context)
	router      *gin.Engine
	GetHealthz  func() int
	GetStatus   func(detail bool) gin.H
}

func (s *HttpService) SetSecret(str string) {
	s.secret = str
}

func (s *HttpService) Init() {
	//define routes
	router := gin.New()
	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, HEAD",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	router.Use(gin.Recovery())
	// url: /debug/pprof/
	ginpprof.Wrapper(router)

	// healthz for service
	router.GET(constant.HealthzURL, func(c *gin.Context) {
		c.String(http.StatusOK, "gin_healthz %v", s.GetHealthz())
	})
	// status for service
	router.GET(constant.StatusURL, func(c *gin.Context) {
		c.JSON(http.StatusOK, CommonResult{
			Status: constant.HTTPResultStatusOK,
			Error:  nil,
			Data:   s.GetStatus(true),
		})
	})

	s.router = router
	s.authorizeFn = func(c *gin.Context) {
		// get token from header
		header := c.Request.Header
		t := header.Get("authorization")
		if t == "" {
			t = header.Get("Authorization")
		}
		// if fail, get token from query
		if t == "" || len(t) < 8 {
			t = c.Query("token")
		} else {
			t = t[7:]
		}

		// if still fail, return 401
		if t == "" {
			glog.Errorln("missing token")
			glog.Errorln(header)
			c.AbortWithStatus(401)
			return
		}
		// Bearer Token
		claim, err := ParseToken([]byte(s.secret), t)
		if err != nil {
			glog.Errorf("parse token failed: %s \n", err.Error())
			c.AbortWithStatus(401)
			return
		}
		// glog.Info(claim)
		c.Set("userID", claim.UserID)
	}
}

func (s *HttpService) GetRouter() *gin.Engine {
	return s.router
}

func (s *HttpService) GetRouterGroup(version string) *gin.RouterGroup {
	return s.router.Group(fmt.Sprintf("/api/%s", version))
}

func (s *HttpService) GetAuthorizeFn() func(c *gin.Context) {
	return s.authorizeFn
}

var zipFn = gzip.Gzip(gzip.DefaultCompression)

func (s *HttpService) SetSwagger(version string, endpoints ...*swagger.Endpoint) {
	api := swag.New(
		swag.Title("Hello World"),
		swag.Version(version),
		swag.Description(""),
		swag.ContactEmail(""),
		swag.License("HelloWorld", ""),
		swag.TermsOfService(""),
		swag.BasePath(fmt.Sprintf("/api/%s", version)),
		swag.Tag("Hello World", "", swag.TagURL("http://localhost")), //must have one or swagger-ui hit issues
		swag.Endpoints(endpoints...),
		swag.SecurityScheme("token", swagger.APIKeySecurity("Authorization", "header")),
	)
	api.Walk(func(path string, endpoint *swagger.Endpoint) {
		h := endpoint.Handler.(func(c *gin.Context))
		s.router.Handle(endpoint.Method, FixGinParam(path), zipFn, h)
	})
	enableCors := true
	s.router.GET(fmt.Sprintf("/api/%s/swagger.json", version), gin.WrapH(api.Handler(enableCors)))
}

func FixGinParam(url string) string {
	strs := strings.Split(url, "/")
	for idx, v := range strs {
		if len(v) < 3 {
			continue
		}
		if v[0] == '{' && v[len(v)-1] == '}' {
			strs[idx] = ":" + v[1:len(v)-1]
		}
	}
	return strings.Join(strs, "/")
}
