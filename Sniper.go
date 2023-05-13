package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
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
