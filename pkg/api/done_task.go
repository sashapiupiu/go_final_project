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
		_, err := w.Write([]byte("405 Method Not Allowed"))
		if err != nil {
			log.Println("failed write HTTP reply:", "405 Method Not Allowed")
		}
		return
	}

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
