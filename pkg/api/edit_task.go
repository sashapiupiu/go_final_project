package api

import (
	"encoding/json"
	"net/http"

	"go_final_project/pkg/db"
)

var task db.Task

func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		writeJSON(w, ErrorResponse{
			Error: "ID is required",
		}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJSON(w, ErrorResponse{
			Error: "title is required",
		}, http.StatusBadRequest)
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		}, http.StatusConflict)
		return
	}

	writeJSON(w, map[string]string{}, http.StatusOK)
}
