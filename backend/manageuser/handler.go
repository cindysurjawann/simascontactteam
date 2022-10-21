package manageuser

import (
	"log"
	"net/http"

	"github.com/bagasalim/simas/custom"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetUser(c *gin.Context) {
	user, status, err := h.Service.GetUser()
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    user,
	})
}

func (h *Handler) GetUserReq(c *gin.Context) {
	var req GetUserRequest
	username := c.Query("username")
	if username == "" {
		messageErr := []string{"Input data not suitable"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	req = GetUserRequest{Username: username}
	user, status, err := h.Service.GetUserReq(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    user,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	

	var req UpdateUserRequest
	username := c.Query("username")

	if username == "" {
		messageErr := []string{"Param data not suitable"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		messageErr := custom.ParseError(err)
		if messageErr == nil {
			messageErr = []string{"Input data not suitable"}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}

	reqFix := UpdateUserRequest{
		Email: req.Email,
		Role: req.Role,
		Name: req.Name,
	}
	user, status, err := h.Service.UpdateUser(reqFix, username)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    user,
	})

}

func (h *Handler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	_, status, err := h.Service.DeleteUser(userId)
	if err != nil {
		log.Println("Error handler Delete : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "deleted success",
		"id":      userId,
	})
}
