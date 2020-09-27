package avgs

import (
	"testing"
)

type testMedian struct {
	list   []float64
	median float64
}

func TestMedian(t *testing.T) {
	tests := []testMedian{
		testMedian{list: []float64{50, 12, 13, 12}, median: 12.5},
		testMedian{list: []float64{50, 52, 51, 53, 50}, median: 51},
		testMedian{list: []float64{10, 20, 40, 25, 30, 31}, median: 27.5},
	}
	for _, test := range tests {
		got := Median(test.list)
		want := test.median
		if want != got {
			t.Error("\ndata:", test.list, "\nwant:", want, "\ngot:", got)
		}
	}
}
