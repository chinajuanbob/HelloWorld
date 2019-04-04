package service

import (
	"context"
	"errors"
	"time"

	todo "github.com/chinajuanbob/helloworld/pb"
	"github.com/golang/protobuf/ptypes"
)

type TodoService struct {
	todos map[int32]*todo.Todo
	next  int32
}

func (s *TodoService) Init() {
	s.todos = map[int32]*todo.Todo{}
	s.next = 1
}

func (s *TodoService) Add(ctx context.Context, req *todo.AddTodoRequest, rsp *todo.AddTodoResponse) error {
	tm, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return err
	}

	t := &todo.Todo{
		Id:          s.next,
		Content:     req.Content,
		LastUpdated: tm,
	}
	rsp.Todo = t
	s.todos[s.next] = t
	s.next++
	return nil
}

func (s *TodoService) Update(ctx context.Context, req *todo.UpdateTodoRequest, rsp *todo.UpdateTodoResponse) error {
	id := req.Id
	todo, ok := s.todos[id]
	if !ok {
		return errors.New("item not found")
	}
	tm, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return err
	}
	todo.Status = req.GetStatus()
	todo.LastUpdated = tm
	rsp.Todo = todo
	return nil
}

func (s *TodoService) List(ctx context.Context, req *todo.ListTodoRequest, rsp *todo.ListTodoResponse) error {
	rsp.Todos = []*todo.Todo{}
	for _, v := range s.todos {
		rsp.Todos = append(rsp.Todos, v)
	}
	return nil
}

func (s *TodoService) Delete(ctx context.Context, req *todo.DeleteTodoRequest, rsp *todo.DeleteTodoResponse) error {
	id := req.Id
	_, ok := s.todos[id]
	if ok {
		delete(s.todos, id)
	}
	rsp.Success = true
	return nil
}
