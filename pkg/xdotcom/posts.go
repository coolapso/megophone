package xdotcom

import (
	"context"
	"fmt"
	"github.com/coolapso/megophone/internal/util"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

func CreatePost(ctx context.Context, client *gotwi.Client, text string) (ID string, err error) {
	if !util.IsXLenght(text) {
		return "", fmt.Errorf("Text is too long for a tweet")
	}

	post := &types.CreateInput{
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
