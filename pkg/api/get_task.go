package api

import (
	"net/http"

	"go_final_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		writeJSON(w, ErrorResponse{
			Error: "id is empty",
		}, http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		}, http.StatusNotFound)
		return
	}

	writeJSON(w, task, http.StatusOK)
}
