package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	content_split := strings.Split(string(content), "\n")
	total_sum := 0
	for _, row := range content_split {
		total_sum += ReadRow(row)
	}
	fmt.Println(total_sum)
}

func ReadRow(row string) int {
	row_split := strings.Split(row, " ")
	did_i_win := DidIWin(row_split)
	shape_point := ShapePointCalcule(row_split)
	return did_i_win + shape_point
}

func DidIWin(row []string) int {
	if row[0] == "A" && row[1] == "X" {
		return 0
	} else if row[0] == "A" && row[1] == "Y" {
		return 3
	} else if row[0] == "A" && row[1] == "Z" {
		return 6
	} else if row[0] == "B" && row[1] == "X" {
		return 0
	} else if row[0] == "B" && row[1] == "Y" {
		return 3
	} else if row[0] == "B" && row[1] == "Z" {
		return 6
	} else if row[0] == "C" && row[1] == "X" {
		return 0
	} else if row[0] == "C" && row[1] == "Y" {
		return 3
	} else if row[0] == "C" && row[1] == "Z" {
		return 6
	}
	return 0
}

func ShapePointCalcule(row []string) int {
	if row[0] == "A" && row[1] == "X" {
		return 3
	} else if row[0] == "A" && row[1] == "Y" {
		return 1
	} else if row[0] == "A" && row[1] == "Z" {
		return 2
	} else if row[0] == "B" && row[1] == "X" {
		return 1
	} else if row[0] == "B" && row[1] == "Y" {
		return 2
	} else if row[0] == "B" && row[1] == "Z" {
		return 3
	} else if row[0] == "C" && row[1] == "X" {
		return 2
	} else if row[0] == "C" && row[1] == "Y" {
		return 3
	} else if row[0] == "C" && row[1] == "Z" {
		return 1
	}
	return 0
}
