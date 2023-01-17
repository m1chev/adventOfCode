package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count, temp, sum int
	var fst, snd, trd string

	for scanner.Scan() {
		count++
		if fst == "" {
			fst = scanner.Text()
		} else if snd == "" {
			snd = scanner.Text()
		} else {
			trd = scanner.Text()
		}

		if count%3 == 0 {
			for _, i := range fst {
				for _, j := range snd {
					if i == j {
						for _, e := range trd {
							if i == e {
								if i > 96 {
									temp += int(i) - 96
								} else {
									temp += int(i) - 38
								}
								goto Next
							}
						}
					}
				}
			}
		Next:
			sum += temp
			temp = 0
			fst = ""
			snd = ""
			trd = ""
		}
	}
	fmt.Println(sum)
}
