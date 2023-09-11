package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"github.com/TikhampornSky/go-auth-verifiedMail/e2e/mock"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
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
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Register
	studentEmail := utils.GenerateRandomString(10) + "@student.chula.ac.th"
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
	v, err := c.VerifyEmailCode(ctx, &pbv1.VerifyEmailCodeRequest{
		StudentId: studentEmail[:10],
		Code:      utils.Encode(studentEmail[:10], mock.NewMockTimeProvider().Now().Unix()),
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
	access_token_wrong, err := utils.CreateToken(config.AccessTokenExpiresIn, 0, config.AccessTokenPrivateKey)
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
	approveMultipleCompanies(t, []int64{c1.Id, c2.Id, c3.Id, c4.Id}, ad, u, "Approve")
	approveMultipleCompanies(t, []int64{c5.Id}, ad, u, "Reject")

	// Create Student
	student_access := createMockStudent(t, ad.AccessToken)

	tests := map[string]struct {
		req    *pbv1.ListApprovedCompaniesRequest
		expect *pbv1.ListApprovedCompaniesResponse
	}{
		"success admin": {
			req: &pbv1.ListApprovedCompaniesRequest{
				AccessToken: ad.AccessToken,
				Search:      "Tech",
			},
			expect: &pbv1.ListApprovedCompaniesResponse{
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
						Status:      "Approve",
					},
					{
						Id:          c3.Id,
						Name:        "Mock Company 3 Tech Group",
						Email:       mail3,
						Description: "Company3 desc",
						Location:    "Bangkok",
						Phone:       "0123456789",
						Category:    "Consultant",
						Status:      "Approve",
					},
					{
						Id:          c2.Id,
						Name:        "Mock Company 2",
						Email:       mail2,
						Description: "Company2 desc",
						Location:    "Bangkok",
						Phone:       "0123456789",
						Category:    "Bank and Tech",
						Status:      "Approve",
					},
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
					{
						Id:          c4.Id,
						Name:        "Mock Company 4",
						Email:       mail4,
						Description: "Company4 desc",
						Location:    "Bangkok",
						Phone:       "0123456789",
						Category:    "Food and Beverage",
						Status:      "Approve",
					},
					{
						Id:          c2.Id,
						Name:        "Mock Company 2",
						Email:       mail2,
						Description: "Company2 desc",
						Location:    "Bangkok",
						Phone:       "0123456789",
						Category:    "Bank and Tech",
						Status:      "Approve",
					},
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
