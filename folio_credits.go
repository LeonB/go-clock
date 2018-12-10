package clock

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewFolioCreditsRequest() FolioCreditsRequest {
	return FolioCreditsRequest{
		client:      c,
		queryParams: c.NewFolioCreditsQueryParams(),
		pathParams:  c.NewFolioCreditsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewFolioCreditsRequestBody(),
	}
}

type FolioCreditsRequest struct {
	client      *Client
	queryParams *FolioCreditsQueryParams
	pathParams  *FolioCreditsPathParams
	method      string
	headers     http.Header
	requestBody FolioCreditsRequestBody
}

func (c *Client) NewFolioCreditsQueryParams() *FolioCreditsQueryParams {
	return &FolioCreditsQueryParams{}
}

type FolioCreditsQueryParams struct {
}

func (p FolioCreditsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FolioCreditsRequest) QueryParams() *FolioCreditsQueryParams {
	return r.queryParams
}

func (c *Client) NewFolioCreditsPathParams() *FolioCreditsPathParams {
	return &FolioCreditsPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type FolioCreditsPathParams struct {
	SubscriptionID int
	AccountID      int
	ID             int
}

func (p *FolioCreditsPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
		"id":              strconv.Itoa(p.ID),
	}
}

func (r *FolioCreditsRequest) PathParams() *FolioCreditsPathParams {
	return r.pathParams
}

func (r *FolioCreditsRequest) SetMethod(method string) {
	r.method = method
}

func (r *FolioCreditsRequest) Method() string {
	return r.method
}

func (s *Client) NewFolioCreditsRequestBody() FolioCreditsRequestBody {
	return FolioCreditsRequestBody{}
}

type FolioCreditsRequestBody struct {
}

func (r *FolioCreditsRequest) RequestBody() *FolioCreditsRequestBody {
	return &r.requestBody
}

func (r *FolioCreditsRequest) SetRequestBody(body FolioCreditsRequestBody) {
	r.requestBody = body
}

func (r *FolioCreditsRequest) NewResponseBody() *FolioCreditsResponseBody {
	return &FolioCreditsResponseBody{}
}

type FolioCreditsResponseBody FolioCredits

func (r *FolioCreditsRequest) URL() url.URL {
	return r.client.GetEndpointURL("/base_api/{{.subscription_id}}/{{.account_id}}/folios/{{.id}}/credit_items.json", r.PathParams())
}

func (r *FolioCreditsRequest) Do() (FolioCreditsResponseBody, error) {
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
