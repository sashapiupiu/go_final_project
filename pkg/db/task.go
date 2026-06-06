package db

import (
	"fmt"
	"time"

	"go_final_project/pkg/consts"
	"go_final_project/pkg/repeat"
)

type Task struct {
	ID      string `json:"id,omitempty"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func Tasks(limit int) ([]*Task, error) {
	rows, err := DB.Query(`
		SELECT id, date, title, comment, repeat
		FROM scheduler
		ORDER BY date
		LIMIT ?
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	now := time.Now()
	result := make([]*Task, 0)

	for rows.Next() {
		var t Task

		err := rows.Scan(
			&t.ID,
			&t.Date,
			&t.Title,
			&t.Comment,
			&t.Repeat,
		)
		if err != nil {
			return nil, err
		}

		taskDate, err := time.Parse(consts.DateLayout, t.Date)
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
func GetTask(id string) (*Task, error) {
	query := `
		SELECT id, date, title, comment,repeat
		FROM scheduler
		WHERE id = ?
	`
	var task Task
	err := DB.QueryRow(query, id).Scan(
		&task.ID,
		&task.Date,
		&task.Title,
		&task.Comment,
		&task.Repeat,
	)
	if err != nil {
		return nil, err
	}
	return &task, nil

}

func UpdateTask(task *Task) error {
	query := `
		UPDATE scheduler
		SET date = ?, title = ?, comment = ?, repeat = ?
		WHERE id = ?
	`
	res, err := DB.Exec(
		query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
		task.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}
func DeleteTask(id string) error {
	res, err := DB.Exec(
		"DELETE FROM scheduler WHERE id = ?",
		id,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}
func UpdateDate(next, id string) error {
	res, err := DB.Exec(
		"UPDATE scheduler SET date = ? WHERE id = ?",
		next,
		id,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}
func AddTask(task *Task) (int64, error) {
	query := `
		INSERT INTO scheduler
		(date, title, comment, repeat)
		VALUES (?, ?, ?, ?)
	`

	res, err := DB.Exec(
		query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
	)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
