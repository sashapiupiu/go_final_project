package main

import (
	"log"
	"net/http"

	"go_final_project/pkg/api"
	"go_final_project/pkg/db"
)

func main() {

	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatal(err)
	}
	api.Init()

	fs := http.FileServer(http.Dir("./web"))

	http.Handle("/", fs)

	log.Println("start server :7540")

	err = http.ListenAndServe(":7540", nil)
	if err != nil {
		log.Fatal(err)
	}
}
