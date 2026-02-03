package main

import "net/http"

func main() {
	http.HandleFunc("/", SeachCEP)
	http.ListenAndServe(":8080", nil)
}

func SeachCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
