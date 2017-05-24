package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func curlURL(url string, headers map[string]string, outputFile string) {
	bodyText := getRequestAsString(url, headers)
	f, err := os.Create(outputFile) // os.O_APPEND|
	check(err)
	defer f.Close()
	if _, err := f.WriteString(bodyText); err != nil {
		panic(err)
	}
	f.Sync()
}

func getRequestAsString(url string, headers map[string]string) string {
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	check(err)

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("Did not receive 200 status code for url %s. Received: %d %s", url, resp.StatusCode, resp.Status))
	}

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return string(body)
}
