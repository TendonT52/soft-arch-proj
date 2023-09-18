package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/TikhampornSky/go-post-service/config"
	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDeletePost(t *testing.T) {
	conn, err := grpc.Dial(":8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, _ := config.LoadConfig("..")
	token, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 1,
		Role:   "company",
	})
	require.NoError(t, err)

	tokenNotOwner, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "company",
	})
	require.NoError(t, err)

	res, err := c.DeletePosts(ctx, &pbv1.DeletePostsRequest{
		AccessToken: token,
	})
	require.NoError(t, err)
	require.Equal(t, int64(200), res.Status)

	lex := `{ "root": {} }`

	CreateRes, err := c.CreatePost(ctx, &pbv1.CreatePostRequest{
		AccessToken: token,
		Post: &pbv1.Post{
			Topic:          "Topic for delete",
			Description:    lex,
			Period:         "04/04/2023 - 04/04/2024",
			HowTo:          lex,
			OpenPositions:  []string{"Test"},
			RequiredSkills: []string{"Test"},
			Benefits:       []string{"Test"},
		},
	})
	require.NoError(t, err)
	require.Equal(t, int64(201), CreateRes.Status)

	CreateRes2, err := c.CreatePost(ctx, &pbv1.CreatePostRequest{
		AccessToken: token,
		Post: &pbv1.Post{
			Topic:          "Topic for tetsting",
			Description:    lex,
			Period:         "05/05/2023 - 05/05/2024",
			HowTo:          lex,
			OpenPositions:  []string{"open position 1"},
			RequiredSkills: []string{"skill 1", "skill 2"},
			Benefits:       []string{"benefit 1", "benefit 2", "benefit 3"},
		},
	})
	require.NoError(t, err)
	require.Equal(t, int64(201), CreateRes2.Status)

	CreateRes3, err := c.CreatePost(ctx, &pbv1.CreatePostRequest{
		AccessToken: token,
		Post: &pbv1.Post{
			Topic:          "Topic for delete 2",
			Description:    lex,
			Period:         "05/05/2023 - 05/05/2024",
			HowTo:          lex,
			OpenPositions:  []string{"open position 1"},
			RequiredSkills: []string{"skill 1", "skill 3"},
			Benefits:       []string{"benefit 1", "benefit 222", "benefit 3"},
		},
	})
	require.NoError(t, err)
	require.Equal(t, int64(201), CreateRes3.Status)

	tests := map[string]struct {
		req    *pbv1.DeletePostRequest
		expect *pbv1.DeletePostResponse
	}{
		"success": {
			req: &pbv1.DeletePostRequest{
				AccessToken: token,
				Id:          CreateRes.Id,
			},
			expect: &pbv1.DeletePostResponse{
				Status:  200,
				Message: "Post deleted successfully",
			},
		},
		"Unauthorized": {
			req: &pbv1.DeletePostRequest{
				AccessToken: tokenNotOwner,
				Id:          CreateRes2.Id,
			},
			expect: &pbv1.DeletePostResponse{
				Status:  401,
				Message: "Unauthorized",
			},
		},
		"delete when some value still in use": {
			req: &pbv1.DeletePostRequest{
				AccessToken: token,
				Id:          CreateRes3.Id,
			},
			expect: &pbv1.DeletePostResponse{
				Status:  200,
				Message: "Post deleted successfully",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.DeletePost(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
		})
	}

	g, err := c.GetPost(ctx, &pbv1.GetPostRequest{
		AccessToken: token,
		Id:          CreateRes.Id,
	})
	require.NoError(t, err)
	require.Nil(t, g.Post)
}