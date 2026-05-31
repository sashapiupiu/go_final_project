package api

import (
	"go_final_project/pkg/repeat"
	"net/http"
	"time"
)

const DateLayout = "20060102"

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeatRule := r.FormValue("repeat")

	now, err := time.Parse(DateLayout, nowStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	next, err := repeat.NextDate(now, date, repeatRule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(next))
}
