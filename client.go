package clock

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"text/template"
	"time"

	digest "github.com/omniboost/go-clock/digest"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-clock/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "sky-eu1.clock-software.com",
		Path:   "",
	}
	requestTimestamps = make(map[int]*timestamps)
)

// NewClient returns a new InvoiceXpress Client client
func NewClient(httpClient *http.Client, userName string, APIKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		http: httpClient,
	}

	client.SetUserName(userName)
	client.SetAPIKey(APIKey)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with InvoiceXpress Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	userName       string
	apiKey         string
	subscriptionID int
	accountID      int

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) UserName() string {
	return c.userName
}

func (c *Client) SetUserName(userName string) {
	c.userName = userName
	t := digest.NewTransport(c.UserName(), c.APIKey())
	c.http.Transport = &t
}

func (c *Client) APIKey() string {
	return c.apiKey
}

func (c *Client) SetAPIKey(apiKey string) {
	c.apiKey = apiKey
	t := digest.NewTransport(c.UserName(), c.APIKey())
	c.http.Transport = &t
}

func (c *Client) SubscriptionID() int {
	return c.subscriptionID
}

func (c *Client) SetSubscriptionID(subscriptionID int) {
	c.subscriptionID = subscriptionID
}

func (c *Client) AccountID() int {
	return c.accountID
}

func (c *Client) SetAccountID(accountID int) {
	c.accountID = accountID

	if requestTimestamps[accountID] == nil {
		requestTimestamps[accountID] = &timestamps{}
	}
}

func (c *Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c *Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c *Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) GetEndpointURL(path string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()
	clientURL.Path = clientURL.Path + path

	tmpl, err := template.New("endpoint_url").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, method string, URL url.URL, body interface{}) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	req, err := http.NewRequest(method, URL.String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// set other headers
	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	req.Header.Add("Accept", c.MediaType())
	req.Header.Add("User-Agent", c.UserAgent())

	return req, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	c.SleepUntilRequestRate()

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	// register timestamp after request has a response
	c.RegisterRequestTimestamp(time.Now())

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	if responseBody != nil {
		dec := json.NewDecoder(httpResp.Body)
		if c.disallowUnknownFields {
			dec.DisallowUnknownFields()
		}
		err = dec.Decode(responseBody)
		if err != nil && err != io.EOF {
			// create a simple error response
			errorResponse := &ErrorResponse{Response: httpResp}
			errorResponse.Errors = append(errorResponse.Errors, err)
			return httpResp, errorResponse
		}
	}

	return httpResp, nil
}

func (c *Client) RegisterRequestTimestamp(t time.Time) {
	if len(*requestTimestamps[c.AccountID()]) >= 5 {
		ts := (*requestTimestamps[c.AccountID()])[1:5]
		requestTimestamps[c.AccountID()] = &ts
	}
	ts := append(*requestTimestamps[c.AccountID()], t)
	requestTimestamps[c.AccountID()] = &ts
}

func (c *Client) SleepUntilRequestRate() {
	// Requestrate is 5r/1s

	// if there are less then 5 registered requests: execute the request
	// immediately
	if len(*requestTimestamps[c.AccountID()]) < 4 {
		return
	}

	// is the first item within 1 second? If it's > 1 second the request can be
	// executed imediately
	c.RegisterRequestTimestamp(time.Now())
	diff := time.Now().Sub((*requestTimestamps[c.AccountID()])[0])
	if diff >= time.Second {
		return
	}

	// Sleep for the time it takes for the first item to be > 1 second old
	// + 1ms to be sure :)
	time.Sleep(time.Second - diff + (1 * time.Millisecond))
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	err := checkContentType(r)
	if err != nil {
		errorResponse.Errors = append(errorResponse.Errors, errors.New(r.Status))
		return errorResponse
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		errorResponse.Errors = append(errorResponse.Errors, err)
		return errorResponse
	}

	return errorResponse
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Errors []error
}

// {"error":"Addon API not installed!"}

func (r *ErrorResponse) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Error string `json:"error"`
	}{}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	if tmp.Error != "" {
		r.Errors = append(r.Errors, errors.New(tmp.Error))
	}

	return nil
}

func (r ErrorResponse) Error() string {
	if len(r.Errors) > 0 {
		str := []string{}
		for _, err := range r.Errors {
			str = append(str, err.Error())
		}
		return strings.Join(str, ", ")
	}

	switch r.Response.StatusCode {
	case 401:
		return "The Client Key parameter is missing or is incorrectly entered."
	case 404:
		return "The requested resource does not exist."
	case 406:
		return "The :document-id provided is in an invalid state."
	case 422:
		return "Some parameters were incorrect."
	}

	return fmt.Sprintf("Unknown status code %d", r.Response.StatusCode)
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

type PathParams interface {
	Params() map[string]string
}

type FilterParams map[string]string

// type FilterParams []Filter

// type Filter struct {
// 	Field    string
// 	Operator string
// 	Value    string
// }

type timestamps []time.Time
