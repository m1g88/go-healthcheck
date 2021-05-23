package main

import (
	"fmt"
	"go-healthcheck/healthcheck"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Printf("Perform website checking... \n")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	start := time.Now()
	// checkList := getWebsiteCheckList()

	checked, err := healthcheck.Do()

	if err != nil {
		fmt.Println(err.Error())
	}
	// r.Total_time = time.Since(start).Nanoseconds()

	fmt.Println("Done!")

	checked.Report.Print()

	fmt.Printf("\n %v  elapsed\n", time.Since(start).Nanoseconds())
}
