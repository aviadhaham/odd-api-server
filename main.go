package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/aviadhaham/odd-api-server/internal/api"
)

func main() {
	if os.Getenv("PORT") == "" {
		log.Fatal("PORT env var is not set")
	}

	oddFilePath := "/tmp"
	oddFileName := "odd-logs.txt"
	oddFile, err := os.Create(filepath.Join(oddFilePath, oddFileName))
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer oddFile.Close()

	log.Printf("Successfully created file: %s", oddFileName)
	s := api.NewServer(os.Getenv("PORT"), oddFile)
	s.Run()
}
