package main

import (
    "net/http"
    "time"
)

type bodyMessage struct {
    Website string
    TimesToHit int
}

func executePerformanceTest(url string, timesToHit int, l *logger) int64 {
	var totalTime time.Duration
	l.Println("[DEBUG] Executing performance test")
	c := make(chan time.Duration)
	for i := 0; i < timesToHit; i++ {
		go measure(url, c, l)
	}
	l.Println("[DEBUG] Goroutines triggered")
	for i := 0; i < timesToHit; i++ {
		l.Println("[DEBUG] Reading channel")
		totalTime = totalTime + <- c
	}
	return int64(totalTime / time.Millisecond) / int64(timesToHit)

}

func measure(url string, c chan time.Duration, l *logger){
	start := time.Now()
	l.Println("[DEBUG] Performing request")
	http.Get(url)
	l.Println("[DEBUG] Request finished succesfully")
	c <- time.Since(start)
	l.Println("[DEBUG] Result written in channel")
}