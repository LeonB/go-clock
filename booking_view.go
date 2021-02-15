package clock

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewBookingViewRequest() BookingViewRequest {
	return BookingViewRequest{
		client:      c,
		queryParams: c.NewBookingViewQueryParams(),
		pathParams:  c.NewBookingViewPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewBookingViewRequestBody(),
	}
}

type BookingViewRequest struct {
	client      *Client
	queryParams *BookingViewQueryParams
	pathParams  *BookingViewPathParams
	method      string
	headers     http.Header
	requestBody BookingViewRequestBody
}

func (c *Client) NewBookingViewQueryParams() *BookingViewQueryParams {
	return &BookingViewQueryParams{}
}

type BookingViewQueryParams struct {
}

func (p BookingViewQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *BookingViewRequest) QueryParams() *BookingViewQueryParams {
	return r.queryParams
}

func (c *Client) NewBookingViewPathParams() *BookingViewPathParams {
	return &BookingViewPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type BookingViewPathParams struct {
	SubscriptionID int
	AccountID      int
	ID             int
}

func (p *BookingViewPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
		"id":              strconv.Itoa(p.ID),
	}
}

func (r *BookingViewRequest) PathParams() *BookingViewPathParams {
	return r.pathParams
}

func (r *BookingViewRequest) SetMethod(method string) {
	r.method = method
}

func (r *BookingViewRequest) Method() string {
	return r.method
}

func (s *Client) NewBookingViewRequestBody() BookingViewRequestBody {
	return BookingViewRequestBody{}
}

type BookingViewRequestBody struct {
}

func (r *BookingViewRequest) RequestBody() *BookingViewRequestBody {
	return &r.requestBody
}

func (r *BookingViewRequest) SetRequestBody(body BookingViewRequestBody) {
	r.requestBody = body
}

func (r *BookingViewRequest) NewResponseBody() *BookingViewResponseBody {
	return &BookingViewResponseBody{}
}

type BookingViewResponseBody Booking

func (r *BookingViewRequest) URL() url.URL {
	return r.client.GetEndpointURL("/pms_api/{{.subscription_id}}/{{.account_id}}/bookings/{{.id}}.json", r.PathParams())
}

func (r *BookingViewRequest) Do() (BookingViewResponseBody, error) {
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
