package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	dbTools "main/database"
)

func GetStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := dbTools.Database()
		defer db.CancelContext()

		if err := dbTools.MongoMainInstance.Ping(db.Context, readpref.Primary()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "DbPingTimeout"})
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
