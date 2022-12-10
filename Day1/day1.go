package Day1

import (
	"fmt"
	"strconv"
)

type Elf struct {
	num int
	calories int
}

func Execute() {
	input := Part1Input()
	
	var elves []Elf
	var currentElf Elf
	var maxCal int
	for i, s := range input {
		if s != 0 {
			currentElf.calories += int(s)
		} else {
			elves = append(elves, currentElf)
			currentElf = Elf{num: i, calories: 0}
		}

		if currentElf.calories > maxCal {
			maxCal = currentElf.calories
		}
	}

	fmt.Printf("Day1:\nPart 1: Max Calories are: %s\n", strconv.Itoa(maxCal))

	
	currentElf = Elf{}
	elves = elves[:0]
	var maxCal1 int
	var maxCal2 int
	var maxCal3 int
	
	for i, s := range input {
		if s != 0 {
			currentElf.calories += int(s)
		} else {
			elves = append(elves, currentElf)
			currentElf = Elf{num: i, calories: 0}
		}

		if currentElf.calories > maxCal1 {
			maxCal3 = maxCal2
			maxCal2 = maxCal1
			maxCal1 = currentElf.calories
		} else if currentElf.calories > maxCal2 {
			maxCal3 = maxCal2
			maxCal2 = currentElf.calories
		}else if currentElf.calories > maxCal3 {
			maxCal3 = currentElf.calories
		}
	}

	sum := maxCal1 + maxCal2 + maxCal3
	
	fmt.Printf("Part 2: Total Calories are: %s", strconv.Itoa(sum))
}