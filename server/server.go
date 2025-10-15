package server

import (
	"os"
	"seleksi-javan/database"
	"seleksi-javan/handler"
	"seleksi-javan/middleware"
	"seleksi-javan/repository"
	uctask "seleksi-javan/usecase/uc_task"
	ucuser "seleksi-javan/usecase/uc_user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if os.Getenv("APP_ENV") == "development" {
		r.Use(cors.Default())
	}

	db := database.GetDBInstance()

	r.Use(middleware.GlobalExceptionHandler())

	//initialize repo and usecase
	userRepository := repository.NewUserRepository(db)
	userUsecase := ucuser.NewUserUsecase(userRepository)

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := uctask.NewTaskUsecase(taskRepository, userRepository)

	//initialize auth
	authMiddleware := middleware.NewAuthMiddleware(userUsecase)

	//initialize handler
	userHandler := handler.NewUserHandler(userUsecase)
	userHandler.Route(r, authMiddleware)

	taskHandler := handler.NewTaskHandler(taskUsecase)
	taskHandler.Route(r, authMiddleware)

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
