package main

import (
	"fmt"
	"log"

	"github.com/Attendly/goxero"
)

func main() {

	// Put your consumer key and secret here...
	consumerKey := "SX43AOAYVY8KD3HBJM3DPILPM0AGWV"
	consumerSecret := "R7M97TDVMWBYTL2F7XJZ66ZTFCLY7K"

	c := goxero.NewConsumer(consumerKey, consumerSecret)

	requestToken, url, err := c.Consumer.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("(1) Go to: " + url)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Println("(3) Enter that verification code here: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := c.Consumer.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Below is your access token:")
	fmt.Printf("%+v\n", accessToken)

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	// Your accessToken should look something like this:
	//    &{
	//    Token:FPX1OHEGVR0FBFV4Q3DCODPBO7V3JM
	//    Secret:QLWICIUCH92QDB7EMQC8QSK2UX648B
	//    AdditionalData:map[oauth_expires_in:1800 xero_org_muid:pAfdKcvMCaYF0JTjIR6dqG]
	//    }
}
