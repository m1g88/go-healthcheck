package main

import (
	"encoding/csv"
	"fmt"
	"go-healthcheck/healthcheck"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func getUrls(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return []string{}, fmt.Errorf("can not convert csv file")
	}

	var checkList []string
	for _, line := range lines {
		checkList = append(checkList, line...)
	}

	return checkList, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: go run main.go test.csv\n")
		os.Exit(0)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Printf("Perform website checking... \n")

	start := time.Now()
	filename := os.Args[1]
	urls, err := getUrls(filename)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	h := &healthcheck.HealthCheck{
		CheckList: urls,
		Config: healthcheck.Config{
			EndpointApi: os.Getenv("HEALTHCEHECK_ENDPOINT"),
			AccessToken: os.Getenv("ACCESS_TOKEN"),
			Timeout:     os.Getenv("TIMEOUT"),
		},
	}

	h.CheckAll()

	elapsed := time.Since(start)
	h.Report.Total_time = elapsed.Nanoseconds()

	if err := h.SendReport(); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Checked webistes: %v\n", h.Report.Total_websites)
	fmt.Printf("Successful websites: %v\n", h.Report.Success)
	fmt.Printf("Failure websites: %v\n", h.Report.Failure)
	fmt.Printf("Total times to finished checking website: %s\n", elapsed)
	fmt.Println("Done!")
}
