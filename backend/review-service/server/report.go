package server

import (
	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
	"github.com/JinnnDamanee/review-service/port"
)

type ReportServer struct {
	ReportService port.ReviewServicePort
	pbv1.UnimplementedReportServiceServer
}

func NewReportServer(reportService port.ReviewServicePort) *ReportServer {
	return &ReportServer{ReportService: reportService}
}
