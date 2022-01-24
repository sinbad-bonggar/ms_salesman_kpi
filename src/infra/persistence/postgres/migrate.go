package postgres

import (
	infra_model "github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/models"
	"gorm.io/gorm"
)

// Migrate represent migration schema models
func Migrate(db *gorm.DB) error {
	Order := infra_model.Order{}

	// auto migrate
	if err := db.AutoMigrate(
		&Order,
	); err != nil {
		return err
	}

	return nil
}
