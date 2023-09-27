package e2e

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/tools"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetCompanyMe(t *testing.T) {
	config, _ := config.LoadConfig("..")
	target := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)
	u := pbv1.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Craete Admin
	aa := utils.GenerateRandomString(10) + "@admin.com"
	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	admin := &pbv1.CreateAdminRequest{
		Email:           aa,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		AccessToken:     admin_access_token,
	}
	a, err := c.CreateAdmin(ctx, admin)
	require.Equal(t, int64(201), a.Status)
	require.NoError(t, err)

	// Admin Sign In
	admin_res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    aa,
		Password: admin.Password,
	})
	require.Equal(t, int64(200), admin_res.Status)
	require.NoError(t, err)

	// Register
	com := &pbv1.CreateCompanyRequest{
		Name:            "Mock Company " + utils.GenerateRandomString(3),
		Email:           utils.GenerateRandomString(10) + "@company.com",
		Password:        "password-test",
		PasswordConfirm: "password-test",
		Description:     "I am a company",
		Location:        "Bangkok",
		Phone:           "0123456789",
		Category:        "IT",
	}
	r, err := c.CreateCompany(ctx, com)
	require.Equal(t, int64(201), r.Status)
	require.NoError(t, err)

	// Approve Company
	result, err := u.UpdateCompanyStatus(ctx, &pbv1.UpdateCompanyStatusRequest{
		AccessToken: admin_res.AccessToken,
		Id:          r.Id,
		Status:      "Approve",
	})
	require.Equal(t, int64(200), result.Status)

	// Sign In
	res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    com.Email,
		Password: com.Password,
	})
	require.Equal(t, int64(200), res.Status)
	require.NoError(t, err)

	// Generate WRONG token
	access_token_wrong, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &domain.Payload{
		UserId: 0,
		Role:   domain.CompanyRole,
	})
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.GetCompanyMeRequest
		expect *pbv1.GetCompanyResponse
	}{
		"success": {
			req: &pbv1.GetCompanyMeRequest{
				AccessToken: res.AccessToken,
			},
			expect: &pbv1.GetCompanyResponse{
				Status: int64(200),
				Company: &pbv1.Company{
					Id:          r.Id,
					Name:        com.Name,
					Email:       com.Email,
					Description: com.Description,
					Location:    com.Location,
					Phone:       com.Phone,
					Category:    com.Category,
					Status:      "Approve",
				},
			},
		},
		"fail: invalid token": {
			req: &pbv1.GetCompanyMeRequest{
				AccessToken: access_token_wrong,
			},
			expect: &pbv1.GetCompanyResponse{
				Status: 500,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := u.GetCompanyMe(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Company, res.Company)
		})
	}
}

func TestGetComapany(t *testing.T) {
	config, _ := config.LoadConfig("..")
	target := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)
	u := pbv1.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Craete Admin
	aa := utils.GenerateRandomString(10) + "@admin.com"
	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	admin := &pbv1.CreateAdminRequest{
		Email:           aa,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		AccessToken:     admin_access_token,
	}
	a, err := c.CreateAdmin(ctx, admin)
	require.Equal(t, int64(201), a.Status)
	require.NoError(t, err)

	// Admin Sign In
	admin_res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    aa,
		Password: admin.Password,
	})
	require.Equal(t, int64(200), admin_res.Status)
	require.NoError(t, err)

	// Register
	com := &pbv1.CreateCompanyRequest{
		Name:            "Mock Company " + utils.GenerateRandomString(3),
		Email:           utils.GenerateRandomString(10) + "@company.com",
		Password:        "password-test",
		PasswordConfirm: "password-test",
		Description:     "I am a company",
		Location:        "Bangkok",
		Phone:           "0123456789",
		Category:        "IT",
	}
	r, err := c.CreateCompany(ctx, com)
	require.Equal(t, int64(201), r.Status)
	require.NoError(t, err)

	// Approve Company
	result, err := u.UpdateCompanyStatus(ctx, &pbv1.UpdateCompanyStatusRequest{
		AccessToken: admin_res.AccessToken,
		Id:          r.Id,
		Status:      "Approve",
	})
	require.Equal(t, int64(200), result.Status)

	// Sign In
	res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    com.Email,
		Password: com.Password,
	})
	require.Equal(t, int64(200), res.Status)
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.GetCompanyRequest
		expect *pbv1.GetCompanyResponse
	}{
		"success": {
			req: &pbv1.GetCompanyRequest{
				Id:          r.Id,
				AccessToken: res.AccessToken,
			},
			expect: &pbv1.GetCompanyResponse{
				Status:  int64(200),
				Message: "success",
				Company: &pbv1.Company{
					Id:          r.Id,
					Name:        com.Name,
					Email:       com.Email,
					Description: com.Description,
					Location:    com.Location,
					Phone:       com.Phone,
					Category:    com.Category,
					Status:      "Approve",
				},
			},
		},
		"fail: company id not found": {
			req: &pbv1.GetCompanyRequest{
				Id:          20000000000000,
				AccessToken: res.AccessToken,
			},
			expect: &pbv1.GetCompanyResponse{
				Status:  404,
				Message: "company id not found",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := u.GetCompany(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			require.Equal(t, tc.expect.Company, res.Company)
		})
	}
}

func TestUpdateCompany(t *testing.T) {
	config, _ := config.LoadConfig("..")
	target := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)
	u := pbv1.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Craete Admin
	aa := utils.GenerateRandomString(10) + "@admin.com"
	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	admin := &pbv1.CreateAdminRequest{
		Email:           aa,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		AccessToken:     admin_access_token,
	}
	a, err := c.CreateAdmin(ctx, admin)
	require.Equal(t, int64(201), a.Status)
	require.NoError(t, err)

	// Admin Sign In
	admin_res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    aa,
		Password: "password-test",
	})
	require.Equal(t, int64(200), admin_res.Status)
	require.NoError(t, err)

	// Register
	com := &pbv1.CreateCompanyRequest{
		Name:            "Mock Company " + utils.GenerateRandomString(3),
		Email:           utils.GenerateRandomString(10) + "@company.com",
		Password:        "password-test",
		PasswordConfirm: "password-test",
		Description:     "I am a company",
		Location:        "Bangkok",
		Phone:           "0123456789",
		Category:        "IT",
	}
	r, err := c.CreateCompany(ctx, com)
	require.Equal(t, int64(201), r.Status)
	require.NoError(t, err)

	// Approve Company
	result, err := u.UpdateCompanyStatus(ctx, &pbv1.UpdateCompanyStatusRequest{
		AccessToken: admin_res.AccessToken,
		Id:          r.Id,
		Status:      "Approve",
	})
	require.Equal(t, int64(200), result.Status)

	// Sign In
	res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    com.Email,
		Password: com.Password,
	})
	require.Equal(t, int64(200), res.Status)
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.UpdateCompanyRequest
		expect *pbv1.UpdateCompanyResponse
	}{
		"success": {
			req: &pbv1.UpdateCompanyRequest{
				AccessToken: res.AccessToken,
				Company: &pbv1.Company{
					Name:        "Mock Company New Name",
					Description: "I am a company New",
					Location:    "Bangkok New",
					Phone:       "0123456780",
					Category:    "IT New",
				},
			},
			expect: &pbv1.UpdateCompanyResponse{
				Status:  int64(200),
				Message: "Update data for Mock Company New Name successfully!",
			},
		},
		"not authorize": {
			req: &pbv1.UpdateCompanyRequest{
				AccessToken: admin_res.AccessToken,
				Company:     &pbv1.Company{},
			},
			expect: &pbv1.UpdateCompanyResponse{
				Status:  403,
				Message: "You are not authorized to update this company",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := u.UpdateCompany(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
		})
	}

	// Get Company
	get_res, _, err := tools.GetCompanyByID(r.Id)
	require.NoError(t, err)
	require.Equal(t, "Approve", get_res.Status)
	require.Equal(t, "Mock Company New Name", get_res.Name)
	require.Equal(t, "I am a company New", get_res.Description)
	require.Equal(t, "Bangkok New", get_res.Location)
	require.Equal(t, "0123456780", get_res.Phone)
	require.Equal(t, "IT New", get_res.Category)
	require.Equal(t, "Approve", get_res.Status)
}
