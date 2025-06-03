package units

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
	}{
		{"1s", time.Second},
		{"1m", time.Minute},
		{"1h", time.Hour},
		{"1d", 24 * time.Hour},
		{"2d", 48 * time.Hour},
		{"3d5h", 3*24*time.Hour + 5*time.Hour},
		{"7d12h30m", 7*24*time.Hour + 12*time.Hour + 30*time.Minute},
		{"", 0}, // 错误输入，期望返回 error
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			dr, err := Duration(test.input)
			if err != nil {
				t.Errorf("Expected no error for input %q, got %v", test.input, err)
			} else if dr != test.expected {
				t.Errorf("For input %q: expected duration %v, got %v", test.input, test.expected, dr)
			}
		})
	}
}
