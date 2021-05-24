package healthcheck

import (
	"fmt"
	"go-healthcheck/api"
	"time"
)

type Healthcheck struct {
	Url    []string
	Report Report
}

func (h *Healthcheck) CheckAll() {
	ch := make(chan api.Response)
	for _, url := range h.Url {
		go api.DoChannelRequest(url, ch)
	}

	for range h.Url {
		res := <-ch
		if res.Success {
			h.Report.Success++
		} else {
			h.Report.Failure++
		}
		h.Report.Total_websites++
	}
}

func (h *Healthcheck) SendReport() {
	resp, err := api.DoRequest(h.Report.ToJson())
	if err != nil {
		fmt.Println(err.Error())
	}

	if resp.StatusCode == 400 {
		fmt.Println("Report api has failure")
	} else if resp.StatusCode != 200 {
		fmt.Printf("Report api status: %v", resp.Status)
	}
}

func Check(url []string) *Healthcheck {
	start := time.Now()
	h := Healthcheck{Url: url}

	h.CheckAll()
	h.Report.Total_time = time.Since(start).Nanoseconds()
	return &h
}
