package api

import (
	"log"
	"net/http"
	"time"

	"go_final_project/pkg/consts"
	"go_final_project/pkg/repeat"
)

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeatRule := r.FormValue("repeat")

	now, err := time.Parse(consts.DateLayout, nowStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	next, err := repeat.NextDate(now, date, repeatRule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write([]byte(next))
	if err != nil {
		log.Println("failed write HTTP reply:", next)
	}
}
