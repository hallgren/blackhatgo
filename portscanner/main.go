package main

import (
	"fmt"
	"net"
	"sort"
)

func main() {
	workChan := make(chan int, 100)
	resultChan := make(chan int)
	workerDoneChan := make(chan struct{}, 10)

	// Setup workers
	for w := 0; w < 10; w++ {
		go worker(workChan, resultChan, workerDoneChan)
	}

	// generate work
	go func() {
		for i := 1; i <= 1024; i++ {
			workChan <- i
		}
		close(workChan)
	}()

	// exit workers
	go func() {
		for y := 0; y < 10; y++ {
			<-workerDoneChan
		}
		// when all workers has exit close the result chan
		close(resultChan)
	}()

	// get result from port scan
	result := []int{}
	for openPort := range resultChan {
		result = append(result, openPort)
	}

	fmt.Println("Open ports")
	sort.Ints(result)
	for _, port := range result {
		fmt.Println(port)
	}
}

func worker(workChan, resultChan chan int, workerDoneChan chan struct{}) {
	for port := range workChan {
		host := fmt.Sprintf("127.0.0.1:%d", port)
		if portOpen(host) {
			resultChan <- port
		}
	}
	workerDoneChan <- struct{}{}
}

func portOpen(host string) bool {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
