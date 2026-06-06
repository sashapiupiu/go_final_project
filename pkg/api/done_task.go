package api

import (
	"go_final_project/pkg/db"
	"go_final_project/pkg/repeat"
	"net/http"
	"time"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, ErrorResponse{
			Error: "id is empty",
		})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
	} else {
		next, err := repeat.NextDate(
			time.Now(),
			task.Date,
			task.Repeat,
		)
		if err != nil {
			writeJSON(w, ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		err = db.UpdateDate(next, id)
	}

	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	writeJSON(w, map[string]string{})
}
