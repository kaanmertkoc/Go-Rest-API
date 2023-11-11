package handlers

import (
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/kaanmertkoc/todo/handlers/types"
)


func DeleteTodoHandler(client *firestore.Client) func(c *gin.Context) {

	return func (c *gin.Context) {
		id := c.Param("id")

		_, err := client.Collection(types.TODO_COLLECTION).Doc(id).Delete(c)

		if err != nil {
			c.JSON(500, gin.H{"error": "error deleting a todo"})
			return
		}
		c.JSON(200, gin.H{"message": "todo deleted successfully"})
	}
}