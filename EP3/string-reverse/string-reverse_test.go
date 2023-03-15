package reverser

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"palindrome", "racecar", "racecar"},
		{"even-length string", "foobar", "raboof"},
		{"odd-length string", "hello", "olleh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.input)
			if got != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}
