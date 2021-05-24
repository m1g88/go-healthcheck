package healthcheck

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"go-healthcheck/api"
	"io/ioutil"
	"net/http"
	"os"
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

func (h *Healthcheck) SendReport() error {
	endpoint := os.Getenv("HEALTHCEHECK_ENDPOINT")
	token := os.Getenv("ACCESS_TOKEN")

	var jsonStr = []byte(h.Report.ToJson())

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr))

	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := api.DoRequest(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	fmt.Println()

	return nil
}

func Do() (*Healthcheck, error) {
	url, err := getWebsiteCheckList()
	start := time.Now()
	if err != nil {
		return nil, err
	}

	h := Healthcheck{Url: url}

	h.CheckAll()
	h.Report.Total_time = time.Since(start).Nanoseconds()
	return &h, nil
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
