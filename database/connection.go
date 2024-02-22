package database

import (
	"fmt"
	"log"

	"github.com/daniir/go_api/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(c *settings.ConfigSrv) (*gorm.DB, error) {
	var DNS = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DB.Host,
		c.DB.User,
		c.DB.Password,
		c.DB.Name,
		c.DB.Port)
	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("Db error connection")
		return nil, err
	} else {
		log.Print("Connected with postgres")
		return db, nil
	}
}
