package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

type flags struct {
	verbose    bool
	longChime  int
	shortChime int
}

func runClock(userFlags flags, secondMessage <-chan string, minuteMessage <-chan string, hourMessage <-chan string) {
	ticker := time.NewTicker(time.Second)
	count := 0

	//Accept the defaults, but mostly declare the message strings
	printableSecond := <-secondMessage
	printableMinute := <-minuteMessage
	printableHour := <-hourMessage

	for {
		select {
		case printableSecond = <-secondMessage:
		case printableMinute = <-minuteMessage:
		case printableHour = <-hourMessage:
		case t := <-ticker.C:
			count++

			verboseOutput := ""
			if userFlags.verbose {
				verboseOutput = fmt.Sprint(" - ", count, " - ", t)
			}

			if count%userFlags.longChime == 0 {
				fmt.Println(printableHour, verboseOutput)
			} else if count%userFlags.shortChime == 0 {
				fmt.Println(printableMinute, verboseOutput)
			} else {
				fmt.Println(printableSecond, verboseOutput)
			}
		}
	}
}

//This extra function is strictly to make the getMessages function testable
func listenForMessages(userFlags flags, secondMessage chan<- string, minuteMessage chan<- string, hourMessage chan<- string) {
	for {
		inputScanner := bufio.NewScanner(os.Stdin)
		getMessages(inputScanner, userFlags, secondMessage, minuteMessage, hourMessage)
	}
}

func getMessages(inputScanner *bufio.Scanner, userFlags flags, secondMessage chan<- string, minuteMessage chan<- string, hourMessage chan<- string) {
	inputScanner.Scan()
	input := inputScanner.Text()

	if len(input) > 7 && input[:7] == "second:" {
		if userFlags.verbose {
			fmt.Println("NEW SECOND MESSAGE: ", input[7:])
		}
		secondMessage <- input[7:]
	} else if len(input) > 7 && input[:7] == "minute:" {
		if userFlags.verbose {
			fmt.Println("NEW MINUTE MESSAGE: ", input[7:])
		}
		minuteMessage <- input[7:]
	} else if len(input) > 5 && input[:5] == "hour:" {
		if userFlags.verbose {
			fmt.Println("NEW HOUR MESSAGE: ", input[5:])
		}
		hourMessage <- input[5:]
	}
}

func main() {

	verbose := flag.Bool("verbose", false, "Add clock counter and timestamps to the output")
	longChime := flag.Int("longChime", 3600, "Long interval to chime")
	shortChime := flag.Int("shortChime", 60, "Short interval to chime")

	flag.Parse()

	var userFlags = flags{*verbose, *longChime, *shortChime}

	if userFlags.verbose {
		fmt.Print("User input: ")
		fmt.Println(userFlags)
	}

	//set Defaults
	secondMessage := make(chan string, 1)
	secondMessage <- "tick"

	minuteMessage := make(chan string, 1)
	minuteMessage <- "tock"

	hourMessage := make(chan string, 1)
	hourMessage <- "dong"

	go runClock(userFlags, secondMessage, minuteMessage, hourMessage)
	go listenForMessages(userFlags, secondMessage, minuteMessage, hourMessage)

	time.Sleep(3 * time.Hour)

}
