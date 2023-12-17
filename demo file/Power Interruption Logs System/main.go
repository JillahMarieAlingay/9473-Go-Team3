/*
created by the Go language group
This project is made for demo in order to introduce the Go language
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// creates a schedule interface with the function is_scheduled
type schedule interface {
	is_scheduled() bool
}

// creates an interruption structure with the values date_time and feeder_number
type interruption struct {
	date_time     string
	feeder_number int
}

// function to see if the interruption exists in the list of scheduled interruptions
func (i interruption) is_scheduled(list []interruption) bool {
	return slices.Contains(list, i)
}

// basic import file function, returns a slice of interruptions
func import_list(filename string) []interruption {
	var content []interruption
	// opens the file
	file, err := os.Open(filename)

	// error handler
	if err != nil {
		// if error is encountered, ends the program
		log.Fatal(err)
	}

	defer file.Close()

	// creates a scanner
	scanner := bufio.NewScanner(file)

	// loops until the scanner can't scan anything
	// scans per line and returns a string
	for scanner.Scan() {
		i := strings.Split(scanner.Text(), ",")

		time := i[0]
		feeder, err := strconv.Atoi(i[1])
		if err != nil {
			log.Fatal(err)
		}

		var pHolder interruption
		pHolder = interruption{time, feeder}
		// adds the new interruption to the content slice
		content = append(content, pHolder)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return content
}

// writes down the interuption to a txt file called results.txt
func export_list(i interruption, list []interruption) {
	time := i.date_time
	feeder := i.feeder_number
	is_scheduled := i.is_scheduled(list)

	to_be_printed := time + "," + strconv.Itoa(feeder) + "," + strconv.FormatBool(is_scheduled)

	file, err := os.OpenFile("results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(to_be_printed + "\n")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Interruption Logged Successfully : " + to_be_printed)
}

func main() {
	var test interruption
	test = interruption{"20 Sep 23 12:50", 13}

	var list []interruption
	list = import_list("list_of_schedule.txt")

	export_list(test, list)
}
