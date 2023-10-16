package tools

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/JinnnDamanee/review-service/config"
	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func DeleteAllRecords() error {
	db, err := createDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM reviews")
	if err != nil {
		return err
	}

	return nil
}

func CreateMockCompany(name, email, password, description, location, phone, category string) (int64, error) {
	config, err := config.LoadConfig("../")
	if err != nil {
		return 0, err
	}

	target := fmt.Sprintf("%s:%s", config.UserServiceHost, config.UserServicePort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	client := pbv1.NewAuthServiceClient(conn)
	res, err := client.CreateCompany(context.Background(), &pbv1.CreateCompanyRequest{
		Name:            name,
		Email:           email,
		Password:        password,
		PasswordConfirm: password,
		Description:     description,
		Location:        location,
		Phone:           phone,
		Category:        category,
	})
	if err != nil {
		return 0, err
	}
	if res.Status != 201 {
		return 0, fmt.Errorf("status code: %d", res.Status)
	}

	return res.Id, nil
}

func CreateMockStudent(name, email, password, description, faculty, major string, year int32) (int64, error) {
	config, err := config.LoadConfig("../")
	if err != nil {
		return 0, err
	}

	target := fmt.Sprintf("%s:%s", config.UserServiceHost, config.UserServicePort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	client := pbv1.NewAuthServiceClient(conn)
	res, err := client.CreateStudent(context.Background(), &pbv1.CreateStudentRequest{
		Name:            name,
		Email:           email,
		Password:        password,
		PasswordConfirm: password,
		Description:     description,
		Faculty:         faculty,
		Major:           major,
		Year:            year,
	})
	if err != nil {
		return 0, err
	}
	if res.Status != 201 {
		return 0, fmt.Errorf("status code: %d", res.Status)
	}

	return res.Id, nil
}
