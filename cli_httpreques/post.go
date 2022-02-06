package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Poster(url string, postString string) string {

	var jsonMap map[string]interface{}

	json.Unmarshal([]byte(postString), &jsonMap)

	postBody, _ := json.Marshal(jsonMap)

	responseBody := bytes.NewBuffer(postBody)

	response, err := http.Post(url, "aplication/json", responseBody)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	sb := string(responseData)
	return sb

}
