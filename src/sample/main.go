package sample

import "net/http"

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
