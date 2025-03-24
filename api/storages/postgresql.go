package storages

import (
	"clean-code/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDbConnection(postgresql config.Postgresql) *gorm.DB {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=public sslmode=disable", postgresql.Host, postgresql.Port, postgresql.User, postgresql.Password, postgresql.Name)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		logrus.Error("Failed to connect to database")
		panic(err)
	}

	return db
}
