package clock

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewUserIndexRequest() UserIndexRequest {
	return UserIndexRequest{
		client:      c,
		queryParams: c.NewUserIndexQueryParams(),
		pathParams:  c.NewUserIndexPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewUserIndexRequestBody(),
	}
}

type UserIndexRequest struct {
	client      *Client
	queryParams *UserIndexQueryParams
	pathParams  *UserIndexPathParams
	method      string
	headers     http.Header
	requestBody UserIndexRequestBody
}

func (c *Client) NewUserIndexQueryParams() *UserIndexQueryParams {
	return &UserIndexQueryParams{Filters: FilterParams{}}
}

type UserIndexQueryParams struct {
	Filters FilterParams `schema:"-"`
}

func (p UserIndexQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	for k, v := range p.Filters {
		params[k] = []string{v}
	}

	return params, nil
}

func (r *UserIndexRequest) QueryParams() *UserIndexQueryParams {
	return r.queryParams
}

func (c *Client) NewUserIndexPathParams() *UserIndexPathParams {
	return &UserIndexPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type UserIndexPathParams struct {
	SubscriptionID int
	AccountID      int
}

func (p *UserIndexPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
	}
}

func (r *UserIndexRequest) PathParams() *UserIndexPathParams {
	return r.pathParams
}

func (r *UserIndexRequest) SetMethod(method string) {
	r.method = method
}

func (r *UserIndexRequest) Method() string {
	return r.method
}

func (s *Client) NewUserIndexRequestBody() UserIndexRequestBody {
	return UserIndexRequestBody{}
}

type UserIndexRequestBody struct {
}

func (r *UserIndexRequest) RequestBody() *UserIndexRequestBody {
	return &r.requestBody
}

func (r *UserIndexRequest) SetRequestBody(body UserIndexRequestBody) {
	r.requestBody = body
}

func (r *UserIndexRequest) NewResponseBody() *UserIndexResponseBody {
	return &UserIndexResponseBody{}
}

type UserIndexResponseBody Users

func (r *UserIndexRequest) URL() url.URL {
	return r.client.GetEndpointURL("/base_api/{{.subscription_id}}/{{.account_id}}/users.json", r.PathParams())
}

func (r *UserIndexRequest) Do() (UserIndexResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return nil, err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
