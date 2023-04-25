package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

// Sıralama yapma, max 3'e kadar payload gönderme, istenilen yeri değiştirme, hazır wordlistler
type RequestData struct {
	RequestNumber  int
	RequestPayload string
	RequestStatus  string
	RequestLength  int
	RequestTime    float64
}

var structMap = make(map[int]interface{})
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

func singleValue(valueName string, wordlist []string, website string) {
	fmt.Println("--- Requests -|---------- Payload ----------|- Status -|- Length -|- Time ---")
	for request, payload := range wordlist {
		Start := time.Now()
		r, err := http.PostForm(website, url.Values{valueName: {payload}})
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
}

func main() {
	var wordlistName string
	var valueName string
	var website string
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
	fmt.Printf("Please enter your web link(https://echoof.me as 0): ")
	fmt.Scan(&website)
	if website == "0" {
		website = "https://echoof.me"
	}
	fmt.Printf("Please enter your value name: ")
	fmt.Scan(&valueName)

	singleValue(valueName, wordlist, website)

}
