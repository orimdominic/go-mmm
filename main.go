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
	return 0.00
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
