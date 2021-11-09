package cookieex

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CookieEx() {
	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s\n", cookie)
	})

	router.Run(":5050")
}
