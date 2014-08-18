package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ae0000/goxero"
)

func main() {
	// Put your consumer key and secret here... or in env.. or something..
	consumerKey := "YOUR_KEY"
	consumerSecret := "YOUR_SECRET"

	token := "MQNPQVFWF3C9TFZPVGFSF8X82XOWMP"
	secret := "7QX8M5MYIEC3RV39BANEFSMHZAEZZZ"
	addData := map[string]string{
		"oauth_expires_in": "1800",
		"xero_org_muid":    "pAfdKchMCaYF0JTjIR6dqG",
	}

	// Create the consumer
	c := goxero.NewConsumer(consumerKey, consumerSecret)

	// Add the accesstoken
	c.SetAccessToken(token, secret, addData)

	response, err := c.Get(
		"invoices/1234",
		map[string]string{"count": "1"})
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	fmt.Println("The invoice is: " + string(bits))

}
