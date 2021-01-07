package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	s "strings"
	"os"
	"bufio"
	"fmt"
)

type UrlinfoController struct{}

var blocklist []string 
var blocklistLen int

func init() {
	var err error
	blocklistFilename := os.Getenv("BLOCKLIST_FILENAME")
	if blocklistFilename == "" {
		blocklistFilename = "blocklist.txt"
	}
	blocklist, err = readLines(blocklistFilename)
	if err != nil {
		fmt.Println("Error: could not open blocklist file", blocklistFilename)
		os.Exit(1)
	}
	blocklistLen = len(blocklist)
}

func (u UrlinfoController) Get(c *gin.Context) {
	originalUrl := parseUrl(c)
	
	c.JSON(http.StatusOK, gin.H{
		"blocklisted": isBlocklisted(originalUrl),
		"originalUrl": originalUrl,
	})
}

func parseUrl(c *gin.Context) (string) {
	originalUrl := &url.URL{}
	originalUrl.Host = c.Param("hostnameAndPort")
	originalUrl.Path = s.TrimSuffix(c.Param("originalPath"), "/")
	originalUrl.RawQuery = c.Request.URL.RawQuery
	urlString := originalUrl.String()
	return urlString[2:len(urlString)]
}

func isBlocklisted(url string) (bool) {
	for i := 0; i < blocklistLen; i++ {
		if blocklist[i] == url {
			return true
		}
	}
	return false
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, nil
}

