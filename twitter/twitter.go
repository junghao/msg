package twitter

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
	"strconv"
)

type Twitter struct {
	api *anaconda.TwitterApi
}

var (
	consumerKey    = os.Getenv("TWITTER_CKEY")
	consumerSecret = os.Getenv("TWITTER_CSECRET")
	oauthToken     = os.Getenv("TWITTER_OSTOKEN")
	oauthSecret    = os.Getenv("TWITTER_OSECRET")
	minMagnitude   = os.Getenv("TWITTER_THRESHOLD")
)

func Init() (Twitter, error) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(oauthToken, oauthSecret)

	var err error

	t := Twitter{
		api: api,
	}

	if api.Credentials == nil {
		err = fmt.Errorf("Credentials are invalid")
	} else {

		err = nil
	}
	return t, err
}

// Note: Twitter will rejects messages longer than 140 chars.
func (a *Twitter) PostTweet(message string, longitude float64, latitude float64) (err error) {
	if a.api.Credentials == nil {
		return fmt.Errorf("Credentials are invalid, cannot post.")
	}

	v := url.Values{}
	v.Set("long", strconv.FormatFloat(longitude, 'f', -1, 64))
	v.Set("lat", strconv.FormatFloat(latitude, 'f', -1, 64))

	_, err = a.api.PostTweet(message, v)

	return err
}
