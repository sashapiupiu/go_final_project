package main

import (
	"go_final_project/pkg/api"
	"go_final_project/pkg/db"
	"log"
	"net/http"
)

func main() {
	db.Init("scheduler.db")
	api.Init()
	fs := http.FileServer(http.Dir("./web"))

	http.Handle("/", fs)

	log.Println("start server :7540")

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		log.Fatal(err)
	}
}
