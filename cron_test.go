package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal("Error creating pipe:", err)
	}
	os.Stdout = w

	f()
	err = w.Close()
	if err != nil {
		log.Fatal("Error closing Pipe:", err)
	}
	os.Stdout = old

	var buf bytes.Buffer
	_, errBuffer := buf.ReadFrom(r) // ✅ Read all data into buf

	if errBuffer != nil {
		fmt.Println("Error:", errBuffer)
	}
	return buf.String()
}

func TestCronIntegration(t *testing.T) {
	input := "*/15 0 1,15 * 1-5 /user/bin/find"
	expectedOutput := fmt.Sprintf(
		`%-15s %s
%-15s %s
%-15s %s
%-15s %s
%-15s %s
%-15s %s
`,
		"minute", "0 15 30 45",
		"hour", "0",
		"day of month", "1 15",
		"month", "1 2 3 4 5 6 7 8 9 10 11 12",
		"day of week", "1 2 3 4 5",
		"command", "/user/bin/find",
	)

	actualOutput := captureOutput(func() {
		cron(input)
	})

	expectedOutput = strings.TrimSpace(expectedOutput)
	actualOutput = strings.TrimSpace(actualOutput)
	expectedOutput = strings.ReplaceAll(expectedOutput, "\r\n", "\n")
	actualOutput = strings.ReplaceAll(actualOutput, "\r\n", "\n")

	if strings.TrimSpace(actualOutput) != strings.TrimSpace(expectedOutput) {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expectedOutput, actualOutput)
	}
}
