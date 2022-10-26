package routers

import (
	"net/http"

	"github.com/mazeem91/trackman-poc/application/controllers"
	"github.com/mazeem91/trackman-poc/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine, repo *repository.SQLite) {
	h := &controllers.Handler{
		Repository: repo,
	}

	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	route.POST("/locations", h.AddLocation)
	route.GET("/locations", h.GetLocations)

	route.POST("/skus", h.AddSku)
	route.GET("/skus", h.GetSkus)

	route.POST("/skus/:sku_id/assign/:location_id", h.AssignSkuLocation)

	//Add All route
	//TestRoutes(route)
}
