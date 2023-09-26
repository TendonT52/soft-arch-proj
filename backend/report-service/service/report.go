package service

import (
	"context"
	gen "github.com/TikhampornSky/report-service/gen/v1"
	"github.com/TikhampornSky/report-service/port"
)

type reportService struct {
	repo port.ReportRepoPort
}

func NewReportService(repo port.ReportRepoPort) port.ReportServicePort {
	return &reportService{repo: repo}
}

func (*reportService) CreateReport(ctx context.Context, token string, post *gen.Report) (int64, error) {
	panic("unimplemented")
}

func (*reportService) GetReport(ctx context.Context, token string, reportId int64) (*gen.Report, error) {
	panic("unimplemented")
}

func (*reportService) GetReports(ctx context.Context, token string) ([]*gen.Report, error) {
	panic("unimplemented")
}
