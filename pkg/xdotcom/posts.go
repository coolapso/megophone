package xdotcom

import (
	"context"
	"fmt"
	"github.com/coolapso/megophone/internal/util"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"

	//Twitter API V2 does not yet support uploading media therefore these "legacy" packages are needed
	// "github.com/dghubble/oauth1"
	twitterv1 "github.com/drswork/go-twitter/twitter"
)

func IsXLenght(s string) bool {
	return len(util.CleanString(s)) <= 280
}

func CreatePost(ctx context.Context, client *gotwi.Client, text string) (ID string, err error) {
	if !IsXLenght(text) {
		return "", fmt.Errorf("Text is too long for a tweet")
	}

	postInput := &types.CreateInput{
		Text: gotwi.String(text),
	}

	resp, err := managetweet.Create(ctx, client, postInput)
	if err != nil {
		return "", fmt.Errorf("failed to post tweet, %v\n", err)
	}

	return *resp.Data.ID, nil
}

func CreatePostWithMedia(ctx context.Context, client *gotwi.Client, clientV1 *twitterv1.Client, text string, media []byte, mediaType string) (ID string, err error) {
	uploadResult, uploadHttpResp, err := clientV1.Media.Upload(media, mediaType)
	if err != nil {
		return "", fmt.Errorf("Failed to upload media, %v\n", err)
	}

	if uploadHttpResp.StatusCode != 200 && uploadHttpResp.StatusCode != 201 {
		return "", fmt.Errorf("Failed to upload media, %v\n", uploadHttpResp.Status)
	}

	mediaInput := &types.CreateInputMedia{
		MediaIDs: []string{uploadResult.MediaIDString},
	}

	postInput := &types.CreateInput{
		Text:  gotwi.String(text),
		Media: mediaInput,
	}

	postResp, err := managetweet.Create(ctx, client, postInput)
	if err != nil {
		return "", fmt.Errorf("Failed to post tweet with media, %v\n", err)
	}

	return *postResp.Data.ID, nil
}

// Placeholder function for future development
func CreateReplyToPost(ctx *context.Context, client *gotwi.Client, text, postID string) (err error) {

	return nil
}

// Placeholder function for future development
func CreatePostWithQuote(ctx *context.Context, client *gotwi.Client, text, postID string) (err error) {

	return nil
}
