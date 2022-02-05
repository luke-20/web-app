package database

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseSettings struct {
	ConnectionString string
}
type Database struct {
	Settings *DatabaseSettings
}

type GoTestModel struct {
	Name string
	Rok  string
}

var dsn = "admin:admin@tcp(192.168.55.9:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func GoDatabaseCreate(w http.ResponseWriter, r *http.Request) {
	goTestModel := GoTestModel{Name: "Luke", Rok: "2022"}
	db.Create(&goTestModel)

	if err := db.Create(&goTestModel).Error; err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(goTestModel)

	fmt.Println("Fields added", goTestModel)
}

func Hlav() {
	http.HandleFunc("/createstuff", GoDatabaseCreate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
