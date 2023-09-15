package httpServer

import (
	"context"
	"jindamanee2544/review-service/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type httpServer struct {
	Server  *echo.Echo
	Service *service.ReviewService
}

func NewHTTPServer(s *service.ReviewService) *httpServer {
	return &httpServer{
		Server:  echo.New(),
		Service: s,
	}
}

func (s *httpServer) Start() {
	log.Printf("server is running on port %s", "8080")
	go func() {
		if err := s.Server.Start(":8080"); err != nil {
			log.Printf("server start failed: %v", err)
		}
	}()
}

func (s *httpServer) SetUpShutdown() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		s.Start()
	}()

	<-sig

	s.Shutdown()
}

func (s *httpServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown failed: %v", err)
	}
	log.Print("Gracefully shutdown the server")
}

func (s *httpServer) InitRouter() {
	s.Server.GET("/review", s.Service.GetAllReviews)
	s.Server.GET("/review/:id", s.Service.GetReviewByID)
	s.Server.POST("/review", s.Service.CreateReview)
	s.Server.PUT("/review/:id", s.Service.UpdateReviewByID)
	s.Server.DELETE("/review/:id", s.Service.DeleteReviewByID)
}
