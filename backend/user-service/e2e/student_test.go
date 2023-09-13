package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	"github.com/TikhampornSky/go-auth-verifiedMail/e2e/mock"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetStudentMe(t *testing.T) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)
	u := pbv1.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Register
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
	r, err := c.CreateStudent(ctx, s)
	require.Equal(t, int64(201), r.Status)
	require.NoError(t, err)

	// Verify Email
	result, err := c.VerifyEmailCode(ctx, &pbv1.VerifyEmailCodeRequest{
		StudentId: id_student,
		Code:      utils.Encode(id_student, mock.NewMockTimeProvider().Now().Unix()),
	})
	require.Equal(t, int64(200), result.Status)
	require.NoError(t, err)

	// Sign In
	res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    s.Email,
		Password: s.Password,
	})
	require.Equal(t, int64(200), res.Status)
	require.NoError(t, err)

	// Generate WRONG token
	config, _ := config.LoadConfig("..")
	access_token_wrong, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &pbv1.Payload{
		UserId: 0,
		Role:   domain.StudentRole,
	}, config.AccessTokenPrivateKey)
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.GetStudentMeRequest
		expect *pbv1.GetStudentResponse
	}{
		"success": {
			req: &pbv1.GetStudentMeRequest{
				AccessToken: res.AccessToken,
			},
			expect: &pbv1.GetStudentResponse{
				Status: 200,
				Student: &pbv1.Student{
					Id:          1,
					Name:        "Mock SignIn",
					Email:       id_student + "@student.chula.ac.th",
					Description: "I am a student",
					Faculty:     "Engineering",
					Major:       "Computer Engineering",
					Year:        4,
				},
			},
		},
		"invalid token": {
			req: &pbv1.GetStudentMeRequest{
				AccessToken: access_token_wrong,
			},
			expect: &pbv1.GetStudentResponse{
				Status: 500,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := u.GetStudentMe(ctx, tc.req)
			if err != nil {
				t.Errorf("could not sign in: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				if tc.expect.Student != nil {
					require.Equal(t, tc.expect.Student.Name, res.Student.Name)
					require.Equal(t, tc.expect.Student.Email, res.Student.Email)
					require.Equal(t, tc.expect.Student.Description, res.Student.Description)
					require.Equal(t, tc.expect.Student.Faculty, res.Student.Faculty)
					require.Equal(t, tc.expect.Student.Major, res.Student.Major)
					require.Equal(t, tc.expect.Student.Year, res.Student.Year)
				}
			}
		})
	}
}

func TestGetStudent(t *testing.T) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)
	u := pbv1.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Register
	id_student := utils.GenerateRandomNumber(10)
	s := &pbv1.CreateStudentRequest{
		Name:            "Mock Get Student",
		Email:           id_student + "@student.chula.ac.th",
		Password:        "password-test",
		PasswordConfirm: "password-test",
		Description:     "I am a student",
		Faculty:         "Engineering",
		Major:           "Computer Engineering",
		Year:            4,
	}
	r, err := c.CreateStudent(ctx, s)
	require.Equal(t, int64(201), r.Status)
	require.NoError(t, err)

	// Verify Email
	result, err := c.VerifyEmailCode(ctx, &pbv1.VerifyEmailCodeRequest{
		StudentId: id_student,
		Code:      utils.Encode(id_student, mock.NewMockTimeProvider().Now().Unix()),
	})
	require.Equal(t, int64(200), result.Status)
	require.NoError(t, err)

	// Sign In
	res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    s.Email,
		Password: s.Password,
	})
	require.Equal(t, int64(200), res.Status)
	require.NoError(t, err)

	// Create Admin
	a := &pbv1.CreateAdminRequest{
		Email:           utils.GenerateRandomString(10) + "@gmail.com",
		Password:        "password-test",
		PasswordConfirm: "password-test",
	}
	_, err = c.CreateAdmin(ctx, a)
	require.NoError(t, err)

	// Sign In Admin
	res, err = c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    a.Email,
		Password: a.Password,
	})

	tests := map[string]struct {
		req    *pbv1.GetStudentRequest
		expect *pbv1.GetStudentResponse
	}{
		"success": {
			req: &pbv1.GetStudentRequest{
				AccessToken: res.AccessToken,
				Id:          r.Id,
			},
			expect: &pbv1.GetStudentResponse{
				Status: 200,
				Student: &pbv1.Student{
					Id:          r.Id,
					Name:        "Mock Get Student",
					Email:       id_student + "@student.chula.ac.th",
					Description: "I am a student",
					Faculty:     "Engineering",
					Major:       "Computer Engineering",
					Year:        4,
				},
			},
		},
		"userID not found": {
			req: &pbv1.GetStudentRequest{
				AccessToken: res.AccessToken,
				Id:          20000000000000,
			},
			expect: &pbv1.GetStudentResponse{
				Status:  404,
				Message: "user id not found",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := u.GetStudent(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			if tc.expect.Student != nil {
				require.Equal(t, tc.expect.Student.Name, res.Student.Name)
				require.Equal(t, tc.expect.Student.Email, res.Student.Email)
				require.Equal(t, tc.expect.Student.Description, res.Student.Description)
				require.Equal(t, tc.expect.Student.Faculty, res.Student.Faculty)
				require.Equal(t, tc.expect.Student.Major, res.Student.Major)
				require.Equal(t, tc.expect.Student.Year, res.Student.Year)
			}
		})
	}
}

func TestUpdateStudent(t *testing.T) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewAuthServiceClient(conn)
	u := pbv1.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Register
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
	r, err := c.CreateStudent(ctx, s)
	require.Equal(t, int64(201), r.Status)
	require.NoError(t, err)

	// Verify Email
	result, err := c.VerifyEmailCode(ctx, &pbv1.VerifyEmailCodeRequest{
		StudentId: id_student,
		Code:      utils.Encode(id_student, mock.NewMockTimeProvider().Now().Unix()),
	})
	require.Equal(t, int64(200), result.Status)
	require.NoError(t, err)

	// Sign In
	res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    s.Email,
		Password: s.Password,
	})
	require.Equal(t, int64(200), res.Status)
	require.NoError(t, err)

	// Generate WRONG token
	config, _ := config.LoadConfig("..")
	access_token_wrong, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &pbv1.Payload{
		UserId: 0,
		Role:   domain.StudentRole,
	}, config.AccessTokenPrivateKey)
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.UpdateStudentRequest
		expect *pbv1.UpdateCompanyResponse
	}{
		"success": {
			req: &pbv1.UpdateStudentRequest{
				AccessToken: res.AccessToken,
				Student: &pbv1.Student{
					Name:        "Mock Update Student",
					Description: "I am a mock student",
					Faculty:     "Mock Engineering",
					Major:       "Mock Computer Engineering",
					Year:        3,
				},
			},
			expect: &pbv1.UpdateCompanyResponse{
				Status:  200,
				Message: "Update data for Mock Update Student successfully!",
			},
		},
		"invalid token": {
			req: &pbv1.UpdateStudentRequest{
				AccessToken: access_token_wrong,
				Student: &pbv1.Student{
					Name:        "Mock Update Student",
					Description: "I am a mock student",
					Faculty:     "Mock Engineering",
					Major:       "Mock Computer Engineering",
					Year:        3,
				},
			},
			expect: &pbv1.UpdateCompanyResponse{
				Status:  404,
				Message: "user id not found",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := u.UpdateStudent(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
		})
	}
}
