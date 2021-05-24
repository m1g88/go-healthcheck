package main

import (
	"fmt"
	"go-healthcheck/healthcheck"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Printf("Perform website checking... \n")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	hch, err := healthcheck.Do()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	hch.SendReport()
	hch.Report.Print()

	fmt.Println("Done!")
}
