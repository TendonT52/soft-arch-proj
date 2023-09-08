package e2e_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/db"
	"github.com/TikhampornSky/go-auth-verifiedMail/email"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/initializers"
	"github.com/TikhampornSky/go-auth-verifiedMail/repo"
	"github.com/TikhampornSky/go-auth-verifiedMail/server"
	"github.com/TikhampornSky/go-auth-verifiedMail/service"
)

func TestMain(m *testing.M) {
	config, err := initializers.LoadConfig("..")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	db, err := db.NewDatabase(&config)
	if err != nil {
		log.Fatalf("Something went wrong. Could not connect to the database. %s", err)
	}

	memphisConn := email.InitMemphisConnection()
	defer memphisConn.Close()

	memphis := email.NewMemphis(memphisConn, config.MemphisStationNameTest)

	userRepo := repo.NewUserRepository(db.GetPostgresqlDB(), db.GetRedisDB())
	authService := service.NewAuthService(userRepo, memphis)
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

	id_student, err := utils.GenerateRandomNumber(10)
	require.NoError(t, err)

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

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
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
