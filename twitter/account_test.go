package twitter

import (
	"testing"
)

// the Twitter handle of your bot
const expectedScreenName = "ComradeMoe"

func TestLogin(t *testing.T) {
	client, err := Login("../credentials.json")
	check(t, err)

	// Verify credentials
	user, _, err := client.Accounts.VerifyCredentials(nil)
	check(t, err)

	if name := user.ScreenName; name != expectedScreenName {
		t.Errorf("Twitter handle should be %s, but is %s instead", expectedScreenName, name)
	}
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
