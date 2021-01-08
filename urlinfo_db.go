package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"bufio"
	"fmt"
)

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

func isBlocklisted(url string) (bool) {
    val, err := rdb.SIsMember(ctx, "blocklist", url).Result()
    if err == redis.Nil {
        fmt.Println("URL does not exist")
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

