package service

import (
	"context"

	todo "github.com/chinajuanbob/helloworld/pb"
)

type TodoService struct{}

func (s *TodoService) Add(ctx context.Context, req *todo.AddTodoRequest, rsp *todo.AddTodoResponse) error {
	rsp.Todo = &todo.Todo{
		Content: req.Content,
	}
	return nil
}

func (s *TodoService) Update(ctx context.Context, req *todo.UpdateTodoRequest, rsp *todo.UpdateTodoResponse) error {
	return nil
}

func (s *TodoService) List(ctx context.Context, req *todo.ListTodoRequest, rsp *todo.ListTodoResponse) error {
	rsp.Todos = []*todo.Todo{{
		Content: "test",
	}}
	return nil
}

func (s *TodoService) Delete(ctx context.Context, req *todo.DeleteTodoRequest, rsp *todo.DeleteTodoResponse) error {
	rsp.Success = true
	return nil
}
