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

func TestUpdatePost(t *testing.T) {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	tokenStudent, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "student",
	})
	require.NoError(t, err)

	lex := `{
		"root": {
		  "children": [],
		  "direction": "ltr",
		  "format": "",
		  "indent": 0,
		  "type": "root",
		  "version": 1
		}
	  }
	`

	updatedLex := `{
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
	openPositions := []string{"Software Engineer", "Data Scientist", "Data Analyst"}
	requiredSkills := []string{"Golang", "Python"}
	benefits := []string{"Free lunch", "Free dinner"}

	updatedTopic := "NEW What to expect from here on out"
	updatedDescription := updatedLex
	updatedPeriod := "NEW 1 month"
	updatedHowTo := updatedLex
	updatedOpenPositions := []string{"NEW Software Engineer", "NEW Data Scientist"}
	updatedRequiredSkills := []string{"NEW Golang", "NEW Python", "NEW Java"}
	updatedBenefits := []string{"NEW Free lunch", "NEW Free dinner"}

	CreateRes, err := c.CreatePost(ctx, &pbv1.CreatePostRequest{
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
	require.Equal(t, int64(201), CreateRes.Status)
	p := &pbv1.Post{
		Topic:          updatedTopic,
		Description:    updatedDescription,
		Period:         updatedPeriod,
		HowTo:          updatedHowTo,
		OpenPositions:  updatedOpenPositions,
		RequiredSkills: updatedRequiredSkills,
		Benefits:       updatedBenefits,
	}

	tests := map[string]struct {
		req    *pbv1.UpdatePostRequest
		expect *pbv1.UpdatePostResponse
	}{
		"success": {
			req: &pbv1.UpdatePostRequest{
				AccessToken: token,
				Id:          CreateRes.Id,
				Post:        p,
			},
			expect: &pbv1.UpdatePostResponse{
				Status:  200,
				Message: "Post updated successfully",
			},
		},
		"fail: unauthorized": {
			req: &pbv1.UpdatePostRequest{
				AccessToken: tokenStudent,
				Id:          CreateRes.Id,
				Post:        p,
			},
			expect: &pbv1.UpdatePostResponse{
				Status:  401,
				Message: "Unauthorized",
			},
		},
		"some field is empty": {
			req: &pbv1.UpdatePostRequest{
				AccessToken: token,
				Id:          CreateRes.Id,
				Post: &pbv1.Post{
					Topic:          "",
					Description:    updatedDescription,
					Period:         updatedPeriod,
					HowTo:          updatedHowTo,
					OpenPositions:  updatedOpenPositions,
					RequiredSkills: updatedRequiredSkills,
					Benefits:       updatedBenefits,
				},
			},
			expect: &pbv1.UpdatePostResponse{
				Status:  400,
				Message: "Please fill all required fields",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.UpdatePost(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			if tc.expect.Status == 200 {
				g, err := c.GetPost(ctx, &pbv1.GetPostRequest{
					AccessToken: token,
					Id:          CreateRes.Id,
				})
				require.NoError(t, err)
				require.Equal(t, updatedTopic, g.Post.Topic)
				require.Equal(t, updatedDescription, g.Post.Description)
				require.Equal(t, updatedPeriod, g.Post.Period)
				require.Equal(t, updatedHowTo, g.Post.HowTo)
				require.Equal(t, updatedOpenPositions, g.Post.OpenPositions)
				require.Equal(t, updatedRequiredSkills, g.Post.RequiredSkills)
				require.Equal(t, updatedBenefits, g.Post.Benefits)
			}
		})
	}
}
