package service

import (
	"context"

	pbUser "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-post-service/port"
)

type userClientMockService struct{}

func NewUserClientMockService() port.UserClientPort {
	return &userClientMockService{}
}

func (u *userClientMockService) GetCompanyProfile(ctx context.Context, req *pbUser.GetCompanyRequest) (*pbUser.GetCompanyResponse, error) {
	return &pbUser.GetCompanyResponse{
		Company: &pbUser.Company{
			Id:          0,
			Name:        "Mock Company Name",
			Email:       "mock@company.com",
			Description: "Mock Company Description",
			Location:    "Mock Company Location",
			Phone:       "0851234455",
			Category:    "Mock Company Category",
		},
		Status:  200,
		Message: "Success",
	}, nil
}
