package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nerigleston/primeiro-crud-go/src/configuration/database/mongodb"
	"github.com/nerigleston/primeiro-crud-go/src/configuration/logger"
	"github.com/nerigleston/primeiro-crud-go/src/controller"
	"github.com/nerigleston/primeiro-crud-go/src/controller/routes"
	"github.com/nerigleston/primeiro-crud-go/src/model/repository"
	"github.com/nerigleston/primeiro-crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	logger.Info("About to start user application")

	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
