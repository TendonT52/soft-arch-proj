package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/port"
)

type PostServer struct {
	PostService port.PostServicePort
	pbv1.UnimplementedPostServiceServer
}

func NewPostServer(postService port.PostServicePort) *PostServer {
	return &PostServer{PostService: postService}
}

func (s *PostServer) CreatePost(ctx context.Context, req *pbv1.CreatePostRequest) (*pbv1.CreatePostResponse, error) {
	postId, err := s.PostService.CreatePost(ctx, req.AccessToken, req.Post)
	if errors.Is(err, domain.ErrFieldsAreRequired) {
		log.Println("Create Post: Please fill all required fields")
		return &pbv1.CreatePostResponse{
			Status:  http.StatusBadRequest,
			Message: "Please fill all required fields",
		}, nil
	}
	if errors.Is(err, domain.ErrUnauthorized) {
		log.Println("Create Post: Unauthorized")
		return &pbv1.CreatePostResponse{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		}, nil
	}
	if err != nil {
		log.Println("Create Post: ", err)
		return &pbv1.CreatePostResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &pbv1.CreatePostResponse{
		Status:  http.StatusCreated,
		Message: "Post created successfully",
		Id:      postId,
	}, nil
}

func (s *PostServer) GetPost(context.Context, *pbv1.GetPostRequest) (*pbv1.GetPostResponse, error) {
	return nil, nil
}

func (s *PostServer) ListPosts(context.Context, *pbv1.ListPostsRequest) (*pbv1.ListPostsResponse, error) {
	return nil, nil
}

func (s *PostServer) UpdatePost(context.Context, *pbv1.UpdatePostRequest) (*pbv1.UpdatePostResponse, error) {
	return nil, nil
}

func (s *PostServer) DeletePost(context.Context, *pbv1.DeletePostRequest) (*pbv1.DeletePostResponse, error) {
	return nil, nil
}
