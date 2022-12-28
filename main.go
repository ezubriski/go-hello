package main

import (
	"fmt"

	//"strings"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

/*
Desired features:
* Listen to incoming response
* Process responses as commands
Command example:
send PHONE_NUMBER text
send --times number PHONE NUMBER

*/

func main() {
	//body := strings.Repeat("♨", 255) fmt.Println(body)
	//client := twilio.NewRestClient()

	//params := &api.CreateMessageParams{}
	//params.SetFrom("+13853965242")
	//params.SetBody(body)
	//params.SetTo("+15415371994")

	//resp, err := client.Api.CreateMessage(params)
	//if err != nil {
	//fmt.Println(err.Error())
	//} else {
	//if resp.Sid != nil {
	//fmt.Println(*resp.Sid)
	//} else {
	//fmt.Println(resp.Sid)
	//}
	//}

	from := "13853965242"
	to := "15032676296"

	client := twilio.NewRestClient()

	params := &api.ListMessageParams{}
	params.SetFrom(from)
	params.SetTo(to)
	params.SetPageSize(20)
	params.SetLimit(100)

	resp, err := client.Api.ListMessage(params)
	if err != nil {
		fmt.Println(err)
	}
	for record := range resp {
		fmt.Println("Body: ", *resp[record].Body)
		fmt.Println("To: ", *resp[record].To)
		fmt.Println("From: ", *resp[record].From)
		fmt.Println("")
	}

	//channel, _ := client.Api.StreamMessage(params)
	//for record := range channel {
	//fmt.Println("Body: ", *record.Body)
	//}
}
