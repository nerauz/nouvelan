package main

import (
	"encoding/json"
	"net/http"
	"victor/new_year/score"
	"victor/new_year/server"
)

func main() {
	s := server.Server{}

	s.SetStaticFiles("/", "./static")
	s.CreateRoute("/score", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		score, err := score.GetScore()

		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(*score)
	})
	s.Start()
}
