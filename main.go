package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// Parse command-line arguments
	inputFile := flag.String("input", "ip_addresses", "Input file containing IP addresses")
	flag.Parse()

	startTime := time.Now()

	file, err := os.Open(*inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	uniqueIPs := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	processedCount := 0

	for scanner.Scan() {
		ip := scanner.Text()
		uniqueIPs[ip] = struct{}{}
		processedCount++
		if processedCount%1000000 == 0 {
			fmt.Printf("Processed IPs: %d, Elapsed time: %s\n", processedCount, time.Since(startTime))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("Total processed IPs: %d\n", processedCount)
	fmt.Printf("Number of unique IP addresses: %d\n", len(uniqueIPs))
	fmt.Printf("Total elapsed time: %s\n", elapsed)
}
