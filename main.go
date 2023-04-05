package main

import (
	"fmt"

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

type smsClient struct {
	From   string
	Client *twilio.RestClient
}

func (c smsClient) Send(to string, body string) error {

	params := &api.CreateMessageParams{}
	params.SetFrom(c.From)
	params.SetBody(body)
	params.SetTo(to)
	resp, err := c.Client.Api.CreateMessage(params)
	if err != nil {
		return err
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
	return nil
}
func (c smsClient) List() error {

	params := &api.ListMessageParams{}
	params.SetFrom(c.From)
	//params.SetTo(to)
	params.SetPageSize(20)
	params.SetLimit(100)
	resp, err := c.Client.Api.ListMessage(params)
	if err != nil {
		fmt.Println(err)
	}
	for record := range resp {
		fmt.Println("Body: ", *resp[record].Body)
		fmt.Println("To: ", *resp[record].To)
		fmt.Println("From: ", *resp[record].From)
		fmt.Println("")
	}

	return nil
}

func main() {
	body := "hello world"
	client := twilio.NewRestClient()
	from := "+13853965242"

	sms := &smsClient{
		From:   from,
		Client: client,
	}
	sms.Send("+15032676296", body)
	sms.List()

	//params.SetTo("+15415371994")

	// Set routing rules

	//Use the default DefaultServeMux.
	//err = http.ListenAndServe(":8080", nil)
	//if err != nil {
	//log.Fatal(err)
	//}

	//from := "13853965242"
	//to := "15032676296"

	//client := twilio.NewRestClient()

	////channel, _ := client.Api.StreamMessage(params)
	////for record := range channel {
	////fmt.Println("Body: ", *record.Body)
	////}

}
