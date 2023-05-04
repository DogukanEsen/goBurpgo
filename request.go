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
	var wordlistName string
	var choice int
	var valueName string
	var website string
	var wordlist []string

	fmt.Printf("Welcome to app.\n 0 - Numbers\n 1 - Simplelist")
	fmt.Scan(&choice)
	if choice == 0 {

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
		website = "https://echoof.me"
	}
	fmt.Printf("Please enter your value name: ")
	fmt.Scan(&valueName)

	singleValue(valueName, wordlist, website)

}
