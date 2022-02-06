package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetUrl(url string) string {

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)

}
