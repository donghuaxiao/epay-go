package main

import (
	"epay"
	"fmt"
	"log"
)

func main() {
	//	fmt.Println("hello world")
	//	resp, err := http.Get("http://www.baidu.com/")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	defer resp.Body.Close()
	//
	//	body, err := ioutil.ReadAll(resp.Body)
	//	fmt.Println(string(body))

	client := epay.NewEpayClient("localhost", 4005)
	resp, err := client.GetPaymentOrder("30", "75635544476038")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
