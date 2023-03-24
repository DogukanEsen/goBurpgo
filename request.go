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

func singleValue(valueName string, wordlist []string) {

	fmt.Println("--- Requests -|---------- Payload ----------|- Status -|- Length -|- Time ---")
	for request, a := range wordlist {
		Start := time.Now()
		r, err := http.PostForm("https://echoof.me", url.Values{valueName: {a}})
		check(err)
		ava := time.Since(Start).Seconds()
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		check(err)
		body := string(b)

		saveMap(request, a, r.Status, len(body), ava)
		/*
			//request
			fmt.Printf("%d \t", request)
			//payload
			fmt.Printf("%s \t", a)
			//status
			fmt.Printf("%s \t", r.Status)
			//length
			a := len(body)
			fmt.Printf("%d \t", a)
			//Time
			fmt.Printf("%f \n", ava)
		*/
	}
}

func saveMap(request int, payload string, status string, length int, time float64) {
	structMap := make(map[int]interface{})
	out := RequestData{request, payload, status, length, time}
	structMap[request] = out

	fmt.Println(structMap)
}

func main() {
	var wordlistName string
	var valueName string
	var wordlist []string

	fmt.Printf("Welcome to app. Please enter your wordlist name: ")
	fmt.Scan(&wordlistName)
	wordlist = kelimeOku(wordlistName)
	fmt.Printf("%s is reading....\n", wordlistName)
	fmt.Printf("Please enter your value name: ")
	fmt.Scan(&valueName)

	singleValue(valueName, wordlist)

}
