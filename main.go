package main

import (
	"fmt"

	"github.com/adedejiosvaldo/terminal_speedtest/cmd"
)

func main() {
	fmt.Println("Testing Speedtest Server...")

	// // Test ping
	// pingServer()

	// // Test file upload
	// helpers.GenerateLargeFile("test_upload.txt", 20)
	// testFile := "test_upload.txt"
	// uploadFile(testFile)

	// // Test file download
	// downloadFile("downloaded_testfile.dat")

	cmd.Execute()
}
