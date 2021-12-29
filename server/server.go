package server

import (
	"fmt"
	"log"
	"net/http"
	"victor/new_year/tools"
)

type Server struct {
	ip   string
	port string
}

func (serv Server) Start() {
	port := tools.Ternary(serv.port == "", "8080", serv.port)
	url := fmt.Sprintf("%s:%s", serv.ip, port)

	log.Printf("Serveur running on: %s", url)

	err := http.ListenAndServe(url, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func (serv Server) CreateRoute(endpoint string, handler func(w http.ResponseWriter, req *http.Request)) {
	http.HandleFunc(endpoint, handler)
}

func (serv Server) SetStaticFiles(endpoint string, dir string) {
	fs := http.FileServer(http.Dir(dir))
	http.Handle(endpoint, fs)
}
