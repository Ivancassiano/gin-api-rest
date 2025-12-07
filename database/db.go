package database

import (
	"github.com/Ivancassiano/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost port=5432 user=root dbname=root password=root sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		panic("Falha ao conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
