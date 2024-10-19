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

func TestCleanString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"New line", "foo\\nbar", "foo\nbar"},
		{"Tab", "foo\\tbar", "foo\tbar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CleanString(tt.input)
			if got != tt.want {
				t.Fatalf("String does not match expected value, want %v, got %v", tt.want, got)
			}
		})
	}
}
