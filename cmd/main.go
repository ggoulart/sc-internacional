package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sc-internacional/internal/clients/mongodb"
	"sc-internacional/internal/teams"
)

func main() {
	r := gin.Default()

	//TODO: refactor this mongo block
	config, err := mongodb.NewConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	mongodbClient, _ := mongodb.NewMongoClient(config)
	defer mongodbClient.MongoClient.Disconnect(context.Background())
	db := mongodbClient.Database()

	teamRepository := teams.NewRepository(db)
	teamService := teams.NewService(teamRepository)
	teamController := teams.NewController(teamService)

	routers(r, teamController)

	r.Run()
}

func routers(r *gin.Engine, controllerTeam *teams.Controller) {
	r.POST("/teams", controllerTeam.PostTeam)
	r.GET("/teams/:id", controllerTeam.GetTeam)
	r.GET("/teams", controllerTeam.GetAllTeams)
	r.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "VAMO COLORADO!!"}) })
}
