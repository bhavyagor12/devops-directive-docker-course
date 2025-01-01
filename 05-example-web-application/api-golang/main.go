package main

import (
	"api-golang/database"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func init() {
	databaseUrl := "postgres://postgres:foobarbaz@localhost:5452/postgres"
	errDB := database.InitDB(databaseUrl)
	if errDB != nil {
		log.Fatalf("⛔ Unable to connect to database: %v\n", errDB)
	} else {
		log.Println("DATABASE CONNECTED 🥇")
	}

}

func main() {

	r := gin.Default()
	var tm time.Time

	r.GET("/", func(c *gin.Context) {
		tm = database.GetTime(c)
		c.JSON(200, gin.H{
			"api": "golang",
			"now": tm,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		tm = database.GetTime(c)
		c.JSON(200, "pong")
	})

	r.Run("0.0.0.0:8084") // listen and serve on 0.0.0.0:8080 (or "PORT" env var)
}
