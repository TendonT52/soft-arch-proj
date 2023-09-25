package service

import (
	"context"
	"fmt"

	pbUser "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-post-service/config"
	"github.com/TikhampornSky/go-post-service/port"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userClientService struct{}

func NewUserClientService() port.UserClientPort {
	return &userClientService{}
}

func (u *userClientService) GetCompanyProfile(ctx context.Context, req *pbUser.GetCompanyRequest) (*pbUser.GetCompanyResponse, error) {
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

	client := pbUser.NewUserServiceClient(conn)
	res, err := client.GetCompany(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userClientService) ListApprovedCompanies(ctx context.Context, req *pbUser.ListApprovedCompaniesRequest) (*pbUser.ListApprovedCompaniesResponse, error) {
	config, err := config.LoadConfig("../")
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(fmt.Sprintf(":%s", config.UserServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbUser.NewUserServiceClient(conn)
	res, err := client.ListApprovedCompanies(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
