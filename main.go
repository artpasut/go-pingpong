package main

import "net/http"

func main() {
	http.HandleFunc("/ht", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{'ok'}`))
	})

	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{'ping'}`))
	})

	http.ListenAndServe(":9000", nil)
}
