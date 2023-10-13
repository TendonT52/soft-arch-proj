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

func TestHealthCheck(t *testing.T) {
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

	res, err := c.ReportHealthCheck(ctx, &pbv1.ReportHealthCheckRequest{})
	require.NoError(t, err)
	require.Equal(t, int64(200), res.Status)
}

func TestCreateReport(t *testing.T) {
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

	token, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 1,
		Role:   "student",
	})
	require.NoError(t, err)

	report_scam_list := &pbv1.Report{
		Topic:       "test-report-scam",
		Type:        domain.REPORT_TYPE_SCAM_LIST,
		Description: "This post is fake and should be deleted",
	}
	report_fake_review := &pbv1.Report{
		Topic:       "test-report-fake-review",
		Type:        domain.REPORT_TYPE_FAKE_REVIEW,
		Description: "This review is fake and should be deleted",
	}
	report_suspicious_user := &pbv1.Report{
		Topic:       "test-report-suspicious-user",
		Type:        domain.REPORT_TYPE_SUSPICIOUS_USER,
		Description: "This user is suspicious and should be deleted",
	}
	report_website_bugs := &pbv1.Report{
		Topic:       "test-report-website-bugs",
		Type:        domain.REPORT_TYPE_WEBSITE_BUGS,
		Description: "This website has bugs and should be fixed",
	}
	report_suggestion := &pbv1.Report{
		Topic:       "test-report-suggestion",
		Type:        domain.REPORT_TYPE_SUGGESTION,
		Description: "This is a suggestion",
	}
	report_other := &pbv1.Report{
		Topic:       "test-report-other",
		Type:        domain.REPORT_TYPE_OTHER,
		Description: "This report is fake and should be deleted",
	}

	tests := map[string]struct {
		req    *pbv1.CreateReportRequest
		expect *pbv1.CreateReportResponse
	}{
		"success with REPORT_TYPE_SCAM_LIST": {
			req: &pbv1.CreateReportRequest{
				AccessToken: token,
				Report:      report_scam_list,
			},
			expect: &pbv1.CreateReportResponse{
				Status:  201,
				Message: "Report created successfully",
			},
		},
		"success with REPORT_TYPE_FAKE_REVIEW": {
			req: &pbv1.CreateReportRequest{
				AccessToken: token,
				Report:      report_fake_review,
			},
			expect: &pbv1.CreateReportResponse{
				Status:  201,
				Message: "Report created successfully",
			},
		},
		"success with REPORT_TYPE_SUSPICIOUS_USER": {
			req: &pbv1.CreateReportRequest{
				AccessToken: token,
				Report:      report_suspicious_user,
			},
			expect: &pbv1.CreateReportResponse{
				Status:  201,
				Message: "Report created successfully",
			},
		},
		"success with REPORT_TYPE_WEBSITE_BUGS": {
			req: &pbv1.CreateReportRequest{
				AccessToken: token,
				Report:      report_website_bugs,
			},
			expect: &pbv1.CreateReportResponse{
				Status:  201,
				Message: "Report created successfully",
			},
		},
		"success with REPORT_TYPE_SUGGESTION": {
			req: &pbv1.CreateReportRequest{
				AccessToken: token,
				Report:      report_suggestion,
			},
			expect: &pbv1.CreateReportResponse{
				Status:  201,
				Message: "Report created successfully",
			},
		},
		"success with REPORT_TYPE_OTHER": {
			req: &pbv1.CreateReportRequest{
				AccessToken: token,
				Report:      report_other,
			},
			expect: &pbv1.CreateReportResponse{
				Status:  201,
				Message: "Report created successfully",
			},
		},
		"topic is empty": {
			req: &pbv1.CreateReportRequest{
				AccessToken: token,
				Report: &pbv1.Report{
					Topic:       "",
					Type:        domain.REPORT_TYPE_OTHER,
					Description: "This report is fake and should be deleted",
				},
			},
			expect: &pbv1.CreateReportResponse{
				Status:  400,
				Message: "Please fill in all required fields",
			},
		},
		"type is empty": {
			req: &pbv1.CreateReportRequest{
				AccessToken: token,
				Report: &pbv1.Report{
					Topic:       "test-report-other",
					Type:        "",
					Description: "This report is fake and should be deleted",
				},
			},
			expect: &pbv1.CreateReportResponse{
				Status:  400,
				Message: "Please fill in all required fields",
			},
		},
		"description is empty": {
			req: &pbv1.CreateReportRequest{
				AccessToken: token,
				Report: &pbv1.Report{
					Topic:       "test-report-other",
					Type:        domain.REPORT_TYPE_OTHER,
					Description: "",
				},
			},
			expect: &pbv1.CreateReportResponse{
				Status:  400,
				Message: "Please fill in all required fields",
			},
		},
		"invalid token": {
			req: &pbv1.CreateReportRequest{
				AccessToken: "",
				Report: &pbv1.Report{
					Topic:       "test-report-other",
					Type:        domain.REPORT_TYPE_OTHER,
					Description: "This report is fake and should be deleted",
				},
			},
			expect: &pbv1.CreateReportResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
	}

	ids := make(map[string]int64)
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.CreateReport(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			if res.Status == 201 {
				ids[name] = res.Id
			}
		})
	}

	require.Equal(t, 6, len(ids))
	tokenAdmin, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "admin",
	})
	require.NoError(t, err)
	// Check if all reports are created
	for key, id := range ids {
		res, err := c.GetReport(ctx, &pbv1.GetReportRequest{
			AccessToken: tokenAdmin,
			Id:          id,
		})
		require.NoError(t, err)
		require.Equal(t, int64(200), res.Status)
		r := tests[key].req.Report
		require.Equal(t, r.Topic, res.Report.Topic)
		require.Equal(t, r.Type, res.Report.Type)
		require.Equal(t, r.Description, res.Report.Description)
	}
}
