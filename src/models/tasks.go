package models

import "database/sql"

type Task struct {
	ID   int
	Body string
}

type TaskModel struct {
	DB *sql.DB
}

func (m *TaskModel) Insert(body string) error {
	stmt := `INSERT INTO tasks (body) VALUES(?)`

	_, err := m.DB.Exec(stmt, body)
	if err != nil {
		return nil
	}

	return nil
}
