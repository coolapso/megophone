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
		name string
		input string
		want string
	}{
		{"New line", "foo\\nbar", "foo\nbar"},
		{"Tab", "foo\\tbar", "foo\tbar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CleanString(tt.input)
			if got != tt.want {
				t.Fatalf("String does not match expected value, want %v, got %v",tt.want, got)
			}
		})
	}
}

func TestIsXlength(t *testing.T) {
	LongPost := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Sed ut perspiciatis unde omnis iste natus error sit voluptatem."

	if IsXLenght(LongPost) {
		t.Fatalf("Expected false, got true")
	}

	if IsToothLenght(LongPost) {
		t.Fatalf("Expected false, got true")
	}

	if !IsXLenght(LongPost[:200]) {
		t.Fatalf("Expected true, got false")
	}

	if !IsToothLenght(LongPost[:500]) {
		t.Fatalf("Expected true, got false")
	}
}
