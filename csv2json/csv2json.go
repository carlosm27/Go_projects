package main

import (
  "encoding/csv"
  "encoding/json"
  "errors"
  "flag"
  "fmt"
  "io"
  "os"
  "path/filepath"
  "strings"
)
func main() {
  flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <csvFile>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	// Getting the file data that was entered by the user
	fileData, err := getFileData()

	if err != nil {
		exitGracefully(err)
	}
	// Validating the file entered
	if _, err := checkIfValidFile(fileData.filepath); err != nil {
		exitGracefully(err)
	}
	// Declaring the channels that our go-routines are going to use
	writerChannel := make(chan map[string]string)
	done := make(chan bool)
	// Running both of our go-routines, the first one responsible for reading and the second one for writing
	go processCsvFile(fileData, writerChannel)
	go writeJSONFile(fileData.filepath, writerChannel, done, fileData.pretty)
	// Waiting for the done channel to receive a value, so that we can terminate the programn execution
	<-done

}

type inputFile struct {
  filepath string
  separator string
  pretty bool
}

func getFileData() (inputFile, error) {
  if len(os.Args) < 2 {
    return inputFile{}, errors.New("A filepath argument is required")
  }

  separator := flag.String("separator", "comma", "Column separator")
  pretty := flag.Bool("pretty", false, "Generate pretty Json")

  flag.Parse()

  fileLocation := flag.Arg(0)

  if !(*separator == "comma" || *separator == "semicolon") {
      return inputFile{}, errors.New("Only comma or semicolon separators are allowed")

  }
  return inputFile{fileLocation, *separator, *pretty},nil
}

func checkIfValidFile (filename string) (bool, error) {
  if fileExtension := filepath.Ext(filename); fileExtension != ".csv" {
    return false, fmt.Errorf("File %s is not CSV", filename)

  }

  if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
    return false, fmt.Errorf("File %s does not exist", filename)
  }
  return true, nil
}

func processCsvFile (fileData inputFile,writerChannel chan <- map[string]string) {
    file, err := os.Open(fileData.filepath)

    check(err)
    defer file.Close()

    var headers, line []string

    reader := csv.NewReader(file)

    if fileData.separator == "semicolon" {
		reader.Comma = ';'
	 }

    headers, err = reader.Read()
    check(err)

    for {
      line, err = reader.Read()
      if err == io.EOF {
        close(writerChannel)
        break
      } else if err != nil {
        exitGracefully(err)
      }
      record, err := processLine(headers, line)
      if err != nil {
        fmt.Printf("Line: %sError: %s\n", line, err)
        continue
      }
      writerChannel <- record
    }

}

func exitGracefully(err error) {
  fmt.Fprintf(os.Stderr,"error: %v\n", err)
  os.Exit(1)
}

func check(e error) {
  if e != nil {
    exitGracefully(e)
  }
}

func processLine(headers [] string, dataList []string) (map[string]string, error) {
  if len(dataList) != len(headers) {
    return nil, errors.New("Line doesn't match headers format.Skipping")

  }
  recordMap := make(map[string]string)
  for i, name := range headers {
    recordMap[name] = dataList[i]
  }
  return recordMap, nil
}

func writeJSONFile(csvPath string, writerChannel <- chan map[string]string, done chan <- bool, pretty bool) {
    writeString := createStringWriter(csvPath)
    jsonFunc, breakLine := getJSONFunc(pretty)

    fmt.Println("Writer JSON file...")

    writeString("[" +breakLine, false)
    first := true

    for {
        record, more := <- writerChannel
        if more {
          if !first {
              writeString("," +breakLine, false)
          } else {
              first = false
          }

          jsonData := jsonFunc(record)
          writeString(jsonData, false)
        } else {
          writeString(breakLine+"]", true)
          fmt.Println("Completed!")
          done <- true
          break
        }
    }
}

func createStringWriter(csvPath string) func(string, bool) {
    jsonDir := filepath.Dir(csvPath)
    jsonName := fmt.Sprintf("%s.json", strings.TrimSuffix(filepath.Base(csvPath), ".csv"))

    finalLocation := filepath.Join(jsonDir, jsonName)

    f, err := os.Create(finalLocation)
    check(err)
    return func( data string, close bool) {
      _, err := f.WriteString(data)
      check(err)
      if close {
          f.Close()
      }
    }
}

func getJSONFunc(pretty bool) (func (map[string]string)string, string) {

    var jsonFunc func (map[string]string) string
    var breakLine string

    if pretty {
      breakLine = "\n"
      jsonFunc = func (record map[string]string) string {
        jsonData, _ := json.MarshalIndent(record, "   ", "   ")
        return "   " + string(jsonData)
      }
    } else {
        breakLine = ""
        jsonFunc = func(record map[string]string) string {
          jsonData, _:= json.Marshal(record)
          return string(jsonData)
        }
    }
    return jsonFunc,breakLine
}
