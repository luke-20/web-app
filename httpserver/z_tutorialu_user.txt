package httpserver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"

*/
// _
// "gorm.io/gorm"

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {

	/*
		db, err := sql.Open("admin:admin@tcp(192.168.55.9)/db?charset=utf8mb4&parseTime=True&loc=Local")
	*/
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	// db, err = sql.Open("mysql", "admin:admin@tcp(192.168.55.9:3308)/db")

	dsn := "admin:admin@tcp(192.168.55.9:3308)/data?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DB!")
	}
	// defer db.Close()

	db.AutoMigrate(&User{})

}

type JsonData struct {
	ID          int
	Json        string
	TimeCreated string
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "All users endpoint hit")
	dsn := "admin:admin@tcp(192.168.55.9:3308)/data?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DB!")
	}
	// defer db.Close()

	var pocetZaznamu int
	err = db.QueryRow("SELECT COUNT(ID) FROM data.Json;").Scan(&pocetZaznamu)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "pocet zaznamu v db: %v\n", pocetZaznamu)

	// json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "New user endpoint hit")
	db, err := sql.Open("mysql", "admin:admin@tcp(192.168.55.9:3308)/db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DB!")
	}
	defer db.Close()

	db.Exec("INSERT INTO data.Json (json) VALUES ('Lukas');")
	fmt.Fprintf(w, "Uživatel vytvořen.")
}

/*
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Delete user endpoint hit")
	db, err := sql.Open("mysql", "admin:admin@tcp(192.168.55.9:3308)/db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DB!")
	}
	defer db.Close()
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Update user endpoint hit")
	db, err := sql.Open("mysql", "admin:admin@tcp(192.168.55.9:3308)/db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DB!")
	}
	defer db.Close()
}
*/
