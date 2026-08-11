package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!, welcome to my website")
		fmt.Println("[LOG] : User mengunjungi halaman utama")
	})

	server := http.Server{
		Addr: ":8080", // containerization (cukup ambil portnya saja)
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}