package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "bufio"
  "strings"
)


func main () {

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Insert your favorite fruit")

  input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

  fruit := strings.TrimSuffix(input, "\r\n")
  response, err := http.Get("https://www.fruityvice.com/api/fruit/" + fruit)

  if err != nil {
    fmt.Print(err.Error())
    os.Exit(1)
  }

  responseData, err := ioutil.ReadAll(response.Body)

  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(responseData))


}
