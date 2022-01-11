package exercism

import "testing"

func TestIsLeap(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{1997, false},
		{1996, true},
		{2000, true},
		{1900, false},
	}

	for i, tc := range tests {
		output := isLeap(tc.input)
		if output != tc.expected {
			t.Errorf("Testcase%v failed - Expected: %v Got: %v", i, tc.expected, output)
		}
	}
}
