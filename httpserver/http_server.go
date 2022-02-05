package httpserver

import (
	"fmt"
	"net/http"
)

type ServerSettings struct {
	Endpoints []string
}
type Server struct {
	Settings *ServerSettings
}

func (server *Server) InitServer() {
	server.Settings = &ServerSettings{Endpoints: []string{"2", "3", "4", "json-data"}}
}

func (server *Server) RunEndpoints() {
	for _, endpoint := range server.Settings.Endpoints {
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

		for _, endpoint := range server.Settings.Endpoints {
			w.Write([]byte("\n"))
			w.Write([]byte(fmt.Sprintf("<a href='/%s'>Strana %s</a>", endpoint, endpoint)))
		}
	})
	http.ListenAndServe(":80", nil)
}
