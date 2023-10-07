package mock

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/TikhampornSky/go-post-service/config"
	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateMockApprovedCompany(ctx context.Context, name, accessToken string) (*pbv1.CreateCompanyResponse, string, error) {
	config, err := config.LoadConfig("../")
	if err != nil {
		return nil, "", err
	}

	target := fmt.Sprintf("%s:%s", config.UserServiceHost, config.UserServicePort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, "", err
	}
	defer conn.Close()

	client := pbv1.NewAuthServiceClient(conn)
	clientUser := pbv1.NewUserServiceClient(conn)

	companyEmail := utils.GenerateRandomString(10) + "@company.com"
	// Create req
	req := &pbv1.CreateCompanyRequest{
		Name:            name,
		Email:           companyEmail,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		Description:     "I am a company",
		Location:        "Bangkok",
		Phone:           "0123456789",
		Category:        "IT",
	}

	res, err := client.CreateCompany(ctx, req)
	if err != nil {
		return nil, "", err
	}
	if res.Status != 201 {
		return nil, "", errors.New(res.Message)
	}

	// Approve Company
	reqApprove := &pbv1.UpdateCompanyStatusRequest{
		AccessToken: accessToken,
		Id:          res.Id,
		Status:      "Approve",
	}

	resCom, err := clientUser.UpdateCompanyStatus(ctx, reqApprove)
	if err != nil {
		return nil, "", err
	}
	if resCom.Status != 200 {
		return nil, "", errors.New(resCom.Message)
	}

	// SignIn Company
	reqSignIn := &pbv1.LoginRequest{
		Email:    companyEmail,
		Password: "password-test",
	}
	resSignIn, err := client.SignIn(ctx, reqSignIn)
	if err != nil {
		return nil, "", err
	}
	if resSignIn.Status != 200 {
		return nil, "", errors.New(resSignIn.Message)
	}

	return res, resSignIn.AccessToken, nil
}

func CreateMockAdmin(ctx context.Context) (string, error) {
	config, err := config.LoadConfig("../")
	if err != nil {
		return "", err
	}

	target := fmt.Sprintf("%s:%s", config.UserServiceHost, config.UserServicePort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pbv1.NewAuthServiceClient(conn)

	// Create Admin
	adminEmail := utils.GenerateRandomString(10) + "@admin.com"
	admin_access_token, err := GenerateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   "admin",
	})
	reqAdmin := &pbv1.CreateAdminRequest{
		Email:           adminEmail,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		AccessToken:     admin_access_token,
	}

	resAdmin, err := client.CreateAdmin(ctx, reqAdmin)
	if err != nil {
		return "", err
	}
	if resAdmin.Status != 201 {
		return "", errors.New(resAdmin.Message)
	}

	// SignIn Admin
	reqSignIn := &pbv1.LoginRequest{
		Email:    adminEmail,
		Password: "password-test",
	}

	resSignIn, err := client.SignIn(ctx, reqSignIn)
	if err != nil {
		return "", err
	}
	if resSignIn.Status != 200 {
		return "", errors.New(resSignIn.Message)
	}

	return resSignIn.AccessToken, nil
}
