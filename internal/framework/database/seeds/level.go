package seeds

import (
	"go-hospital-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func levelSeeder() Seed {
	seeds := []models.Level{
		{Name: "Admin"},
		{Name: "Doctor"},
		{Name: "Nurse"},
	}
	model := &models.Level{}

	return Seed{
		models: model,
		run: func(db *gorm.DB) (err error) {
			for _, v := range seeds {
				err = db.Create(&v).Error
			}
			return
		},
	}
}
