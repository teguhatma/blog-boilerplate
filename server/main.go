package main

import (
	"context"
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/teguhatma/blog-boilerplate/container"
	"github.com/teguhatma/blog-boilerplate/server/http"
)

func main() {
	ctx := context.Background()
	router := mux.NewRouter()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := container.CreateHTTPContainer(router); err != nil {
		log.Fatal(ctx, "Error in initRoutes Error %v", err)
	}

	httpServer, err := http.InitServer(router)
	if err != nil {
		log.Fatal(ctx, "Error while initialising server %s", err)
	}

	if err = httpServer.Start(ctx); err != nil {
		log.Fatal(ctx, "Error on starting up Http Server. %v", err)
	}
}
