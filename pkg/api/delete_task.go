package api

import (
	"net/http"

	"go_final_project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		writeJSON(w, ErrorResponse{
			Error: "id is empty",
		})
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	writeJSON(w, map[string]string{})
}
