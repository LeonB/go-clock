package clock

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewFolioIndexRequest() FolioIndexRequest {
	return FolioIndexRequest{
		client:      c,
		queryParams: c.NewFolioIndexQueryParams(),
		pathParams:  c.NewFolioIndexPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewFolioIndexRequestBody(),
	}
}

type FolioIndexRequest struct {
	client      *Client
	queryParams *FolioIndexQueryParams
	pathParams  *FolioIndexPathParams
	method      string
	headers     http.Header
	requestBody FolioIndexRequestBody
}

func (c *Client) NewFolioIndexQueryParams() *FolioIndexQueryParams {
	return &FolioIndexQueryParams{Filters: FilterParams{}}
}

type FolioIndexQueryParams struct {
	Filters FilterParams `schema:"-"`
}

func (p FolioIndexQueryParams) ToURLValues() (url.Values, error) {
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

func (r *FolioIndexRequest) QueryParams() *FolioIndexQueryParams {
	return r.queryParams
}

func (c *Client) NewFolioIndexPathParams() *FolioIndexPathParams {
	return &FolioIndexPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type FolioIndexPathParams struct {
	SubscriptionID int
	AccountID      int
}

func (p *FolioIndexPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
	}
}

func (r *FolioIndexRequest) PathParams() *FolioIndexPathParams {
	return r.pathParams
}

func (r *FolioIndexRequest) SetMethod(method string) {
	r.method = method
}

func (r *FolioIndexRequest) Method() string {
	return r.method
}

func (s *Client) NewFolioIndexRequestBody() FolioIndexRequestBody {
	return FolioIndexRequestBody{}
}

type FolioIndexRequestBody struct {
}

func (r *FolioIndexRequest) RequestBody() *FolioIndexRequestBody {
	return &r.requestBody
}

func (r *FolioIndexRequest) SetRequestBody(body FolioIndexRequestBody) {
	r.requestBody = body
}

func (r *FolioIndexRequest) NewResponseBody() *FolioIndexResponseBody {
	return &FolioIndexResponseBody{}
}

type FolioIndexResponseBody []int

func (r *FolioIndexRequest) URL() url.URL {
	return r.client.GetEndpointURL("/base_api/{subscription_id}/{account_id}/folios.json", r.PathParams())
}

func (r *FolioIndexRequest) Do() (FolioIndexResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, true)
	if err != nil {
		return nil, err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
