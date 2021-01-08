package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"bufio"
	"fmt"
	"context"
	"github.com/go-redis/redis/v8"

)

type UrlinfoController struct{}

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
        	Addr:     "localhost:6379",
         	Password: "",
         	DB:       0,
   	})

	var err error
	blocklistFilename := os.Getenv("BLOCKLIST_FILENAME")
	if blocklistFilename == "" {
		blocklistFilename = "blocklist.txt"
	}
	err = readLines(blocklistFilename)
	if err != nil {
		fmt.Println("Error: could not load blocklist")
		panic(err)
		os.Exit(1)
	}
}

func (u UrlinfoController) Get(c *gin.Context) {
	originalUrl := parseUrl(c)
	
	c.JSON(http.StatusOK, gin.H{
		"blocklisted": isBlocklisted(originalUrl),
	})
}

func parseUrl(c *gin.Context) (string) {
	originalUrl := &url.URL{}
	originalUrl.Host = c.Param("hostnameAndPort")
	originalUrl.Path = c.Param("originalPath")
	originalUrl.RawQuery = c.Request.URL.RawQuery
	urlString := originalUrl.String()
	return urlString[2:len(urlString)]
}

func isBlocklisted(url string) (bool) {
    val, err := rdb.SIsMember(ctx, "blocklist", url).Result()
    fmt.Println("looked up key2: %s %s", val, url)
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } 
	return val 
}

func readLines(path string) (error) {
    file, err := os.Open(path)
    if err != nil {
        return err
    }

    defer file.Close()

	var line string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line = scanner.Text()
		_, err := rdb.SAdd(ctx, "blocklist", line).Result()
		if err != nil {
			return err
		}
    }
    return nil
}

