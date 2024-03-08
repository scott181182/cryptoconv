package main

import (
	"io"
	"os"
	"testing"
)

func TestParseJson(t *testing.T) {
	file, openErr := os.Open("./testdata/crypto.json")
	if openErr != nil {
		t.Fatal("Failed to open test file")
	}

	filedata, readErr := io.ReadAll(file)
	if readErr != nil {
		t.Fatalf("Error reading test file:\n%+v", readErr)
	}

	rates, err := parseExchangeJson(filedata)
	if err != nil {
		t.Fatalf("Error parsing JSON:\n%+v", err)
	}

	// You wouldn't normally want to check equality with floats, but this is a test with known inputs.
	if rates["BTC"] != 0.0000149123003886 {
		t.Fatalf("Unexpected BTC rate: '%v'", rates["BTC"])
	}
	if rates["ETH"] != 0.0002578286445046 {
		t.Fatalf("Unexpected ETH rate: '%v'", rates["BTC"])
	}
}
