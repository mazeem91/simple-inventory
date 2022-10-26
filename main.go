package main

import (
	"time"

	"github.com/mazeem91/trackman-poc/application/routers"
	"github.com/mazeem91/trackman-poc/config"
	"github.com/mazeem91/trackman-poc/infrastructure/database"
	"github.com/mazeem91/trackman-poc/infrastructure/logger"
	"github.com/mazeem91/trackman-poc/infrastructure/migrations"

	"github.com/spf13/viper"
)

func main() {

	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "UTC")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	masterDSN, replicaDSN := config.DbConfiguration()

	if err := database.DbConnection(masterDSN, replicaDSN); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}
	//later separate migration
	migrations.Migrate()

	router := routers.SetupRoute()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
