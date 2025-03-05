package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type time_type struct {
	name        string
	start_value int
	max_value   int
}

func parsePart(part string, start_value, max_value int) string {
	values := []string{}
	if part == "*" {
		for i := start_value; i <= max_value; i++ {
			values = append(values, strconv.Itoa(i))
		}
		return strings.Join(values, " ")
	}
	if strings.Contains(part, ",") {
		comma_splits := strings.Split(part, ",")
		for _, comma_split := range comma_splits {
			values = append(values, comma_split)
		}
		return strings.Join(values, " ")
	}
	if strings.Contains(part, "-") {
		hyphen_splits := strings.Split(part, "-")
		start_value, err := strconv.Atoi(hyphen_splits[0])
		if err != nil {
			log.Fatalf("Сonversion error start_value: %v", err)
		}

		end_value, err := strconv.Atoi(hyphen_splits[1])
		if err != nil {
			log.Fatalf("Сonversion error end_value: %v", err)
		}
		for i := start_value; i <= end_value; i++ {
			values = append(values, strconv.Itoa(i))
		}
		return strings.Join(values, " ")
	}
	if strings.Contains(part, "/") {
		next_value := start_value
		slash_value := strings.Split(part, "/")[1]
		int_slash_value, err := strconv.Atoi(slash_value)
		if err != nil {
			log.Fatalf("Сonversion error int_slash_value: %v", err)
		}
		for ; next_value <= max_value; next_value += int_slash_value {
			values = append(values, strconv.Itoa(next_value))
		}
		return strings.Join(values, " ")
	}
	// if just one value
	return part
}

func printPart(width, index, start_value, max_value int, name string, parts []string) {
	if name == "command" {
		fmt.Printf("%-*s %s\n", width, name, parts[index])
	} else {
		fmt.Printf("%-*s %s\n", width, name, parsePart(parts[index], start_value, max_value))
	}
}

func cron(str string) {
	width := -15
	parts := strings.Split(str, " ")
	time_slice := [6]time_type{
		{"minute", 0, 59},
		{"hour", 0, 23},
		{"day of month", 1, 31},
		{"month", 1, 12},
		{"day of week", 0, 6},
		{"command", 0, 0},
	}
	for index, time_type := range time_slice {
		printPart(width, index, time_type.start_value, time_type.max_value, time_type.name, parts)
	}
}

func main() {
	cron("*/15 0 1,15 * 1-5 /usr/bin/find")
}
