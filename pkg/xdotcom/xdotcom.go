package xdotcom

import ( 
	"fmt"
	"github.com/coolapso/megophone/internal/util"
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

type Secrets struct {
	oauthToken string
	oauthTokenSecret string
	apiKey string
	apiKeySecret string
}

func(x *Secrets) SetOauthToken(s string) { 
	x.oauthToken = s
}

func(x *Secrets) SetOauthTokenSecret (s string) {
	x.oauthTokenSecret = s
}

func(x *Secrets) SetApiKey(s string) { 
	x.apiKey = s
}

func(x *Secrets) SetApiKeySecret (s string) {
	x.apiKeySecret = s
}

func(x *Secrets) GetOauthToken() string {
	return x.oauthToken
}

func(x *Secrets) GetOauthTokenSecret() string { 
	return x.oauthTokenSecret
}

func(x *Secrets) GetApiKey() string {
	return x.apiKey
}

func(x *Secrets) GetApiKeySecret() string {
	return x.apiKeySecret
}



func CreatePost(ctx context.Context, client *gotwi.Client, text string) (ID string, err error) { 
	if !util.IsXLenght(text) {
		return "", fmt.Errorf("Text is too long for a tweet")
	}

	post := &types.CreateInput {
		Text: gotwi.String(text),
	}

	resp, err := managetweet.Create(ctx, client, post)
	if err != nil { 
		return "", fmt.Errorf("failed to post tweet: %v\n", err)
	}

	return *resp.Data.ID, nil
}

func CreatePostWithMedia(ctx *context.Context, client *gotwi.Client, text string, media []byte) (err error) {

	return nil
}

// Placeholder function for future development
func CreateReplyToPost(ctx *context.Context, client *gotwi.Client, text, postID string) (err error) {

	return nil
}

// Placeholder function for future development
func CreatePostWithQuote(ctx *context.Context, client *gotwi.Client, text, postID string) (err error) {

	return nil
}
