package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"github.com/go-gota/gota/dataframe"
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


}
func OpenCsvFile(name string) {
    CsvFile, err := os.Open(name)
    if err != nil {
      log.Fatal(err)
    }

    df := dataframe.ReadCSV(CsvFile)
    fmt.Println(df)
}
