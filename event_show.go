package clock

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewEventShowRequest() EventShowRequest {
	return EventShowRequest{
		client:      c,
		queryParams: c.NewEventShowQueryParams(),
		pathParams:  c.NewEventShowPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewEventShowRequestBody(),
	}
}

type EventShowRequest struct {
	client      *Client
	queryParams *EventShowQueryParams
	pathParams  *EventShowPathParams
	method      string
	headers     http.Header
	requestBody EventShowRequestBody
}

func (c *Client) NewEventShowQueryParams() *EventShowQueryParams {
	return &EventShowQueryParams{}
}

type EventShowQueryParams struct {
}

func (p EventShowQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *EventShowRequest) QueryParams() *EventShowQueryParams {
	return r.queryParams
}

func (c *Client) NewEventShowPathParams() *EventShowPathParams {
	return &EventShowPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type EventShowPathParams struct {
	SubscriptionID int
	AccountID      int
	ID             int
}

func (p *EventShowPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
		"id":              strconv.Itoa(p.ID),
	}
}

func (r *EventShowRequest) PathParams() *EventShowPathParams {
	return r.pathParams
}

func (r *EventShowRequest) SetMethod(method string) {
	r.method = method
}

func (r *EventShowRequest) Method() string {
	return r.method
}

func (s *Client) NewEventShowRequestBody() EventShowRequestBody {
	return EventShowRequestBody{}
}

type EventShowRequestBody struct {
}

func (r *EventShowRequest) RequestBody() *EventShowRequestBody {
	return &r.requestBody
}

func (r *EventShowRequest) SetRequestBody(body EventShowRequestBody) {
	r.requestBody = body
}

func (r *EventShowRequest) NewResponseBody() *EventShowResponseBody {
	return &EventShowResponseBody{}
}

type EventShowResponseBody Event

func (r *EventShowRequest) URL() url.URL {
	return r.client.GetEndpointURL("/pms_api/{{.subscription_id}}/{{.account_id}}/events/{{.id}}.json", r.PathParams())
}

func (r *EventShowRequest) Do() (EventShowResponseBody, error) {
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
