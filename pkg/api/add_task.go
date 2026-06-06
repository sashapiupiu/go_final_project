package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"go_final_project/pkg/consts"
	"go_final_project/pkg/db"
	"go_final_project/pkg/repeat"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type IdResponse struct {
	ID string `json:"id"`
}

func writeJSON(w http.ResponseWriter, data any, responseCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
w.WriteHeader(responseCode)
	err := 	json.NewEncoder(w).Encode(data)
}

func afterNow(a, b time.Time) bool {
	a = time.Date(a.Year(), a.Month(), a.Day(), 0, 0, 0, 0, a.Location())
	b = time.Date(b.Year(), b.Month(), b.Day(), 0, 0, 0, 0, b.Location())

	return a.After(b)
}

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(consts.DateLayout)
	}

	t, err := time.Parse(consts.DateLayout, task.Date)
	if err != nil {
		return err
	}

	var next string

	if task.Repeat != "" {
		next, err = repeat.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if afterNow(now, t) {

		if task.Repeat == "" {
			task.Date = now.Format(consts.DateLayout)
		} else {
			task.Date = next
		}
	}

	return nil
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
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

	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		}, http.StatusConflict)
		return
	}

	writeJSON(w, IdResponse{
		ID: strconv.FormatInt(id, 10),
	}, http.StatusCreated)
}
