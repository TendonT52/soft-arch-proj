package port

import (
	"context"

	pbv1 "github.com/TikhampornSky/report-service/gen/v1"
)

type ReportServicePort interface {
	CreateReport(ctx context.Context, token string, report *pbv1.Report) (int64, error)
	GetReport(ctx context.Context, token string, reportId int64) (*pbv1.Report, error)
	GetReports(ctx context.Context, token string) ([]*pbv1.Report, error)
}
