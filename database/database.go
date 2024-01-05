package database

import (
	"chg-gera-script-brad/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComDB() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Fatal("Erro ao conectar com DB: ", err.Error())
	}
	var chg models.Chg
	var firewalls models.Firewall
	DB.AutoMigrate(&chg)
	DB.AutoMigrate(&firewalls)
}
