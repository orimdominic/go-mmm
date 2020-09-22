package main

import (
	"testing"
)

type testMode struct {
	list []float64
	mode interface{}
}

func TestMode(t *testing.T) {
	tests := []testMode{
		testMode{list: []float64{50, 12, 13, 12}, mode: float64(12)},
		testMode{list: []float64{50, 52, 51, 53, 50}, mode: float64(50)},
		testMode{list: []float64{10, 20, 40, 25, 30, 31}, mode: nil},
	}
	for _, test := range tests {
		got := mode(test.list)
		want := test.mode
		if want != got {
			t.Error("\ndata:", test.list, "\nwant:", want, "\ngot:", got)
		}
	}
}
