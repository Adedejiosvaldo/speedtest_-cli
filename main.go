package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

// Server URL
const serverURL = "https://go-speed-server.onrender.com"

// Ping the server
func pingServer() {
	start := time.Now()

	resp, err := http.Get(serverURL + "/ping")
	if err != nil {
		fmt.Println("Ping failed:", err)
		return
	}
	defer resp.Body.Close()
	latency := time.Since(start)
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Ping Response: %s (Latency: %v)\n", string(body), latency)
}

func uploadFile(filePath string) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffer to store form data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create form file field
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}

	// Copy file data into form field
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file data:", err)
		return
	}
	writer.Close()
	start := time.Now()

	// Send POST request
	req, err := http.NewRequest("POST", serverURL+"/upload", &requestBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Upload failed:", err)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Upload Response: %s\nUpload Time: %v\n", string(body), elapsed)
}

// Download a file
func downloadFile(savePath string) {
	start := time.Now()

	resp, err := http.Get(serverURL + "/download")
	if err != nil {
		fmt.Println("Download failed:", err)
		return
	}
	defer resp.Body.Close()

	// Create a new file to save the downloaded data
	outFile, err := os.Create(savePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	// Copy response body to file
	size, err := io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	timeElapsed := time.Since(start)
	downloadSpeed := float64(size) / 1024 / 1024 / timeElapsed.Seconds()

	fmt.Printf("Download Complete: %s\nSize: %.2f MB\nSpeed: %.2f MB/s\n", savePath, float64(size)/1024/1024, downloadSpeed)
}

func main() {
	fmt.Println("Testing Speedtest Server...")

	// Test ping
	pingServer()

	// Test file upload
	GenerateLargeFile("test_upload.txt", 20)
	testFile := "test_upload.txt"
	uploadFile(testFile)

	// Test file download
	downloadFile("downloaded_testfile.dat")
}

func GenerateLargeFile(filePath string, sizeInMB int) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a random large file
	data := make([]byte, 1024*1024) // 1MB
	rand.Read(data)
	for i := 0; i < sizeInMB; i++ {
		_, err := file.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}
