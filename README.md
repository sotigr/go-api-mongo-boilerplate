## Mongo find example

```go
import (  
    "fmt"
	"net/http" 
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	dbTools "main/database" 
)
 
type DbUser struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    string             `bson:"firstName"`
	EmailAddress string             `bson:"emailAddress"`
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := dbTools.Database()
		defer db.CancelContext()

		var user DbUser

		err := db.Current.Collection("user").FindOne(db.Context, bson.M{"emailAddress": "user@example.com"}).Decode(&user)

		if err != nil {
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, user)
	}
} 
```