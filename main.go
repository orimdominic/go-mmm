package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func logErr(err error) {
	fmt.Errorf(err.Error())
}

func getFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
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
	return 0.00
}

func mode(nums []float64) float64 {
	return 0.00
}

func main() {
	file := getFile("data.txt")
	_ = getNums(file)
}
