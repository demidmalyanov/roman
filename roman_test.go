package roman

import (
	"testing"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type romanTest struct {
	arg1, arg2 string
	want       int
}

var romanTests = []romanTest{
	{"MMLXII", "MMLX", 1},
	{"XX|XX", "XX|XXI", -1},
	{"X|N|I", "V|VI", 1},
}

func TestRoman(t *testing.T) {

	for _, test := range romanTests {
		if output := compareTwoRomans(test.arg1, test.arg2); output != test.want {
			t.Errorf("Output %v not equal to expected %v", output, test.want)
		}
	}
}
