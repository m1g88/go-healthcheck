package main

import (
	"encoding/csv"
	"fmt"
	"go-healthcheck/api"
	"go-healthcheck/report"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func getWebsiteCheckList() []string {
	lines, err := readCsv(os.Args[1])
	if err != nil {
		panic(err)
	}

	var checkList []string
	for _, line := range lines {
		checkList = append(checkList, line...)
	}

	return checkList
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

func main() {
	fmt.Printf("Perform website checking... \n")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// accessToken := os.Getenv("AccessToken")

	// fmt.Println(accessToken)

	start := time.Now()
	checkList := getWebsiteCheckList()
	r := report.Report{}

	ch := make(chan api.Response)
	for _, url := range checkList {
		go api.DoChannelRequest(url, ch)
	}

	for range checkList {
		res := <-ch
		if res.IsSuccess() {
			r.Success()
		} else {
			r.Failure()
		}

	}
	// r.Total_time = time.Since(start).Nanoseconds()

	fmt.Println("Done!")

	r.Print()

	fmt.Printf("\n %v  elapsed\n", time.Since(start).Nanoseconds())
}
