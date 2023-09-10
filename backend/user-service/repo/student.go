package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
)

func (r *userRepository) CreateStudent(ctx context.Context, student *pbv1.CreateStudentRequest, createTime int64) (int64, error) {
	// Start a transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into Table users
	query := "INSERT INTO users (email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int64
	err = tx.QueryRowContext(ctx, query, student.Email, student.Password, "student", createTime, createTime).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into Table students
	query = "INSERT INTO students (sid, name, description, faculty, major, year, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = tx.ExecContext(ctx, query, id, student.Name, student.Description, student.Faculty, student.Major, student.Year, createTime, createTime)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Commit the transaction if all insertions were successful
	err = tx.Commit()
	if err != nil {
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	return id, nil
}

func (r *userRepository) GetSalt(ctx context.Context, email string) (string, error) {
	query := "SELECT created_at FROM users WHERE email = $1"
	var salt string
	err := r.db.QueryRowContext(ctx, query, email).Scan(&salt)
	if err != nil {
		return "", domain.ErrInternal.From(err.Error(), err)
	}
	return salt, nil
}

func (r *userRepository) GetStudentByID(ctx context.Context, id int64) (*pbv1.Student, error) {
	query := "SELECT users.id, students.name, users.email, students.description, students.faculty, students.major, students.year FROM users INNER JOIN students ON users.id = students.sid WHERE users.id = $1"
	var student pbv1.Student
	err := r.db.QueryRowContext(ctx, query, id).Scan(&student.Id, &student.Name, &student.Email, &student.Description, &student.Faculty, &student.Major, &student.Year)
	if errors.Is(err, sql.ErrNoRows){
		return &pbv1.Student{}, domain.ErrUserIDNotFound.From(err.Error(), err)
	}
	if err != nil {
		return &pbv1.Student{}, domain.ErrInternal.From(err.Error(), err)
	}
	return &student, nil
}

func (r *userRepository) UpdateStudentByID(ctx context.Context, id int64, req *pbv1.Student) error {
	current_timestamp := time.Now().Unix()
	query := "UPDATE students SET name = $1, description = $2, faculty = $3, major = $4, year = $5, updated_at = $6 WHERE sid = $7"
	_, err := r.db.ExecContext(ctx, query, req.Name, req.Description, req.Faculty, req.Major, req.Year, current_timestamp, id)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}

func (r *userRepository) DeleteStudent(ctx context.Context, id int64) error {
	// Start a transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Delete from Table students
	query := "DELETE FROM students WHERE sid = $1"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Delete from Table users
	query = "DELETE FROM users WHERE id = $1"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Commit the transaction if all insertions were successful
	err = tx.Commit()
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}

func (r *userRepository) UpdateStudentStatus(ctx context.Context, email string, verified bool) error {
	query := "UPDATE users SET verified = $1 WHERE email = $2"
	_, err := r.db.ExecContext(ctx, query, verified, email)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}
