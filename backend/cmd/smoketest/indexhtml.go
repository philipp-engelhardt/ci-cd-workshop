package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type IndexHtmlTest struct {
}

func (test *IndexHtmlTest) GetTestName() string {
	return "Frontend - Verify presence of index.html"
}

func (test *IndexHtmlTest) RunTest(baseUrl string) bool {
	resp, err := http.Get(baseUrl + "/")
	if err != nil {
		fmt.Printf("Error: Failed to send request to root path: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Expected status 200, got %d\n", resp.StatusCode)
		return false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read response body: %v\n", err)
		return false
	}

	if !containsIndexHTML(string(body)) {
		fmt.Println("Error: index.html not found in the root path response")
		return false
	}

	return true
}

// Helper function to check if index.html is present in the response body
func containsIndexHTML(body string) bool {
	return body != "" && strings.Contains(strings.ToLower(body), "<!doctype html>") && strings.Contains(body, "<html")
}
