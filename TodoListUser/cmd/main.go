package main

import (
	"github.com/North-al/grpcTodoList/TodoListUser/config"
	"github.com/North-al/grpcTodoList/TodoListUser/internal/server"
)

func main() {
	config.Init()
	server.InitDatabase()

}
