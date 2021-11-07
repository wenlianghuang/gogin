package parameterinpath

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParameterInPath() {
	router := gin.Default()

	// This handler will match /user/matt but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" //true
		c.String(http.StatusOK, "%t", b)
	})

	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available group are [...]")
	})

	router.Run(":5050")
}
