package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type City struct {
	Pop2023 int64   `json:"pop2023"`
	Pop2022 int64   `json:"pop2022"`
	City    string  `json:"city"`
	Country string  `json:"country"`
	Growth  float64 `json:"growthRate"`
	Type    string  `json:"type"`
	Rank    int64   `json:"rank"`
}

type Cities struct {
	Cities []City `json:"cities"`
}

type CallResponse struct {
	City    City
	Status  string
	Time    time.Time
	Elapsed time.Duration
}

func loadCities() (*Cities, error) {
	var cities Cities
	f, err := os.Open("cities.json")
	if err != nil {
		return &cities, err
	}
	defer f.Close()

	byteValue, err := io.ReadAll(f)
	if err != nil {
		return &cities, err
	}

	json.Unmarshal(byteValue, &cities)
	return &cities, nil
}

func getWeather(address *string, city City, ch chan *CallResponse) error {
	var callResp CallResponse
	start := time.Now()
	endpoint := fmt.Sprintf("%s/weather?city=%s", *address, city.City)
	endpoint = strings.Replace(endpoint, " ", "+", -1)
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Error - %s\n", city.City)
		return err
	}
	defer resp.Body.Close()

	callResp.City = city
	callResp.Status = resp.Status
	callResp.Time = time.Now()
	callResp.Elapsed = time.Since(start)
	fmt.Printf("%s - %s\n", strings.Split(callResp.Status, " "), city.City)

	ch <- &callResp
	// return &callResp, nil
	return nil
}

func pullFromChannel(ch chan *CallResponse) []*CallResponse {
	ss := make([]*CallResponse, 0)
	for s := range ch {
		ss = append(ss, s)
	}
	return ss
}

func main() {
	address := flag.String("addr", "http://127.0.0.1:8080", "Address to test")
	num := flag.Int("n", 500, "Number of cities to test")
	flag.Parse()
	cities, err := loadCities()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	start := time.Now()

	var wg sync.WaitGroup

	partialCities := cities.Cities[0:*num]
	if len(partialCities) < *num {
		*num = len(partialCities)
	}
	fmt.Printf("Testing %d cities\n", *num)

	ch := make(chan *CallResponse)
	wg.Add(*num)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for _, city := range partialCities {
		go func(city City) {
			defer wg.Done()
			getWeather(address, city, ch)
		}(city)
	}

	ss := pullFromChannel(ch)
	// pullFromChannel(ch)
	calls := len(ss)

	end := time.Since(start)
	rate := float64(calls) / end.Seconds()
	fmt.Printf("\n%d calls happened at a rate of %.2f calls/second\n", calls, rate)

}
