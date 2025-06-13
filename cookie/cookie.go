package cookie

import (
	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, name, value string, domain string, maxAge int) {
	c.SetCookie(name, value, maxAge, "/", domain, false, true)
}

func DeleteCookie(c *gin.Context, name string) {
	c.SetCookie(name, "", -1, "/", "localhost", false, true)
}
