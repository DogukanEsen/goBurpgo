package main

import (
	"fmt"
	"strconv"
)

func check(e error) bool {
	err := false
	if e != nil {
		fmt.Println(e)
		err = true
	}
	return err
}

func numbers(from int, to int, step int) []int {
	//Return ederken pointer ediyor düzelt
	var tempNumberlist []int
	for i := from; i < to; i = i + step {
		tempNumberlist = append(tempNumberlist, i)
	}
	return tempNumberlist
}

func numbers2(basamak int, from int, to int, step int) {
	var tempNumberlist []int

	switch basamak {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		for i := 0; i < 10; i++ {
			if basamak > 1 {
				temp := fmt.Sprintf("%0*d", basamak-1, i)
				temp2, err := strconv.Atoi(temp)
				check(err)
				tempNumberlist = append(tempNumberlist, temp2)
			} else {
				tempNumberlist = append(tempNumberlist, i)
			}
		}
	}
	fmt.Print(tempNumberlist)
}
