package clock_test

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/aodin/date"
	clock "github.com/omniboost/go-clock"
)

func TestFolioLedgerGet(t *testing.T) {
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
	baseURL, err := url.Parse(os.Getenv("CLOCK_BASE_URL"))
	if err != nil {
		t.Error(err)
	}

	client := clock.NewClient(nil, userName, apiKey)
	client.SetDebug(true)
	client.SetSubscriptionID(subscriptionID)
	client.SetAccountID(accountID)
	client.SetDisallowUnknownFields(true)
	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	req := client.NewFolioLedgerGetRequest()

	req.QueryParams().ToDate = clock.Date{date.New(2021, 12, 27)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
