package handlers

import (
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/kaanmertkoc/todo/handlers/types"
	"google.golang.org/api/iterator"
)



func ListTodosHandler(client *firestore.Client) func(c *gin.Context) {

	return func (c *gin.Context) {

		var result []types.Todo

		if(result == nil) {
			c.JSON(http.StatusOK, result)
			return
		}

		iter := client.Collection(types.TODO_COLLECTION).Documents(c)

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, "")
				return
			}

			result = append(result, types.Todo{
				ID: doc.Ref.ID,
				Title: doc.Data()["title"].(string),
				Description: doc.Data()["description"].(string),
				Completed: doc.Data()["completed"].(bool),
				CreatedAt: doc.Data()["createdAt"].(time.Time),
				UpdatedAt: doc.Data()["updatedAt"].(time.Time),
			})
		}
		c.JSON(http.StatusOK, result)
	}
}