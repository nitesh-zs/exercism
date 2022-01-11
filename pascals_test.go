package exercism

import "testing"

func TestPascalsT(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, ""}, {1, "1 \n"}, {2, "1 \n1 1 \n"}, {5, "1 \n1 1 \n1 2 1 \n1 3 3 1 \n1 4 6 4 1 \n"},
	}

	for i, tc := range tests {
		output := PascalsT(tc.input)
		if output != tc.expected {
			t.Errorf("Testcase%v failed - Expected:\n%v\n Got:\n%v\n", i, tc.expected, output)
		}
	}
}
