package main

import (
	"context"
	"log"

	"github.com/gorilla/mux"
	"github.com/teguhatma/blog-boilerplate/container"
	"github.com/teguhatma/blog-boilerplate/server/http"
)

func main() {
	ctx := context.Background()
	router := mux.NewRouter()

	err := container.CreateHTTPContainer(router)
	if err != nil {
		log.Fatal(ctx, "Error in initRoutes Error %v", err)
	}

	httpServer, err := http.InitServer(router)
	if err != nil {
		log.Fatal(ctx, "Error while initialising server %s", err)
	}

	err = httpServer.Start(ctx)
	if err != nil {
		log.Fatal(ctx, "Error on starting up Http Server. %v", err)
	}
}
