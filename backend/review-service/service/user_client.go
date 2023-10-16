package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JinnnDamanee/review-service/config"
	"github.com/JinnnDamanee/review-service/port"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
)

type userClientService struct{}

func NewUserClientService() port.UserClientPort {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatalln("Could not load environment variables", err)
	}

	target := fmt.Sprintf("%s:%s", config.UserServiceHost, config.UserServicePort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Could not connect to the user service", err)
	}
	defer conn.Close()

	clientAuth := pbv1.NewAuthServiceClient(conn)
	clientUser := pbv1.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	healthAuth, err := clientAuth.AuthHealthCheck(ctx, &pbv1.AuthHealthCheckRequest{})
	if err != nil || healthAuth.Status != 200 {
		log.Fatalln("Could not connect to the user service (AUTH)", err)
	}

	healthUser, err := clientUser.UserHealthCheck(ctx, &pbv1.UserHealthCheckRequest{})
	if err != nil || healthUser.Status != 200 {
		log.Fatalln("Could not connect to the user service (USER)", err)
	}
	log.Println("Connected to the user service successfully")

	return &userClientService{}
}

func (s *userClientService) GetUserProfile(ctx context.Context, req *pbv1.GetStudentRequest) (*pbv1.GetStudentResponse, error) {
	config, err := config.LoadConfig("../")
	if err != nil {
		return nil, err
	}

	target := fmt.Sprintf("%s:%s", config.UserServiceHost, config.UserServicePort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbv1.NewUserServiceClient(conn)
	res, err := client.GetStudent(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userClientService) GetCompanyProfile(ctx context.Context, req *pbv1.GetCompanyRequest) (*pbv1.GetCompanyResponse, error) {
	config, err := config.LoadConfig("../")
	if err != nil {
		return nil, err
	}

	target := fmt.Sprintf("%s:%s", config.UserServiceHost, config.UserServicePort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbv1.NewUserServiceClient(conn)
	res, err := client.GetCompany(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}