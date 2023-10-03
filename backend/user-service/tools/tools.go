package tools

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func createDBConnection() (*sql.DB, error) {
	config, err := config.LoadConfig("..")
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName, "disable")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func DeleteAll() error {
	db, err := createDBConnection()
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}
	ctx := context.Background()

	// Start a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Delete all data from Table students
	query := "DELETE FROM students"
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Delete all data from Table companies
	query = "DELETE FROM companies"
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Delete all data from Table users
	query = "DELETE FROM users"
	_, err = tx.ExecContext(ctx, query)
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

func GetCompanyByID(id int64) (*pbv1.Company, *domain.UserStatus, error) {
	db, err := createDBConnection()
	if err != nil {
		return nil, nil, domain.ErrInternal.From(err.Error(), err)
	}
	ctx := context.Background()

	query := "SELECT users.id, users.role, users.verified, companies.name, users.email, companies.description, companies.location, companies.phone, companies.category, companies.status FROM users INNER JOIN companies ON users.id = companies.cid WHERE users.id = $1"
	var company pbv1.Company
	var status domain.UserStatus
	err = db.QueryRowContext(ctx, query, id).Scan(&company.Id, &status.Role, &status.Verified, &company.Name, &company.Email, &company.Description, &company.Location, &company.Phone, &company.Category, &company.Status)
	if err != nil {
		return nil, nil, domain.ErrInternal.From(err.Error(), err)
	}

	return &company, &status, nil
}

func GetStudentByID(id int64) (*pbv1.Student, *domain.UserStatus, error) {
	db, err := createDBConnection()
	if err != nil {
		return nil, nil, domain.ErrInternal.From(err.Error(), err)
	}
	ctx := context.Background()

	query := "SELECT users.id, users.role, users.verified, students.name, users.email, students.description, students.faculty, students.major, students.year FROM users INNER JOIN students ON users.id = students.sid WHERE users.id = $1"
	var student pbv1.Student
	var status domain.UserStatus
	err = db.QueryRowContext(ctx, query, id).Scan(&student.Id, &status.Role, &status.Verified, &student.Name, &student.Email, &student.Description, &student.Faculty, &student.Major, &student.Year)
	if err != nil {
		return nil, nil, domain.ErrInternal.From(err.Error(), err)
	}

	return &student, &status, nil
}

func GetUserByID(id int64) (*pbv1.User, error) {
	db, err := createDBConnection()
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	ctx := context.Background()
	query := "SELECT verified, password, email, role, created_at FROM users WHERE id = $1;"
	u := &pbv1.User{}
	err = db.QueryRowContext(ctx, query, id).Scan(&u.Verified, &u.Password, &u.Email, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}

	return u, nil
}

func GetCreateTime(id int64) (int64, error) {
	db, err := createDBConnection()
	if err != nil {
		return 0, domain.ErrInternal.From(err.Error(), err)
	}
	ctx := context.Background()
	query := "SELECT created_at FROM users WHERE id = $1;"
	var createTime int64
	err = db.QueryRowContext(ctx, query, id).Scan(&createTime)
	if err != nil {
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	return createTime, nil
}

func GetValueFromRedis(refreshToken string) (string, error) {
	config, err := config.LoadConfig("..")
	redis := redis.NewClient(&redis.Options{
		Addr:     config.REDISHost + ":" + config.REDISPort,
		Password: config.REDISPassword,
		DB:       config.REDISDB,
	})
	ctx := context.Background()

	value, err := redis.Get(ctx, refreshToken).Result()
	if err != nil {
		return "", domain.ErrInternal.From(err.Error(), err)
	}
	return value, nil
}