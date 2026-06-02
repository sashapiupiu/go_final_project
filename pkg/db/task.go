package db

import (
	"go_final_project/pkg/repeat"
	"time"
)

type Task struct {
	ID      string `json:"id,omitempty"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// query := `
// SELECT id, date, title, comment, repeat
// FROM scheduler
// WHERE date >= ?  или тудэй?
// ORDER BY date
// LIMIT ?
// `
func Tasks(limit int) ([]*Task, error) {
	rows, err := DB.Query(`
		SELECT id, date, title, comment, repeat
		FROM scheduler
		ORDER BY date
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	now := time.Now()
	result := make([]*Task, 0)

	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)

		taskDate, err := time.Parse("20060102", t.Date)
		if err != nil {
			continue
		}

		if taskDate.After(now) {
			result = append(result, &t)
			continue
		}

		if t.Repeat != "" {
			next, err := repeat.NextDate(now, t.Date, t.Repeat)
			if err == nil {
				t.Date = next
				result = append(result, &t)
			}
		}
	}

	return result, nil
}
