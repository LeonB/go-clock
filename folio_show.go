package clock

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewFolioShowRequest() FolioShowRequest {
	return FolioShowRequest{
		client:      c,
		queryParams: c.NewFolioShowQueryParams(),
		pathParams:  c.NewFolioShowPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewFolioShowRequestBody(),
	}
}

type FolioShowRequest struct {
	client      *Client
	queryParams *FolioShowQueryParams
	pathParams  *FolioShowPathParams
	method      string
	headers     http.Header
	requestBody FolioShowRequestBody
}

func (c *Client) NewFolioShowQueryParams() *FolioShowQueryParams {
	return &FolioShowQueryParams{}
}

type FolioShowQueryParams struct {
}

func (p FolioShowQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FolioShowRequest) QueryParams() *FolioShowQueryParams {
	return r.queryParams
}

func (c *Client) NewFolioShowPathParams() *FolioShowPathParams {
	return &FolioShowPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type FolioShowPathParams struct {
	SubscriptionID int
	AccountID      int
	ID             int
}

func (p *FolioShowPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
		"id":              strconv.Itoa(p.ID),
	}
}

func (r *FolioShowRequest) PathParams() *FolioShowPathParams {
	return r.pathParams
}

func (r *FolioShowRequest) SetMethod(method string) {
	r.method = method
}

func (r *FolioShowRequest) Method() string {
	return r.method
}

func (s *Client) NewFolioShowRequestBody() FolioShowRequestBody {
	return FolioShowRequestBody{}
}

type FolioShowRequestBody struct {
}

func (r *FolioShowRequest) RequestBody() *FolioShowRequestBody {
	return &r.requestBody
}

func (r *FolioShowRequest) SetRequestBody(body FolioShowRequestBody) {
	r.requestBody = body
}

func (r *FolioShowRequest) NewResponseBody() *FolioShowResponseBody {
	return &FolioShowResponseBody{}
}

type FolioShowResponseBody Folio

func (r *FolioShowRequest) URL() url.URL {
	return r.client.GetEndpointURL("/base_api/{{.subscription_id}}/{{.account_id}}/folios/{{.id}}.json", r.PathParams())
}

func (r *FolioShowRequest) Do() (FolioShowResponseBody, error) {
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
