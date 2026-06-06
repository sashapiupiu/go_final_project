package api

import (
	"net/http"
	"time"

	"go_final_project/pkg/consts"
	"go_final_project/pkg/db"
	"go_final_project/pkg/repeat"
)

type TasksResponse struct {
	Tasks []*db.Task `json:"tasks"`
}

func filteringTasks(inp []*db.Task) (out []*db.Task) {

	out = make([]*db.Task, 0)
	now := time.Now()

	for _, t := range inp {
		taskDate, err := time.Parse(consts.DateLayout, t.Date)
		if err != nil {
			continue
		}

		if taskDate.After(now) {
			out = append(out, t)
			continue
		}

		if t.Repeat != "" {
			next, err := repeat.NextDate(now, t.Date, t.Repeat)
			if err == nil {
				t.Date = next
				out = append(out, t)
			}
		}
	}
	return
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(consts.TaskLimit)
	if err != nil {
		writeJSON(w, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	tasks = filteringTasks(tasks)

	writeJSON(w, TasksResponse{
		Tasks: tasks,
	})
}
