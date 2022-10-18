package config

import (
	"fmt"
	"test/mnc/migration"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbUsername := "root"
	dbPassword := "Sql24"
	dbPort := "3306"
	dbHost := "127.0.0.1"
	dbName := "db_test_mnc"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migration.InitMigrate(db)
	return db
}
