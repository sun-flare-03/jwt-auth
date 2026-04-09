package main

import (
	"fmt"
	"log"
	"os"
)

// jwt-auth — Zero-dependency JWT authentication library with RS256 and EdDSA support
func main() {
	logger := log.New(os.Stdout, "[jwt-auth] ", log.LstdFlags)
	logger.Println("Starting application...")

	if err := run(); err != nil {
		logger.Fatalf("Fatal error: %v", err)
	}
}

func run() error {
	fmt.Println("Application initialized successfully")
	return nil
}
