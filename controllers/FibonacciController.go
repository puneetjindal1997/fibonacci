package controllers

import (
	"fibonacci/helper"
	"fibonacci/redis"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var Db redis.Database
var valueStartFrom = 0
var ttl = time.Minute * 15

// init function to first initialize the number
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	//	redis init
	databaseImplementation := os.Getenv("REDIS_DATABASE_NAME")
	log.Println(databaseImplementation)
	db, redisErr := redis.Factory(databaseImplementation)
	if redisErr != nil {
		panic(redisErr)
	}
	Db = db
	log.Println(db)
	redis.InsertDataTORedis(Db, "value", valueStartFrom, ttl)
}

// controller function containing business
// logic for getting fibonacci number
func FibonacciPrint(c *gin.Context) {

	// geting the value from reedis for fibonacci
	respValue, err := redis.GetDataTORedis(Db, "value")
	if err != nil {
		c.JSON(400, gin.H{"error": true, "data": err})
		return
	}
	respValueInt, err := strconv.Atoi(respValue.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": true, "data": err})
		return
	}
	// getting fibonacci number from helper
	value := helper.FibonacciRecursion(respValueInt)
	respValueInt += 1
	redis.InsertDataTORedis(Db, "value", respValueInt, ttl)
	c.JSON(http.StatusOK, gin.H{"value": value})
}
