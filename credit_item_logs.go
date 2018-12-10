package clock

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (c *Client) NewCreditItemLogsRequest() CreditItemLogsRequest {
	return CreditItemLogsRequest{
		client:      c,
		queryParams: c.NewCreditItemLogsQueryParams(),
		pathParams:  c.NewCreditItemLogsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewCreditItemLogsRequestBody(),
	}
}

type CreditItemLogsRequest struct {
	client      *Client
	queryParams *CreditItemLogsQueryParams
	pathParams  *CreditItemLogsPathParams
	method      string
	headers     http.Header
	requestBody CreditItemLogsRequestBody
}

func (c *Client) NewCreditItemLogsQueryParams() *CreditItemLogsQueryParams {
	return &CreditItemLogsQueryParams{}
}

type CreditItemLogsQueryParams struct {
}

func (p CreditItemLogsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreditItemLogsRequest) QueryParams() *CreditItemLogsQueryParams {
	return r.queryParams
}

func (c *Client) NewCreditItemLogsPathParams() *CreditItemLogsPathParams {
	return &CreditItemLogsPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type CreditItemLogsPathParams struct {
	SubscriptionID int
	AccountID      int
	Date           time.Time
}

func (p *CreditItemLogsPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
		"date":            p.Date.Format("2006-01-02"),
	}
}

func (r *CreditItemLogsRequest) PathParams() *CreditItemLogsPathParams {
	return r.pathParams
}

func (r *CreditItemLogsRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreditItemLogsRequest) Method() string {
	return r.method
}

func (s *Client) NewCreditItemLogsRequestBody() CreditItemLogsRequestBody {
	return CreditItemLogsRequestBody{}
}

type CreditItemLogsRequestBody struct {
}

func (r *CreditItemLogsRequest) RequestBody() *CreditItemLogsRequestBody {
	return &r.requestBody
}

func (r *CreditItemLogsRequest) SetRequestBody(body CreditItemLogsRequestBody) {
	r.requestBody = body
}

func (r *CreditItemLogsRequest) NewResponseBody() *CreditItemLogsResponseBody {
	return &CreditItemLogsResponseBody{}
}

type CreditItemLogsResponseBody FolioCredits

func (r *CreditItemLogsRequest) URL() url.URL {
	return r.client.GetEndpointURL("/base_api/{{.subscription_id}}/{{.account_id}}/{{.date}}/credit_item_logs.json", r.PathParams())
}

func (r *CreditItemLogsRequest) Do() (CreditItemLogsResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
