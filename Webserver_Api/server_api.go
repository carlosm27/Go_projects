package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "html/template"
)

type Response struct {
    Name  string  `json:"name"`
    Nutrition    []Nutrition    `json:"nutritions"`
}

type Nutrition struct {
    Carbs int `json: "carbohydrates"`
    Protein int `json: "protein"`
    Fat int `json: "fat"`
    Calories int `json: "calories"`
    Sugar int `json: "sugar"`

}

func main() {

    http.Handle("/", http.FileServer(http.Dir("./static")))

    log.Fatal(http.ListenAndServe(":8081", nil))
}

func RetrievingJson() {

  response, err := http.Get("https://www.fruityvice.com/api/fruit/all")

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
