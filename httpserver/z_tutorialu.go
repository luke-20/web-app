package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloMainpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello mainpage!")
}

func handleRequests() {
	// router zpracovava requesty a routuje na funkce ktere maji resit jednotlive requesty(asi na tech endpointech)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloMainpage).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/create/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/delete/{name}/{email}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/update/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":80", myRouter))

}

func Hlavni() {
	fmt.Printf("Hlavni funkce z tutorialu")
	// po spusteni se pusti migrace - vytvori to tabulku
	// InitialMigration()
	// apak se pusti rest
	handleRequests()
}
