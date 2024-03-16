package main

import (
	"io"
	"net/http"
)

func receive(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(body)
}

func main() {
	http.HandleFunc("/receive", receive)
	http.ListenAndServe(":8000", nil)
}
