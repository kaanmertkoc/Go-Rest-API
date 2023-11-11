package handlers

import (
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/kaanmertkoc/todo/handlers/types"
)



func GetTodoHandler(client *firestore.Client) func(c *gin.Context) {

	return func (c *gin.Context) {

		id := c.Param("id")

		dsnap, err := client.Collection(types.TODO_COLLECTION).Doc(id).Get(c)

		if(dsnap == nil) {
			c.JSON(http.StatusNotFound, "")
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
			return
		}

		doc := dsnap.Data()
		todo := &types.Todo{
			ID: id,
			Title: doc["title"].(string),
			Description: doc["description"].(string),
			Completed: doc["completed"].(bool),
			CreatedAt: doc["createdAt"].(time.Time),
			UpdatedAt: doc["updatedAt"].(time.Time),
		}
		c.JSON(http.StatusOK, todo)
	}
}