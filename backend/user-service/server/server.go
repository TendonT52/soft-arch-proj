package server

import (
	"log"
	"net"

	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/port"
	"google.golang.org/grpc"
)

func NewServer(port string, authService port.AuthServicePort, userService port.UserServicePort) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("tcp connection failed: %v", err)
	}
	log.Printf("listening at %v", lis.Addr())

	s := grpc.NewServer()
	a := NewAuthServer(authService)
	pbv1.RegisterAuthServiceServer(s, a)
	u := NewUserServer(userService)
	pbv1.RegisterUserServiceServer(s, u)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc server failed: %v", err)
	}
}
