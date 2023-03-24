package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type RequestData struct {
	RequestNumber  int
	RequestPayload string
	RequestStatus  string
	RequestLength  int
	RequestTime    float64
}

var structMap = make(map[int]interface{})

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

func singleValue(valueName string, wordlist []string) {
	fmt.Println("--- Requests -|---------- Payload ----------|- Status -|- Length -|- Time ---")
	for request, payload := range wordlist {
		Start := time.Now()
		r, err := http.PostForm("https://echoof.me", url.Values{valueName: {payload}})
		check(err)
		totalTime := time.Since(Start).Seconds()
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		check(err)
		body := string(b)

		saveMap(request, payload, r.Status, len(body), totalTime)

		//request
		fmt.Printf("%d \t", request)
		//payload
		fmt.Printf("%s \t", payload)
		//status
		fmt.Printf("%s \t", r.Status)
		//length
		fmt.Printf("%d \t", len(body))
		//Time
		fmt.Printf("%f \n", totalTime)

	}
}

func saveMap(request int, payload string, status string, length int, time float64) {
	out := RequestData{request, payload, status, length, time}
	structMap[request] = out
}

func main() {
	var wordlistName string
	var valueName string
	var wordlist []string

	fmt.Printf("Welcome to app. Please enter your wordlist name: ")
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
	fmt.Printf("Please enter your value name: ")
	fmt.Scan(&valueName)

	singleValue(valueName, wordlist)

}
