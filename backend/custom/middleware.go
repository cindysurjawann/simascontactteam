package custom

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MiddleWareClass interface {
	Auth(c *gin.Context)
	IsAdmin(c *gin.Context)
	IsCS(c *gin.Context)
}
type MiddleWare struct {
}

const (
	message = "You not have access"
)

func (m *MiddleWare) Auth(c *gin.Context) {
	auth := c.Request.Header["Authorization"]
	if len(auth) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": message,
		})
		c.Abort()
		fmt.Println("no auth")
		return
	}
	token := auth[0]
	dataUser, err := ClaimToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": message,
		})
		c.Abort()
		return
	}
	c.Set("user", dataUser)
	c.Next()
}
func (m *MiddleWare) IsAdmin(c *gin.Context) {
	data, exist := c.Get("user")
	if exist == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": message,
		})
		c.Abort()
		return
	}
	var dataUser DataJWT = data.(DataJWT)
	if dataUser.Role != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You not have access, for admin only",
		})
		c.Abort()
	}
	c.Next()
}
func (m *MiddleWare) IsCS(c *gin.Context) {
	data, exist := c.Get("user")
	if exist == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": message,
		})
		c.Abort()
		return
	}
	var dataUser DataJWT = data.(DataJWT)
	if dataUser.Role != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You not have access, for cs only",
		})
		c.Abort()
		return
	}
	c.Next()
}
