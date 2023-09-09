package e2e

import (
	"context"
	"fmt"
	"testing"
	"time"

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
	id_student, err := utils.GenerateRandomNumber(10)
	require.NoError(t, err)
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

	res, err := c.SignIn(ctx, &pbv1.LoginRequest{
		Email:    s.Email,
		Password: s.Password,
	})
	require.NoError(t, err)

	fmt.Println("==> ", res)
	fmt.Println("--> ", res.AccessToken)

	tests := map[string]struct {
		req    *pbv1.GetStudentMeRequest
		expect *pbv1.GetStudentResponse
	}{
		"success": {
			req: &pbv1.GetStudentMeRequest{
				AccessToken: res.AccessToken,
			},
			expect: &pbv1.GetStudentResponse{
				Status:  200,
				Message: "success",
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
		// "invalid token": {
		// 	req: &pbv1.GetStudentMeRequest{
		// 		AccessToken: "invalid token",
		// 	},
		// 	expect: &pbv1.GetStudentResponse{
		// 		Status:  400,
		// 		Message: "invalid token",
		// 	},
		// },
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := u.GetStudentMe(ctx, tc.req)
			if err != nil {
				t.Errorf("could not sign in: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
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
