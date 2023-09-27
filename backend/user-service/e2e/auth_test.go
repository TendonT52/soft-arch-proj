package e2e_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	"github.com/TikhampornSky/go-auth-verifiedMail/tools"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
)

func TestHealthCheck(t *testing.T) {
}

func TestCreateStudent(t *testing.T) {
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

	id_student := utils.GenerateRandomNumber(10)

	tests := map[string]struct {
		req    *pbv1.CreateStudentRequest
		expect *pbv1.CreateStudentResponse
	}{
		"success": {
			req: &pbv1.CreateStudentRequest{
				Name:            "Name Test",
				Email:           id_student + "@student.chula.ac.th",
				Password:        "password-test",
				PasswordConfirm: "password-test",
				Description:     "I am a student",
				Faculty:         "Engineering",
				Major:           "Computer Engineering",
				Year:            4,
			},
			expect: &pbv1.CreateStudentResponse{
				Status:  201,
				Message: "Your account has been created. Please verify your email",
			},
		},
		"email already exists": {
			req: &pbv1.CreateStudentRequest{
				Name:            "Name Test",
				Email:           id_student + "@student.chula.ac.th",
				Password:        "password-test",
				PasswordConfirm: "password-test",
				Description:     "I am a student",
				Faculty:         "Engineering",
				Major:           "Computer Engineering",
				Year:            4,
			},
			expect: &pbv1.CreateStudentResponse{
				Status:  400,
				Message: "Email already exists",
			},
		},
		"password and password confirm not match": {
			req: &pbv1.CreateStudentRequest{
				Name:            "Name Test",
				Email:           id_student + "@student.chula.ac.th",
				Password:        "password-test",
				PasswordConfirm: "password-test-not-match",
				Description:     "I am a student",
				Faculty:         "Engineering",
				Major:           "Computer Engineering",
				Year:            4,
			},
			expect: &pbv1.CreateStudentResponse{
				Status:  400,
				Message: "Passwords do not match",
			},
		},
		"email is not student.chula.ac.th": {
			req: &pbv1.CreateStudentRequest{
				Name:            "Name Test",
				Email:           id_student + "@gmail.com",
				Password:        "password-test",
				PasswordConfirm: "password-test",
				Description:     "I am a student",
				Faculty:         "Engineering",
				Major:           "Computer Engineering",
				Year:            4,
			},
			expect: &pbv1.CreateStudentResponse{
				Status:  400,
				Message: "Email must be studentID with @student.chula.ac.th",
			},
		},
		"email length is less than 20": {
			req: &pbv1.CreateStudentRequest{
				Name:            "Name Test",
				Email:           id_student + "@g.com",
				Password:        "password-test",
				PasswordConfirm: "password-test",
				Description:     "I am a student",
				Faculty:         "Engineering",
				Major:           "Computer Engineering",
				Year:            4,
			},
			expect: &pbv1.CreateStudentResponse{
				Status:  400,
				Message: "Email must be studentID with @student.chula.ac.th",
			},
		},
		"year must be greater than zero": {
			req: &pbv1.CreateStudentRequest{
				Name:            "Name Test",
				Email:           id_student + "@g.com",
				Password:        "password-test",
				PasswordConfirm: "password-test",
				Description:     "I am a student",
				Faculty:         "Engineering",
				Major:           "Computer Engineering",
				Year:            0,
			},
			expect: &pbv1.CreateStudentResponse{
				Status:  400,
				Message: "Year must be greater than zero",
			},
		},
	}
	testOrder := []string{"success", "email already exists", "password and password confirm not match", "email is not student.chula.ac.th", "email length is less than 20", "year must be greater than zero"}
	studentID := 0
	for _, testName := range testOrder {
		tc := tests[testName]
		t.Run(testName, func(t *testing.T) {
			res, err := c.CreateStudent(ctx, tc.req)
			if err != nil {
				t.Errorf("could not create student: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
				if tc.expect.Status == 201 {
					studentID = int(res.Id)
				}
			}
		})
	}

	// Check if student is created
	require.NotEqual(t, 0, studentID)
	student, status, err := tools.GetStudentByID(int64(studentID))
	require.NoError(t, err)
	require.Equal(t, id_student+"@student.chula.ac.th", student.Email)
	require.Equal(t, "student", status.Role)
	require.Equal(t, false, status.Verified)
	require.Equal(t, "Name Test", student.Name)
	require.Equal(t, "I am a student", student.Description)
	require.Equal(t, "Engineering", student.Faculty)
	require.Equal(t, "Computer Engineering", student.Major)
	require.Equal(t, int32(4), student.Year)
}

func TestCreateCompany(t *testing.T) {
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

	companyTestEmail := utils.GenerateRandomString(10) + "@gmail.com"

	tests := map[string]struct {
		req    *pbv1.CreateCompanyRequest
		expect *pbv1.CreateCompanyResponse
	}{
		"success": {
			req: &pbv1.CreateCompanyRequest{
				Name:            "Comapany Name Test",
				Email:           companyTestEmail,
				Password:        "password-test",
				PasswordConfirm: "password-test",
				Description:     "I am a company",
				Location:        "Bangkok",
				Phone:           "0123456789",
				Category:        "Technology",
			},
			expect: &pbv1.CreateCompanyResponse{
				Status:  201,
				Message: "The Approval process will take 1-2 days. Thank you for your patience",
			},
		},
		"email already exists": {
			req: &pbv1.CreateCompanyRequest{
				Name:            "Comapany Name Test",
				Email:           companyTestEmail,
				Password:        "password-test",
				PasswordConfirm: "password-test",
				Description:     "I am a company",
				Location:        "Bangkok",
				Phone:           "0123456789",
				Category:        "Technology",
			},
			expect: &pbv1.CreateCompanyResponse{
				Status:  400,
				Message: "Email already exists",
			},
		},
		"password and password confirm not match": {
			req: &pbv1.CreateCompanyRequest{
				Name:            "Comapany Name Test",
				Email:           companyTestEmail,
				Password:        "password-test",
				PasswordConfirm: "password-test-not-match",
				Description:     "I am a company",
				Location:        "Bangkok",
				Phone:           "0123456789",
				Category:        "Technology",
			},
			expect: &pbv1.CreateCompanyResponse{
				Status:  400,
				Message: "Passwords do not match",
			},
		},
	}
	testOrder := []string{"success", "email already exists", "password and password confirm not match"}

	companyID := 0
	for _, testName := range testOrder {
		tc := tests[testName]
		t.Run(testName, func(t *testing.T) {
			res, err := c.CreateCompany(ctx, tc.req)
			if err != nil {
				t.Errorf("could not create company: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
				if tc.expect.Status == 201 {
					companyID = int(res.Id)
				}
			}
		})
	}

	// Check if company is created
	require.NotEqual(t, 0, companyID)
	company, status, err := tools.GetCompanyByID(int64(companyID))
	require.NoError(t, err)
	require.Equal(t, companyTestEmail, company.Email)
	require.Equal(t, "Comapany Name Test", company.Name)
	require.Equal(t, "I am a company", company.Description)
	require.Equal(t, "Bangkok", company.Location)
	require.Equal(t, "0123456789", company.Phone)
	require.Equal(t, "Technology", company.Category)
	require.Equal(t, "Pending", company.Status)
	require.Equal(t, "company", status.Role)
	require.Equal(t, false, status.Verified)
}

func TestCreateAdmin(t *testing.T) {
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

	adminTestEmail := utils.GenerateRandomString(10) + "@gmail.com"
	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	student_access_token, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &domain.Payload{
		UserId: 0,
		Role:   domain.StudentRole,
	})

	tests := map[string]struct {
		req    *pbv1.CreateAdminRequest
		expect *pbv1.CreateAdminResponse
	}{
		"success": {
			req: &pbv1.CreateAdminRequest{
				Email:           adminTestEmail,
				Password:        "password-test",
				PasswordConfirm: "password-test",
				AccessToken:     admin_access_token,
			},
			expect: &pbv1.CreateAdminResponse{
				Status:  201,
				Message: "Welcome to admin world!",
			},
		},
		"email already exists": {
			req: &pbv1.CreateAdminRequest{
				Email:           adminTestEmail,
				Password:        "password-test",
				PasswordConfirm: "password-test",
				AccessToken:     admin_access_token,
			},
			expect: &pbv1.CreateAdminResponse{
				Status:  400,
				Message: "Email already exists",
			},
		},
		"password and password confirm not match": {
			req: &pbv1.CreateAdminRequest{
				Email:           adminTestEmail,
				Password:        "password-test",
				PasswordConfirm: "password-test-not-match",
				AccessToken:     admin_access_token,
			},
			expect: &pbv1.CreateAdminResponse{
				Status:  400,
				Message: "Passwords do not match",
			},
		},
		"not admin to create": {
			req: &pbv1.CreateAdminRequest{
				Email:           adminTestEmail,
				Password:        "password-test",
				PasswordConfirm: "password-test",
				AccessToken:     student_access_token,
			},
			expect: &pbv1.CreateAdminResponse{
				Status:  403,
				Message: "You are not admin",
			},
		},
	}
	testOrder := []string{"success", "email already exists", "password and password confirm not match"}
	adminId := 0
	for _, testName := range testOrder {
		tc := tests[testName]
		t.Run(testName, func(t *testing.T) {
			res, err := c.CreateAdmin(ctx, tc.req)
			if err != nil {
				t.Errorf("could not create admin: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
				if tc.expect.Status == 201 {
					adminId = int(res.Id)
				}
			}
		})
	}

	// Check if admin is created
	require.NotEqual(t, 0, adminId)
	admin, err := tools.GetUserByID(int64(adminId))
	require.NoError(t, err)
	require.Equal(t, adminTestEmail, admin.Email)
	require.Equal(t, "admin", admin.Role)
	require.Equal(t, true, admin.Verified)
}

func TestSignIn(t *testing.T) {
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

	id_student := utils.GenerateRandomNumber(10)
	s := &pbv1.CreateStudentRequest{
		Name:            "Mock SignIn",
		Email:           id_student + "@student.chula.ac.th",
		Password:        "password-test",
		PasswordConfirm: "password-test",
		Description:     "I am a student",
		Faculty:         "Engineering",
		Major:           "Computer Engineering",
		Year:            4,
	}
	_, err = c.CreateStudent(ctx, s)
	require.NoError(t, err)

	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	a := &pbv1.CreateAdminRequest{
		Email:           utils.GenerateRandomString(10) + "@gmail.com",
		Password:        "password-test",
		PasswordConfirm: "password-test",
		AccessToken:     admin_access_token,
	}
	_, err = c.CreateAdmin(ctx, a)
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.LoginRequest
		expect *pbv1.LoginResponse
	}{
		"success": {
			req: &pbv1.LoginRequest{
				Email:    a.Email,
				Password: a.Password,
			},
			expect: &pbv1.LoginResponse{
				Status:  200,
				Message: "Login success",
			},
		},
		"password not match": {
			req: &pbv1.LoginRequest{
				Email:    a.Email,
				Password: "password-test-not-match",
			},
			expect: &pbv1.LoginResponse{
				Status:  400,
				Message: "Passwords do not match",
			},
		},
		"account not verified": {
			req: &pbv1.LoginRequest{
				Email:    s.Email,
				Password: s.Password,
			},
			expect: &pbv1.LoginResponse{
				Status:  400,
				Message: "Your account is not verified",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.SignIn(ctx, tc.req)
			if err != nil {
				t.Errorf("could not sign in: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
			}
		})
	}
}

func TestRefreshToken(t *testing.T) {
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

	// Person 1 OK
	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	mock_email := utils.GenerateRandomString(10) + "@gmail.com"
	_, err = c.CreateAdmin(ctx, &pbv1.CreateAdminRequest{
		Email:           mock_email,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		AccessToken:     admin_access_token,
	})
	require.NoError(t, err)

	res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    mock_email,
		Password: "password-test",
	})

	// Wrong token (Unknown person)
	refresh_token_wrong, err := utils.CreateRefreshToken(config.RefreshTokenExpiresIn, 0)
	require.NoError(t, err)

	// Person 2 Already logged out
	mock_email2 := utils.GenerateRandomString(10) + "@gmail.com"
	_, err = c.CreateAdmin(ctx, &pbv1.CreateAdminRequest{
		Email:           mock_email2,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		AccessToken:     admin_access_token,
	})
	u, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    mock_email2,
		Password: "password-test",
	})
	require.NoError(t, err)
	_, err = c.LogOut(ctx, &pbv1.LogOutRequest{
		RefreshToken: u.RefreshToken,
	})
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.RefreshTokenRequest
		expect *pbv1.RefreshTokenResponse
	}{
		"success": {
			req: &pbv1.RefreshTokenRequest{
				RefreshToken: res.RefreshToken,
			},
			expect: &pbv1.RefreshTokenResponse{
				Status:  200,
				Message: "Refresh token success",
			},
		},
		"user not found": {
			req: &pbv1.RefreshTokenRequest{
				RefreshToken: refresh_token_wrong,
			},
			expect: &pbv1.RefreshTokenResponse{
				Status:  403,
				Message: "the user belonging to this token no logger exists",
			},
		},

		"already logged out": {
			req: &pbv1.RefreshTokenRequest{
				RefreshToken: u.RefreshToken,
			},
			expect: &pbv1.RefreshTokenResponse{
				Status:  500,
				Message: "your token has been logged out!",
			},
		},
	}
	testOrder := []string{"success", "user not found", "already logged out"}

	for _, testName := range testOrder {
		tc := tests[testName]
		t.Run(testName, func(t *testing.T) {
			res, err := c.RefreshToken(ctx, tc.req)
			if err != nil {
				t.Errorf("could not refresh token: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
			}
		})
	}
}

func TestLogOut(t *testing.T) {
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

	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	mock_email := utils.GenerateRandomString(10) + "@gmail.com"
	_, err = c.CreateAdmin(ctx, &pbv1.CreateAdminRequest{
		Email:           mock_email,
		Password:        "password-test",
		PasswordConfirm: "password-test",
		AccessToken:     admin_access_token,
	})
	require.NoError(t, err)

	res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    mock_email,
		Password: "password-test",
	})

	tests := map[string]struct {
		req    *pbv1.LogOutRequest
		expect *pbv1.LogOutResponse
	}{
		"success": {
			req: &pbv1.LogOutRequest{
				RefreshToken: res.RefreshToken,
			},
			expect: &pbv1.LogOutResponse{
				Status:  200,
				Message: "Logout success",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.LogOut(ctx, tc.req)
			if err != nil {
				t.Errorf("could not log out: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
			}
		})
	}

	// Check if refresh token is deleted
	str, err := tools.GetValueFromRedis(res.RefreshToken)
	require.NoError(t, err)
	require.Equal(t, "logged_out", str)
}

func TestVerifyEmailCode(t *testing.T) {
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

	id_student := utils.GenerateRandomNumber(10)
	s := &pbv1.CreateStudentRequest{
		Name:            "Mock SignIn",
		Email:           id_student + "@student.chula.ac.th",
		Password:        "password-test",
		PasswordConfirm: "password-test",
		Description:     "I am a student",
		Faculty:         "Engineering",
		Major:           "Computer Engineering",
		Year:            4,
	}

	res, err := c.CreateStudent(ctx, s)
	require.NoError(t, err)
	require.Equal(t, int64(201), res.Status)

	timeNow, err := tools.GetCreateTime(res.Id)
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.VerifyEmailCodeRequest
		expect *pbv1.VerifyEmailCodeResponse
	}{
		"success": {
			req: &pbv1.VerifyEmailCodeRequest{
				Code:      utils.Encode(id_student, timeNow),
				StudentId: id_student,
			},
			expect: &pbv1.VerifyEmailCodeResponse{
				Status:  200,
				Message: "verify success",
			},
		},

		"wrong code": {
			req: &pbv1.VerifyEmailCodeRequest{
				Code: "1234567",
			},
			expect: &pbv1.VerifyEmailCodeResponse{
				Status:  400,
				Message: "Invalid verification code or user doesn't exists",
			},
		},
	}
	testOrder := []string{"success", "wrong code"}

	for _, testName := range testOrder {
		tc := tests[testName]
		t.Run(testName, func(t *testing.T) {
			res, err := c.VerifyEmailCode(ctx, tc.req)
			if err != nil {
				t.Errorf("could not verify email code: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
			}
		})
	}

	// Check if student is verified
	_, status, err := tools.GetStudentByID(res.Id)
	require.NoError(t, err)
	require.Equal(t, true, status.Verified)
	require.Equal(t, "student", status.Role)
}
