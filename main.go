package main

import (
	"github.com/gin-gonic/gin"
)

func setRedirectHTTPtoHTTPS() {
	httpRouter := gin.Default()
	httpRouter.GET("/*path", func(c *gin.Context) {
		c.Redirect(302, "https://localhost:8081"+c.Request.RequestURI)
	})
	go httpRouter.Run(":8080")
}

func main() {

	setRedirectHTTPtoHTTPS()

	httpsRouter := gin.Default()

	httpsRouter.GET("/", home)
	httpsRouter.RunTLS(":8081", "server.crt", "server.key")
}

func home(c *gin.Context) {
	c.JSON(200, "hola")
}
