package clock_test

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"testing"

	clock "github.com/omniboost/go-clock"
)

func TestFolioIndex(t *testing.T) {
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

	req := client.NewFolioIndexRequest()
	req.QueryParams().Filters["closed_at.gteq"] = "2018-12-06"
	req.QueryParams().Filters["closed_at.lt"] = "2018-12-10"
	req.QueryParams().Filters["payed"] = "false"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	for _, id := range resp {
		req := client.NewFolioShowRequest()
		req.PathParams().ID = id
		resp, err := req.Do()
		if err != nil {
			t.Error(err)
		}

		b, _ := json.MarshalIndent(resp, "", "  ")
		log.Println(string(b))
	}

	os.Exit(21)

	for _, id := range resp {
		req := client.NewFolioChargesRequest()
		req.PathParams().ID = id
		resp, err := req.Do()
		if err != nil {
			t.Error(err)
		}

		b, _ := json.MarshalIndent(resp, "", "  ")
		log.Println(string(b))
		break
	}

	for _, id := range resp {
		req := client.NewFolioCreditsRequest()
		req.PathParams().ID = id
		resp, err := req.Do()
		if err != nil {
			t.Error(err)
		}

		b, _ := json.MarshalIndent(resp, "", "  ")
		log.Println(string(b))
	}
}
