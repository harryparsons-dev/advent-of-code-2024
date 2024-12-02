package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//part1()
	part2()
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var list1 []int
	var list2 []int
	var simScore int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, "   ")
		num1, err := strconv.Atoi(tmp[0])
		if err != nil {
			fmt.Println(err)
		}
		list1 = append(list1, num1)

		num2, err := strconv.Atoi(tmp[1])
		if err != nil {
			fmt.Println(err)
		}

		list2 = append(list2, num2)

	}

	for i := range len(list1) {
		tmpCount := 0
		for j := range len(list2) {
			if list1[i] == list2[j] {
				tmpCount = tmpCount + 1
			}
		}
		increase := list1[i] * tmpCount
		simScore = simScore + increase

	}

	fmt.Println(simScore)

}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var list1 []int
	var list2 []int
	var distance int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, "   ")
		num1, err := strconv.Atoi(tmp[0])
		if err != nil {
			fmt.Println(err)
		}
		list1 = append(list1, num1)

		num2, err := strconv.Atoi(tmp[1])
		if err != nil {
			fmt.Println(err)
		}

		list2 = append(list2, num2)

	}

	sort.Ints(list1)
	sort.Ints(list2)

	for i := range len(list1) {
		tmpDist := math.Abs(float64(list1[i] - list2[i]))
		distance = distance + int(tmpDist)
	}

	fmt.Println(distance)
}
