package factory

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DbConnection return *gorm.DB to use gorm obj
func DbConnection() *gorm.DB {
	CONNECT := "%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql",
		fmt.Sprintf(CONNECT,
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_PROTCOL"),
			os.Getenv("DB_NAME")))

	if err != nil {
		panic("cannot connect db")
	}
	return db
}
