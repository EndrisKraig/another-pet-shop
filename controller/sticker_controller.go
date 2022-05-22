package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"playground.io/another-pet-store/service"
)

type StickerController interface {
	GetKits(c *gin.Context)
}

type SimpleStickerController struct {
	stickerService service.StickerService
}

func NewStickerController(stickerService service.StickerService) StickerController {
	return &SimpleStickerController{stickerService: stickerService}
}

func (controller *SimpleStickerController) GetKits(c *gin.Context) {
	service := controller.stickerService
	response, err := service.GetAllStickers()
	if err != nil {
		fmt.Print(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error during getting stickers!"})
		return
	}
	c.IndentedJSON(http.StatusOK, response)
}
