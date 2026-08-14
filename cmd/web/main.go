package main

import (
	"fmt"
	"log"
	"go-financial/internal/config"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	// mysql connection
	db, err := config.OpenConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!, welcome to my website")
		fmt.Println("[LOG] : User mengunjungi halaman utama")
	})

	server := http.Server{
		Addr: ":8080", // containerization (cukup ambil portnya saja)
		Handler: mux,
	}

	log.Printf("Server berjalan di port : %s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}