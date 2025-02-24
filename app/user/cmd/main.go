package main

import (
	"github.com/North-al/grpcTodoList/app/user/config"
	"github.com/North-al/grpcTodoList/app/user/internal/data/migration"
	"github.com/North-al/grpcTodoList/app/user/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	server.InitDatabase()
	if gin.Mode() == gin.DebugMode {
		migration.Migration()
	}
}
