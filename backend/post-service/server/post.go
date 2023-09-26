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

func (s *PostServer) PostHealthCheck(ctx context.Context, req *pbv1.PostHealthCheckRequest) (*pbv1.PostHealthCheckResponse, error) {
	log.Println("Post HealthCheck success: ", http.StatusOK)
	return &pbv1.PostHealthCheckResponse{
		Status: http.StatusOK,
	}, nil
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

func (s *PostServer) GetPost(ctx context.Context, req *pbv1.GetPostRequest) (*pbv1.GetPostResponse, error) {
	post, err := s.PostService.GetPost(ctx, req.AccessToken, req.Id)
	if errors.Is(err, domain.ErrUnauthorized) {
		log.Println("Get Post: Unauthorized")
		return &pbv1.GetPostResponse{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		}, nil
	}
	if errors.Is(err, domain.ErrPostNotFound) {
		log.Println("Get Post: Post not found")
		return &pbv1.GetPostResponse{
			Status:  http.StatusNotFound,
			Message: "Post not found",
		}, nil
	}
	if err != nil {
		log.Println("Get Post: ", err)
		return &pbv1.GetPostResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}
	return &pbv1.GetPostResponse{
		Status:  http.StatusOK,
		Message: "Post retrieved successfully",
		Post:    post,
	}, nil
}

func (s *PostServer) ListPosts(ctx context.Context, req *pbv1.ListPostsRequest) (*pbv1.ListPostsResponse, error) {
	posts, err := s.PostService.GetPosts(ctx, req.AccessToken, req.SearchOptions)
	if errors.Is(err, domain.ErrPostNotFound) {
		log.Println("List Posts: Posts not found")
		return &pbv1.ListPostsResponse{
			Status:  http.StatusNotFound,
			Message: "Posts not found",
		}, nil
	}
	if err != nil {
		log.Println("List Posts: ", err)
		return &pbv1.ListPostsResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}
	return &pbv1.ListPostsResponse{
		Status:  http.StatusOK,
		Message: "Posts retrieved successfully",
		Posts:   posts,
		Total:   int64(len(posts)),
	}, nil
}

func (s *PostServer) UpdatePost(ctx context.Context, req *pbv1.UpdatePostRequest) (*pbv1.UpdatePostResponse, error) {
	err := s.PostService.UpdatePost(ctx, req.AccessToken, req.Id, req.Post)
	if errors.Is(err, domain.ErrFieldsAreRequired) {
		log.Println("Update Post: Please fill all required fields")
		return &pbv1.UpdatePostResponse{
			Status:  http.StatusBadRequest,
			Message: "Please fill all required fields",
		}, nil
	}
	if errors.Is(err, domain.ErrUnauthorized) {
		log.Println("Update Post: Unauthorized")
		return &pbv1.UpdatePostResponse{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		}, nil
	}
	if err != nil {
		log.Println("Update Post: ", err)
		return &pbv1.UpdatePostResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &pbv1.UpdatePostResponse{
		Status:  http.StatusOK,
		Message: "Post updated successfully",
	}, nil
}

func (s *PostServer) DeletePost(ctx context.Context, req *pbv1.DeletePostRequest) (*pbv1.DeletePostResponse, error) {
	err := s.PostService.DeletePost(ctx, req.AccessToken, req.Id)
	if errors.Is(err, domain.ErrUnauthorized) {
		log.Println("Delete Post: Unauthorized")
		return &pbv1.DeletePostResponse{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		}, nil
	}
	if err != nil {
		log.Println("Delete Post: ", err)
		return &pbv1.DeletePostResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}
	return &pbv1.DeletePostResponse{
		Status:  http.StatusOK,
		Message: "Post deleted successfully",
	}, nil
}

func (s *PostServer) GetOpenPositions(ctx context.Context, req *pbv1.GetOpenPositionsRequest) (*pbv1.GetOpenPositionsResponse, error) {
	openPositions, err := s.PostService.GetOpenPositions(ctx, req.AccessToken, req.Search)
	if errors.Is(err, domain.ErrUnauthorized) {
		log.Println("Get Open Positions: Unauthorized")
		return &pbv1.GetOpenPositionsResponse{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		}, nil
	}
	if err != nil {
		log.Println("Get Open Positions: ", err)
		return &pbv1.GetOpenPositionsResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &pbv1.GetOpenPositionsResponse{
		Status:        http.StatusOK,
		Message:       "Open positions retrieved successfully",
		OpenPositions: openPositions,
	}, nil
}

func (s *PostServer) GetRequiredSkills(ctx context.Context, req *pbv1.GetRequiredSkillsRequest) (*pbv1.GetRequiredSkillsResponse, error) {
	requiredSkills, err := s.PostService.GetRequiredSkills(ctx, req.AccessToken, req.Search)
	if errors.Is(err, domain.ErrUnauthorized) {
		log.Println("Get Required Skills: Unauthorized")
		return &pbv1.GetRequiredSkillsResponse{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		}, nil
	}
	if err != nil {
		log.Println("Get Required Skills: ", err)
		return &pbv1.GetRequiredSkillsResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &pbv1.GetRequiredSkillsResponse{
		Status:         http.StatusOK,
		Message:        "Required skills retrieved successfully",
		RequiredSkills: requiredSkills,
	}, nil
}

func (s *PostServer) GetBenefits(ctx context.Context, req *pbv1.GetBenefitsRequest) (*pbv1.GetBenefitsResponse, error) {
	benefits, err := s.PostService.GetBenefits(ctx, req.AccessToken, req.Search)
	if errors.Is(err, domain.ErrUnauthorized) {
		log.Println("Get Benefits: Unauthorized")
		return &pbv1.GetBenefitsResponse{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		}, nil
	}
	if err != nil {
		log.Println("Get Benefits: ", err)
		return &pbv1.GetBenefitsResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &pbv1.GetBenefitsResponse{
		Status:   http.StatusOK,
		Message:  "Benefits retrieved successfully",
		Benefits: benefits,
	}, nil
}
