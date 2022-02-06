package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	url := getCmd.String("url", "No url", "Api Url")

	postCmd := flag.NewFlagSet("post", flag.ExitOnError)

	postUrl := postCmd.String("url", "", "port of your localhost")
	

	reader := bufio.NewReader(os.Stdin)

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'post' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, url)
		response := HandleGet(getCmd, url)
		fmt.Println(response)
	case "post":
		fmt.Println("Enter body:")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}

		fmt.Println("This is the body that have been posted:")
		fmt.Println(HandlePost(postCmd, postUrl, input))
	default:

	}

}

func HandleGet(getCmd *flag.FlagSet, url *string) string {
	getCmd.Parse(os.Args[2:])

	if *url == "" {
		fmt.Print("url is required")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	urlResponse := GetUrl(*url)

	return string(urlResponse)

}

func HandlePost(postCmd *flag.FlagSet, postUrl *string, postString string) string {

	postCmd.Parse(os.Args[2:])

	if *postUrl == "" || postString == "" {
		fmt.Println("All fields required")
		postCmd.PrintDefaults()
		os.Exit(1)
	}
	post := Poster(*postUrl, postString)
	return string(post)

}
