package main

import (
	"fmt"
	"strconv"
	"strings"
)

type time_type struct {
	name        string
	start_value int
	max_value   int
}

func parsePart(part string, start_value, max_value int) string {
	result := ""
	if part == "*" {
		for i := 0; i <= max_value; i++ {
			result += strconv.Itoa(i)
			if i < max_value {
				result += " "
			}
		}
		return result
	}
	if strings.Contains(part, ",") {
		comma_splits := strings.Split(part, ",")
		for i, comma_split := range comma_splits {
			result += comma_split
			if i+1 < len(comma_splits) {
				result += " "
			}
		}
		return result
	}
	if strings.Contains(part, "-") {
		hyphen_splits := strings.Split(part, "-")
		integer1, _ := strconv.Atoi(hyphen_splits[0])
		integer2, _ := strconv.Atoi(hyphen_splits[1])
		for i := integer1; i <= integer2; i++ {
			result += strconv.Itoa(i)
			if i < integer2 {
				result += " "
			}
		}
		return result
	}
	if strings.Contains(part, "/") {
		next_value := start_value
		slash_value := strings.Split(part, "/")[1]
		int_slash_value, _ := strconv.Atoi(slash_value)
		for ; next_value <= max_value; next_value += int_slash_value {
			result += strconv.Itoa(next_value)
			if next_value != max_value {
				result += " "
			}
		}
		return result
	}
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
