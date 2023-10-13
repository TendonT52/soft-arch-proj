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
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createMockReport(t *testing.T, c pbv1.ReportServiceClient, ctx context.Context, token, topic, reporttype, description string) *pbv1.Report {
	report := &pbv1.Report{
		Topic:       topic,
		Type:        reporttype,
		Description: description,
	}
	resCreate, err := c.CreateReport(ctx, &pbv1.CreateReportRequest{
		AccessToken: token,
		Report:      report,
	})
	require.NoError(t, err)
	require.Equal(t, int64(201), resCreate.Status)

	return report
}

func TestListReports(t *testing.T) {
	config, _ := config.LoadConfig("..")
	target := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewReportServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokenAdmin, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 1,
		Role:   "admin",
	})
	require.NoError(t, err)

	tokenStudent, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "student",
	})
	require.NoError(t, err)

	tokenCompany, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 3,
		Role:   "company",
	})
	require.NoError(t, err)

	// Delete all records
	err = tools.DeleteAllRecords()
	require.NoError(t, err)

	// Create reports
	report1 := createMockReport(t, c, ctx, tokenStudent, "test-report-1", domain.REPORT_TYPE_FAKE_REVIEW, "This review is fake and should be deleted!!!")
	report2 := createMockReport(t, c, ctx, tokenCompany, "test-report-2", domain.REPORT_TYPE_OTHER, "Something is wrong to this post")
	report3 := createMockReport(t, c, ctx, tokenStudent, "test-report-3", domain.REPORT_TYPE_SCAM_LIST, "This post is fake and should be deleted!!!")
	report4 := createMockReport(t, c, ctx, tokenCompany, "test-report-4", domain.REPORT_TYPE_SUGGESTION, "Your website shoud have more feature")
	report5 := createMockReport(t, c, ctx, tokenStudent, "test-report-5", domain.REPORT_TYPE_SUSPICIOUS_USER, "User name CYZ is suspicious and should be deleted")
	report6 := createMockReport(t, c, ctx, tokenAdmin, "test-report-6", domain.REPORT_TYPE_WEBSITE_BUGS, "I found a bug in your website")

	tests := map[string]struct {
		req    *pbv1.ListReportsRequest
		expect *pbv1.ListReportsResponse
	}{
		"Admin list reports": {
			req: &pbv1.ListReportsRequest{
				AccessToken: tokenAdmin,
			},
			expect: &pbv1.ListReportsResponse{
				Status:  200,
				Message: "Reports retrieved successfully",
				Reports: []*pbv1.Report{
					report1,
					report2,
					report3,
					report4,
					report5,
					report6,
				},
			},
		},
		"Student list reports": {
			req: &pbv1.ListReportsRequest{
				AccessToken: tokenStudent,
			},
			expect: &pbv1.ListReportsResponse{
				Status:  403,
				Message: "You don't have permission to access this resource",
			},
		},
		"Company list reports": {
			req: &pbv1.ListReportsRequest{
				AccessToken: tokenCompany,
			},
			expect: &pbv1.ListReportsResponse{
				Status:  403,
				Message: "You don't have permission to access this resource",
			},
		},
		"invalid token": {
			req: &pbv1.ListReportsRequest{
				AccessToken: "invalid-token",
			},
			expect: &pbv1.ListReportsResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.ListReports(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			require.Equal(t, len(tc.expect.Reports), len(res.Reports))
			if res.Status == 200 {
				for i, report := range res.Reports {
					require.Equal(t, tc.expect.Reports[i].Topic, report.Topic)
					require.Equal(t, tc.expect.Reports[i].Type, report.Type)
					require.Equal(t, tc.expect.Reports[i].Description, report.Description)
				}
			}
		})
	}
}
