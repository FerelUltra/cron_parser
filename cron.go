package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type timeType struct {
	name       string
	startValue int
	maxValue   int
}

func parsePart(part string, startValue, maxValue int) (string, error) {
	values := []string{}

	if part == "*" {
		for i := startValue; i <= maxValue; i++ {
			values = append(values, strconv.Itoa(i))
		}
		return strings.Join(values, " "), nil
	}

	if strings.Contains(part, ",") {
		commaSplits := strings.Split(part, ",")
		return strings.Join(commaSplits, " "), nil
	}

	if strings.Contains(part, "-") {
		hyphenSplits := strings.Split(part, "-")

		start, err := strconv.Atoi(hyphenSplits[0])
		if err != nil {
			return "", fmt.Errorf("parsePart: failed to convert start_value: %w", err)
		}

		end, err := strconv.Atoi(hyphenSplits[1])
		if err != nil {
			return "", fmt.Errorf("parsePart: failed to convert end_value: %w", err)
		}

		for i := start; i <= end; i++ {
			values = append(values, strconv.Itoa(i))
		}
		return strings.Join(values, " "), nil
	}

	if strings.Contains(part, "/") {
		slashSplits := strings.Split(part, "/")
		step, err := strconv.Atoi(slashSplits[1])
		if err != nil {
			return "", fmt.Errorf("parsePart: failed to convert step value: %w", err)
		}

		for i := startValue; i <= maxValue; i += step {
			values = append(values, strconv.Itoa(i))
		}
		return strings.Join(values, " "), nil
	}

	// Если просто одно значение, возвращаем его
	return part, nil
}

func printPart(width, index, startValue, maxValue int, name string, parts []string) error {
	if name == "command" {
		fmt.Printf("%-*s %s\n", width, name, parts[index])
	} else {
		parsed, err := parsePart(parts[index], startValue, maxValue)
		if err != nil {
			return fmt.Errorf("printPart: error parsing %q: %w", name, err)
		}
		fmt.Printf("%-*s %s\n", width, name, parsed)
	}
	return nil
}

func cron(str string) error {
	width := -15
	parts := strings.Split(str, " ")
	timeSlice := [6]timeType{
		{"minute", 0, 59},
		{"hour", 0, 23},
		{"day of month", 1, 31},
		{"month", 1, 12},
		{"day of week", 0, 6},
		{"command", 0, 0},
	}

	for index, t := range timeSlice {
		err := printPart(width, index, t.startValue, t.maxValue, t.name, parts)
		if err != nil {
			return fmt.Errorf("cron: failed at %q: %w", t.name, err)
		}
	}

	return nil
}

func main() {
	err := cron("*/15 0 1,15 * 1-5 /usr/bin/find")
	if err != nil {
		log.Fatal(err)
	}
}
