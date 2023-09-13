package e2e

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestUpdateCompanyStatus(t *testing.T) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)
	u := pbv1.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create Admin
	admin := &pbv1.CreateAdminRequest{
		Email:           utils.GenerateRandomString(10) + "@admin.com",
		Password:        "password-test",
		PasswordConfirm: "password-test",
	}
	a, err := c.CreateAdmin(ctx, admin)
	require.Equal(t, int64(201), a.Status)
	require.NoError(t, err)

	// Admin Sign In
	ad, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    admin.Email,
		Password: admin.Password,
	})

	// Delete all companies in table
	d, err := u.DeleteCompanies(ctx, &pbv1.DeleteCompaniesRequest{
		AccessToken: ad.AccessToken,
	})
	require.Equal(t, int64(200), d.Status)
	require.NoError(t, err)

	// Register
	companyEmail := utils.GenerateRandomString(10) + "@company.com"
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
	companyEmail2 := utils.GenerateRandomString(10) + "@company.com"
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
	config, _ := config.LoadConfig("..")
	access_token_wrong, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &pbv1.Payload{
		UserId: 0,
		Role:   domain.CompanyRole,
	}, config.AccessTokenPrivateKey)
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.UpdateCompanyStatusRequest
		expect *pbv1.UpdateCompanyStatusResponse
	}{
		"success approve": {
			req: &pbv1.UpdateCompanyStatusRequest{
				AccessToken: ad.AccessToken,
				Id:          com.Id,
				Status:      "Approve",
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
				Status:      "Reject",
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
				Status:      "verified",
			},
			expect: &pbv1.UpdateCompanyStatusResponse{
				Status:  401,
				Message: "Only admin can approve",
			},
		},
		"fail: company already approved": {
			req: &pbv1.UpdateCompanyStatusRequest{
				AccessToken: ad.AccessToken,
				Id:          com.Id,
				Status:      "Reject",
			},
			expect: &pbv1.UpdateCompanyStatusResponse{
				Status:  400,
				Message: "company already approved or rejected",
			},
		},
	}
	testOrder := []string{"success approve", "success reject", "fail: not admin", "fail: company already approved"}

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

func createMockComapny(t *testing.T, name, email, description, location, phone, category string) *pbv1.CreateCompanyResponse {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	return com
}

func TestListCompanies(t *testing.T) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)
	u := pbv1.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create Admin
	admin := &pbv1.CreateAdminRequest{
		Email:           utils.GenerateRandomString(10) + "@admin.com",
		Password:        "password-test",
		PasswordConfirm: "password-test",
	}
	a, err := c.CreateAdmin(ctx, admin)
	require.Equal(t, int64(201), a.Status)
	require.NoError(t, err)

	// Admin Sign In
	ad, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    admin.Email,
		Password: admin.Password,
	})

	// Delete all companies in table
	d, err := u.DeleteCompanies(ctx, &pbv1.DeleteCompaniesRequest{
		AccessToken: ad.AccessToken,
	})
	require.Equal(t, int64(200), d.Status)
	require.NoError(t, err)

	// Generate WRONG token
	config, _ := config.LoadConfig("..")
	access_token_wrong, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &pbv1.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	}, config.AccessTokenPrivateKey)
	require.NoError(t, err)

	// name, email, description, location, phone, category
	mail1 := utils.GenerateRandomString(10) + "@company.com"
	mail2 := utils.GenerateRandomString(10) + "@company.com"
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
					{
						Id:          c1.Id,
						Name:        "Mock Company 1",
						Email:       mail1,
						Description: "Company1 desc",
						Location:    "Bangkok",
						Phone:       "0123456789",
						Category:    "Technology",
						Status:      "Pending",
					},
					{
						Id:          c2.Id,
						Name:        "Mock Company 2",
						Email:       mail2,
						Description: "Company2 desc",
						Location:    "Bangkok",
						Phone:       "0123456789",
						Category:    "Technical Finanace",
						Status:      "Pending",
					},
				},
			},
		},
		"fail: not admin": {
			req: &pbv1.ListCompaniesRequest{
				AccessToken: access_token_wrong,
			},
			expect: &pbv1.ListCompaniesResponse{
				Status:  401,
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
