package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int

	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		elf0 := strings.Split(pair[0], "-")
		elf1 := strings.Split(pair[1], "-")
		elf0s, _ := strconv.Atoi(elf0[0])
		elf0e, _ := strconv.Atoi(elf0[1])
		elf1s, _ := strconv.Atoi(elf1[0])
		elf1e, _ := strconv.Atoi(elf1[1])

		if (elf0s >= elf1s && elf0e <= elf1e) || (elf1s >= elf0s && elf1e <= elf0e) {
			sum++
		}
	}
	fmt.Println(sum)
}

