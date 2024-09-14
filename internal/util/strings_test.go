package util

import (
	"testing"
)

func TestMaskString(t *testing.T) {
	t.Run("test 3 char string", func(t *testing.T) {
		want := "bar"
		got := MaskString("bar")
		if got != want { 
			t.Fatalf("Strings do not match, want %v, got %v", want, got)
		}
	})

	t.Run("test masked string", func(t *testing.T) {
		want := "******baz"
		got := MaskString("foobarbaz")
		if got != want { 
			t.Fatalf("Strings do not match, want %v, got %v", want, got)
		}

	})
}

//TODO: Add tests for lengths
