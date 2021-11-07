package securejson

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SecureJSON() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		name := []string{"lena", "austin", "foo"}

		c.SecureJSON(http.StatusOK, name)
	})

	r.Run(":5050")
}
