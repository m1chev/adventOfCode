// The line below solves the task in terminal
// cat ../input | paste -sd + | sed 's/++/\n/g' | bc | sort -n | tail -3 | paste -sd + | bc

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	var elves []int

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			calories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Print(err)
			}
			sum += calories
		} else {
			elves = append(elves, sum)
			sum = 0
		}
	}
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})
	total := elves[0] + elves[1] + elves[2]
	fmt.Println(total)
}
