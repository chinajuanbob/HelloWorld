package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	todo "github.com/chinajuanbob/helloworld/pb"
	"github.com/chinajuanbob/helloworld/pkg/common"
	"github.com/chinajuanbob/helloworld/pkg/constant"
	_ "github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/client/grpc"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"
)

const (
	TodoURL   = "/todo"          // post, put
	TodosURL  = "/todos"         // get
	TodoIDURL = "/todo/{todoID}" // get, delete
)

type TodoHttpService struct {
	common.HttpService
	address string
	client  todo.TodoService
}

func NewTodoHttpService(address string) *TodoHttpService {
	service := micro.NewService(
		micro.Name("todo"),
		micro.Client(grpc.NewClient()),
	)
	service.Init()

	s := TodoHttpService{
		address: address,
		client:  todo.NewTodoService("todo", service.Client()),
	}
	s.HttpService.GetStatus = s.GetStatus
	s.HttpService.GetHealthz = s.GetHealthz
	s.Init()

	s.SetSwagger(
		"v1",
		s.addTodoEndpoint(),
		s.listTodosEndpoint(),
		s.updateTodoEndpoint(),
		s.deleteTodoEndpoint(),
	)
	return &s
}

func (s *TodoHttpService) GetHealthz() int {
	return 1
}

func (s *TodoHttpService) GetStatus(detail bool) gin.H {
	return gin.H{
		"a": 100,
	}
}

type TodoResult struct {
	common.CommonResult
	Data *todo.Todo `json:"data,omitempty"`
}

type TodosResult struct {
	common.CommonResult
	Data []*todo.Todo `json:"data,omitempty"`
}

type SuccessResult struct {
	common.CommonResult
	Data bool `json:"data,omitempty"`
}

func (s *TodoHttpService) addTodoEndpoint() *swagger.Endpoint {
	return endpoint.New("post", TodoURL, "add todo",
		endpoint.Tags("Todo"),
		endpoint.Description("the description of addTodo()"),
		endpoint.OperationID("addTodo"),
		endpoint.Body(todo.AddTodoRequest{}, "new content", true),
		endpoint.Handler(func(c *gin.Context) {
			bytes, err := c.GetRawData()
			if err != nil {
				common.ReturnError(c, err)
				return
			}
			var form todo.AddTodoRequest
			err = json.Unmarshal(bytes, &form)
			if err != nil {
				common.ReturnError(c, err)
				return
			}
			rsp, err := s.client.Add(context.Background(), &todo.AddTodoRequest{
				Content: form.Content,
			}, client.WithAddress(s.address))
			if err != nil {
				common.ReturnError(c, err)
				return
			}
			c.JSON(http.StatusOK, TodoResult{
				CommonResult: common.CommonResult{
					Status: constant.HTTPResultStatusOK,
				},
				Data: rsp.GetTodo(),
			})
			return
		}),
		endpoint.Response(http.StatusOK, TodoResult{}, "successful operation"),
	)
}

func (s *TodoHttpService) listTodosEndpoint() *swagger.Endpoint {
	return endpoint.New("get", TodosURL, "list todos",
		endpoint.Tags("Todo"),
		endpoint.Description("the description of listTodos()"),
		endpoint.OperationID("listTodos"),
		endpoint.Handler(func(c *gin.Context) {
			rsp, err := s.client.List(context.Background(), &todo.ListTodoRequest{}, client.WithAddress(s.address))
			if err != nil {
				common.ReturnError(c, err)
				return
			}
			todos := rsp.GetTodos()
			c.JSON(http.StatusOK, TodosResult{
				CommonResult: common.CommonResult{
					Status: constant.HTTPResultStatusOK,
				},
				Data: todos,
			})
		}),
		endpoint.Response(http.StatusOK, TodosResult{}, "successful operation"),
	)
}

func (s *TodoHttpService) updateTodoEndpoint() *swagger.Endpoint {
	return endpoint.New("put", TodoURL, "update todo",
		endpoint.Tags("Todo"),
		endpoint.Description("the description of updateTodo()"),
		endpoint.OperationID("updateTodo"),
		endpoint.Body(todo.UpdateTodoRequest{}, "new content", true),
		endpoint.Handler(func(c *gin.Context) {
			bytes, err := c.GetRawData()
			if err != nil {
				common.ReturnError(c, err)
				return
			}
			var form todo.UpdateTodoRequest
			err = json.Unmarshal(bytes, &form)
			if err != nil {
				common.ReturnError(c, err)
				return
			}
			rsp, err := s.client.Update(context.Background(), &todo.UpdateTodoRequest{
				Id:     form.Id,
				Status: form.Status,
			}, client.WithAddress(s.address))
			if err != nil {
				common.ReturnError(c, err)
				return
			}
			c.JSON(http.StatusOK, TodoResult{
				CommonResult: common.CommonResult{
					Status: constant.HTTPResultStatusOK,
				},
				Data: rsp.GetTodo(),
			})
		}),
		endpoint.Response(http.StatusOK, TodoResult{}, "successful operation"),
	)
}

func (s *TodoHttpService) deleteTodoEndpoint() *swagger.Endpoint {
	return endpoint.New("delete", TodoIDURL, "delete todo",
		endpoint.Tags("Todo"),
		endpoint.Description("the description of deleteTodo()"),
		endpoint.OperationID("deleteTodo"),
		endpoint.Path("todoID", "integer", "the todo id", true),
		endpoint.Handler(func(c *gin.Context) {
			todoID := c.Param("todoID")
			idx, err := strconv.Atoi(todoID)
			if err != nil {
				common.ReturnError(c, err)
				return
			}
			rsp, err := s.client.Delete(context.Background(), &todo.DeleteTodoRequest{
				Id: int32(idx),
			}, client.WithAddress(s.address))
			if err != nil {
				common.ReturnError(c, err)
				return
			}
			c.JSON(http.StatusOK, SuccessResult{
				CommonResult: common.CommonResult{
					Status: constant.HTTPResultStatusOK,
				},
				Data: rsp.Success,
			})
		}),
		endpoint.Response(http.StatusOK, SuccessResult{}, "successful operation"),
	)
}
