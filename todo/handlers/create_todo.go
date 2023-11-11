package handlers

import (
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/kaanmertkoc/todo/handlers/types"
)


func CreateTodoHandler(client *firestore.Client)  func(c *gin.Context) {

	return func(c *gin.Context) {
		var todo types.Todo

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		now := time.Now()
		todo.CreatedAt = now
		todo.UpdatedAt = now
	
		ref := client.Collection(types.TODO_COLLECTION).NewDoc()

		_, err := ref.Set(c, map[string]interface{}{
			"title": todo.Title,
			"description": todo.Description,
			"completed": false,
			"createdAt": todo.CreatedAt,
			"updatedAt": todo.UpdatedAt,
		})
		
		if err != nil {
				log.Printf("error adding a todo: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding a todo"})
				return
		}	

		c.JSON(http.StatusCreated, todo)

	}
}