package mastodon

import (
	"testing"
	"github.com/coolapso/megophone/internal/util"
)

func TestIsXlength(t *testing.T) {

	if IsToothLenght(util.LongPost) {
		t.Fatalf("Expected false, got true")
	}

	if !IsToothLenght(util.LongPost[:200]) {
		t.Fatalf("Expected true, got false")
	}
}
