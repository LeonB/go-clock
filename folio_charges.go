package clock

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewFolioChargesRequest() FolioChargesRequest {
	return FolioChargesRequest{
		client:      c,
		queryParams: c.NewFolioChargesQueryParams(),
		pathParams:  c.NewFolioChargesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewFolioChargesRequestBody(),
	}
}

type FolioChargesRequest struct {
	client      *Client
	queryParams *FolioChargesQueryParams
	pathParams  *FolioChargesPathParams
	method      string
	headers     http.Header
	requestBody FolioChargesRequestBody
}

func (c *Client) NewFolioChargesQueryParams() *FolioChargesQueryParams {
	return &FolioChargesQueryParams{}
}

type FolioChargesQueryParams struct {
}

func (p FolioChargesQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FolioChargesRequest) QueryParams() *FolioChargesQueryParams {
	return r.queryParams
}

func (c *Client) NewFolioChargesPathParams() *FolioChargesPathParams {
	return &FolioChargesPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type FolioChargesPathParams struct {
	SubscriptionID int
	AccountID      int
	ID             int
}

func (p *FolioChargesPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
		"id":              strconv.Itoa(p.ID),
	}
}

func (r *FolioChargesRequest) PathParams() *FolioChargesPathParams {
	return r.pathParams
}

func (r *FolioChargesRequest) SetMethod(method string) {
	r.method = method
}

func (r *FolioChargesRequest) Method() string {
	return r.method
}

func (s *Client) NewFolioChargesRequestBody() FolioChargesRequestBody {
	return FolioChargesRequestBody{}
}

type FolioChargesRequestBody struct {
}

func (r *FolioChargesRequest) RequestBody() *FolioChargesRequestBody {
	return &r.requestBody
}

func (r *FolioChargesRequest) SetRequestBody(body FolioChargesRequestBody) {
	r.requestBody = body
}

func (r *FolioChargesRequest) NewResponseBody() *FolioChargesResponseBody {
	return &FolioChargesResponseBody{}
}

type FolioChargesResponseBody FolioCharges

func (r *FolioChargesRequest) URL() url.URL {
	return r.client.GetEndpointURL("/base_api/{{.subscription_id}}/{{.account_id}}/folios/{{.id}}/charges.json", r.PathParams())
}

func (r *FolioChargesRequest) Do() (FolioChargesResponseBody, error) {
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
