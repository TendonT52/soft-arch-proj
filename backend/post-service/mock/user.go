package mock

import (
	"context"
	"fmt"

	"github.com/TikhampornSky/go-post-service/config"
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

	// Approve Company
	reqApprove := &pbv1.UpdateCompanyStatusRequest{
		AccessToken: accessToken,
		Id:          res.Id,
		Status:      "Approve",
	}

	_, err = clientUser.UpdateCompanyStatus(ctx, reqApprove)
	if err != nil {
		return nil, "", err
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
	reqAdmin := &pbv1.CreateAdminRequest{
		Email:           adminEmail,
		Password:        "password-test",
		PasswordConfirm: "password-test",
	}

	_, err = client.CreateAdmin(ctx, reqAdmin)
	if err != nil {
		return "", err
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

	return resSignIn.AccessToken, nil
}
