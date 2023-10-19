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

func TestHealthCheck(t *testing.T) {
	config, _ := config.LoadConfig("..")
	target := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewReviewServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := c.ReviewHealthCheck(ctx, &pbv1.ReviewHealthCheckRequest{})
	require.NoError(t, err)
	require.Equal(t, int64(200), res.Status)
}

func TestCreateReview(t *testing.T) {
	conf, _ := config.LoadConfig("..")
	configTest, _ := config.LoadConfigTest("..")
	target := fmt.Sprintf("%s:%s", conf.ServerHost, conf.ServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewReviewServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokenCompany, err := mock.GenerateAccessToken(configTest.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 1,
		Role:   "company",
	})
	require.NoError(t, err)

	tokenStudent, err := mock.GenerateAccessToken(configTest.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "student",
	})
	require.NoError(t, err)

	companyId, err := tools.CreateMockCompany("Test Company", utils.GenerateRandomString(5)+"@gmail.com", "123456", "Test Description", "Test Location", "0123456789", "Test Category")
	require.NoError(t, err)

	// Clear all reviews
	err = tools.DeleteAllRecords()
	require.NoError(t, err)

	lex := `{
		"root": {
		}
	}`
	r := &pbv1.CreatedReview{
		Cid:         companyId,
		Title:       "Created Review Test Title 1",
		Description: lex,
		Rating:      5,
		IsAnonymous: true,
	}

	tests := map[string]struct {
		req    *pbv1.CreateReviewRequest
		expect *pbv1.CreateReviewResponse
	}{
		"Successfull": {
			req: &pbv1.CreateReviewRequest{
				AccessToken: tokenStudent,
				Review:      r,
			},
			expect: &pbv1.CreateReviewResponse{
				Status:  201,
				Message: "Review created successfully",
			},
		},
		"Invalid access token": {
			req: &pbv1.CreateReviewRequest{
				AccessToken: "invalid access token",
				Review: &pbv1.CreatedReview{
					Cid:         companyId,
					Title:       "Created Review Test Title 2",
					Description: lex,
					Rating:      5,
					IsAnonymous: false,
				},
			},
			expect: &pbv1.CreateReviewResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
		"Create review with invalid company id": {
			req: &pbv1.CreateReviewRequest{
				AccessToken: tokenStudent,
				Review: &pbv1.CreatedReview{
					Cid:         0,
					Title:       "Created Review Test Title 3-5",
					Description: lex,
					Rating:      5,
					IsAnonymous: false,
				},
			},
			expect: &pbv1.CreateReviewResponse{
				Status:  400,
				Message: "Company not found",
			},
		},
		"Not allowed to create review": {
			req: &pbv1.CreateReviewRequest{
				AccessToken: tokenCompany,
				Review: &pbv1.CreatedReview{
					Cid:         companyId,
					Title:       "Created Review Test Title 3",
					Description: lex,
					Rating:      5,
					IsAnonymous: false,
				},
			},
			expect: &pbv1.CreateReviewResponse{
				Status:  403,
				Message: "You are not allowed to create review",
			},
		},
		"Rating must be between 1 and 5": {
			req: &pbv1.CreateReviewRequest{
				AccessToken: tokenStudent,
				Review: &pbv1.CreatedReview{
					Cid:         companyId,
					Title:       "Created Review Test Title 4",
					Description: lex,
					Rating:      6,
					IsAnonymous: false,
				},
			},
			expect: &pbv1.CreateReviewResponse{
				Status:  400,
				Message: "Rating must be between 1 and 5",
			},
		},
		"Rating can not be less than 1": {
			req: &pbv1.CreateReviewRequest{
				AccessToken: tokenStudent,
				Review: &pbv1.CreatedReview{
					Cid:         companyId,
					Title:       "Created Review Test Title 5",
					Description: lex,
					Rating:      -1,
					IsAnonymous: false,
				},
			},
			expect: &pbv1.CreateReviewResponse{
				Status:  400,
				Message: "Rating must be between 1 and 5",
			},
		},
		"Title are required": {
			req: &pbv1.CreateReviewRequest{
				AccessToken: tokenStudent,
				Review: &pbv1.CreatedReview{
					Cid:         companyId,
					Title:       "",
					Description: lex,
					Rating:      5,
					IsAnonymous: false,
				},
			},
			expect: &pbv1.CreateReviewResponse{
				Status:  400,
				Message: "Please fill in all required fields",
			},
		},
		"Description are required": {
			req: &pbv1.CreateReviewRequest{
				AccessToken: tokenStudent,
				Review: &pbv1.CreatedReview{
					Cid:         companyId,
					Title:       "Created Review Test Title 8",
					Description: "",
					Rating:      5,
					IsAnonymous: false,
				},
			},
			expect: &pbv1.CreateReviewResponse{
				Status:  400,
				Message: "Please fill in all required fields",
			},
		},
		"Rating are required": {
			req: &pbv1.CreateReviewRequest{
				AccessToken: tokenStudent,
				Review: &pbv1.CreatedReview{
					Cid:         companyId,
					Title:       "Created Review Test Title 9",
					Description: lex,
					Rating:      0,
					IsAnonymous: false,
				},
			},
			expect: &pbv1.CreateReviewResponse{
				Status:  400,
				Message: "Please fill in all required fields",
			},
		},
	}

	var id int64
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.CreateReview(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			if res.Status == 201 {
				id = res.Id
				require.Greater(t, res.Id, int64(0))
			}
		})
	}

	// Check if review created
	res, err := c.GetReview(ctx, &pbv1.GetReviewRequest{
		AccessToken: tokenStudent,
		Id:          id,
	})
	require.NoError(t, err)
	require.Equal(t, int64(200), res.Status)
	require.Equal(t, id, res.Review.Id)
	require.Equal(t, r.Title, res.Review.Title)
	require.Equal(t, r.Description, res.Review.Description)
	require.Equal(t, r.Rating, res.Review.Rating)
	require.Equal(t, companyId, res.Review.Company.Id)
}
