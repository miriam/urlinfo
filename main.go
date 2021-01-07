package main

import "github.com/gin-gonic/gin"
import "magoldbe/urlinfo/controllers"

func setupRouter() *gin.Engine {
	router := gin.Default()
	urlinfo := new(controllers.UrlinfoController)
	router.GET("/urlinfo/1/:hostnameAndPort/*originalPath", urlinfo.Get)
	router.GET("/urlinfo/1/:hostnameAndPort", urlinfo.Get)

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}

