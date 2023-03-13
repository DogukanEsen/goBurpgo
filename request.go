package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func kelimeOku(filename string) []string {
	var words []string
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}

func main() {
	fmt.Printf("Welcome to app. Please enter your wordlist name: ")
	var wordlistName string
	fmt.Scan(&wordlistName)
	var wordlist []string = kelimeOku(wordlistName)
	fmt.Printf("%s is reading....", wordlistName)
	fmt.Println(wordlist)

	/*
		data := url.Values{"name": {"John Doe"}, "aa": {"aa"}}
		r, err := http.PostForm("https://echoof.me", data)
		check(err)
		defer r.Body.Close()
		io.Copy(os.Stdout, r.Body)
		print("Deneme")
	*/

}
