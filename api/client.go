package api

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

// const Timeout = os.Getenv("TIMEOUT") || "5"

type Response struct {
	url          string
	Success      bool
	HttpResponse *http.Response
}

func DoRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(request)
}

func DoChannelRequest(url string, ch chan Response) {
	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))

	if err != nil {
		timeout = 5
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	resp, err := client.Get(url)

	// Success: any http status code that is returned from the website
	// Failure: cannot reach the website (request timeout)
	success := true
	if err != nil {
		success = false
	}

	ch <- Response{
		url:          url,
		Success:      success,
		HttpResponse: resp,
	}
}
