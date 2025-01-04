package model

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"type:datetime(0)"`
	UpdatedAt time.Time      `gorm:"type:datetime(0)"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index" json:"-"`
}

var models = []interface{}{
	User{},
	Channel{},
}

func AutoMigrate(db *gorm.DB) {
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Printf("gorm autoMigrate err %s", err)
		}
	}
}
