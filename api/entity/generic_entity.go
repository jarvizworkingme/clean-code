package entity

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func DropAndCreateDb(db *gorm.DB) {
	migrator := db.Migrator()

	_ = migrator.DropTable(&Customer{})
	_ = migrator.CreateTable(&Customer{})

	logrus.Info("Create customer example dummy data")
	for i := 1; i < 7; i++ {
		db.Create(&Customer{
			Email: fmt.Sprintf("%d@dummy.dummy", i),
		})
	}

	_ = migrator.DropTable(&Loan{})
	_ = migrator.CreateTable(&Loan{})

	_ = migrator.DropTable(&Schedule{})
	_ = migrator.CreateTable(&Schedule{})

	_ = migrator.DropTable(&Payment{})
	_ = migrator.CreateTable(&Payment{})

}
