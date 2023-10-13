package repo

import (
	"context"
	"database/sql"
	"time"

	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
	"github.com/JinnnDamanee/review-service/port"
	_ "github.com/lib/pq"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type reportRepository struct {
	db DBTX
}

func NewReportRepository(db DBTX) port.ReportRepoPort {
	return &reportRepository{db: db}
}

func (r *reportRepository) CreateReport(ctx context.Context, userId int64, report *pbv1.Report) (int64, error) {
	current_timestamp := time.Now().Unix()

	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO reports (uid, topic, type, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRowContext(ctx, userId, report.Topic, report.Type, report.Description, current_timestamp, current_timestamp).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *reportRepository) GetReport(ctx context.Context, reportId int64) (*pbv1.Report, error) {
	query := "SELECT topic, type, description, updated_at FROM reports WHERE id = $1"
	var topic, reportType, description string
	var updatedAt int64
	err := r.db.QueryRowContext(ctx, query, reportId).Scan(&topic, &reportType, &description, &updatedAt)
	if err != nil {
		return nil, err
	}

	return &pbv1.Report{
		Topic:       topic,
		Type:        reportType,
		Description: description,
		UpdatedAt:   updatedAt,
	}, nil
}

func (r *reportRepository) GetReports(ctx context.Context) ([]*pbv1.Report, error) {
	query := "SELECT topic, type, description, updated_at FROM reports ORDER BY updated_at DESC"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var reports []*pbv1.Report
	for rows.Next() {
		var topic, reportType, description string
		var updatedAt int64
		err = rows.Scan(&topic, &reportType, &description, &updatedAt)
		if err != nil {
			return nil, err
		}

		reports = append(reports, &pbv1.Report{
			Topic:       topic,
			Type:        reportType,
			Description: description,
			UpdatedAt:   updatedAt,
		})
	}

	return reports, nil
}
