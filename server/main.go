package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/teguhatma/blog-boilerplate/container"
	"github.com/teguhatma/blog-boilerplate/server/http"
)

func main() {
	ctx := context.Background()
	router := mux.NewRouter()

	if err := godotenv.Load();  err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println(".env file loaded!")
	}

	if err := container.CreateHTTPContainer(router); err != nil {
		log.Fatal(ctx, "Error in initRoutes Error %v", err)
	} else {
		fmt.Println("routes loaded!")
	}

	fmt.Println("Initialising a server ...")
	httpServer, err := http.InitServer(router)
	if err != nil {
		log.Fatal(ctx, "Error while initialising server %s", err)
	} else {
		fmt.Println("Initialising a server done!")
	}
	
	fmt.Println("Server is running ...")
	if err = httpServer.Start(ctx); err != nil {
		log.Fatal(ctx, "Error on starting up Http Server. %v", err)
	}
}
