package migration

import (
	"fmt"

	"github.com/North-al/grpcTodoList/TodoListUser/internal/data/models"
	"github.com/North-al/grpcTodoList/TodoListUser/internal/server"
)

func Migration() {
	err := server.DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&models.UserModel{},
	)
	if err != nil {
		panic(fmt.Errorf("auto migrate failed: %v", err))
	}
}
