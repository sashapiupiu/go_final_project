package api

import (
	"go_final_project/pkg/db"
	"go_final_project/pkg/repeat"
	"log"
	"net/http"
	"time"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		if err != nil {
			log.Println("failed write HTTP reply:", http.StatusText(http.StatusMethodNotAllowed))
		}
		return
	}

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

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeJSON(w, ErrorResponse{
				Error: err.Error(),
			}, http.StatusConflict)
			return
		}
	} else {
		next, err := repeat.NextDate(
			time.Now(),
			task.Date,
			task.Repeat,
		)
		if err != nil {
			writeJSON(w, ErrorResponse{
				Error: err.Error(),
			}, http.StatusInternalServerError)
			return
		}

		err = db.UpdateDate(next, id)
		if err != nil {
			writeJSON(w, ErrorResponse{
				Error: err.Error(),
			}, http.StatusConflict)
			return
		}
	}

	writeJSON(w, map[string]string{}, http.StatusOK)
}
