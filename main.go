package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// DB & context initialisation
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
	router.GET("/newurl", newURL)

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
	// get short url param
	shorturl := c.Param("shorturl")

	// retrieve full url from DB
	fullurl, err := rdb.Get(ctx, shorturl).Result()
	if err != nil {
		// return full url
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
		})
	} else {
		// redirect to full URL
		c.Redirect(http.StatusFound, fullurl)
	}
}

// make shorturl for provided fullurl
func newURL(c *gin.Context) {
	// get fullurl provided
	fullurl := c.Query("url")

	if fullurl != "" {
		// get new short link
		shorturl := makeShortLink()

		// if the key already exists, make a new one
		exists, err := rdb.Exists(ctx, shorturl).Result()
		if err != nil {
			panic(err)
		}
		for exists > 0 {
			shorturl = makeShortLink()
		}

		// test DB functionality and add google.com as placeholder for now
		err = rdb.Set(ctx, shorturl, fullurl, 0).Err()
		if err != nil {
			panic(err)
		}

		// expire the link within 1 week
		err = rdb.Expire(ctx, shorturl, 7*24*time.Hour).Err()
		if err != nil {
			panic(err)
		}

		// return new shortlink
		c.JSON(http.StatusOK, gin.H{
			"shorturl": shorturl,
		})
	} else {
		// return error message
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
	}
}
