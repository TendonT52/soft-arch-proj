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
	"github.com/TikhampornSky/go-post-service/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestHealthCheck(t *testing.T) {
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

	res, err := c.PostHealthCheck(ctx, &pbv1.PostHealthCheckRequest{})
	require.NoError(t, err)
	require.Equal(t, int64(200), res.Status)
}

func TestCreatePost(t *testing.T) {
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

	token, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 1,
		Role:   "company",
	})
	require.NoError(t, err)

	tokenStudent, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "student",
	})
	require.NoError(t, err)

	// Clear all posts
	err = tools.DeleteAllPosts()
	require.NoError(t, err)

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

	tests := map[string]struct {
		req    *pbv1.CreatePostRequest
		expect *pbv1.CreatePostResponse
	}{
		"success": {
			req: &pbv1.CreatePostRequest{
				Post: &pbv1.CreatedPost{
					Topic:          "Topic Test",
					Description:    lex,
					Period:         "01/01/2023 - 02/02/2023",
					HowTo:          lex,
					OpenPositions:  []string{"OpenPositions Test 1", "OpenPositions Test 2", "OpenPositions Test 3"},
					RequiredSkills: []string{"RequiredSkills Test 1", "RequiredSkills Test 2"},
					Benefits:       []string{"Benefits Test 1"},
				},
				AccessToken: token,
			},
			expect: &pbv1.CreatePostResponse{
				Status:  201,
				Message: "Post created successfully",
			},
		},
		"success with same title": {
			req: &pbv1.CreatePostRequest{
				Post: &pbv1.CreatedPost{
					Topic:          "Topic Test 2",
					Description:    lex,
					Period:         "01/01/2023 - 02/02/2023",
					HowTo:          lex,
					OpenPositions:  []string{"OpenPositions Test 1", "OpenPositions Test 2", "OpenPositions Test 3"},
					RequiredSkills: []string{"RequiredSkills Test 1", "RequiredSkills Test 2"},
					Benefits:       []string{"Benefits Test 2"},
				},
				AccessToken: token,
			},
			expect: &pbv1.CreatePostResponse{
				Status:  201,
				Message: "Post created successfully",
			},
		},
		"Some fields are empty": {
			req: &pbv1.CreatePostRequest{
				Post: &pbv1.CreatedPost{
					Topic:          "",
					Description:    lex,
					Period:         "01/01/2023 - 02/02/2023",
					HowTo:          lex,
					OpenPositions:  []string{"OpenPositions Test"},
					RequiredSkills: []string{"RequiredSkills Test"},
					Benefits:       []string{"Benefits Test"},
				},
				AccessToken: token,
			},
			expect: &pbv1.CreatePostResponse{
				Status:  400,
				Message: "Please fill all required fields",
			},
		},
		"Not company": {
			req: &pbv1.CreatePostRequest{
				Post: &pbv1.CreatedPost{
					Topic:          "Topic Test",
					Description:    lex,
					Period:         "01/01/2023 - 02/02/2023",
					HowTo:          lex,
					OpenPositions:  []string{"OpenPositions Test"},
					RequiredSkills: []string{"RequiredSkills Test"},
					Benefits:       []string{"Benefits Test"},
				},
				AccessToken: tokenStudent,
			},
			expect: &pbv1.CreatePostResponse{
				Status:  403,
				Message: "You are not allowed to create post",
			},
		},
		"Invalid token": {
			req: &pbv1.CreatePostRequest{
				Post: &pbv1.CreatedPost{
					Topic:          "Topic Test",
					Description:    lex,
					Period:         "01/01/2023 - 02/02/2023",
					HowTo:          lex,
					OpenPositions:  []string{"OpenPositions Test"},
					RequiredSkills: []string{"RequiredSkills Test"},
					Benefits:       []string{"Benefits Test"},
				},
				AccessToken: "invalid token",
			},
			expect: &pbv1.CreatePostResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.CreatePost(ctx, tc.req)
			if err != nil {
				t.Errorf("could not create student: %v", err)
			} else {
				require.Equal(t, tc.expect.Status, res.Status)
				require.Equal(t, tc.expect.Message, res.Message)
			}
		})
	}

	// Check open positions
	resOpenPositions, err := c.GetOpenPositions(ctx, &pbv1.GetOpenPositionsRequest{
		AccessToken: token,
	})
	require.NoError(t, err)
	require.Equal(t, int64(200), resOpenPositions.Status)
	require.Equal(t, true, utils.CheckArrayEqual(&[]string{"OpenPositions Test 1", "OpenPositions Test 2", "OpenPositions Test 3"}, &resOpenPositions.OpenPositions))

	// Check required skills
	resRequiredSkills, err := c.GetRequiredSkills(ctx, &pbv1.GetRequiredSkillsRequest{
		AccessToken: token,
	})
	require.NoError(t, err)
	require.Equal(t, int64(200), resRequiredSkills.Status)
	require.Equal(t, true, utils.CheckArrayEqual(&[]string{"RequiredSkills Test 1", "RequiredSkills Test 2"}, &resRequiredSkills.RequiredSkills))

	// Check benefits
	resBenefits, err := c.GetBenefits(ctx, &pbv1.GetBenefitsRequest{
		AccessToken: token,
	})
	require.NoError(t, err)
	require.Equal(t, int64(200), resBenefits.Status)
	require.Equal(t, true, utils.CheckArrayEqual(&[]string{"Benefits Test 1", "Benefits Test 2"}, &resBenefits.Benefits))
}
