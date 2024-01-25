package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	gUnZip()
}

func gZip() {
	newGzip, err := os.Create("new.txt.gz")
	if err != nil {
		log.Fatalf("Failed to create gzip file: %v", err)
	}
	defer newGzip.Close()

	gzipWriter := gzip.NewWriter(newGzip)

	fileToGzip, err := os.Open("./gzip/test2.txt")
	if err != nil {
		log.Fatalf("Failed to open source file: %v", err)
	}
	defer fileToGzip.Close()

	_, err = io.Copy(gzipWriter, fileToGzip)
	if err != nil {
		log.Fatalf("Failed to write to gzip file: %v", err)
	}

	if err := gzipWriter.Close(); err != nil {
		log.Fatalf("Failed to close gzip writer: %v", err)
	}
}

func gUnZip() {
	gzipFileName := "new.txt.gz"  // Should be replace with your gzip file path
	destFileName := "example.txt" // This will be your output file path

	gzipFile, err := os.Open(gzipFileName)
	if err != nil {
		fmt.Println("Error opening gzip file:", err)
		return
	}
	defer gzipFile.Close()

	reader, err := gzip.NewReader(gzipFile)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}
	defer reader.Close()

	destFile, err := os.Create(destFileName)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, reader)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("Successfully decompressed gzip file")
}
