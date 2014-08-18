package goxero

import (
	"net/http"

	"github.com/mrjones/oauth"
)

const (
	XeroRequestTokenURL   = "https://api.xero.com/oauth/RequestToken"
	XeroAuthorizeTokenURL = "https://api.xero.com/oauth/Authorize"
	XeroAccessTokenURL    = "https://api.xero.com/oauth/AccessToken"
	XeroApiEndpointURL    = "https://api.xero.com/api.xro/2.0/"
)

type XeroConsumer struct {
	Consumer    *oauth.Consumer
	AccessToken *oauth.AccessToken
}

func NewConsumer(consumerKey, consumerSecret string) *XeroConsumer {

	x := XeroConsumer{}
	x.Consumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   XeroRequestTokenURL,
			AuthorizeTokenUrl: XeroAuthorizeTokenURL,
			AccessTokenUrl:    XeroAccessTokenURL,
		})
	return &x
}

func (x *XeroConsumer) RefreshToken(accessToken *oauth.AccessToken) (atoken *oauth.AccessToken, err error) {
	return x.Consumer.RefreshToken(accessToken)
}

func (x *XeroConsumer) SetAccessToken(token, secret string, data map[string]string) {
	a := CreateAccessToken(token, secret, data)
	x.AccessToken = &a
}

func (x *XeroConsumer) Get(path string, userParams map[string]string) (resp *http.Response, err error) {
	return x.Consumer.Get(
		XeroApiEndpointURL+path,
		userParams,
		x.AccessToken)
}

func CreateAccessToken(token, secret string, data map[string]string) oauth.AccessToken {
	return oauth.AccessToken{
		Token:          token,
		Secret:         secret,
		AdditionalData: data,
	}
}
