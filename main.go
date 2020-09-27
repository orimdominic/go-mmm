package main

import (
	"bufio"
	"fmt"
	"github.com/sudo-kaizen/go-mmm/avgs"
	"log"
	"os"
	"strconv"
)

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

func main() {
	file := getFile("data.txt")
	nums := getNums(file)
	fmt.Printf("%-6v %.2f\n", "Mean", avgs.Mean(nums))
	fmt.Printf("%-6v %.2f\n", "Median", avgs.Median(nums))
	fmt.Printf("%-6v %.2f\n", "Mode", avgs.Mode(nums))
}
