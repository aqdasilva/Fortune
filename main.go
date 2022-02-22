package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var fortuneList = []string{"Fortunes.txt"}

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}

func teller(link chan<- string) {
	for _, m := range fortuneList {
		link <- m
	}
	close(link)
}

func consumer(link <-chan string, done chan<- bool) {
	for b := range link {
		fmt.Println(b)
	}
	done <- true
}

func main() {
	link := make(chan string)
	done := make(chan bool)
	go teller(link)
	go consumer(link, done)
	<-done
	// Loop over lines in file.
	for index, line := range LinesInFile(`Fortunes.txt`) {
		fmt.Printf("Index = %v, line = %v\n", index, line)
	}

	// Get count of lines.
	lines := LinesInFile(`Fortunes.txt`)
	fmt.Println(len(lines))

	randsource := rand.NewSource(time.Now().UnixNano())
	randgenerator := rand.New(randsource)
	firstLoc := randgenerator.Intn(10)

	candidate1 := ""

	dat, err := ioutil.ReadFile("Fortunes.txt")
	if err == nil {
		ascii := string(dat)
		splt := strings.Split(ascii, "\n")
		candidate1 = splt[firstLoc]

	}
	fmt.Println(candidate1)

	fmt.Println("Do you want your fortune told? (Y/N) ?")
	var answer string
	fmt.Scanln(&answer)

	if answer == "y" {
		fmt.Printf("Your Fortune: %v\n", candidate1)
	} else {
		fmt.Println("Learn how to type")
	}

}
