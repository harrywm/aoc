package main

import (
	"testing"
)

func TestDigitize(t *testing.T) {
    var tests = []struct {
        input string
        want string
    }{
        {"1two1", "121"},
        {"nine8aoibnlkjnxone001", "98aoibnlkjnx1001"},
        {"asnlfkdn88eight662", "asnlfkdn888662"},
        {"six8sixfivesicasnfjkn88", "6865sicasnfjkn88"},
        {"eighteighteighteighteighteight", "888888"},
    }
    for _, test := range tests {
        if got := digitize(test.input); got != test.want {
            t.Errorf("digitize(%q) = %q", test.input, got)
        }
    }
}
