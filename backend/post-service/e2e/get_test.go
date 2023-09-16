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

func TestGetPost(t *testing.T) {
	conn, err := grpc.Dial(":8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	lex := `{
		"root": {
		  "children": [
			{
			  "children": [
				{
				  "detail": 0,
				  "format": 0,
				  "mode": "normal",
				  "style": "",
				  "text": "What to expect from here on out",
				  "type": "text",
				  "version": 1
				}
			  ],
			  "direction": "ltr",
			  "format": "start",
			  "indent": 0,
			  "type": "paragraph",
			  "version": 1
			}
		  ],
		  "direction": "ltr",
		  "format": "",
		  "indent": 0,
		  "type": "root",
		  "version": 1
		}
	  }
	`

	topic := "What to expect from here on out"
	description := lex
	period := "1 month"
	howTo := lex
	openPositions := []string{"Software Engineer", "Data Scientist"}
	requiredSkills := []string{"Golang", "Python"}
	benefits := []string{"Free lunch", "Free dinner"}

	config, _ := config.LoadConfig("..")
	token, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 6,
		Role:   "company",
	})
	require.NoError(t, err)

	res, err := c.CreatePost(ctx, &pbv1.CreatePostRequest{
		AccessToken: token,
		Post: &pbv1.Post{
			Topic:          topic,
			Description:    description,
			Period:         period,
			HowTo:          howTo,
			OpenPositions:  openPositions,
			RequiredSkills: requiredSkills,
			Benefits:       benefits,
		},
	})
	require.NoError(t, err)

	tests := map[string]struct {
		req    *pbv1.GetPostRequest
		expect *pbv1.GetPostResponse
	}{
		"success": {
			req: &pbv1.GetPostRequest{
				AccessToken: token,
				Id:          res.Id,
			},
			expect: &pbv1.GetPostResponse{
				Status:  200,
				Message: "Post retrieved successfully",
				Post: &pbv1.Post{
					Topic:          topic,
					Description:    description,
					Period:         period,
					HowTo:          howTo,
					OpenPositions:  openPositions,
					RequiredSkills: requiredSkills,
					Benefits:       benefits,
					Owner: &pbv1.PostOwner{
						Id:   6,
						Name: "Mock Company Name",
					},
				},
			},
		},
		"token wrong": {
			req: &pbv1.GetPostRequest{
				AccessToken: "wrong",
				Id:          res.Id,
			},
			expect: &pbv1.GetPostResponse{
				Status:  500,
				Message: "Internal server error",
			},
		},
		"post not found": {
			req: &pbv1.GetPostRequest{
				AccessToken: token,
				Id:          999,
			},
			expect: &pbv1.GetPostResponse{
				Status:  404,
				Message: "Post not found",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetPost(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			if tc.expect.Post != nil {
				require.Equal(t, tc.expect.Post.Topic, res.Post.Topic)
				require.Equal(t, tc.expect.Post.Description, res.Post.Description)
				require.Equal(t, tc.expect.Post.Period, res.Post.Period)
				require.Equal(t, tc.expect.Post.HowTo, res.Post.HowTo)
				require.Equal(t, tc.expect.Post.OpenPositions, res.Post.OpenPositions)
				require.Equal(t, tc.expect.Post.RequiredSkills, res.Post.RequiredSkills)
				require.Equal(t, tc.expect.Post.Benefits, res.Post.Benefits)
				require.Equal(t, tc.expect.Post.Owner.Name, res.Post.Owner.Name)
			}
		})
	}
}
