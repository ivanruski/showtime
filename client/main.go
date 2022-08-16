package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var client *http.Client

func main() {
	host := flag.String("host", "http://server:8080", "Server to ping")

	flag.Parse()

	client = &http.Client{}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	for {
		select {
		case <-sigint:
			return
		default:
			wg := &sync.WaitGroup{}
			for i := 0; i < 20; i++ {
				wg.Add(1)
				go sendRequest(i, *host, wg)
			}
			wg.Wait()
		}
	}
}

func sendRequest(i int, host string, wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequest("GET", host, nil)
	if err != nil {
		log.Printf("error: creating request: %s", err)
	} else {
		res, err := client.Do(req)
		if err != nil {
			log.Printf("error: sending request: %s", err)
		} else {
			data, err := readResponse(res)
			if err != nil {
				log.Printf("error (%d): %s", i, err)
			} else {
				log.Printf("success (%d): %s", i, string(data))
			}
		}
	}
}

func readResponse(r *http.Response) (string, error) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("error: reading response body: %s", err)
	}

	return string(data), nil
}
