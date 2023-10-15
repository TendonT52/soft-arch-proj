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
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestUpdateReview(t *testing.T) {
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

	tokenStudent, err := mock.GenerateAccessToken(configTest.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "student",
	})
	require.NoError(t, err)

	tokenStudent2, err := mock.GenerateAccessToken(configTest.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 22,
		Role:   "student",
	})
	require.NoError(t, err)

	// Create review
	lex := `{
		"root": {
		}
	}`
	updatedlex := `{
		"update_root": {
		}
	}`
	r := &pbv1.CreatedReview{
		Cid:         1,
		Title:       "Test Title for Update Review",
		Description: lex,
		Rating:      3,
		IsAnonymous: false,
	}
	res, err := c.CreateReview(ctx, &pbv1.CreateReviewRequest{
		AccessToken: tokenStudent,
		Review:      r,
	})
	require.NoError(t, err)
	require.Equal(t, int64(201), res.Status)

	u := &pbv1.UpdatedReview{
		Title:       "UPDATED Test Title for Update Review",
		Description: updatedlex,
		Rating:      4,
		IsAnonymous: true,
	}
	
	tests := map[string]struct {
		req *pbv1.UpdateReviewRequest
		res *pbv1.UpdateReviewResponse
	}{
		"Successful": {
			req: &pbv1.UpdateReviewRequest{
				AccessToken: tokenStudent,
				Id:          res.Id,
				Review: u,
			},
			res: &pbv1.UpdateReviewResponse{
				Status: 200,
				Message: "Update review successfully",
			},
		},
		"Invalid access token": {
			req: &pbv1.UpdateReviewRequest{
				AccessToken: "invalid access token",
				Id:          res.Id,
				Review: u,
			},
			res: &pbv1.UpdateReviewResponse{
				Status: 401,
				Message: "Your access token is invalid",
			},
		},
		"Not owner": {
			req: &pbv1.UpdateReviewRequest{
				AccessToken: tokenStudent2,
				Id:          2,
				Review: u,
			},
			res: &pbv1.UpdateReviewResponse{
				Status: 403,
				Message: "You are not allowed to update review",
			},
		},
		"Title is required": {
			req: &pbv1.UpdateReviewRequest{	
				AccessToken: tokenStudent,
				Id:          res.Id,
				Review: &pbv1.UpdatedReview{
					Title:       "",
					Description: updatedlex,
					Rating:      4,
					IsAnonymous: true,
				},
			},
			res: &pbv1.UpdateReviewResponse{
				Status: 400,
				Message: "Please fill in all required fields",
			},
		},
		"Description is required": {
			req: &pbv1.UpdateReviewRequest{
				AccessToken: tokenStudent,
				Id:          res.Id,
				Review: &pbv1.UpdatedReview{
					Title:       "UPDATED Test Title for Update Review",
					Description: "",
					Rating:      4,
					IsAnonymous: true,
				},
			},
			res: &pbv1.UpdateReviewResponse{
				Status: 400,
				Message: "Please fill in all required fields",
			},
		},
		"Rating is required": {
			req: &pbv1.UpdateReviewRequest{
				AccessToken: tokenStudent,
				Id:          res.Id,
				Review: &pbv1.UpdatedReview{
					Title:       "UPDATED Test Title for Update Review",
					Description: updatedlex,
					Rating:      0,
					IsAnonymous: true,
				},
			},
			res: &pbv1.UpdateReviewResponse{
				Status: 400,
				Message: "Please fill in all required fields",
			},
		},
		"Rating must be between 1 and 5": {
			req: &pbv1.UpdateReviewRequest{
				AccessToken: tokenStudent,
				Id:          res.Id,
				Review: &pbv1.UpdatedReview{
					Title:       "UPDATED Test Title for Update Review",
					Description: updatedlex,
					Rating:      6,
					IsAnonymous: true,
				},
			},
			res: &pbv1.UpdateReviewResponse{
				Status: 400,
				Message: "Rating must be between 1 and 5",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.UpdateReview(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.res.Status, res.Status)
			require.Equal(t, tc.res.Message, res.Message)
		})
	}

	// Check updated review
	result, err := c.GetReview(ctx, &pbv1.GetReviewRequest{
		AccessToken: tokenStudent,
		Id:          res.Id,
	})
	require.NoError(t, err)
	require.Equal(t, int64(200), result.Status)
	require.Equal(t, u.Title, result.Review.Title)
	require.Equal(t, u.Description, result.Review.Description)
	require.Equal(t, u.Rating, result.Review.Rating)
	require.Equal(t, int64(0), result.Review.Owner.Id)
	require.Equal(t, "Anonymous", result.Review.Owner.Name)
}
