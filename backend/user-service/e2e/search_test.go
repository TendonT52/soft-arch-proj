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

func approveMultipleCompanies(t *testing.T, ids []int64, ad *pbv1.LoginResponse, u pbv1.UserServiceClient, com_status string) {
	for _, id := range ids {
		aa, err := u.UpdateCompanyStatus(context.Background(), &pbv1.UpdateCompanyStatusRequest{
			AccessToken: ad.AccessToken,
			Id:          id,
			Status:      com_status,
		})
		require.Equal(t, int64(200), aa.Status)
		require.NoError(t, err)
	}
}

func createMockStudent(t *testing.T, admin_access string) string {
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

	// Register
	studentEmail := utils.GenerateRandomNumber(10) + "@student.chula.ac.th"
	student := &pbv1.CreateStudentRequest{
		Name:            "Mock Student" + utils.GenerateRandomString(3),
		Email:           studentEmail,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		Faculty:         "Medical",
		Major:           "Doctor",
		Year:            5,
	}
	stu, err := c.CreateStudent(ctx, student)
	require.Equal(t, int64(201), stu.Status)
	require.NoError(t, err)

	// Verify Student
	timeNow, err := tools.GetCreateTime(stu.Id)
	require.NoError(t, err)

	v, err := c.VerifyEmailCode(ctx, &pbv1.VerifyEmailCodeRequest{
		StudentId: studentEmail[:10],
		Code:      utils.Encode(studentEmail[:10], timeNow),
	})
	require.Equal(t, int64(200), v.Status)
	require.NoError(t, err)

	// Student Sign In
	st, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    studentEmail,
		Password: student.Password,
	})
	require.Equal(t, int64(200), st.Status)
	require.NoError(t, err)

	return st.AccessToken
}

func TestListApprovedCompanies(t *testing.T) {
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
		Email:           utils.GenerateRandomString(18) + "@admin.com",
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
	require.Equal(t, int64(200), ad.Status)
	require.NoError(t, err)

	// Generate WRONG token
	access_token_wrong, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &domain.Payload{
		UserId: 0,
		Role:   domain.StudentRole,
	})
	require.NoError(t, err)

	// name, email, description, location, phone, category
	mail1 := utils.GenerateRandomString(10) + "@company.com"
	mail2 := utils.GenerateRandomString(10) + "@company.com"
	mail3 := utils.GenerateRandomString(10) + "@company.com"
	mail4 := utils.GenerateRandomString(10) + "@company.com"
	mail5 := utils.GenerateRandomString(10) + "@company.com"

	c1 := createMockComapny(t, "Mock Company 1", mail1, "Company1 desc", "Bangkok", "0123456789", "Technology")
	c2 := createMockComapny(t, "Mock Company 2", mail2, "Company2 desc", "Bangkok", "0123456789", "Bank and Tech")
	c3 := createMockComapny(t, "Mock Company 3 Tech Group", mail3, "Company3 desc", "Bangkok", "0123456789", "Consultant")
	c4 := createMockComapny(t, "Mock Company 4", mail4, "Company4 desc", "Bangkok", "0123456789", "Food and Beverage")
	c5 := createMockComapny(t, "Mock Company 5 Tech Company", mail5, "Company5 desc", "Bangkok", "0123456789", "Food")

	// Approve Companies
	approveMultipleCompanies(t, []int64{c1.Id, c2.Id, c3.Id, c4.Id}, ad, u, domain.ComapanyStatusApprove)
	approveMultipleCompanies(t, []int64{c5.Id}, ad, u, domain.ComapanyStatusReject)

	// Create Student
	student_access := createMockStudent(t, ad.AccessToken)

	tests := map[string]struct {
		req    *pbv1.ListApprovedCompaniesRequest
		expect *pbv1.ListApprovedCompaniesResponse
	}{
		"success admin": {
			req: &pbv1.ListApprovedCompaniesRequest{
				AccessToken: ad.AccessToken,
				Search:      "Tech!!",
			},
			expect: &pbv1.ListApprovedCompaniesResponse{
				Status:  200,
				Message: "success",
				Companies: []*pbv1.Company{
					c1,
					c3,
					c2,
				},
			},
		},
		"success student": {
			req: &pbv1.ListApprovedCompaniesRequest{
				AccessToken: student_access,
				Search:      "Food and Beverage",
			},
			expect: &pbv1.ListApprovedCompaniesResponse{
				Status:  200,
				Message: "success",
				Companies: []*pbv1.Company{
					c4,
					c2,
				},
			},
		},
		"userID not found": {
			req: &pbv1.ListApprovedCompaniesRequest{
				AccessToken: access_token_wrong,
			},
			expect: &pbv1.ListApprovedCompaniesResponse{
				Status:  500,
				Message: "the user belonging to this token no logger exists",
			},
		},
		"search value empty": {
			req: &pbv1.ListApprovedCompaniesRequest{
				AccessToken: ad.AccessToken,
				Search:      "",
			},
			expect: &pbv1.ListApprovedCompaniesResponse{
				Status:  200,
				Message: "success",
				Companies: []*pbv1.Company{
					c1,
					c2,
					c3,
					c4,
				},
			},
		},
		"invalidate token": {
			req: &pbv1.ListApprovedCompaniesRequest{
				AccessToken: "wrong token",
			},
			expect: &pbv1.ListApprovedCompaniesResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := u.ListApprovedCompanies(ctx, tc.req)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			require.NoError(t, err)
			require.Equal(t, len(tc.expect.Companies), len(res.Companies))
			for i, s := range res.Companies {
				require.Equal(t, tc.expect.Companies[i].Name, s.Name)
				require.Equal(t, tc.expect.Companies[i].Email, s.Email)
			}
		})
	}
}
