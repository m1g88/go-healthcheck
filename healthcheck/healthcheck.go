package healthcheck

import (
	"fmt"
	"go-healthcheck/api"
)

type HealthCheck struct {
	CheckList CheckList
	Config    Config
	Report    Report
}

type Config struct {
	EndpointApi string
	AccessToken string
	Timeout     string
}

type CheckList []string

func (h *HealthCheck) CheckAll() {
	ch := make(chan api.Response)
	for _, url := range h.CheckList {
		go api.DoChannelRequest(url, h.Config.Timeout, ch)
	}

	for range h.CheckList {
		res := <-ch
		if res.Success {
			h.Report.Success++
		} else {
			h.Report.Failure++
		}
		h.Report.Total_websites++
	}
}

func (h *HealthCheck) SendReport() error {
	resp, err := api.DoRequest(h.Config.EndpointApi, h.Config.AccessToken, h.Report.ToJson())
	if err != nil {
		fmt.Println(err.Error())
	}

	if resp.StatusCode == 400 {
		return fmt.Errorf("Report api has failure")
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("Report api status: %v", resp.Status)
	}

	return nil
}
