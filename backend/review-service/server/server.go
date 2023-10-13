package server

import (
	"log"
	"net"

	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
	"github.com/JinnnDamanee/review-service/port"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	customFunc grpc_recovery.RecoveryHandlerFunc
)

func NewServer(port string, reportService port.ReportServicePort) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("tcp connection failed: %v", err)
	}
	log.Printf("listening at %v", lis.Addr())

	customFunc = func(p interface{}) (err error) {
		log.Println("panic triggered: ", p)
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
	)

	p := NewReportServer(reportService)
	pbv1.RegisterReportServiceServer(s, p)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc server failed: %v", err)
	}
}
