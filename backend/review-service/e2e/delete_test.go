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

func TestDeleteReview(t *testing.T) {
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

	tokenStudent, err := mock.GenerateAccessToken(conf_test.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 1,
		Role:   "student",
	})
	require.NoError(t, err)

	tokenCompany, err := mock.GenerateAccessToken(conf_test.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "company",
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
	successReview := &pbv1.CreatedReview{
		Cid:         companyId,
		Title:       "Created Review Test Title 1",
		Description: lex,
		Rating:      5,
		IsAnonymous: true,
	}
	res, err := c.CreateReview(ctx, &pbv1.CreateReviewRequest{
		AccessToken: tokenStudent,
		Review:      successReview,
	})
	require.NoError(t, err)
	require.Equal(t, int64(201), res.Status)
	require.Greater(t, res.Id, int64(0))

	// Delete Review
	tests := map[string]struct {
		req    *pbv1.DeleteReviewRequest
		expect *pbv1.DeleteReviewResponse
	}{
		"success": {
			req: &pbv1.DeleteReviewRequest{
				AccessToken: tokenStudent,
				Id:          companyId,
			},
			expect: &pbv1.DeleteReviewResponse{
				Status:  200,
				Message: "Delete review successfully",
			},
		},
		"Invalid access token": {
			req: &pbv1.DeleteReviewRequest{
				AccessToken: "invalid access token",
				Id:          companyId,
			},
			expect: &pbv1.DeleteReviewResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
		"Invalid review ID": {
			req: &pbv1.DeleteReviewRequest{
				AccessToken: tokenStudent,
				Id:          0,
			},
			expect: &pbv1.DeleteReviewResponse{
				Status:  403,
				Message: "You are not allowed to delete review",
			},
		},
		"Not allowed to delete review": {
			req: &pbv1.DeleteReviewRequest{
				AccessToken: tokenCompany,
				Id:          companyId,
			},
			expect: &pbv1.DeleteReviewResponse{
				Status:  403,
				Message: "You are not allowed to delete review",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.DeleteReview(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
		})
	}
}
