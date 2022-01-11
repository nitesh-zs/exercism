package exercism

import "testing"

func TestTriType(t *testing.T) {
	tests := []struct {
		side1    int
		side2    int
		side3    int
		expected string
	}{
		{0, 0, 0, "Invalid Triangle"},
		{3, 5, -1, "Invalid Triangle"},
		{1, 1, 3, "Invalid Triangle"},
		{4, 4, 5, "Isosceles Triangle"},
		{5, 5, 5, "Equilateral Triangle"},
		{3, 4, 5, "Scalene Triangle"},
	}

	for i, tc := range tests {
		output := TriType(tc.side1, tc.side2, tc.side3)
		if output != tc.expected {
			t.Errorf("Testcase%v failed - Expected: %v Got: %v", i, tc.expected, output)
		}
	}
}
