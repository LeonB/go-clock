package clock

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewCompanyShowRequest() CompanyShowRequest {
	return CompanyShowRequest{
		client:      c,
		queryParams: c.NewCompanyShowQueryParams(),
		pathParams:  c.NewCompanyShowPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewCompanyShowRequestBody(),
	}
}

type CompanyShowRequest struct {
	client      *Client
	queryParams *CompanyShowQueryParams
	pathParams  *CompanyShowPathParams
	method      string
	headers     http.Header
	requestBody CompanyShowRequestBody
}

func (c *Client) NewCompanyShowQueryParams() *CompanyShowQueryParams {
	return &CompanyShowQueryParams{}
}

type CompanyShowQueryParams struct {
}

func (p CompanyShowQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CompanyShowRequest) QueryParams() *CompanyShowQueryParams {
	return r.queryParams
}

func (c *Client) NewCompanyShowPathParams() *CompanyShowPathParams {
	return &CompanyShowPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type CompanyShowPathParams struct {
	SubscriptionID int
	AccountID      int
	ID             int
}

func (p *CompanyShowPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
		"id":              strconv.Itoa(p.ID),
	}
}

func (r *CompanyShowRequest) PathParams() *CompanyShowPathParams {
	return r.pathParams
}

func (r *CompanyShowRequest) SetMethod(method string) {
	r.method = method
}

func (r *CompanyShowRequest) Method() string {
	return r.method
}

func (s *Client) NewCompanyShowRequestBody() CompanyShowRequestBody {
	return CompanyShowRequestBody{}
}

type CompanyShowRequestBody struct {
}

func (r *CompanyShowRequest) RequestBody() *CompanyShowRequestBody {
	return &r.requestBody
}

func (r *CompanyShowRequest) SetRequestBody(body CompanyShowRequestBody) {
	r.requestBody = body
}

func (r *CompanyShowRequest) NewResponseBody() *CompanyShowResponseBody {
	return &CompanyShowResponseBody{}
}

type CompanyShowResponseBody Company

func (r *CompanyShowRequest) URL() url.URL {
	return r.client.GetEndpointURL("/base_api/{{.subscription_id}}/{{.account_id}}/companies/{{.id}}.json", r.PathParams())
}

func (r *CompanyShowRequest) Do() (CompanyShowResponseBody, error) {
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
