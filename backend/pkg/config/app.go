package config


import (
	"github.com/jinzhu/gorm"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "user:mypassword@tcp(db)/testdb?parseTime=true")
	if err != nil {
		// panic(err)
	}
	
	for d.DB().Ping() != nil {
		d, err = gorm.Open("mysql", "user:mypassword@tcp(db)/testdb?parseTime=true")
		fmt.Println("Attempting connection to db")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Connected")
	db = d
}

func GetDb() * gorm.DB {
	return db
}