package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type doer interface {
	doStuff()
}

type myTest struct{}

func (myTest) doStuff() {
	fmt.Println("Doing some stuff...")
}

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

	io.Copy(os.Stdout, resp.Body)

	t := myTest{}
	doSomething(t)
}

func doSomething(d doer) {
	d.doStuff()
}

func printBody(resp http.Response) {
	bodyData := make([]byte, 99999)
	resp.Body.Read(bodyData)
	fmt.Println("Response:", string(bodyData))
}
