package httpServer

import (
	"JinnnDamanee/review-service/internal/handler"
	"JinnnDamanee/review-service/internal/service"
	"context"
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
	Handler *handler.ReviewHandler
}

func NewHTTPServer(s *service.ReviewService, h *handler.ReviewHandler) *httpServer {
	return &httpServer{
		Server:  echo.New(),
		Service: s,
		Handler: h,
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
	s.Server.GET("/review", s.Handler.GetAllReview)
	s.Server.GET("/review/:id", s.Handler.GetReviewByID)
	s.Server.POST("/review", s.Handler.CreateReview)
	s.Server.PUT("/review/:id", s.Handler.UpdateReviewByID)
	s.Server.DELETE("/review/:id", s.Handler.DeleteReviewByID)
}
