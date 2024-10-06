package championships

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type service interface {
	createChampionship(ctx context.Context, championship Championship) (Championship, error)
	getChampionship(ctx context.Context, id string) (Championship, error)
}

type Controller struct {
	service service
}

func NewController(service service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) PostChampionship(ctx *gin.Context) {
	var req Championship
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	championship, err := c.service.createChampionship(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, championship)
}

func (c *Controller) GetChampionship(ctx *gin.Context) {
	championship, err := c.service.getChampionship(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if championship.isEmpty() {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, championship)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
