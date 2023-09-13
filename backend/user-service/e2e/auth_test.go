package e2e_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/db"
	"github.com/TikhampornSky/go-auth-verifiedMail/e2e/mock"
	"github.com/TikhampornSky/go-auth-verifiedMail/email"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/repo"
	"github.com/TikhampornSky/go-auth-verifiedMail/server"
	"github.com/TikhampornSky/go-auth-verifiedMail/service"
)

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("..")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	os.Chdir("../")
	db, err := db.NewDatabase(&config)
	if err != nil {
		log.Fatalf("Something went wrong. Could not connect to the database. %s", err)
	}

	memphisConn := email.InitMemphisConnection()
	defer memphisConn.Close()

	memphis := email.NewMemphis(memphisConn, config.MemphisStationNameTest)

	userRepo := repo.NewUserRepository(db.GetPostgresqlDB(), db.GetRedisDB())
	timeService := mock.NewMockTimeProvider()
	authService := service.NewAuthService(userRepo, memphis, timeService)
	userService := service.NewUserService(userRepo, memphis)

	// gRPC Zone
	go server.NewServer(config.ServerPort, authService, userService)

	code := m.Run()
	os.Exit(code)
}

func TestCreateStudent(t *testing.T) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
				Message: "Email must be @student.chula.ac.th",
			},
		},
	}
	testOrder := []string{"success", "email already exists", "password and password confirm not match", "email is not student.chula.ac.th"}

	for _, testName := range testOrder {
		tc := tests[testName]
		t.Run(testName, func(t *testing.T) {
			res, err := c.CreateStudent(ctx, tc.req)
			if err != nil {
				t.Errorf("could not create student: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
			}
		})
	}
}

func TestCreateCompany(t *testing.T) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		expect *pbv1.CreateAdminResponse
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
			expect: &pbv1.CreateAdminResponse{
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
			expect: &pbv1.CreateAdminResponse{
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
			expect: &pbv1.CreateAdminResponse{
				Status:  400,
				Message: "Passwords do not match",
			},
		},
	}
	testOrder := []string{"success", "email already exists", "password and password confirm not match"}

	for _, testName := range testOrder {
		tc := tests[testName]
		t.Run(testName, func(t *testing.T) {
			res, err := c.CreateCompany(ctx, tc.req)
			if err != nil {
				t.Errorf("could not create company: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
			}
		})
	}
}

func TestCreateAdmin(t *testing.T) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	adminTestEmail := utils.GenerateRandomString(10) + "@gmail.com"

	tests := map[string]struct {
		req    *pbv1.CreateAdminRequest
		expect *pbv1.CreateAdminResponse
	}{
		"success": {
			req: &pbv1.CreateAdminRequest{
				Email:           adminTestEmail,
				Password:        "password-test",
				PasswordConfirm: "password-test",
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
			},
			expect: &pbv1.CreateAdminResponse{
				Status:  400,
				Message: "Passwords do not match",
			},
		},
	}
	testOrder := []string{"success", "email already exists", "password and password confirm not match"}

	for _, testName := range testOrder {
		tc := tests[testName]
		t.Run(testName, func(t *testing.T) {
			res, err := c.CreateAdmin(ctx, tc.req)
			if err != nil {
				t.Errorf("could not create admin: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
			}
		})
	}
}

func TestSignIn(t *testing.T) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	a := &pbv1.CreateAdminRequest{
		Email:           utils.GenerateRandomString(10) + "@gmail.com",
		Password:        "password-test",
		PasswordConfirm: "password-test",
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
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Person 1 OK
	mock_email := utils.GenerateRandomString(10) + "@gmail.com"
	_, err = c.CreateAdmin(ctx, &pbv1.CreateAdminRequest{
		Email:           mock_email,
		Password:        "password-test",
		PasswordConfirm: "password-test",
	})
	require.NoError(t, err)

	res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    mock_email,
		Password: "password-test",
	})

	// Wrong token (Unknown person)
	config, _ := config.LoadConfig("..")
	refresh_token_wrong, err := utils.CreateToken(config.RefreshTokenExpiresIn, 0, config.RefreshTokenPrivateKey)
	require.NoError(t, err)

	// Person 2 Already logged out
	mock_email2 := utils.GenerateRandomString(10) + "@gmail.com"
	_, err = c.CreateAdmin(ctx, &pbv1.CreateAdminRequest{
		Email:           mock_email2,
		Password:        "password-test",
		PasswordConfirm: "password-test",
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
				Status:  403,
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
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mock_email := utils.GenerateRandomString(10) + "@gmail.com"
	_, err = c.CreateAdmin(ctx, &pbv1.CreateAdminRequest{
		Email:           mock_email,
		Password:        "password-test",
		PasswordConfirm: "password-test",
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
}

func TestVerifyEmailCode(t *testing.T) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	tests := map[string]struct {
		req    *pbv1.VerifyEmailCodeRequest
		expect *pbv1.VerifyEmailCodeResponse
	}{
		"success": {
			req: &pbv1.VerifyEmailCodeRequest{
				Code:      utils.Encode(id_student, mock.NewMockTimeProvider().Now().Unix()),
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
}
