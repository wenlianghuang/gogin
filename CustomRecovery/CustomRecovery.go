package customrecovery

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomRecovery() {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.CustomRecovery(func(c *gin.Context, recoverd interface{}) {
		if err, ok := recoverd.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.GET("/panic", func(c *gin.Context) {
		panic("foo")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ohai")
	})

	r.Run(":5050")
}
