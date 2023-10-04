package e2e

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/TikhampornSky/report-service/config"
	"github.com/TikhampornSky/report-service/domain"
	pbv1 "github.com/TikhampornSky/report-service/gen/v1"
	"github.com/TikhampornSky/report-service/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetReport(t *testing.T) {
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

	// Create report
	report := &pbv1.Report{
		Topic:       "test-report",
		Type:        domain.REPORT_TYPE_SCAM_LIST,
		Description: "This post is fake and should be deleted!!!",
	}
	resCreate, err := c.CreateReport(ctx, &pbv1.CreateReportRequest{
		AccessToken: tokenStudent,
		Report:      report,
	})
	require.NoError(t, err)
	require.Equal(t, int64(201), resCreate.Status)

	tests := map[string]struct {
		req    *pbv1.GetReportRequest
		expect *pbv1.GetReportResponse
	}{
		"Admin get report": {
			req: &pbv1.GetReportRequest{
				AccessToken: tokenAdmin,
				Id:          resCreate.Id,
			},
			expect: &pbv1.GetReportResponse{
				Status:  200,
				Message: "Report retrieved successfully",
				Report: report,
			},
		},
		"Student get report": {
			req: &pbv1.GetReportRequest{
				AccessToken: tokenStudent,
				Id:          resCreate.Id,
			},
			expect: &pbv1.GetReportResponse{
				Status:  403,
				Message: "You don't have permission to access this resource",
			},
		},
		"Company get report": {
			req: &pbv1.GetReportRequest{
				AccessToken: tokenCompany,
				Id:          resCreate.Id,
			},
			expect: &pbv1.GetReportResponse{
				Status:  403,
				Message: "You don't have permission to access this resource",
			},
		},
		"Report not found": {
			req: &pbv1.GetReportRequest{
				AccessToken: tokenAdmin,
				Id:          0,
			},
			expect: &pbv1.GetReportResponse{
				Status:  404,
				Message: "Report not found",
			},
		},
		"Invalid token": {
			req: &pbv1.GetReportRequest{
				AccessToken: "",
				Id:          resCreate.Id,
			},
			expect: &pbv1.GetReportResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
	}
		
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetReport(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			if res.Status == 200 {
				require.Equal(t, report.Topic, res.Report.Topic)
				require.Equal(t, report.Type, res.Report.Type)
				require.Equal(t, report.Description, res.Report.Description)
				require.NotEmpty(t, res.Report.UpdatedAt)
			}
		})
	}
}
