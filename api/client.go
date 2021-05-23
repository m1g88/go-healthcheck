package api

import (
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	url     string
	success bool
	time    float64
}

func (r *Response) IsSuccess() bool {
	return r.success
}

func DoChannelRequest(url string, ch chan<- Response) {
	start := time.Now()
	client := http.Client{
		Timeout: 4 * time.Second,
	}
	_, err := client.Get(url)

	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed\n", secs)

	success := true
	if err != nil {
		success = false
	}
	ch <- Response{
		url:     url,
		success: success,
		time:    secs,
	}
}
