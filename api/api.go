package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	url          string
	Success      bool
	HttpResponse *http.Response
}

func DoRequest(endpoint string, token string, requestBody string) (*http.Response, error) {

	jsonStr := []byte(requestBody)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resp, nil
}

func DoChannelRequest(url string, timeout string, ch chan Response) {
	duration, err := strconv.Atoi(timeout)
	if err != nil {
		duration = 5
	}

	client := &http.Client{
		Timeout: time.Duration(duration) * time.Second,
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
