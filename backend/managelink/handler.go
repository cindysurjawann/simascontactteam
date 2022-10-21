package managelink

import (
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

func (h *Handler) GetLink(c *gin.Context) {
	var req GetLinkRequest
	linktype := c.Query("linktype")
	if linktype == "" {
		messageErr := []string{"Input data not suitable"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	req = GetLinkRequest{LinkType: linktype}
	link, status, err := h.Service.GetLink(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    link,
	})
}

func (h *Handler) UpdateLink(c *gin.Context) {
	

	var req UpdateLinkRequest
	linktype := c.Query("linktype")

	if linktype == "" {
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

	reqFix := UpdateLinkRequest{
		LinkType:  linktype,
		LinkValue: req.LinkValue,
		UpdatedBy: req.UpdatedBy,
	}
	link, status, err := h.Service.UpdateLink(reqFix)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    link,
	})

}
func (h *Handler) GetLinkRequest(c *gin.Context) {
	linktype, _ := c.Params.Get("type")
	req := GetLinkRequest{LinkType: linktype}
	link, status, err := h.Service.GetLink(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    link,
	})
}
