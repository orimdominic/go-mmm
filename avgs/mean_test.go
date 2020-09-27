package avgs

import (
	"testing"
)

type testMean struct {
	list []float64
	mean float64
}

func TestMean(t *testing.T) {
	tests := []testMean{
		testMean{list: []float64{50, 12, 13, 12}, mean: 21.75},
		testMean{list: []float64{50, 52, 51, 53, 50}, mean: 51.2},
		testMean{list: []float64{10, 20, 40, 25, 30, 31}, mean: 26},
	}
	for _, test := range tests {
		got := Mean(test.list)
		want := test.mean
		if want != got {
			t.Error("\ndata:", test.list, "\nwant:", want, "\ngot:", got)
		}
	}
}
