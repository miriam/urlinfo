package main

import "github.com/gin-gonic/gin"

func init() {
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	urlinfo := new(UrlinfoController)
	router.GET("/urlinfo/1/:hostnameAndPort/*originalPath", urlinfo.Get)
	router.GET("/urlinfo/1/:hostnameAndPort", urlinfo.Get)

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
