package mock

import (
	"context"
	"fmt"

	pbUser "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
	"github.com/TikhampornSky/go-post-service/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateMockApprovedCompany(ctx context.Context, name, accessToken string) (*pbUser.CreateCompanyResponse, string, error) {
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

	client := pbUser.NewAuthServiceClient(conn)
	clientUser := pbUser.NewUserServiceClient(conn)

	companyEmail := utils.GenerateRandomString(10) + "@company.com"
	// Create req
	req := &pbUser.CreateCompanyRequest{
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
	reqApprove := &pbUser.UpdateCompanyStatusRequest{
		AccessToken: accessToken,
		Id:          res.Id,
		Status:      "Approve",
	}

	_, err = clientUser.UpdateCompanyStatus(ctx, reqApprove)
	if err != nil {
		return nil, "", err
	}

	// SignIn Company
	reqSignIn := &pbUser.LoginRequest{
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

	client := pbUser.NewAuthServiceClient(conn)

	// Create Admin
	adminEmail := utils.GenerateRandomString(10) + "@admin.com"
	reqAdmin := &pbUser.CreateAdminRequest{
		Email:           adminEmail,
		Password:        "password-test",
		PasswordConfirm: "password-test",
	}

	_, err = client.CreateAdmin(ctx, reqAdmin)
	if err != nil {
		return "", err
	}

	// SignIn Admin
	reqSignIn := &pbUser.LoginRequest{
		Email:    adminEmail,
		Password: "password-test",
	}

	resSignIn, err := client.SignIn(ctx, reqSignIn)
	if err != nil {
		return "", err
	}

	return resSignIn.AccessToken, nil
}
