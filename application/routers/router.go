package routers

import (
	"github.com/mazeem91/trackman-poc/application/routers/middleware"
	"github.com/mazeem91/trackman-poc/infrastructure/database"
	"github.com/mazeem91/trackman-poc/infrastructure/repository"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoute() *gin.Engine {

	environment := viper.GetBool("DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	allowedHosts := viper.GetString("ALLOWED_HOSTS")
	router := gin.New()
	router.SetTrustedProxies([]string{allowedHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	repo := &repository.SQLite{DB: database.DB}
	RegisterRoutes(router, repo) //routes register

	return router
}
