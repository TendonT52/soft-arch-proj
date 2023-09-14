package server

import (
	"log"
	"net"

	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/port"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	customFunc grpc_recovery.RecoveryHandlerFunc
)

func NewServer(port string, authService port.AuthServicePort, userService port.UserServicePort) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("tcp connection failed: %v", err)
	}
	log.Printf("listening at %v", lis.Addr())

	// Define customfunc to handle panic
	customFunc = func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
	)

	a := NewAuthServer(authService)
	pbv1.RegisterAuthServiceServer(s, a)
	u := NewUserServer(userService)
	pbv1.RegisterUserServiceServer(s, u)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc server failed: %v", err)
	}
}
