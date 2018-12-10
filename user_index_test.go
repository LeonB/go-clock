package clock_test

import (
	"os"
	"strconv"
	"testing"

	clock "github.com/omniboost/go-clock"
)

func TestUserIndex(t *testing.T) {
	userName := os.Getenv("CLOCK_USER_NAME")
	apiKey := os.Getenv("CLOCK_API_KEY")
	subscriptionID, err := strconv.Atoi(os.Getenv("CLOCK_SUBSCRIPTION_ID"))
	if err != nil {
		t.Error(err)
	}
	accountID, err := strconv.Atoi(os.Getenv("CLOCK_ACCOUNT_ID"))
	if err != nil {
		t.Error(err)
	}

	client := clock.NewClient(nil, userName, apiKey)
	client.SetDebug(true)
	client.SetSubscriptionID(subscriptionID)
	client.SetAccountID(accountID)
	client.SetDisallowUnknownFields(true)

	req := client.NewUserIndexRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	for _, user := range resp {
		req := client.NewUserShowRequest()
		req.PathParams().ID = user.ID
		_, err := req.Do()
		if err != nil {
			t.Error(err)
		}
		break
	}
}
