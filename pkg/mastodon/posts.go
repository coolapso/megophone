package mastodon

import (
	"github.com/coolapso/megophone/internal/util"
)

func IsToothLenght(s string) bool {
	return len(util.CleanString(s)) <= 500
}
