package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx context.Context

func init() {
	// redis DB connection
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	ctx = context.Background()
}

func main() {
	// router
	router := gin.Default()

	// endpoints
	router.GET("/", getHome)
	router.GET("/:shorturl", getURL)
	router.GET("/newurl/", newURL)
	router.POST("/make:fullurl", setURL)

	// start server
	router.Run()
}

// root for basic functionality test
func getHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

// get full URL from short URL
func getURL(c *gin.Context) {

}

// testing functionality in browser for short url and
// DB connection
func newURL(c *gin.Context) {
	shorturl := makeShortLink()

	err := rdb.Set(ctx, shorturl, "https://www.google.com", 0).Err()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"shorturl": shorturl,
	})
}

// make short URL from full URL
func setURL(c *gin.Context) {
	shorturl := makeShortLink()

	c.JSON(http.StatusOK, gin.H{
		"shorturl": shorturl,
	})
}
