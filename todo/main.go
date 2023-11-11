package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/kaanmertkoc/todo/handlers"
)

func main() {

	client, err := createClient()
	if err != nil {
		log.Printf("erro creating client %v", err)
		return
	}

	r := gin.Default()

	r.GET("/api/health", handlers.HealthCheckHandler())

	// create a todo

	r.POST("/api/todo", handlers.CreateTodoHandler(client))

	r.GET("/api/todo", handlers.ListTodosHandler(client))

	r.GET("/api/todo/:id", handlers.GetTodoHandler(client))

	r.PUT("/api/todo/:id", handlers.UpdateTodoHandler(client))

	r.DELETE("/api/todo/:id", handlers.DeleteTodoHandler(client))

	r.Run("127.0.0.1:9090")

	defer client.Close()
}


func createClient() (*firestore.Client, error) {

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "gotutorial-6644c"}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Printf("error initializing app: %v\n", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Printf("error getting firestore client: %v\n", err)
		log.Fatalln(err)
	}

	return client, nil
}