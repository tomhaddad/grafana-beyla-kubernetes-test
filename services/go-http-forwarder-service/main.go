package main

import (
	"io"
	"net/http"
)

func forward(w http.ResponseWriter, req *http.Request) {
	client := http.Client{}

	requestToForward, err := http.NewRequest("POST", "http://go-http-receiver-service:8000/receive", req.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := client.Do(requestToForward)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(body)
}

func main() {
	http.HandleFunc("/forward", forward)
	http.ListenAndServe(":8000", nil)
}
