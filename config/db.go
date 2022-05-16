package config

import (
	"Todoapp/model/entity"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB{
	// errEnv := godotenv.Load()
	// if errEnv != nil {
	// 	panic("fail to load .env")
	// }
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",host,user,pass,dbname,port,sslmode)
	
	fmt.Println(dsn)
	db, errDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("--------fail connect to DB-------")
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		panic(errDB)
	}
	fmt.Println("====================================")
	fmt.Println("--------success connect to DB-------")
	fmt.Println("====================================")
	db.AutoMigrate(&entity.Passnote{},&entity.Todolist{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB){
	dbcon, err := db.DB()
	if err != nil {
		panic("fail to close con database")
	}
	dbcon.Close()
}