package api

import (
	"encoding/json"
	"net/http"
)

//const DateLayout = "20060102"

type Task struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}
type Response struct {
	ID    int64  `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, Response{
			Error: "invalid json",
		})
		return
	}
	if task.Title == "" {
		writeJSON(w, Response{
			Error: "title is required",
		})
		return
	}

	if task.Date != now.Format(DateLayout) {
		writeJSON(w, Response{
			Error: "invalid date",
		})
		return
	}

	if task.Date == "" {

	}

}
