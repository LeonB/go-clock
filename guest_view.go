package clock

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewGuestViewRequest() GuestViewRequest {
	return GuestViewRequest{
		client:      c,
		queryParams: c.NewGuestViewQueryParams(),
		pathParams:  c.NewGuestViewPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGuestViewRequestBody(),
	}
}

type GuestViewRequest struct {
	client      *Client
	queryParams *GuestViewQueryParams
	pathParams  *GuestViewPathParams
	method      string
	headers     http.Header
	requestBody GuestViewRequestBody
}

func (c *Client) NewGuestViewQueryParams() *GuestViewQueryParams {
	return &GuestViewQueryParams{}
}

type GuestViewQueryParams struct {
}

func (p GuestViewQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GuestViewRequest) QueryParams() *GuestViewQueryParams {
	return r.queryParams
}

func (c *Client) NewGuestViewPathParams() *GuestViewPathParams {
	return &GuestViewPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type GuestViewPathParams struct {
	SubscriptionID int
	AccountID      int
	ID             int
}

func (p *GuestViewPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
		"id":              strconv.Itoa(p.ID),
	}
}

func (r *GuestViewRequest) PathParams() *GuestViewPathParams {
	return r.pathParams
}

func (r *GuestViewRequest) SetMethod(method string) {
	r.method = method
}

func (r *GuestViewRequest) Method() string {
	return r.method
}

func (s *Client) NewGuestViewRequestBody() GuestViewRequestBody {
	return GuestViewRequestBody{}
}

type GuestViewRequestBody struct {
}

func (r *GuestViewRequest) RequestBody() *GuestViewRequestBody {
	return &r.requestBody
}

func (r *GuestViewRequest) SetRequestBody(body GuestViewRequestBody) {
	r.requestBody = body
}

func (r *GuestViewRequest) NewResponseBody() *GuestViewResponseBody {
	return &GuestViewResponseBody{}
}

type GuestViewResponseBody Guest

func (r *GuestViewRequest) URL() url.URL {
	return r.client.GetEndpointURL("/pms_api/{{.subscription_id}}/{{.account_id}}/guests/{{.id}}.json", r.PathParams())
}

func (r *GuestViewRequest) Do() (GuestViewResponseBody, error) {
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
