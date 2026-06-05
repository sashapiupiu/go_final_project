package api

import (
	"encoding/json"
	"go_final_project/pkg/db"
	"net/http"
)

var task db.Task

func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if task.ID == "" {
		writeJSON(w, ErrorResponse{
			Error: "ID is required",
		})
		return
	}

	if task.Title == "" {
		writeJSON(w, ErrorResponse{
			Error: "title is required",
		})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	writeJSON(w, map[string]string{})
}
