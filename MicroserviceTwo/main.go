package main

import (
	"log"
	"net/http"
	"os"
	s "weather/service"
)

func teaHandler(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(s.Call() + host))
}

func main() {
	http.HandleFunc("/", teaHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
