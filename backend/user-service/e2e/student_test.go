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

func TestGetStudentMe(t *testing.T) {
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

	timeNow, err := tools.GetCreateTime(r.Id)
	require.NoError(t, err)

	// Verify Email
	result, err := c.VerifyEmailCode(ctx, &pbv1.VerifyEmailCodeRequest{
		StudentId: id_student,
		Code:      utils.Encode(id_student, timeNow),
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
	access_token_wrong, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &domain.Payload{
		UserId: 0,
		Role:   domain.StudentRole,
	})
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
		"not correct student": {
			req: &pbv1.GetStudentMeRequest{
				AccessToken: access_token_wrong,
			},
			expect: &pbv1.GetStudentResponse{
				Status:  500,
				Message: "Something went wrong",
			},
		},
		"invalid token": {
			req: &pbv1.GetStudentMeRequest{
				AccessToken: "invalid token",
			},
			expect: &pbv1.GetStudentResponse{
				Status:  401,
				Message: "Your access token is invalid",
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

	timeNow, err := tools.GetCreateTime(r.Id)
	require.NoError(t, err)

	// Verify Email
	result, err := c.VerifyEmailCode(ctx, &pbv1.VerifyEmailCodeRequest{
		StudentId: id_student,
		Code:      utils.Encode(id_student, timeNow),
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
	admin_access_token, err := utils.CreateAccessToken(365*24*time.Hour, &domain.Payload{
		UserId: 0,
		Role:   domain.AdminRole,
	})
	a := &pbv1.CreateAdminRequest{
		Email:           utils.GenerateRandomString(20) + "@gmail.com",
		Password:        "password-test",
		PasswordConfirm: "password-test",
		AccessToken:     admin_access_token,
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
		"invalid token": {
			req: &pbv1.GetStudentRequest{
				AccessToken: "invalid token",
				Id:          r.Id,
			},
			expect: &pbv1.GetStudentResponse{
				Status:  401,
				Message: "Your access token is invalid",
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

	timeNow, err := tools.GetCreateTime(r.Id)
	require.NoError(t, err)

	// Verify Email
	result, err := c.VerifyEmailCode(ctx, &pbv1.VerifyEmailCodeRequest{
		StudentId: id_student,
		Code:      utils.Encode(id_student, timeNow),
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
	access_token_wrong, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &domain.Payload{
		UserId: 0,
		Role:   domain.StudentRole,
	})
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.UpdateStudentRequest
		expect *pbv1.UpdateCompanyResponse
	}{
		"success": {
			req: &pbv1.UpdateStudentRequest{
				AccessToken: res.AccessToken,
				Student: &pbv1.UpdatedStudent{
					Name:        "UPADATED Mock Update Student",
					Description: "UPADATED I am a mock student",
					Faculty:     "UPADATED Mock Engineering",
					Major:       "UPADATED Mock Computer Engineering",
					Year:        3,
				},
			},
			expect: &pbv1.UpdateCompanyResponse{
				Status:  200,
				Message: "Update data for UPADATED Mock Update Student successfully!",
			},
		},
		"Not correct student": {
			req: &pbv1.UpdateStudentRequest{
				AccessToken: access_token_wrong,
				Student: &pbv1.UpdatedStudent{
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
		"invalid token": {
			req: &pbv1.UpdateStudentRequest{
				AccessToken: "invalid token",
				Student:     &pbv1.UpdatedStudent{},
			},
			expect: &pbv1.UpdateCompanyResponse{
				Status:  401,
				Message: "Your access token is invalid",
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

	// Get Student
	resGet, _, err := tools.GetStudentByID(r.Id)
	require.NoError(t, err)
	require.Equal(t, "UPADATED Mock Update Student", resGet.Name)
	require.Equal(t, "UPADATED I am a mock student", resGet.Description)
	require.Equal(t, "UPADATED Mock Engineering", resGet.Faculty)
	require.Equal(t, "UPADATED Mock Computer Engineering", resGet.Major)
	require.Equal(t, int32(3), resGet.Year)
}

func TestGetStudents(t *testing.T) {
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

	_, id1 := createMockStudent(t, "Student - 1", ad.AccessToken)
	_, id2 := createMockStudent(t, "Student - 2", ad.AccessToken)
	_, id3 := createMockStudent(t, "Student - 3", ad.AccessToken)

	tests := map[string]struct {
		req    *pbv1.GetStudentsRequest
		expect *pbv1.GetStudentsResponse
	}{
		"success": {
			req: &pbv1.GetStudentsRequest{
				AccessToken: ad.AccessToken,
				Ids:         []int64{id1, id2, id3},
			},
			expect: &pbv1.GetStudentsResponse{
				Status: 200,
				Students: []*pbv1.StudentInfo{
					{
						Id:   id1,
						Name: "Student - 1",
					},
					{
						Id:   id2,
						Name: "Student - 2",
					},
					{
						Id:   id3,
						Name: "Student - 3",
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := u.GetStudents(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			if tc.expect.Students != nil {
				for i, s := range tc.expect.Students {
					require.Equal(t, s.Id, res.Students[i].Id)
					require.Equal(t, s.Name, res.Students[i].Name)
				}
			}
		})
	}
}
