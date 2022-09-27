package methods

import (
	"fmt"

	proto "example.com/todo/todo"
	"google.golang.org/grpc"
)

func Connection() proto.TodoClient {
	Conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	Client := proto.NewTodoClient(Conn)
	fmt.Println("Server started on port 8080")

	return Client
}

var Option string

func Test(op string) {
	Option = op

}
