package e2e

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/TikhampornSky/go-post-service/config"
	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/mock"
	"github.com/TikhampornSky/go-post-service/tools"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetPost(t *testing.T) {
	config, _ := config.LoadConfig("..")
	target := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	howTo := "Ypu can apply via our facebook page"
	openPositions := []string{"Software Engineer", "Data Scientist"}
	requiredSkills := []string{"Golang", "Python"}
	benefits := []string{"Free lunch", "Free dinner"}

	a, err := mock.CreateMockAdmin(ctx)
	require.NoError(t, err)
	companyName := "Mock Company Name"
	userRes, _, err := mock.CreateMockApprovedCompany(ctx, companyName, a)
	require.NoError(t, err)

	token, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: userRes.Id,
		Role:   "company",
	})
	require.NoError(t, err)

	err = tools.DeleteAllPosts()
	require.NoError(t, err)

	res, err := c.CreatePost(ctx, &pbv1.CreatePostRequest{
		AccessToken: token,
		Post: &pbv1.CreatedPost{
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
						Id:   userRes.Id,
						Name: companyName,
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
				Status:  401,
				Message: "Your access token is invalid",
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
				require.NotEmpty(t, res.Post.Owner.Id)
				require.NotEmpty(t, res.Post.UpdatedAt)
			}
		})
	}
}
