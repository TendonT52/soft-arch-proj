package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/TikhampornSky/report-service/domain"
	pbv1 "github.com/TikhampornSky/report-service/gen/v1"
	"github.com/TikhampornSky/report-service/port"
)

type ReportServer struct {
	ReportService port.ReportServicePort
	pbv1.UnimplementedReportServiceServer
}

func NewReportServer(reportService port.ReportServicePort) *ReportServer {
	return &ReportServer{ReportService: reportService}
}

func (s *ReportServer) CreateReport(ctx context.Context, req *pbv1.CreateReportRequest) (*pbv1.CreateReportResponse, error) {
	reportId, err := s.ReportService.CreateReport(ctx, req.AccessToken, req.Report)
	if errors.Is(err, domain.ErrFieldsAreRequired) {
		log.Println("Create Report: Fields are required")
		return &pbv1.CreateReportResponse{
			Status:  http.StatusBadRequest,
			Message: "Please fill in all required fields",
		}, nil
	}
	if err != nil {
		log.Println("Create Report: ", err)
		return &pbv1.CreateReportResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &pbv1.CreateReportResponse{
		Status:  http.StatusCreated,
		Message: "Report created successfully",
		Id:      reportId,
	}, nil
}

func (s *ReportServer) GetReport(ctx context.Context, req *pbv1.GetReportRequest) (*pbv1.GetReportResponse, error) {
	report, err := s.ReportService.GetReport(ctx, req.AccessToken, req.Id)
	if errors.Is(err, domain.ErrUnauthorized) {
		log.Println("Get Report: Unauthorized")
		return &pbv1.GetReportResponse{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		}, nil
	}
	if errors.Is(err, domain.ErrReportNotFound) {
		log.Println("Get Report: Report not found")
		return &pbv1.GetReportResponse{
			Status:  http.StatusNotFound,
			Message: "Report not found",
		}, nil
	}
	if err != nil {
		log.Println("Get Report: ", err)
		return &pbv1.GetReportResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &pbv1.GetReportResponse{
		Status:  http.StatusOK,
		Message: "Report retrieved successfully",
		Report:  report,
	}, nil
}

func (s *ReportServer) ListReports(ctx context.Context, req *pbv1.ListReportsRequest) (*pbv1.ListReportsResponse, error) {
	reports, err := s.ReportService.GetReports(ctx, req.AccessToken)
	if errors.Is(err, domain.ErrUnauthorized) {
		log.Println("List Reports: Unauthorized")
		return &pbv1.ListReportsResponse{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		}, nil
	}
	if err != nil {
		log.Println("List Reports: ", err)
		return &pbv1.ListReportsResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &pbv1.ListReportsResponse{
		Status:  http.StatusOK,
		Message: "Reports retrieved successfully",
		Reports: reports,
		Total:   int64(len(reports)),
	}, nil
}
