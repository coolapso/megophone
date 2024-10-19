package xdotcom

import (
	"github.com/coolapso/megophone/internal/util"
	"testing"
)

func TestIsXlength(t *testing.T) {

	if IsXLenght(util.LongPost) {
		t.Fatalf("Expected false, got true")
	}

	if !IsXLenght(util.LongPost[:200]) {
		t.Fatalf("Expected true, got false")
	}
}
