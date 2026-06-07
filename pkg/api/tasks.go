package api

import (
	"net/http"

	"go_final_project/pkg/consts"
	"go_final_project/pkg/db"
)

type TasksResponse struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(consts.TaskLimit)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		}, http.StatusNotFound)
		return
	}

	writeJSON(w, TasksResponse{
		Tasks: tasks,
	}, http.StatusOK)
}
