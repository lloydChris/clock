package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSecondMessage(t *testing.T) {

	// Setup
	var userFlags = flags{false, 10, 5}
	secondMessage := make(chan string, 1)
	minuteMessage := make(chan string, 1)
	hourMessage := make(chan string, 1)

	input := "second:s"
	inputScanner := bufio.NewScanner(strings.NewReader(input))

	//Test
	go getMessages(inputScanner, userFlags, secondMessage, minuteMessage, hourMessage)

	result := <-secondMessage
	if result != "s" {
		t.Errorf(result, " does not equal ", "s")
	}
}

func TestMinuteMessage(t *testing.T) {

	// Setup
	var userFlags = flags{false, 10, 5}
	secondMessage := make(chan string, 1)
	minuteMessage := make(chan string, 1)
	hourMessage := make(chan string, 1)

	input := "minute:m"
	inputScanner := bufio.NewScanner(strings.NewReader(input))

	//Test
	go getMessages(inputScanner, userFlags, secondMessage, minuteMessage, hourMessage)

	result := <-minuteMessage
	if result != "m" {
		t.Errorf(result, " does not equal ", "m")
	}

}

func TestHourMessage(t *testing.T) {

	// Setup
	var userFlags = flags{false, 10, 5}
	secondMessage := make(chan string, 1)
	minuteMessage := make(chan string, 1)
	hourMessage := make(chan string, 1)

	input := "hour:h"
	inputScanner := bufio.NewScanner(strings.NewReader(input))

	//Test
	go getMessages(inputScanner, userFlags, secondMessage, minuteMessage, hourMessage)

	result := <-hourMessage
	if result != "h" {
		t.Errorf(result, " does not equal ", "h")
	}

}
