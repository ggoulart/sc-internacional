package teams

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type service interface {
	createTeam(ctx context.Context, team Team) (Team, error)
}

type Controller struct {
	service service
}

func NewController(service service) *Controller {
	return &Controller{service: service}
}

func (c Controller) PostTeam(ctx *gin.Context) {
	var req Team
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	team, err := c.service.createTeam(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, team)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
