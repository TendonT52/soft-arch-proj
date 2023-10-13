package port

import (
	"context"

	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
)

type ReportRepoPort interface {
	CreateReport(ctx context.Context, userId int64, report *pbv1.Report) (int64, error)
	GetReport(ctx context.Context, reportId int64) (*pbv1.Report, error)
	GetReports(ctx context.Context) ([]*pbv1.Report, error)
}
