// main.go (web-server)

package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {
	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()

	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	s.GET("/count", count)

	s.Run(":8000")

}

// Handler for the count request
func count(c *gin.Context) {
	val, err := client.Incr("count").Result()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"count": val})
}