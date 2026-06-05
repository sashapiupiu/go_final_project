package api

import (
	"log"
	"net/http"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("taskHandler:", r.Method)
	switch r.Method {
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodPut:
		editTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	default:
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
	}
}
