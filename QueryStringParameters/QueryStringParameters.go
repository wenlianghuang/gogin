package querystringparameters

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryStringParameters() {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Matt&lastname=Huang
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Matt")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)

	})

	router.Run(":5050")
}
