package main

import (
	"fmt"
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
	var tempNumberlist []string

	switch digit {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		for i := 0; i < 10; i++ {
			if digit > 1 {
				temp := fmt.Sprintf("%0*d", digit-1, i)
				tempNumberlist = append(tempNumberlist, temp)
			} else {
				tempNumberlist = append(tempNumberlist, strconv.Itoa(i))
			}
		}
	}
	return tempNumberlist
}
