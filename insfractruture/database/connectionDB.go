package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/safe_msvc_user/insfractruture/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println(errEnv.Error())
	}
	DB_NAME := os.Getenv("DB_NAME")

	strConnection := CreateDatabase()

	dsn := fmt.Sprintf(strConnection+" dbname=%s", DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}
	MigrateDatabase(db)
	return db
}
func CreateDatabase() string {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println(errEnv.Error())
	}
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")
	//DB_NAME := os.Getenv("DB_NAME")
	strConnection := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_PORT, DB_SSLMODE)

	_, err := gorm.Open(postgres.Open(strConnection), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}
	/*query := fmt.Sprintf("SELECT 1 FROM  pg_database WHERE datname ='%s'", DB_NAME)
	errd := db.Exec(query)
	if errd.RowsAffected == 0 {
		// Crear la base de datos usando una consulta SQL cruda
		createDBSQL := fmt.Sprintf("CREATE DATABASE %s", DB_NAME)
		err = db.Exec(createDBSQL).Error
		if err != nil {
			log.Println(err.Error())
		}
		log.Printf("Base de datos '%s' creada exitosamente.\n", DB_NAME)
	}*/
	return strConnection
}

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&entities.School{})
}
func CloseConnection() {
	var db *gorm.DB = DatabaseConnection()
	dbSQL, err := db.DB()
	if err != nil {
		log.Println(err.Error())
	}
	dbSQL.Close()
}
