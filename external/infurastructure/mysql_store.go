package infurastructure

import (
	"context"
	"database/sql"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/model/task"
	ifmodel "github.com/ArtefactGitHub/Go_T_Clean/external/infurastructure/model"
)

type mySqlTaskRepository struct {
	db *sql.DB
}

func NewMySqlTaskRepository(setting ifmodel.MySqlSetting) (task.TaskRepository, error) {
	db, err := sql.Open(setting.DriverName(), setting.DataSourceName())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	r := mySqlTaskRepository{
		db: db,
	}

	return &r, nil
}

func (r *mySqlTaskRepository) Finalize() {
	if r.db != nil {
		r.db.Close()
	}
}

func (r *mySqlTaskRepository) GetAll(ctx context.Context) ([]task.Task, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var m task.Task
	var result []task.Task
	for rows.Next() {
		err := rows.Scan(
			&m.Id,
			&m.Name,
			&m.CreatedAt,
			&m.UpdatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mySqlTaskRepository) Get(ctx context.Context, id int) (*task.Task, error) {
	m := task.Task{}
	err := r.db.QueryRowContext(ctx, "SELECT * FROM tasks WHERE id = ?", id).Scan(
		&m.Id,
		&m.Name,
		&m.CreatedAt,
		&m.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *mySqlTaskRepository) Create(ctx context.Context, m task.Task) (int, error) {
	result, err := r.db.ExecContext(ctx, `
		INSERT INTO tasks(id, name, createdAt, updatedAt) values(?, ?, ?, ?)`,
		nil,
		m.Name,
		m.CreatedAt,
		m.UpdatedAt)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	m.Id = int(id)
	return m.Id, nil
}

func (r *mySqlTaskRepository) Update(ctx context.Context, m task.Task) (*task.Task, error) {
	result, err := r.db.ExecContext(ctx, `
		UPDATE tasks
		SET name = ?, updatedAt = ?
		WHERE id = ?`,
		m.Name,
		m.UpdatedAt,
		m.Id)
	if err != nil {
		return &m, err
	}

	if num, err := result.RowsAffected(); err != nil {
		return nil, err
	} else if num == 0 {
		return nil, nil
	}
	return &m, nil
}

func (r *mySqlTaskRepository) Delete(ctx context.Context, id int) (bool, error) {
	result, err := r.db.ExecContext(ctx, `
		DELETE FROM tasks
		WHERE id = ?`,
		id)
	if err != nil {
		return false, err
	}

	if num, err := result.RowsAffected(); err != nil {
		return false, err
	} else if num == 0 {
		return false, nil
	}
	return true, nil
}
