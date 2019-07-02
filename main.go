package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

//Simple check directly pinging the address provided
func simpleCheck(proxyAddress string) *http.Response {
	response, err := http.Get(proxyAddress)
	if err != nil {
		fmt.Println("Couldn't reach the address")
	}
	return response

}

//Isn't detecting the proxy and getting through, doing a direct connection
//to the checking address, fix this ASAP.
func completeCheck(proxyAddress string, checkingAddress string) *http.Response {
	proxyUrl, err := url.Parse(proxyAddress)
	if err != nil {
		fmt.Println(proxyUrl) // <nil> ...this is the problem... parsing not working?
		fmt.Println("Couldn't reach the proxy")
	}
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	response, err := myClient.Get(checkingAddress)
	if err != nil {
		fmt.Println("Couldn't reach the address")
	}
	return response
}

//Ok, this is the main problem right now
//this should retrieve a string with the Status Code
//of the http.Response assigned to "response"
func statusCheck(response *http.Response) bool {
	if response.StatusCode == http.StatusOK {
		return true
	} else {
		return false
	}
	//This doesn't work for some reason... figure it out for a more
	//specific status checking using the recieved response body
	/*
		var bodyString string
		if response.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(response.Body)
			bodyString := string(bodyBytes)
			if err != nil {
				fmt.Println("Error reading body of the response:", err)
			}
			//statusString := string(response.StatusCode)
		}
		return bodyString
	*/
}

//Fixed, now turn it into a proper .txt checker and writter
func main() {

	var proxyAddress string = "149.56.133.165:3128"
	var checkingAddress string = "http://google.com"
	timerStarted := time.Now()
	//resp := simpleCheck(proxyAddress)
	resp := completeCheck(proxyAddress, checkingAddress)
	chk := statusCheck(resp)
	elapsedTime := time.Since(timerStarted)

	//var checkstatus bool = strings.Contains(response, "200 OK")

	if chk {
		fmt.Printf("200 OK --> Working proxy | %s\n", elapsedTime)
	} else {
		fmt.Printf("Proxy not Working | %d\n", resp.StatusCode)
	}

}
