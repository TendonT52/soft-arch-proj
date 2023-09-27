package e2e

import (
	"context"
	"fmt"
	"strconv"
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

func TestUpdateCompanyStatus(t *testing.T) {
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

	// Delete all companies in table
	err = tools.DeleteAll()
	require.NoError(t, err)

	// Create Admin
	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	admin := &pbv1.CreateAdminRequest{
		Email:           utils.GenerateRandomString(2) + "@admin.com",
		Password:        "password-test",
		PasswordConfirm: "password-test",
		AccessToken:     admin_access_token,
	}
	a, err := c.CreateAdmin(ctx, admin)
	require.Equal(t, int64(201), a.Status)
	require.NoError(t, err)

	// Admin Sign In
	ad, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    admin.Email,
		Password: admin.Password,
	})

	// Register
	companyEmail := utils.GenerateRandomString(3) + "@company.com"
	company := &pbv1.CreateCompanyRequest{
		Name:            "Mock Company",
		Email:           companyEmail,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		Description:     "I am a company",
		Location:        "Bangkok",
		Phone:           "0123456789",
		Category:        "IT",
	}
	com, err := c.CreateCompany(ctx, company)
	require.Equal(t, int64(201), com.Status)
	require.NoError(t, err)

	// Register
	companyEmail2 := utils.GenerateRandomString(4) + "@company.com"
	company2 := &pbv1.CreateCompanyRequest{
		Name:            "Mock Company 2",
		Email:           companyEmail2,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		Description:     "I am a company2",
		Location:        "Bangkok",
		Phone:           "0123456789",
		Category:        "IT",
	}
	com2, err := c.CreateCompany(ctx, company2)
	require.Equal(t, int64(201), com2.Status)
	require.NoError(t, err)

	// Generate WRONG token
	access_token_wrong, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &domain.Payload{
		UserId: 0,
		Role:   domain.CompanyRole,
	})
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.UpdateCompanyStatusRequest
		expect *pbv1.UpdateCompanyStatusResponse
	}{
		"success approve": {
			req: &pbv1.UpdateCompanyStatusRequest{
				AccessToken: ad.AccessToken,
				Id:          com.Id,
				Status:      domain.ComapanyStatusApprove,
			},
			expect: &pbv1.UpdateCompanyStatusResponse{
				Status:  200,
				Message: "Update status for company id " + strconv.FormatInt(com.Id, 10) + " successfully!",
			},
		},
		"success reject": {
			req: &pbv1.UpdateCompanyStatusRequest{
				AccessToken: ad.AccessToken,
				Id:          com2.Id,
				Status:      domain.ComapanyStatusReject,
			},
			expect: &pbv1.UpdateCompanyStatusResponse{
				Status:  200,
				Message: "Update status for company id " + strconv.FormatInt(com2.Id, 10) + " successfully!",
			},
		},
		"fail: not admin": {
			req: &pbv1.UpdateCompanyStatusRequest{
				AccessToken: access_token_wrong,
				Id:          com.Id,
				Status:      domain.ComapanyStatusApprove,
			},
			expect: &pbv1.UpdateCompanyStatusResponse{
				Status:  403,
				Message: "Only admin can approve",
			},
		},
		"fail: company already approved": {
			req: &pbv1.UpdateCompanyStatusRequest{
				AccessToken: ad.AccessToken,
				Id:          com.Id,
				Status:      domain.ComapanyStatusReject,
			},
			expect: &pbv1.UpdateCompanyStatusResponse{
				Status:  400,
				Message: "company already approved or rejected",
			},
		},
		"invalidate status": {
			req: &pbv1.UpdateCompanyStatusRequest{
				AccessToken: ad.AccessToken,
				Id:          com.Id,
				Status:      "verified",
			},
			expect: &pbv1.UpdateCompanyStatusResponse{
				Status:  400,
				Message: "status must be Approve or Reject",
			},
		},
	}
	testOrder := []string{"success approve", "success reject", "fail: not admin", "fail: company already approved", "invalidate status"}

	for _, testName := range testOrder {
		tc := tests[testName]
		t.Run(testName, func(t *testing.T) {
			res, err := u.UpdateCompanyStatus(ctx, tc.req)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			require.NoError(t, err)
		})
	}
}

func createMockComapny(t *testing.T, name, email, description, location, phone, category string) *pbv1.Company {
	config, _ := config.LoadConfig("..")
	target := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	company := &pbv1.CreateCompanyRequest{
		Name:            name,
		Email:           email,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		Description:     description,
		Location:        location,
		Phone:           phone,
		Category:        category,
	}
	com, err := c.CreateCompany(ctx, company)
	require.Equal(t, int64(201), com.Status)
	require.NoError(t, err)

	return &pbv1.Company{
		Id:          com.Id,
		Name:        name,
		Email:       email,
		Description: description,
		Location:    location,
		Phone:       phone,
		Category:    category,
		Status:      domain.ComapanyStatusPending,
	}
}

func TestListCompanies(t *testing.T) {
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

	// Delete all companies in table
	err = tools.DeleteAll()
	require.NoError(t, err)

	// Create Admin
	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	admin := &pbv1.CreateAdminRequest{
		Email:           utils.GenerateRandomString(11) + "@admin.com",
		Password:        "password-test",
		PasswordConfirm: "password-test",
		AccessToken:     admin_access_token,
	}
	a, err := c.CreateAdmin(ctx, admin)
	require.Equal(t, int64(201), a.Status)
	require.NoError(t, err)

	// Admin Sign In
	ad, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    admin.Email,
		Password: admin.Password,
	})

	// Generate WRONG token
	access_token_wrong, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	require.NoError(t, err)

	// name, email, description, location, phone, category
	mail1 := utils.GenerateRandomString(12) + "@company.com"
	mail2 := utils.GenerateRandomString(13) + "@company.com"
	c1 := createMockComapny(t, "Mock Company 1", mail1, "Company1 desc", "Bangkok", "0123456789", "Technology")
	c2 := createMockComapny(t, "Mock Company 2", mail2, "Company2 desc", "Bangkok", "0123456789", "Technical Finanace")

	tests := map[string]struct {
		req    *pbv1.ListCompaniesRequest
		expect *pbv1.ListCompaniesResponse
	}{
		"success list all companies": {
			req: &pbv1.ListCompaniesRequest{
				AccessToken: ad.AccessToken,
			},
			expect: &pbv1.ListCompaniesResponse{
				Status:  200,
				Message: "success",
				Companies: []*pbv1.Company{
					c1,
					c2,
				},
			},
		},
		"fail: not admin": {
			req: &pbv1.ListCompaniesRequest{
				AccessToken: access_token_wrong,
			},
			expect: &pbv1.ListCompaniesResponse{
				Status:  403,
				Message: "Only admin can view",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := u.ListCompanies(ctx, tc.req)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			require.NoError(t, err)
			require.Equal(t, len(tc.expect.Companies), len(res.Companies))
			for i, s := range res.Companies {
				require.Equal(t, tc.expect.Companies[i].Name, s.Name)
				require.Equal(t, tc.expect.Companies[i].Email, s.Email)
				require.Equal(t, tc.expect.Companies[i].Description, s.Description)
				require.Equal(t, tc.expect.Companies[i].Location, s.Location)
				require.Equal(t, tc.expect.Companies[i].Phone, s.Phone)
				require.Equal(t, tc.expect.Companies[i].Category, s.Category)
				require.Equal(t, tc.expect.Companies[i].Status, s.Status)
			}
		})
	}
}
