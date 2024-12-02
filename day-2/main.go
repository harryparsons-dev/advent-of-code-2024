package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part2() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	reports := strings.Split(strings.TrimSpace(string(file)), "\n")
	safeReportCount := 0
	for _, line := range reports {
		report := strings.Fields(line)

		if len(report) == 0 {
			continue
		}

		if checkSafe(report) {
			safeReportCount++
			continue
		}

		for i := 0; i < len(report); i++ {
			tmp := append([]string{}, report[:i]...)
			tmp = append(tmp, report[i+1:]...)
			if checkSafe(tmp) {
				safeReportCount++
				break
			}

		}

	}

	fmt.Printf("Number of safe reports: %+v", safeReportCount)
}

func part1() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	data := strings.Split(string(file), "\n")

	safeReportCount := 0
	for _, line := range data {

		report := strings.Fields(line)

		if checkSafe(report) {
			safeReportCount++
		}

	}
	fmt.Printf("Number of safe reports: %+v", safeReportCount)

}

func checkSafe(report []string) bool {
	if len(report) < 2 {
		return false
	}

	increasing := true
	first, _ := strconv.Atoi(report[0])
	last, _ := strconv.Atoi(report[len(report)-1])
	if first > last {
		increasing = false
	}

	for i := 0; i < len(report)-1; i++ {

		ptr1, err := strconv.Atoi(report[i])
		if err != nil {
			panic(err)
		}
		ptr2, err := strconv.Atoi(report[i+1])
		if err != nil {
			panic(err)
		}

		difference := math.Abs(float64(ptr1 - ptr2))
		if difference < 1 || difference > 3 {
			return false // Adjacent levels differ by less than 1 or more than 3.
		}

		if increasing && ptr2 < ptr1 {
			return false

		}
		if !increasing && ptr2 > ptr1 {
			return false

		}

	}
	return true
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func checkDifference(num1, num2 float64) bool {

	difference := math.Abs(float64(num1 - num2))
	//fmt.Println(difference)
	if difference < 1 || difference > 3 {
		return false
	}
	return true
}
