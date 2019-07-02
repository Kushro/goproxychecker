package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// TODO:
// ROUND MS TO 2 DECIMALS
// OPEN A FILE, READ ITS CONTENTS AND DO A QUEUE/LIST AND THEN CHECK

func main() {

	var proxy string = "http://google.com"
	timerStarted := time.Now()
	resp, err := http.Get(proxy)
	elapsedTime := time.Since(timerStarted)

	var response string = fmt.Sprintf("%v", resp)
	if err != nil {
		fmt.Println("Couldn't reach the address")
	}
	var checkstatus bool = strings.Contains(response, "200 OK")

	if checkstatus {
		fmt.Printf("200 OK --> Working proxy with a speed of %s\n", elapsedTime)
	} else {
		fmt.Println("Unknown Response --> Unknown proxy state")
	}

}
