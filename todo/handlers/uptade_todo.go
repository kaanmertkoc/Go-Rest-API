package handlers

import (
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/kaanmertkoc/todo/handlers/types"
)


func UpdateTodoHandler(client *firestore.Client) func(c *gin.Context)  {
	return func(c *gin.Context) {
		id := c.Param("id")
		var todo types.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		todo.ID = id
		todo.UpdatedAt = time.Now()	

		_, err := client.Collection(types.TODO_COLLECTION).Doc(todo.ID).Set(c, map[string]interface{}{
			"title": todo.Title,
			"description": todo.Description,
			"completed": todo.Completed,
			"updatedAt": todo.UpdatedAt,
		}, firestore.MergeAll)
		
		if err != nil {
			log.Printf("error updating a todo: %v", err)
			c.JSON(500, gin.H{"error": "error updating a todo"})
			return
		}
		c.JSON(200, todo)
	}
}