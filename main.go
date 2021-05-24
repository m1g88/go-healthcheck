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

func getCheckList(filename string) ([]string, error) {
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
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "usage: go run main.go test.csv")
		os.Exit(0)
	}

	fmt.Printf("Perform website checking... \n")
	start := time.Now()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	filename := os.Args[1]
	checkList, err := getCheckList(filename)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	hck := healthcheck.Check(checkList)
	hck.SendReport()
	hck.Report.Print()

	fmt.Printf("Total times to finished checking website: %s\n", time.Since(start))
	fmt.Println("Done!")
}
