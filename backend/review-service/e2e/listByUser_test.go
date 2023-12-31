package e2e

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/JinnnDamanee/review-service/config"
	"github.com/JinnnDamanee/review-service/domain"
	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
	"github.com/JinnnDamanee/review-service/mock"
	"github.com/JinnnDamanee/review-service/tools"
	"github.com/JinnnDamanee/review-service/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createMockMyReviews(c pbv1.ReviewServiceClient, cid, uid int64, companyName, ownerName, title, description string, rating int32, isAnnonymous bool) (*pbv1.MyReview, error) {
	configTest, _ := config.LoadConfigTest("..")
	token, err := mock.GenerateAccessToken(configTest.AccessTokenExpiredInTest, &domain.Payload{
		UserId: uid,
		Role:   "student",
	})
	res, err := c.CreateReview(context.Background(), &pbv1.CreateReviewRequest{
		AccessToken: token,
		Review: &pbv1.CreatedReview{
			Cid:         cid,
			Title:       title,
			Description: description,
			Rating:      rating,
			IsAnonymous: isAnnonymous,
		},
	})
	if err != nil {
		return nil, err
	}
	if res.Status != 201 {
		return nil, fmt.Errorf("create review failed with status %d", res.Status)
	}

	r := &pbv1.MyReview{
		Id:          res.Id,
		Title:       title,
		Description: description,
		Rating:      rating,
		Company: &pbv1.ReviewdCompany{
			Id:   cid,
			Name: companyName,
		},
	}

	return r, err
}

func TestListByUser(t *testing.T) {
	conf, _ := config.LoadConfig("..")
	conf_test, _ := config.LoadConfigTest("..")
	target := fmt.Sprintf("%s:%s", conf.ServerHost, conf.ServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewReviewServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokenStudent1, err := mock.GenerateAccessToken(conf_test.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 1,
		Role:   "student",
	})
	require.NoError(t, err)

	tokenStudent2, err := mock.GenerateAccessToken(conf_test.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "student",
	})
	require.NoError(t, err)

	companyId, err := tools.CreateMockCompany("Test Company", utils.GenerateRandomString(5)+"@gmail.com", "123456", "Test Description", "Test Location", "0123456789", "Test Category")
	require.NoError(t, err)

	// Create Review
	err = tools.DeleteAllRecords()
	require.NoError(t, err)

	lex := `{
		"root": {
		}
	}`

	r1, err := createMockMyReviews(c, companyId, 1, "Test Company", "Test Student 1", "Test Title 1-1", lex, 5, true)
	require.NoError(t, err)
	r2, err := createMockMyReviews(c, companyId, 1, "Test Company", "Test Student 1", "Test Title 1-2", lex, 4, false)

	r3, err := createMockMyReviews(c, companyId, 2, "Test Company", "Test Student 2", "Test Title 2-1", lex, 3, true)
	require.NoError(t, err)

	// List Review By User
	tests := map[string]struct {
		req    *pbv1.ListReviewsByUserRequest
		expect *pbv1.ListReviewsByUserResponse
	}{
		"User 1 Success": {
			req: &pbv1.ListReviewsByUserRequest{
				AccessToken: tokenStudent1,
			},
			expect: &pbv1.ListReviewsByUserResponse{
				Status:  200,
				Message: "List reviews by user successfully",
				Reviews: []*pbv1.MyReview{r1, r2},
				Total:   2,
			},
		},
		"User 2 Success": {
			req: &pbv1.ListReviewsByUserRequest{
				AccessToken: tokenStudent2,
			},
			expect: &pbv1.ListReviewsByUserResponse{
				Status:  200,
				Message: "List reviews by user successfully",
				Reviews: []*pbv1.MyReview{r3},
				Total:   1,
			},
		},
		"Invalid Access Token": {
			req: &pbv1.ListReviewsByUserRequest{
				AccessToken: "invalid token",
			},
			expect: &pbv1.ListReviewsByUserResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.ListReviewsByUser(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			require.Equal(t, int32(len(tc.expect.Reviews)), res.Total)

			if tc.expect.Reviews != nil {
				for i := 0; i < len(tc.expect.Reviews); i++ {
					require.Equal(t, tc.expect.Reviews[i].Id, res.Reviews[i].Id)
					require.Equal(t, tc.expect.Reviews[i].Title, res.Reviews[i].Title)
					require.Equal(t, tc.expect.Reviews[i].Description, res.Reviews[i].Description)
					require.Equal(t, tc.expect.Reviews[i].Rating, res.Reviews[i].Rating)
					require.Equal(t, tc.expect.Reviews[i].Company.Id, res.Reviews[i].Company.Id)
					require.Equal(t, tc.expect.Reviews[i].Company.Name, res.Reviews[i].Company.Name)
				}
			}
		})
	}
}
