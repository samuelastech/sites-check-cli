package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const sleepTime = 5 * time.Second
const monitoringTimes = 5

func main() {
	for {
		showMenu()
		option := getInput()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			exit()
		default:
			fmt.Println("Unknown typed option")
			os.Exit(1)
		}
	}
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	urls := readFile()

	for i := 0; i < monitoringTimes; i++ {
		for _, url := range urls {
			checkSite(url)
		}
		time.Sleep(sleepTime)
	}
}

func readFile() []string {
	file, error := os.Open("sites.txt")
	urls := []string{}

	if error != nil {
		fmt.Println(error)
	} else {
		defer file.Close()
		reader := bufio.NewScanner(file)

		for reader.Scan() {
			url := reader.Text()
			urls = append(urls, url)
		}
	}

	return urls
}

func checkSite(url string) {
	response, error := http.Get(url)

	if error != nil {
		fmt.Println(error)
		return
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", url, "was loaded successfully, status code:", response.StatusCode)
	} else {
		fmt.Println("Site:", url, "faced a issue, status code:", response.StatusCode)
	}
}

func showLogs() {
	fmt.Println("Showing logs...")
}

func exit() {
	fmt.Println("Shutting down the program...")
	os.Exit(0)
}

func showMenu() {
	fmt.Println("\n1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Close the program")
}

func getInput() int {
	fmt.Print("Type an option: ")
	var option int
	fmt.Scan(&option)
	return option
}
