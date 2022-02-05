package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

// var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func GoDatabaseCreate(w http.ResponseWriter, r *http.Request) {
	/*
		db.AutoMigrate(&GoTestModel{})

		goTestModel := GoTestModel{Name: "Luke", Rok: "2022"}
		db.Create(&goTestModel)

		if err := db.Create(&goTestModel).Error; err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(goTestModel)

		fmt.Println("Fields added", goTestModel)
	*/
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var pocet int
	db.Exec("INSERT INTO data.Json(json) VALUES ('dalsi zaznam');")
	err = db.QueryRow("SELECT COUNT(id) FROM data.Json;").Scan(&pocet)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Zapis do db.")

}

/*
func Hlav() {
	http.HandleFunc("/createstuff", GoDatabaseCreate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/
