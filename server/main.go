package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"example.com/todo/cmd"
	"example.com/todo/data"
	"example.com/todo/todo"
	proto "example.com/todo/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedTodoServer
}

func (s *server) MarkTask(ctx context.Context, r *proto.Args) (*proto.Response, error) {
	id := r.GetId()
	var rec data.Task

	option := cmd.Selection()
	//var option string
	rec = data.Test2(id)
	a := rec.Id
	b := rec.Todo
	var c int64
	fmt.Printf("Option selected %+v", option)
	if option == "Pending" {
		c = 0
	} else {
		c = 1
	}
	task := data.Task{Id: a, Todo: b, Mark: c}
	data.Database.Save(task)
	return &todo.Response{
		Id:     task.Id,
		Task:   task.Todo,
		Status: task.Mark,
	}, nil
}
func (s *server) ListTask(ctx context.Context, r *proto.List) (*proto.ListResponse, error) {
	log.Println("ListTask() called")
	var records []data.Task

	var List []*proto.Task
	var rows int64

	records, rows = data.Test()
	for i := 0; i < int(rows); i++ {
		todos := &proto.Task{Id: records[i].Id, Task: records[i].Todo, Status: records[i].Mark}
		List = append(List, todos)
	}
	// for i := 0; i < len(List); i++ {
	// 	log.Println(List[i])
	// }
	return &todo.ListResponse{
		Items: List,
	}, nil

}

func (s *server) GetTask(ctx context.Context, r *proto.Args) (*proto.Response, error) {
	log.Println("GetTask() called")
	id := r.GetId()
	var rec data.Task
	rec = data.Test2(id)
	return &todo.Response{
		Id:     rec.Id,
		Task:   rec.Todo,
		Status: rec.Mark,
	}, nil
}

func (s *server) DeleteTask(ctx context.Context, r *proto.Args) (*proto.List, error) {
	id := r.GetId()
	// var rec data.Task
	data.DeleteTask(id)
	return &todo.List{
		Message: "Task Deleted",
	}, nil
}
func (s *server) AddTask(ctx context.Context, r *proto.Task) (*proto.Response, error) {
	log.Println("CreateItem() called")

	id, newTask, st :=
		r.GetId(),
		r.GetTask(),
		r.GetStatus()

		//todos.TodoList.AddItem(newTask)
	// rec := data.Task{}
	data.CreateTask(id, newTask, st)
	a, b, c := data.GetRecord()
	j := int64(a)
	return &todo.Response{
		Id:     j,
		Task:   b,
		Status: c,
	}, nil
}
func tests() {
	var id int64 = 7
	res := data.Test2(id)
	log.Printf("ReadAll result: <%+v>\n\n", res)

}

func main() {
	data.DataMigration()
	data.CreateTable()
	// tests()
	//data.GetRecord()
	//data.Test()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterTodoServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}
