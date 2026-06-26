package store

import (
	"database/sql"
	"task-api/models"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {

	return &Store{
		db: db,
	}

}

func (s *Store) Create(task models.Task) error {
	query := `INSERT INTO tasks(title, description, status, created_at)
		VALUES (?, ?, ?, ?)`

	_, err := s.db.Exec(
		query,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
	)
	return err
}

func (s *Store) GetAll() ([]models.Task, error) {
	rows, err := s.db.Query("select id, title, description, status, created_at from tasks")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
		)

		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *Store) GetByID(id int) (*models.Task, error) {

	row := s.db.QueryRow("select id, title, description, status, created_at from tasks where id=?", id)
	var task models.Task

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *Store) DeleteByID(id int) error {
	query := `
	DELETE FROM tasks
	WHERE id = ?
	`

	result, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (s *Store) UpdateByID(id int, task models.Task) error {
	query := `UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?`

	result, err := s.db.Exec(query, task.Title, task.Description, task.Status, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
