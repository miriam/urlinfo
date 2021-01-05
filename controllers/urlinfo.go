package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	s "strings"
)

type UrlinfoController struct{}

var BlockList = [3]string {"google.com", "google.com/foobar", "yahoo.com?123=4"}

func (u UrlinfoController) Retrieve(c *gin.Context) {
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
	for i := 0; i < 3; i++ {
		if BlockList[i] == url {
			return true
		}
	}
	return false
}

