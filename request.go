package main

import (
	"bufio"
	"fmt"
	"os"
)

// Sıralama yapma, max 3'e kadar payload gönderme, istenilen yeri değiştirme, hazır wordlistler

var tempWordlist []string

func check(e error) bool {
	err := false
	if e != nil {
		fmt.Println(e)
		err = true
	}
	return err
}

func kelimeOku(filename string) []string {
	var words []string
	file, err := os.Open(filename)
	if check(err) {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}

func main() {
	var wordlistName, valueName, website string
	var choice int
	var wordlist []string
	fmt.Printf("Welcome to app.\n 0 - Numbers\n 1 - Simplelist\nPlease enter a number: ")
	fmt.Scan(&choice)
	if choice == 0 {
		var choice2 int
		fmt.Printf("Welcome to app.\n 0 - Numbers\n 1 - Number2\nPlease enter a number: ")
		fmt.Scan(&choice2)
		if choice2 == 0 {
			var from, to, step int
			fmt.Printf("Please enter the value from: ")
			fmt.Scan(&from)
			fmt.Printf("Please enter the value to: ")
			fmt.Scan(&to)
			fmt.Printf("Please enter the value step: ")
			fmt.Scan(&step)
			wordlist = numbers(from, to, step)
		}
		if choice2 == 1 {
			var digit int
			fmt.Printf("Please enter the value digit: ")
			fmt.Scan(&digit)
			wordlist = numbers2(digit)
		}

	}
	if choice == 1 {
		fmt.Printf("Please enter your wordlist name: ")
		fmt.Scan(&wordlistName)
		wordlist = kelimeOku(wordlistName)
		for {
			if wordlist == nil {
				fmt.Printf("Wrong wordlist name. Please enter your wordlist again: ")
				fmt.Scan(&wordlistName)
				wordlist = kelimeOku(wordlistName)
			} else {
				break
			}
		}
		fmt.Printf("%s is reading....\n", wordlistName)
	}

	fmt.Printf("Please enter your web link(https://echoof.me as 0): ")
	fmt.Scan(&website)
	if website == "0" {
		// website = "" Unfortunately echoof.me is closed
	}
	fmt.Printf("Please enter your value name: ")
	fmt.Scan(&valueName)
	singleValue(valueName, wordlist, website)

}
