package Day2

import (
	"fmt"
	"strconv"
)

func Execute() {
	scores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	oppToYou := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	input := MainInput()

	score := 0

	for _, s := range input {
		if (s.Opponent == "A" && s.Yours == "X") ||
		(s.Opponent == "B" && s.Yours == "Y") ||
		(s.Opponent == "C" && s.Yours == "Z") {
			score += scores[s.Yours] + 3
		} else if s.Opponent == "A" && s.Yours == "Y" {
			score += scores[s.Yours] + 6
		} else if s.Opponent == "B" && s.Yours == "Z" {
			score += scores[s.Yours] + 6
		} else if s.Opponent == "C" && s.Yours == "X" {
			score += scores[s.Yours] + 6
		} else {
			score += scores[s.Yours] + 0
		}
	}

	fmt.Printf("\nDay2 \n Part1: %s", strconv.Itoa(score))

	score = 0

	for _, s := range input {
		if s.Yours == "X" {
			if s.Opponent == "A" {
				score += scores["Z"]
			} else if s.Opponent == "B" {
				score += scores["X"]
			} else if s.Opponent == "C" {
				score += scores["Y"]
			}
		} else if s.Yours == "Y" {
			score += scores[oppToYou[s.Opponent]] + 3
		} else {
			if s.Opponent == "A" {
				score += scores["Y"] + 6
			} else if s.Opponent == "B" {
				score += scores["Z"]  + 6
			} else if s.Opponent == "C" {
				score += scores["X"] + 6
			}
		}
	}

	fmt.Printf("\n Part2: %s", strconv.Itoa(score))
}

