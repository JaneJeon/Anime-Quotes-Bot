package twitter

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"io/ioutil"
)

func Login(credentialsFile string) (*twitter.Client, error) {
	// read the credentials file
	bytes, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		return nil, err
	}

	// convert json->string map
	var cred map[string]string
	if err := json.Unmarshal(bytes, &cred); err != nil {
		return nil, err
	}

	// generate OAuth1 credentials
	config := oauth1.NewConfig(cred["consumerKey"], cred["consumerSecret"])
	token := oauth1.NewToken(cred["accessToken"], cred["accessSecret"])

	return twitter.NewClient(config.Client(oauth1.NoContext, token)), nil
}
