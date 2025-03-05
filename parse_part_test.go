package main

import (
	"strings"
	"testing"
)

func TestParsePart(t *testing.T) {
	tests := []struct {
		part       string
		start, max int
		expected   string
	}{
		{"*", 0, 5, "0 1 2 3 4 5"},
		{"1,3,5", 0, 5, "1 3 5"},
		{"2-4", 0, 5, "2 3 4"},
		{"*/2", 0, 5, "0 2 4"},
		{"*/3", 0, 9, "0 3 6 9"},
	}

	for _, tt := range tests {
		result, err := parsePart(tt.part, tt.start, tt.max)
		if err != nil {
			t.Errorf("parsePart(%q, %d, %d) returned error: %v", tt.part, tt.start, tt.max, err)
			continue
		}

		if strings.TrimSpace(result) != tt.expected {
			t.Errorf("parsePart(%q, %d, %d) = %q, want %q",
				tt.part, tt.start, tt.max, result, tt.expected)
		}
	}
}
