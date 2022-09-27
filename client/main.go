package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/todo/methods"
	proto "example.com/todo/todo"
)

func main() {
	Client := methods.Connection()
	// _=Client
	// 	var task string
	// 	var status int64 =0
	// 	fmt.Print("Enter the task: ")
	// 	fmt.Scan(&task)
	// 	// fmt.Print("Enter the status: ")
	// 	// fmt.Scan(&status)

	// //-----------------------------------------------------------------

	// 	// Call Create
	// 	req1 := proto.Task{
	// 		Task:   task,
	// 		Status: status,
	// 	}
	// 	a, err := Client.AddTask(ctx, &req1)
	// 	if err != nil {
	// 		log.Fatalf("Create failed: %v", err)
	// 	}
	// 	log.Printf("Create result: ID: %+v--- Task: %+v---Status: %+v\n\n", a.Id,a.Task,a.Status)
	//-----------------------------------------------------------------
	//call list

	// log.Printf("Create result: <%+v>\n\n", )

	//-----------------------------------------------------------------
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var id int64
	fmt.Println("Enter the ID of Task: ")
	fmt.Scan(&id)
	req := proto.Args{
		Id: id,
	}
	// Option := Selection()
	response, err := Client.MarkTask(ctx, &req)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	// log.Printf("Create result: ID: %+v--- Task: %+v---Status: %+v\n\n", response.Id,response.Task,response.Status)
	log.Println(response)
	// req := proto.Args{
	// 	Id: id,
	// }
	// response, err := Client.GetTask(ctx, &req)
	// if err != nil {
	// 	log.Fatalf("Create failed: %v", err)
	// }
	// log.Printf("Create result: ID: %+v--- Task: %+v---Status: %+v\n\n", response.Id,response.Task,response.Status)
	log.Println(response)

}
