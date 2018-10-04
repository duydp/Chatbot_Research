package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://chatbot.baovietnhantho.com.vn/webhook?hub.mode=subscribe&hub.verify_token=baoviet@123"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Loi 1")
	}

	req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("postman-token", "80155789-7843-9123-ac1d-1b23cc4ab1ab")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Loi 2")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Loi 3")
	}

	fmt.Println(res)
	fmt.Println(string(body))

}