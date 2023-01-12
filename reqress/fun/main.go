package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// The function will perform a http get request and will return the response bytes.
func MakeHttpGet(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error:" + err.Error())
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println("Error:" + err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error:" + err.Error())
	}
	return body, err
}

// func MakeHttpPost(url string, request interface{}) ([]byte, error) {
// 	jsonBody, err := json.Marshal(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

// 	req.Header.Add("Content-Type", "application/json")
// 	response, err := client.Do(req)
// 	defer response.Body.Close()
// 	body, err := ioutil.ReadAll(response.Body)
// 	return body, err
// }

func main() {
	MakeHttpGet("https://reqres.in/apdh")
	fmt.Println("Macruz")

}
