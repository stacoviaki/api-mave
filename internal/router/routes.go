package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stacoviaki/api-mave/db"
	"github.com/stacoviaki/api-mave/internal/controller"
	"github.com/stacoviaki/api-mave/internal/repositories"
	"github.com/stacoviaki/api-mave/internal/usecases"
)

func Routes() {
	server := gin.Default()
	server.Use(cors.Default())

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Users
	// Camada de Repositories
	UserRepositories := repositories.NewUserRepositories(dbConnection)

	// Camada de Usecase
	UserUseCase := usecases.NewUserUseCase(UserRepositories)

	// Camada de Controller
	UserController := controller.NewUserController(UserUseCase)

	// Chamadas (Endpoints)
	server.GET("/users", UserController.GetUsers)
	server.GET("/user/:userId", UserController.GetUserById)
	server.POST("/user", UserController.CreateUser)
	server.PUT("/users/:userId", UserController.UpdateUser)
	server.DELETE("/user/:userId", UserController.DeleteUser)

	server.Run(":9000")
}
