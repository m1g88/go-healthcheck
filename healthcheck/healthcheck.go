package healthcheck

import (
	"encoding/csv"
	"fmt"
	"go-healthcheck/api"
	"os"
)

// accessToken := os.Getenv("AccessToken")

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
		if res.IsSuccess() {
			h.Report.Success++
		} else {
			h.Report.Failure++
		}
		h.Report.Total_websites++
	}
}

func Do() (*Healthcheck, error) {
	url, err := getWebsiteCheckList()
	if err != nil {
		return nil, err
	}

	h := Healthcheck{Url: url}

	h.CheckAll()

	return &h, nil
}

func SendReport() {

}

func getWebsiteCheckList() ([]string, error) {
	lines, err := readCsv(os.Args[1])
	if err != nil {
		return nil, fmt.Errorf("can not convert csv file")
	}

	var checkList []string
	for _, line := range lines {
		checkList = append(checkList, line...)
	}

	return checkList, nil
}

func readCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
