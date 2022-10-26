package migrations

import (
	"github.com/mazeem91/trackman-poc/domain/models"
	"github.com/mazeem91/trackman-poc/infrastructure/database"
)

// Migrate Add list of model add for migrations
// TODO later separate migration each models
func Migrate() {
	var migrationModels = []interface{}{&models.Location{}, &models.Area{}, &models.Sku{}, models.SkuLocation{}}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
