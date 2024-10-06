package teams

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type service interface {
	createTeam(ctx context.Context, team Team) (Team, error)
	getTeam(ctx context.Context, id string) (Team, error)
	getAllTeams(ctx context.Context) ([]Team, error)
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

func (c Controller) GetTeam(ctx *gin.Context) {
	team, err := c.service.getTeam(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if team.isEmpty() {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("team not found")))
		return
	}

	ctx.JSON(http.StatusOK, team)
}

func (c Controller) GetAllTeams(ctx *gin.Context) {
	teams, err := c.service.getAllTeams(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teams)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
