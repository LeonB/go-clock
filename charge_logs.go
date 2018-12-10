package clock

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (c *Client) NewChargeLogsRequest() ChargeLogsRequest {
	return ChargeLogsRequest{
		client:      c,
		queryParams: c.NewChargeLogsQueryParams(),
		pathParams:  c.NewChargeLogsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewChargeLogsRequestBody(),
	}
}

type ChargeLogsRequest struct {
	client      *Client
	queryParams *ChargeLogsQueryParams
	pathParams  *ChargeLogsPathParams
	method      string
	headers     http.Header
	requestBody ChargeLogsRequestBody
}

func (c *Client) NewChargeLogsQueryParams() *ChargeLogsQueryParams {
	return &ChargeLogsQueryParams{}
}

type ChargeLogsQueryParams struct {
}

func (p ChargeLogsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ChargeLogsRequest) QueryParams() *ChargeLogsQueryParams {
	return r.queryParams
}

func (c *Client) NewChargeLogsPathParams() *ChargeLogsPathParams {
	return &ChargeLogsPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type ChargeLogsPathParams struct {
	SubscriptionID int
	AccountID      int
	Date           time.Time
}

func (p *ChargeLogsPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
		"date":            p.Date.Format("2006-01-02"),
	}
}

func (r *ChargeLogsRequest) PathParams() *ChargeLogsPathParams {
	return r.pathParams
}

func (r *ChargeLogsRequest) SetMethod(method string) {
	r.method = method
}

func (r *ChargeLogsRequest) Method() string {
	return r.method
}

func (s *Client) NewChargeLogsRequestBody() ChargeLogsRequestBody {
	return ChargeLogsRequestBody{}
}

type ChargeLogsRequestBody struct {
}

func (r *ChargeLogsRequest) RequestBody() *ChargeLogsRequestBody {
	return &r.requestBody
}

func (r *ChargeLogsRequest) SetRequestBody(body ChargeLogsRequestBody) {
	r.requestBody = body
}

func (r *ChargeLogsRequest) NewResponseBody() *ChargeLogsResponseBody {
	return &ChargeLogsResponseBody{}
}

type ChargeLogsResponseBody FolioCharges

func (r *ChargeLogsRequest) URL() url.URL {
	return r.client.GetEndpointURL("/base_api/{{.subscription_id}}/{{.account_id}}/{{.date}}/charge_logs.json", r.PathParams())
}

func (r *ChargeLogsRequest) Do() (ChargeLogsResponseBody, error) {
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
