package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bodyData := make([]byte, 99999)
	// resp.Body.Read(bodyData)
	// fmt.Println("Response:", string(bodyData))
	printBody(*resp)
}

func printBody(resp http.Response) {
	bodyData := make([]byte, 99999)
	resp.Body.Read(bodyData)
	fmt.Println("Response:", string(bodyData))
}
