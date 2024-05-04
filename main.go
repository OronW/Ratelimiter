package main

import "net/http"

func main() {
	http.HandleFunc("/test-server", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server Works"))
	})
	http.ListenAndServe(":8080", nil)
}
