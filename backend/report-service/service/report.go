package service

import (
	"context"

	"github.com/TikhampornSky/report-service/domain"
	pbv1 "github.com/TikhampornSky/report-service/gen/v1"
	"github.com/TikhampornSky/report-service/port"
	"github.com/TikhampornSky/report-service/utils"
)

const (
	AdminRole = "admin"
)

type reportService struct {
	repo port.ReportRepoPort
}

func NewReportService(repo port.ReportRepoPort) port.ReportServicePort {
	return &reportService{repo: repo}
}

func (s *reportService) CreateReport(ctx context.Context, token string, report *pbv1.Report) (int64, error) {
	if !domain.CheckRequireFields(report) {
		return 0, domain.ErrFieldsAreRequired
	}

	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return 0, err
	}

	reportId, err := s.repo.CreateReport(ctx, payload.UserId, report)
	if err != nil {
		return 0, err
	}

	return reportId, nil
}

func (s *reportService) GetReport(ctx context.Context, token string, reportId int64) (*pbv1.Report, error) {
	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return nil, err
	}

	if payload.Role != AdminRole {
		return nil, domain.ErrUnauthorized
	}

	report, err := s.repo.GetReport(ctx, reportId)
	if report == nil {
		return nil, domain.ErrReportNotFound
	}
	if err != nil {
		return nil, err
	}

	return report, nil
}

func (s *reportService) GetReports(ctx context.Context, token string) ([]*pbv1.Report, error) {
	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return nil, err
	}

	if payload.Role != AdminRole {
		return nil, domain.ErrUnauthorized
	}

	reports, err := s.repo.GetReports(ctx)
	if err != nil {
		return nil, err
	}

	return reports, nil
}
