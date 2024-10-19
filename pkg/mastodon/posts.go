package mastodon

import (
	"context"
	"fmt"
	"github.com/coolapso/megophone/internal/util"
	gomasto "github.com/mattn/go-mastodon"
)

func IsToothLenght(s string) bool {
	return len(util.CleanString(s)) <= 500
}

func CreatePost(ctx context.Context, client *gomasto.Client, text, visibility string) (ID string, err error) {
	if !IsToothLenght(text) {
		return "", fmt.Errorf("Text is too long for a toot")
	}

	toot := &gomasto.Toot{
		Status:     text,
		Visibility: visibility,
	}

	post, err := client.PostStatus(ctx, toot)
	if err != nil {
		return "", fmt.Errorf("failed to post toot, %v\n", err)
	}

	return string(post.ID), nil
}
