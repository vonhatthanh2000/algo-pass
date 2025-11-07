package main

import (
	"fmt"
	"net"
)

func findAvailablePort(start, end int) (int, error) {
	for port := start; port <= end; port++ {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			ln.Close() // Close the listener if successful
			return port, nil
		}
	}
	return 0, fmt.Errorf("no available ports found in range %d-%d", start, end)
}

func main() {
	port, err := findAvailablePort(20000, 30000)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Available port found: %d\n", port)
	}
}
