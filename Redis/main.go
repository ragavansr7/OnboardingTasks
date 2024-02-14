package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Create a Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis default address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Ping Redis to ensure the connection is established
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}
	// Route for posting a regular Redis key
	router.POST("/redis-post-key", func(c *gin.Context) {
		var request struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}
		err := rdb.Set(ctx, request.Key, request.Value, 0).Err()
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(200, gin.H{"message": "Key posted successfully"})
	})

	// Route for posting a Redis hash key
	router.POST("/redis-post-hashkey", func(c *gin.Context) {
		var request struct {
			HashName string `json:"hash_name"`
			Key      string `json:"key"`
			Value    string `json:"value"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}
		err := rdb.HSet(ctx, request.HashName, request.Key, request.Value).Err()
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(200, gin.H{"message": "Hash key posted successfully"})
	})

	// Route for fetching all Redis keys
	router.GET("/redis-keys", func(c *gin.Context) {
		keys, err := rdb.Keys(ctx, "*").Result()
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(200, gin.H{"keys": keys})
	})

	// Route for fetching hash keys from a specific hash
	router.GET("/redis-hash-keys/:hashName", func(c *gin.Context) {
		hashName := c.Param("hashName")
		hashKeys, err := rdb.HGetAll(ctx, hashName).Result()
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(200, gin.H{"hash_keys": hashKeys})
	})
	// Route for posting data to a Redis stream
	// Route for posting data to a Redis stream
	router.POST("/redis-post-stream/:streamName", func(c *gin.Context) {
		var request struct {
			StreamName string                 `json:"stream_name"`
			Data       map[string]interface{} `json:"data"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}

		// Convert the map[string]interface{} to a slice of string pairs
		var streamValues []string
		for k, v := range request.Data {
			streamValues = append(streamValues, k, fmt.Sprintf("%v", v))
		}

		// Add the data to the Redis stream
		_, err := rdb.XAdd(ctx, &redis.XAddArgs{
			Stream: request.StreamName,
			Values: streamValues,
		}).Result()
		if err != nil {
			fmt.Println("error adding data to stream:", err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(200, gin.H{"message": "Data posted to stream successfully"})
	})

	// Route for fetching entries from a Redis stream
	router.GET("/redis-streams/:streamName", func(c *gin.Context) {
		streamName := c.Param("streamName")
		streamData, err := rdb.XRange(ctx, streamName, "-", "+").Result()
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(200, gin.H{"stream_data": streamData})
	})

	// Route for setting a Redis key with expiry using POST
	router.POST("/redis-set-expiry", func(c *gin.Context) {
		var request struct {
			Key   string `json:"key"`
			Value string `json:"value"`
			TTL   int    `json:"ttl"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}
		err := rdb.Set(ctx, request.Key, request.Value, time.Duration(request.TTL)*time.Second).Err()
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(200, gin.H{"message": "Key set with expiry"})
	})

	// Route for deleting a Redis key using DELETE
	router.DELETE("/redis-delete/:key", func(c *gin.Context) {
		key := c.Param("key")
		deleted, err := rdb.Del(ctx, key).Result()
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(200, gin.H{"deleted": deleted})
	})

	// Run the Gin server
	router.Run(":8080")
}
