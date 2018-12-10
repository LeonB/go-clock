package clock

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewUserShowRequest() UserShowRequest {
	return UserShowRequest{
		client:      c,
		queryParams: c.NewUserShowQueryParams(),
		pathParams:  c.NewUserShowPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewUserShowRequestBody(),
	}
}

type UserShowRequest struct {
	client      *Client
	queryParams *UserShowQueryParams
	pathParams  *UserShowPathParams
	method      string
	headers     http.Header
	requestBody UserShowRequestBody
}

func (c *Client) NewUserShowQueryParams() *UserShowQueryParams {
	return &UserShowQueryParams{}
}

type UserShowQueryParams struct {
}

func (p UserShowQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *UserShowRequest) QueryParams() *UserShowQueryParams {
	return r.queryParams
}

func (c *Client) NewUserShowPathParams() *UserShowPathParams {
	return &UserShowPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type UserShowPathParams struct {
	SubscriptionID int
	AccountID      int
	ID             int
}

func (p *UserShowPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
		"id":              strconv.Itoa(p.ID),
	}
}

func (r *UserShowRequest) PathParams() *UserShowPathParams {
	return r.pathParams
}

func (r *UserShowRequest) SetMethod(method string) {
	r.method = method
}

func (r *UserShowRequest) Method() string {
	return r.method
}

func (s *Client) NewUserShowRequestBody() UserShowRequestBody {
	return UserShowRequestBody{}
}

type UserShowRequestBody struct {
}

func (r *UserShowRequest) RequestBody() *UserShowRequestBody {
	return &r.requestBody
}

func (r *UserShowRequest) SetRequestBody(body UserShowRequestBody) {
	r.requestBody = body
}

func (r *UserShowRequest) NewResponseBody() *UserShowResponseBody {
	return &UserShowResponseBody{}
}

type UserShowResponseBody User

func (r *UserShowRequest) URL() url.URL {
	return r.client.GetEndpointURL("/base_api/{{.subscription_id}}/{{.account_id}}/users/{{.id}}.json", r.PathParams())
}

func (r *UserShowRequest) Do() (UserShowResponseBody, error) {
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
