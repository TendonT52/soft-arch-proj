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

func createMockReviews(c pbv1.ReviewServiceClient, cid, uid int64, ownerName, title, description string, rating int32, isAnnonymous bool) (*pbv1.ReviewCompany, error) {
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

	r := &pbv1.ReviewCompany{
		Id:          res.Id,
		Title:       title,
		Description: description,
		Rating:      rating,
		Owner:       &pbv1.Owner{},
	}

	if isAnnonymous {
		r.Owner.Id = 0
		r.Owner.Name = "Anonymous"
	} else {
		r.Owner.Id = uid
		r.Owner.Name = ownerName
	}

	return r, err
}

func TestListReviewByCompany(t *testing.T) {
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

	// Clear all reviews
	err = tools.DeleteAllRecords()
	require.NoError(t, err)

	companyId1, err := tools.CreateMockCompany("Test Company for List Review By Company - 1", utils.GenerateRandomString(5)+"@gmail.com", "123456", "Test Description", "Test Location", "0123456789", "Test Category")
	require.NoError(t, err)
	companyId2, err := tools.CreateMockCompany("Test Company for List Review By Company - 2", utils.GenerateRandomString(5)+"@gmail.com", "123456", "Test Description", "Test Location", "0123456789", "Test Category")
	require.NoError(t, err)

	studentId1, err := tools.CreateMockStudent("Test Student for List Review By Company - 1", utils.GenerateRandomNumber(10)+"@student.chula.ac.th", "123456", "Test Description", "Engineering", "Computer", 3)
	require.NoError(t, err)
	studentId2, err := tools.CreateMockStudent("Test Student for List Review By Company - 2", utils.GenerateRandomNumber(10)+"@student.chula.ac.th", "123456", "Test Description", "Engineering", "Computer", 4)
	require.NoError(t, err)

	tokenCompany, err := mock.GenerateAccessToken(configTest.AccessTokenExpiredInTest, &domain.Payload{
		UserId: companyId1,
		Role:   "company",
	})
	require.NoError(t, err)

	// Create review
	lex := `{
		"root": {
		}
	}`
	r1, err := createMockReviews(c, companyId1, studentId1, "Test Student for List Review By Company - 1", "Test Title for List Review By Company - 1", lex, 3, true)
	require.NoError(t, err)
	r2, err := createMockReviews(c, companyId1, studentId2, "Test Student for List Review By Company - 2", "Test Title for List Review By Company - 2", lex, 4, false)
	require.NoError(t, err)
	r3, err := createMockReviews(c, companyId1, studentId1, "Test Student for List Review By Company - 1", "Test Title for List Review By Company - 3", lex, 5, true)
	require.NoError(t, err)
	r4, err := createMockReviews(c, companyId2, studentId1, "Test Student for List Review By Company - 1", "Test Title for List Review By Company - 4", lex, 2, false)
	require.NoError(t, err)
	r5, err := createMockReviews(c, companyId2, studentId2, "Test Student for List Review By Company - 2", "Test Title for List Review By Company - 5", lex, 4, true)
	require.NoError(t, err)

	_ = r4
	_ = r5

	tests := map[string]struct {
		req    *pbv1.ListReviewsByCompanyRequest
		expect *pbv1.ListReviewsByCompanyResponse
	}{
		"sucessful": {
			req: &pbv1.ListReviewsByCompanyRequest{
				AccessToken: tokenCompany,
				Cid:         companyId1,
			},
			expect: &pbv1.ListReviewsByCompanyResponse{
				Status:  200,
				Message: "List reviews by company successfully",
				Reviews: []*pbv1.ReviewCompany{
					r1, r2, r3,
				},
				Total: 3,
			},
		},
		"invalide access token": {
			req: &pbv1.ListReviewsByCompanyRequest{
				AccessToken: "invalid token",
				Cid:         companyId1,
			},
			expect: &pbv1.ListReviewsByCompanyResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
		"company not found": {
			req: &pbv1.ListReviewsByCompanyRequest{
				AccessToken: tokenCompany,
				Cid:         0,
			},
			expect: &pbv1.ListReviewsByCompanyResponse{
				Status:  200,
				Message: "List reviews by company successfully",
				Reviews: []*pbv1.ReviewCompany{},
				Total:   0,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.ListReviewsByCompany(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			require.Equal(t, int32(len(tc.expect.Reviews)), res.Total)
			for i := 0; i < len(tc.expect.Reviews); i++ {
				require.Equal(t, tc.expect.Reviews[i].Id, res.Reviews[i].Id)
				require.Equal(t, tc.expect.Reviews[i].Title, res.Reviews[i].Title)
				require.Equal(t, tc.expect.Reviews[i].Description, res.Reviews[i].Description)
				require.Equal(t, tc.expect.Reviews[i].Rating, res.Reviews[i].Rating)
				require.Equal(t, tc.expect.Reviews[i].Owner.Id, res.Reviews[i].Owner.Id)
				require.Equal(t, tc.expect.Reviews[i].Owner.Name, res.Reviews[i].Owner.Name)
			}
		})
	}
}
