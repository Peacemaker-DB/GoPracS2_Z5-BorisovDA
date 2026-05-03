package student

import (
	"context"
	"database/sql"
	"errors"
)

var ErrStudentNotFound = errors.New("student not found")

type Repo struct {
	db             *sql.DB
	getByIDStmt    *sql.Stmt
	getByEmailStmt *sql.Stmt
}

func NewRepo(db *sql.DB) (*Repo, error) {
	getByIDStmt, err := db.Prepare("SELECT id, full_name, study_group, email FROM students WHERE id = $1")
	if err != nil {
		return nil, err
	}

	getByEmailStmt, err := db.Prepare("SELECT id, full_name, study_group, email FROM students WHERE email = $1")
	if err != nil {
		_ = getByIDStmt.Close()
		return nil, err
	}

	return &Repo{
		db:             db,
		getByIDStmt:    getByIDStmt,
		getByEmailStmt: getByEmailStmt,
	}, nil
}

func (r *Repo) GetByID(ctx context.Context, id int64) (Student, error) {
	row := r.getByIDStmt.QueryRowContext(ctx, id)
	return scanStudent(row)
}

func (r *Repo) GetByEmail(ctx context.Context, email string) (Student, error) {
	row := r.getByEmailStmt.QueryRowContext(ctx, email)
	return scanStudent(row)
}

func (r *Repo) Close() error {
	var result error

	if r.getByIDStmt != nil {
		if err := r.getByIDStmt.Close(); err != nil {
			result = err
		}
	}

	if r.getByEmailStmt != nil {
		if err := r.getByEmailStmt.Close(); err != nil && result == nil {
			result = err
		}
	}

	return result
}

func UnsafeBuildGetByIDQuery(rawID string) string {
	return "SELECT id, full_name, study_group, email FROM students WHERE id = " + rawID
}

func scanStudent(row *sql.Row) (Student, error) {
	var st Student

	err := row.Scan(&st.ID, &st.FullName, &st.StudyGroup, &st.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Student{}, ErrStudentNotFound
		}

		return Student{}, err
	}

	return st, nil
}