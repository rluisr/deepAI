package main

import (
	"deepAI"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	c := deepAI.Client{
		Token: os.Getenv("DEEPAI_API"),
	}

	fmt.Println("DetectWithURL")
	nudityResponse, err := c.DetectWithURL("https://img2.thejournal.ie/inline/2442982/original/?width=516&version=2442982")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n\n", nudityResponse)

	fmt.Println("DetectWithFile")
	data, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		panic(err)
	}
	nudityResponse, err = c.DetectWithFile(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", nudityResponse)
}
