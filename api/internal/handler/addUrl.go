package handler

import (
	"github.com/gin-gonic/gin"
	"go_pg_s3_efk/domain"
	"net/http"
)

func (h *Handler) addUrl(c *gin.Context) {
	var input domain.EncoreTab

	err := c.BindJSON(&input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.InsertDBurl(input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "OK"})
}
