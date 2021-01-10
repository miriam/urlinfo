package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

type UrlinfoController struct{}

func (u UrlinfoController) Get(c *gin.Context) {
	db := new(UrlinfoDb)
	originalUrl := parseUrl(c)

	c.JSON(http.StatusOK, gin.H{
		"blocklisted": db.isBlocklisted(originalUrl),
	})
}

func parseUrl(c *gin.Context) string {
	originalUrl := &url.URL{}
	originalUrl.Host = c.Param("hostnameAndPort")
	originalUrl.Path = c.Param("originalPath")
	originalUrl.RawQuery = c.Request.URL.RawQuery
	urlString := originalUrl.String()
	return urlString[2:len(urlString)]
}
