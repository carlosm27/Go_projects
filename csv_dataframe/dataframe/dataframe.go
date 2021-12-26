package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"github.com/tobgu/qframe"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name of a file with .csv extension or '0.csv' to end the program: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	CsvName := strings.TrimSuffix(input,"\r\n")
	if strings.HasSuffix(CsvName, ".csv") {
		OpenCsvFile(CsvName)
	}
	fmt.Println("Enter the name of a file with .csv extension or anything else to end the program: ")

	for strings.HasSuffix(CsvName, ".csv") {

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}
		CsvName := strings.TrimSuffix(input,"\r\n")
		OpenCsvFile(CsvName)
		fmt.Println("Enter the name of a file with .csv extension or anything else to end the program: ")
	}
	fmt.Println("Bye!")


}

func OpenCsvFile(name string) {
	csvfile, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	dataframe := qframe.ReadCSV(csvfile)
	fmt.Println(dataframe)
}
