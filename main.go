package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func logErr(err error) {
	fmt.Errorf(err.Error())
}

func getFile(path string) *os.File {
	f, err := os.Open(path)
	if err !=
		nil {
		log.Fatal(err)
		return nil
	}
	return f
}

func getNums(f *os.File) (nums []float64) {
	defer f.Close()
	scnr := bufio.NewScanner(f)
	for scnr.Scan() {
		num, err := strconv.ParseFloat(scnr.Text(), 64)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}
	return
}

func mean(nums []float64) float64 {
	if len(nums) == 0 {
		return 0.00
	}
	sum := 0.00
	for _, num := range nums {
		sum += float64(num)
	}
	return sum / float64(len(nums))
}

func median(nums []float64) float64 {
	sort.Float64s(nums)
	if len(nums)%2 == 0 {
		id := (len(nums) - 1) / 2
		return float64((nums[id] + nums[id+1]) / 2)
	}
	return float64(nums[(len(nums)-1)/2])
}

func mode(nums []float64) interface{} {
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

func main() {
	file := getFile("data.txt")
	nums:= getNums(file)
	fmt.Printf("%-6v %.2f\n","Mean", mean(nums))
	fmt.Printf("%-6v %.2f\n","Median", median(nums))
	fmt.Printf("%-6v %.2f\n", "Mode", mode(nums))
}
