package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	httpsRouter := gin.Default()
	httpRouter := gin.Default()

	httpRouter.GET("/*path", func(c *gin.Context) {
		c.Redirect(302, "https://localhost:8081"+c.Request.RequestURI)
	})

	httpsRouter.GET("/", home)
	go httpsRouter.RunTLS(":8081", "server.crt", "server.key")
	httpRouter.Run(":8080")
}

func home(c *gin.Context) {
	c.JSON(200, "hola")
}
