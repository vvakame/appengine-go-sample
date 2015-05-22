package sample

import "net/http"

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	if param == "bye" {
		w.Write([]byte("Good bye..."))
	} else {
		w.Write([]byte("Hello!"))
	}
}
