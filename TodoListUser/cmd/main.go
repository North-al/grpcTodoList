package main

import (
	"github.com/North-al/grpcTodoList/TodoListUser/config"
	"github.com/North-al/grpcTodoList/TodoListUser/internal/data/migration"
	"github.com/North-al/grpcTodoList/TodoListUser/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	server.InitDatabase()
	if gin.Mode() == gin.DebugMode {
		migration.Migration()
	}

}
