package main

import (
	"amandeepsTlgt/REST-API/controllers"
	"amandeepsTlgt/REST-API/services"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*
@Author: DevProblems(Sarang Kumar)
@YTChannel: https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ
*/
var (
	server         *gin.Engine
	userService    services.UserServiceInterface
	userController controllers.UserController
	ctx            context.Context
	userCollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	ENV := os.Getenv("ENV")
	// checking currentActive environement
	currentActiveEnvironment := getCurrentEnv(ENV)
	envFile := currentActiveEnvironment + ".env"

	// loading env file
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userCollection = mongoclient.Database("REST-API-GIN").Collection("users")
	userService = services.NewUserServiceInstance(userCollection, ctx)
	userController = controllers.NewUserControllerInstance(userService)
	server = gin.Default()
	// enabled cors
	server.Use(cors.Default())
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/")
	userController.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":" + os.Getenv("APP_PORT")))

}

func getCurrentEnv(env string) string {
	switch env {
	case "dev":
		return "dev"
	case "stage":
		return "stage"
	case "prod":
		return "prod"
	default:
		return "dev"
	}
}
