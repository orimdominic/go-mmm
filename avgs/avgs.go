package avgs

import (
	"sort"
)

// Mean alculates the mean of a slice of numbers
func Mean(nums []float64) float64 {
	if len(nums) == 0 {
		return 0.00
	}
	sum := 0.00
	for _, num := range nums {
		sum += float64(num)
	}
	return sum / float64(len(nums))
}

// Median alculates the median  of a slice of numbers
func Median(nums []float64) float64 {
	sort.Float64s(nums)
	if len(nums)%2 == 0 {
		id := (len(nums) - 1) / 2
		return float64((nums[id] + nums[id+1]) / 2)
	}
	return float64(nums[(len(nums)-1)/2])
}

// Mode alculates the mode  of a slice of numbers
func Mode(nums []float64) interface{} {
	modMap := make(map[float64]int)
	for _, num := range nums {
		modMap[num]++
	}
	// create a slice of the freqs
	freqSlc := make([]int, len(nums))
	for _, v := range modMap {
		freqSlc = append(freqSlc, v)
	}
	sort.Ints(freqSlc) // sort freqs asc
	if freqSlc[len(freqSlc)-1] == freqSlc[len(freqSlc)-2] {
		// if the last two highest frequencies are the same
		return nil
	}
	topF := freqSlc[len(freqSlc)-1]
	var res float64
	for num, f := range modMap {
		if f == topF {
			res = num
			break
		}
	}
	return res
}
