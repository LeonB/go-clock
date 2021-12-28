package clock

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewFolioLedgerGetRequest() FolioLedgerGetRequest {
	return FolioLedgerGetRequest{
		client:      c,
		queryParams: c.NewFolioLedgerGetQueryParams(),
		pathParams:  c.NewFolioLedgerGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewFolioLedgerGetRequestBody(),
	}
}

type FolioLedgerGetRequest struct {
	client      *Client
	queryParams *FolioLedgerGetQueryParams
	pathParams  *FolioLedgerGetPathParams
	method      string
	headers     http.Header
	requestBody FolioLedgerGetRequestBody
}

func (c *Client) NewFolioLedgerGetQueryParams() *FolioLedgerGetQueryParams {
	return &FolioLedgerGetQueryParams{}
}

type FolioLedgerGetQueryParams struct {
	ToDate Date `schema:"to_date"`
}

func (p FolioLedgerGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *FolioLedgerGetRequest) QueryParams() *FolioLedgerGetQueryParams {
	return r.queryParams
}

func (c *Client) NewFolioLedgerGetPathParams() *FolioLedgerGetPathParams {
	return &FolioLedgerGetPathParams{
		SubscriptionID: c.SubscriptionID(),
		AccountID:      c.AccountID(),
	}
}

type FolioLedgerGetPathParams struct {
	SubscriptionID int
	AccountID      int
}

func (p *FolioLedgerGetPathParams) Params() map[string]string {
	return map[string]string{
		"subscription_id": strconv.Itoa(p.SubscriptionID),
		"account_id":      strconv.Itoa(p.AccountID),
	}
}

func (r *FolioLedgerGetRequest) PathParams() *FolioLedgerGetPathParams {
	return r.pathParams
}

func (r *FolioLedgerGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *FolioLedgerGetRequest) Method() string {
	return r.method
}

func (s *Client) NewFolioLedgerGetRequestBody() FolioLedgerGetRequestBody {
	return FolioLedgerGetRequestBody{}
}

type FolioLedgerGetRequestBody struct {
}

func (r *FolioLedgerGetRequest) RequestBody() *FolioLedgerGetRequestBody {
	return &r.requestBody
}

func (r *FolioLedgerGetRequest) SetRequestBody(body FolioLedgerGetRequestBody) {
	r.requestBody = body
}

func (r *FolioLedgerGetRequest) NewResponseBody() *FolioLedgerGetResponseBody {
	return &FolioLedgerGetResponseBody{}
}

type FolioLedgerGetResponseBody FolioLedgers

func (r *FolioLedgerGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/base_api/{{.subscription_id}}/{{.account_id}}/folio_ledger.json", r.PathParams())
}

func (r *FolioLedgerGetRequest) Do() (FolioLedgerGetResponseBody, error) {
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
