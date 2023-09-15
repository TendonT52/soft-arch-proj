package service

import (
	"context"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/port"
)

type postService struct {
	PostRepo port.PostRepoPort
}

func NewPostService(postRepo port.PostRepoPort) port.PostServicePort {
	return &postService{PostRepo: postRepo}
}

func (*postService) CreatePost(ctx context.Context, userId int64, post *pbv1.Post) (int64, error) {
	panic("unimplemented")
}

func (*postService) GetPost(ctx context.Context, userId int64, postId int64) (*pbv1.Post, error) {
	return nil, nil
}

func (*postService) GetPosts(ctx context.Context, userId int64, search string) ([]*pbv1.Post, error) {
	return nil, nil
}

func (*postService) UpdatePost(ctx context.Context, userId int64, postId int64, post *pbv1.Post) error {
	return nil
}

func (*postService) DeletePost(ctx context.Context, userId int64, postId int64) error {
	return nil
}