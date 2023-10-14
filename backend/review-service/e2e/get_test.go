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

func TestGetReviewByID(t *testing.T) {
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

	companyId, err := tools.CreateMockCompany("Test Company for Get Review", utils.GenerateRandomString(5)+"@gmail.com", "123456", "Test Description", "Test Location", "0123456789", "Test Category")
	require.NoError(t, err)

	studentId, err := tools.CreateMockStudent("Test Student for Get Review", utils.GenerateRandomNumber(10)+"@student.chula.ac.th", "123456", "Test Description", "Engineering", "Computer", 3)
	require.NoError(t, err)

	tokenCompany, err := mock.GenerateAccessToken(configTest.AccessTokenExpiredInTest, &domain.Payload{
		UserId: companyId,
		Role:   "company",
	})
	require.NoError(t, err)

	tokenStudent, err := mock.GenerateAccessToken(configTest.AccessTokenExpiredInTest, &domain.Payload{
		UserId: studentId,
		Role:   "student",
	})
	require.NoError(t, err)

	// Create review
	lex := `{
		"root": {
		}
	}`

	r_anonymous := &pbv1.CreatedReview{
		Cid:         companyId,
		Title:       "Test Title for Get Review",
		Description: lex,
		Rating:      3,
		IsAnonymous: true,
	}
	reqCreateReview := &pbv1.CreateReviewRequest{
		AccessToken: tokenStudent,
		Review:      r_anonymous,
	}
	resCreateReviewAno, err := c.CreateReview(ctx, reqCreateReview)
	require.NoError(t, err)
	require.Equal(t, int64(201), resCreateReviewAno.Status)

	r_unanonymous := &pbv1.CreatedReview{
		Cid:         companyId,
		Title:       "Test Title for Get Review",
		Description: lex,
		Rating:      4,
		IsAnonymous: false,
	}
	reqCreateReview = &pbv1.CreateReviewRequest{
		AccessToken: tokenStudent,
		Review:      r_unanonymous,
	}
	resCreateReviewUnAno, err := c.CreateReview(ctx, reqCreateReview)
	require.NoError(t, err)
	require.Equal(t, int64(201), resCreateReviewUnAno.Status)

	tests := map[string]struct {
		req    *pbv1.GetReviewRequest
		expect *pbv1.GetReviewResponse
	}{
		"Successfull with Annonymous": {
			req: &pbv1.GetReviewRequest{
				AccessToken: tokenCompany,
				Id:          resCreateReviewAno.Id,
			},
			expect: &pbv1.GetReviewResponse{
				Status:  200,
				Message: "Get review successfully",
				Review: &pbv1.Review{
					Id:          resCreateReviewAno.Id,
					Title:       r_anonymous.Title,
					Description: r_anonymous.Description,
					Rating:      r_anonymous.Rating,
					Owner: &pbv1.Owner{
						Id:   0,
						Name: "Anonymous",
					},
					Company: &pbv1.ReviewdCompany{
						Id:   companyId,
						Name: "Test Company for Get Review",
					},
				},
			},
		},
		"Successfull with UnAnnonymous": {
			req: &pbv1.GetReviewRequest{
				AccessToken: tokenCompany,
				Id:          resCreateReviewUnAno.Id,
			},
			expect: &pbv1.GetReviewResponse{
				Status:  200,
				Message: "Get review successfully",
				Review: &pbv1.Review{
					Id:          resCreateReviewUnAno.Id,
					Title:       r_unanonymous.Title,
					Description: r_unanonymous.Description,
					Rating:      r_unanonymous.Rating,
					Owner: &pbv1.Owner{
						Id:   studentId,
						Name: "Test Student for Get Review",
					},
					Company: &pbv1.ReviewdCompany{
						Id:   companyId,
						Name: "Test Company for Get Review",
					},
				},
			},
		},
		"Invalid access token": {
			req: &pbv1.GetReviewRequest{
				AccessToken: "invalid access token",
				Id:          resCreateReviewAno.Id,
			},
			expect: &pbv1.GetReviewResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
		"Get review with invalid review id": {
			req: &pbv1.GetReviewRequest{
				AccessToken: tokenCompany,
				Id:          0,
			},
			expect: &pbv1.GetReviewResponse{
				Status:  404,
				Message: "Review not found",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetReview(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			if res.Status == 200 {
				require.Equal(t, tc.expect.Review.Id, res.Review.Id)
				require.Equal(t, tc.expect.Review.Title, res.Review.Title)
				require.Equal(t, tc.expect.Review.Description, res.Review.Description)
				require.Equal(t, tc.expect.Review.Rating, res.Review.Rating)
				require.NotEmpty(t, res.Review.UpdatedAt)
				require.Equal(t, tc.expect.Review.Owner.Id, res.Review.Owner.Id)
				require.Equal(t, tc.expect.Review.Owner.Name, res.Review.Owner.Name)
				require.Equal(t, tc.expect.Review.Company.Id, res.Review.Company.Id)
				require.Equal(t, tc.expect.Review.Company.Name, res.Review.Company.Name)
			}
		})
	}
}
