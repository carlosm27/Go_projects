package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "encoding/csv"
  "encoding/json"
  "reflect"
)
type Fruits struct {
    Genus       string      `json:"genus"`
    Name        string      `json:"name"`
    Family      string       `json:"family"`
    Order       string       `json:"order"`
    Id          int          `json:"id"`
    Nutrition   []Nutrition  `json:"nutritions"`

}
type Nutrition struct {
    Carbs     int `json: "carbohydrates"`
    Protein   int `json: "protein"`
    Fat       int `json: "fat"`
    Calories  int `json: "calories"`
    Sugar     int `json: "sugar"`


}


func main() {
  response, err := http.Get("https://www.fruityvice.com/api/fruit/apple" )

  if err != nil {
    fmt.Print(err.Error())
    os.Exit(1)
  }

  responseData, err := ioutil.ReadAll(response.Body)

  if err != nil {
    log.Fatal(err)
  }

  var fruitObject map[string]interface{}
  err = json.Unmarshal([]byte(responseData), &fruitObject)

  if err != nil {
    fmt.Println(err)
  }

  csvFile, err := os.Create("./fruits.csv")
  if err != nil {
    fmt.Println(err)
  }
  defer csvFile.Close()

  writer := csv.NewWriter(csvFile)

  for _,v := range fruitObject {
    var row []string
    var s string
    var d string
    var carb string
    var pro string
    var cal string
    var fa string
    var sug string


    c := reflect.ValueOf(v)
    switch c.Kind()  {
    case reflect.String:
      s = v.(string)

    case reflect.Float64:
      d = fmt.Sprintf("%.0f", v)

    case reflect.Map:
      iter := c.MapRange()
      for iter.Next() {
          carbs := c.MapIndex(reflect.ValueOf("carbohydrates"))
          protein := c.MapIndex(reflect.ValueOf("protein"))
          fat := c.MapIndex(reflect.ValueOf("fat"))
          calories := c.MapIndex(reflect.ValueOf("calories"))
          sugar := c.MapIndex(reflect.ValueOf("sugar"))

          carb = fmt.Sprintf("%.1f", carbs)
          pro = fmt.Sprintf("%.1f", protein)
          cal = fmt.Sprintf("%.1f", calories)
          fa = fmt.Sprintf("%.1f", fat)
          sug = fmt.Sprintf("%.1f", sugar)
      }


    }
    row = append(row, d, s, carb, pro, cal, fa, sug)
    fmt.Println(row)


    writer.Write(row)
  }

  writer.Flush()
  fmt.Println("csv file created")
}
