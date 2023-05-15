package main

import (
	"math"
	"strconv"
)

func numbers(from int, to int, step int) []string {
	//Return ederken pointer ediyor düzelt
	var tempNumberlist []string
	for i := from; i < to; i = i + step {
		tempNumberlist = append(tempNumberlist, strconv.Itoa(i))
	}
	return tempNumberlist
}

func numbers2(digit int) []string {
	result := make([]string, 0, int(math.Pow(float64(10), float64(digit))))

	numbers2_Recursive("", 10, digit, &result)
	return result
}

func numbers2_Recursive(prefix string, base int, digit int, result *[]string) {
	if digit == 0 {
		*result = append(*result, prefix)
		return
	}

	for i := 0; i < base; i++ {
		numbers2_Recursive(prefix+strconv.Itoa(i), base, digit-1, result)
	}
}
