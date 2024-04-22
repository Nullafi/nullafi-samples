package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const API_URL = "http://api_url_here.com/api/scan-dynamic"
const API_KEY = "your_api_key_here"

func testUpload(fileName string, contentType string) {
	// Path to the file to upload

	filePath := fmt.Sprintf("../sample-files/%s", fileName)

	// Create a new file upload request
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// Add file to upload
	fileWriter, err := bodyWriter.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		log.Fatalf("Error creating form file: %v", err)
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		log.Fatalf("Error copying file data: %v", err)
	}

	bodyWriter.Close()

	// Construct query parameters using url.Values
	queryParams := url.Values{}
	queryParams.Set("namespace", "go.apiclient.upload.sample")
	queryParams.Set("obfuscatedDataTypes", "PHONE_NUMBER,EMAIL_ADDRESS")
	queryParams.Set("maskFormats", "LEAVE_LAST_FOUR,FULLY_OBFUSCATED")
	queryParams.Set("username", "golang-sample-user")
	queryParams.Set("usergroup", "golang-sample-usergroup")
	queryParams.Set("storeOriginalValues", "true")

	// Construct the full URL with query parameters
	urlWithParams := API_URL + "?" + queryParams.Encode()

	// Create a new HTTP request with authorization header and query strings
	req, err := http.NewRequest("POST", urlWithParams, bodyBuf)
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	// Add authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", API_KEY))
	req.Header.Set("Content-Type", contentType)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body Length: %v\n", len(respBody))

	// Save the response body as a file
	err = os.WriteFile(fmt.Sprintf("response-files/response-%s", fileName), respBody, 0644)
	if err != nil {
		log.Fatalf("Error saving response file: %v", err)
	}
	fmt.Printf("Response saved as response-%s", fileName)
}

func main() {
	testUpload("sample-test.xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	testUpload("sample-test.docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
}
