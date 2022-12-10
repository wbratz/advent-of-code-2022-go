package Day3

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func Execute() {
	fmt.Println("\nDay3")

	sum := 0

	input := MainInput()
	for _, s := range input {
		firstSack := s[:len(s)/2]
		secondSack := s[len(s)/2:]

		commons := FindCommon(firstSack, secondSack)
		sum += GetValueFromByte(commons)
	}

	fmt.Printf("\nPart 1: %s", strconv.Itoa(sum))

	sum = 0

	for i := 0; i < len(input); i++ {
		firstCommon := FindCommon(input[i], input[i + 1])
		allCommon := FindCommon(firstCommon, input[i + 2])

		sum += GetValueFromByte(allCommon)
		i = i + 2
	}

	fmt.Printf("\nPart 2: %s", strconv.Itoa(sum))
}

func FindCommon(firstSack string, secondSack string) string {
	set := map[byte]struct{}{}
	dict := map[byte]int{}
	str := ""
	buff := bytes.NewBufferString(str)
	for _, s := range firstSack {
		dict[byte(s)] = 1
	}

	for _, s := range secondSack {
		_, ok := dict[byte(s)]
		if ok {
			set[byte(s)] = struct{}{}
		}
	}

	for b := range set {
		buff.WriteByte(b)
	}

	return buff.String()
}

func GetValueFromByte(str string) int {
	val := 0
	upper := strings.ToUpper(str)

	for i := 0; i < len(str); i++ {
		if str[i] == upper[i] {
			val += int(byte(str[i])) - int(byte('A')) + 27
		} else {
			val += int(byte(str[i])) - int(byte('a')) + 1
		}
	}

	return val
}