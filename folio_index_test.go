package clock_test

import (
	"log"
	"os"
	"strconv"
	"testing"

	clock "github.com/omniboost/go-clock"
)

func TestClientsCreate(t *testing.T) {
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

	req := client.NewFolioIndexRequest()
	req.QueryParams().Filters["id.gt"] = "7071185"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}
