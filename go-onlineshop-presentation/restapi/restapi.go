package restapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type HttpClientFactory func() *http.Client

type BaseAPIResponse struct {
	Success bool `json:"success"`
	Result  json.RawMessage
	Message string
}

type StringResponse struct {
	BaseAPIResponse
	Result string
}

type BoolResponse struct {
	BaseAPIResponse
	Result bool
}

type GenericMapResponse struct {
	BaseAPIResponse
	Result map[string]interface{}
}

type HttpError struct {
	error
	StatusCode int
}

type RestClientMode uint32

type RestClient struct {
	Service         string
	Client          *http.Client
	Headers         map[string]string
	SourceHeader    string
	ResponseHeaders http.Header
}

func GetNewRestClient(service string, httpFactory HttpClientFactory) (*RestClient, error) {
	jar, err := cookiejar.New(nil)

	if err != nil {
		return nil, err
	}

	url, err := url.Parse(service)
	if err != nil {
		return nil, err
	}
	url.Scheme = "https"
	url.Path = ""

	client := &RestClient{}
	client.Service = url.String()
	if httpFactory != nil {
		client.Client = httpFactory()
	} else {
		client.Client = &http.Client{}
	}
	client.Client.Jar = jar
	client.Headers = make(map[string]string)
	client.SourceHeader = "cloud-golang-sdk"
	return client, nil
}

func (r *RestClient) CallRawAPI(method string, args map[string]interface{}) ([]byte, error) {
	return r.postAndGetBody(method, args)
}

func (r *RestClient) CallBaseAPI(method string, args map[string]interface{}) (*BaseAPIResponse, error) {
	body, err := r.postAndGetBody(method, args)
	if err != nil {
		return nil, err
	}
	return bodyToBaseAPIResponse(body)
}

func (r *RestClient) CallGenericMapAPI(method string, args map[string]interface{}) (*GenericMapResponse, error) {
	body, err := r.postAndGetBody(method, args)
	if err != nil {
		return nil, err
	}
	return bodyToGenericMapResponse(body)
}

func (r *RestClient) CallStringAPI(method string, args map[string]interface{}) (*StringResponse, error) {
	body, err := r.postAndGetBody(method, args)
	if err != nil {
		return nil, err
	}
	return bodyToStringResponse(body)
}

func (r *RestClient) CallBoolAPI(method string, args map[string]interface{}) (*BoolResponse, error) {
	body, err := r.postAndGetBody(method, args)
	if err != nil {
		return nil, err
	}
	return bodyToBoolResponse(body)
}

func (r *RestClient) postAndGetBody(method string, args map[string]interface{}) ([]byte, error) {
	service := strings.TrimSuffix(r.Service, "/")
	method = strings.TrimPrefix(method, "/")
	postdata := strings.NewReader(payloadFromMap(args))
	postreq, err := http.NewRequest("POST", service+"/"+method, postdata)

	if err != nil {
		return nil, err
	}

	postreq.Header.Add("Content-Type", "application/json")
	postreq.Header.Add("X-CENTRIFY-NATIVE-CLIENT", "Yes")
	postreq.Header.Add("X-CFY-SRC", r.SourceHeader)

	for k, v := range r.Headers {
		postreq.Header.Add(k, v)
	}

	httpresp, err := r.Client.Do(postreq)
	if err != nil {
		r.ResponseHeaders = nil
		return nil, err
	}

	defer httpresp.Body.Close()

	r.ResponseHeaders = httpresp.Header

	if httpresp.StatusCode == 200 {
		return ioutil.ReadAll(httpresp.Body)
	}

	body, _ := ioutil.ReadAll(httpresp.Body)
	return nil, &HttpError{error: fmt.Errorf("POST to %s failed with code %d, body: %s", method, httpresp.StatusCode, body), StatusCode: httpresp.StatusCode}
}

func (r *RestClient) GetLastResponseHeaders() http.Header {
	return r.ResponseHeaders
}

func payloadFromMap(input map[string]interface{}) string {
	if input != nil {
		p, _ := json.Marshal(input)
		return string(p)
	}

	return ""
}

func bodyToBaseAPIResponse(body []byte) (*BaseAPIResponse, error) {
	reply := &BaseAPIResponse{}
	err := json.Unmarshal(body, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func bodyToGenericMapResponse(body []byte) (*GenericMapResponse, error) {
	reply := &GenericMapResponse{}
	err := json.Unmarshal(body, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func bodyToStringResponse(body []byte) (*StringResponse, error) {
	reply := &StringResponse{}
	err := json.Unmarshal(body, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func bodyToBoolResponse(body []byte) (*BoolResponse, error) {
	reply := &BoolResponse{}
	err := json.Unmarshal(body, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// /https://github.com/centrify/cloud-golang-sdk/blob/ed5f25b01f45967d8347753053eb3be704afdae1/sample-app/sample-app.go
func GetRestClient(host string) (*RestClient, error) {
	restClient, err := GetNewRestClient(host, nil)
	if err != nil {
		return nil, err
	}

	restClient.SourceHeader = "golang-sdk-sample"

	return restClient, nil
}
