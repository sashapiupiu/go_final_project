package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web"))

	http.Handle("/", fs)

	log.Println("start server :7540")

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		log.Fatal(err)
	}
}
