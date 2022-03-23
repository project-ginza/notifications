package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/project-ginza/notifications/config"
	logging "github.com/project-ginza/notifications/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnector *gorm.DB
var loggingCh = logging.LoggingCh

func Connect(host string, user string, password string, dbName string, port int) *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func Initialise(c *config.Config) {
	host := c.DBHost
	username := c.DBUsername
	password := c.DBPassword
	dbName := c.DBName
	port := c.DBPort

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, username, password, dbName)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		fmt.Printf("Failed to connect to DB. %s\n", err.Error())
		return
	}
	DBConnector = db
}
