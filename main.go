package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Main!")

	endpoints := []string{"2", "3"}

	for _, endpoint := range endpoints {
		http.HandleFunc(("/" + endpoint), func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte(fmt.Sprintf("%v stranka\n", endpoint)))
			w.Write([]byte("<a href='home'>Domů</a>"))
		})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("Pozdrav z main page!\n"))
		w.Write([]byte("dalsi stranky na tomto webu na :\n"))

		for _, endpoint := range endpoints {
			w.Write([]byte("\n"))
			w.Write([]byte(fmt.Sprintf("<a href='/%s'>Strana %s</a>", endpoint, endpoint)))
		}
	})
	http.ListenAndServe(":80", nil)
}
