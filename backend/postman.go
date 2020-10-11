package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"

	dashboard "github.com/aclisp/godashboard/proto"
)

const (
	// ProductionEnvironment key
	ProductionEnvironment = "product"
	// TestingEnvironment key
	TestingEnvironment = "testing"
	// India key
	India = "IN"
	// Indonesia key
	Indonesia = "ID"
	// Singapore key
	Singapore = "SG"
	// Arab key
	Arab = "AE"
	// USA key
	USA = "US"
	// Brazil key
	Brazil = "BR"
	// Russia key
	Russia = "RU"
)

var (
	// RegionalURL mappings
	RegionalURL = map[string]map[string]string{
		ProductionEnvironment: {
			India:     "http://api-inner-mgr-875.ihago.net/ymicro/api",
			Indonesia: "http://api-inner-mgr-863.ihago.net/ymicro/api",
			Singapore: "http://api-inner-mgr-881.ihago.net/ymicro/api",
			Arab:      "http://api-inner-mgr-894.ihago.net/ymicro/api",
			USA:       "http://api-inner-mgr-872.ihago.net/ymicro/api",
			Brazil:    "http://api-inner-mgr-889.ihago.net/ymicro/api",
			Russia:    "http://api-inner-mgr-892.ihago.net/ymicro/api",
		},
		TestingEnvironment: {
			Indonesia: "http://api-inner-test-863.ihago.net/ymicro/api",
			India:     "http://api-inner-test-in.ihago.net/ymicro/api",
			Singapore: "http://api-inner-test-sg.ihago.net/ymicro/api",
		},
	}

	// HTTPClient is a global HTTP client
	HTTPClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 10000 * time.Millisecond,
			}).DialContext,
			MaxIdleConns:        2,
			MaxIdleConnsPerHost: 2,
			IdleConnTimeout:     10 * time.Minute,
		},
		Timeout: 10000 * time.Millisecond,
	}
)

// Link is the calling parameters
type Link struct {
	FromService string
	URL         string
	ToService   string
	Method      string
}

// WithMethod sets the Link's Method
func (l Link) WithMethod(m string) Link {
	l.Method = m
	return l
}

// RoundTrip calls the backend service in one HTTP round trip
func RoundTrip(req proto.Message, link Link, res proto.Message) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var (
		reqBody  []byte
		respBody []byte
		reqObj   *http.Request
		respObj  *http.Response
	)
	if reqBody, err = json.Marshal(req); err != nil {
		return fmt.Errorf("marshal request to json: %w", err)
	}
	if reqObj, err = http.NewRequestWithContext(ctx, "POST", link.URL, bytes.NewReader(reqBody)); err != nil {
		return fmt.Errorf("new http request: %w", err)
	}
	reqObj.Header.Set("Content-Type", "application/json")
	reqObj.Header.Set("X-Micro-From-Service", link.FromService)
	reqObj.Header.Set("X-Ymicro-Api-Service-Name", link.ToService)
	reqObj.Header.Set("X-Ymicro-Api-Method-Name", link.Method)
	//logger.Debugf("send http request with header: %v , body: %v", reqObj.Header, string(reqBody))

	if respObj, err = HTTPClient.Do(reqObj); err != nil {
		return fmt.Errorf("do http roundtrip: %w", err)
	}
	defer respObj.Body.Close()

	if respBody, err = ioutil.ReadAll(respObj.Body); err != nil {
		return fmt.Errorf("read http response: %w", err)
	}
	//logger.Debugf("got http response body: %v", string(respBody))

	if err = json.Unmarshal(respBody, res); err != nil {
		return fmt.Errorf("unmarshal response from json: %w", err)
	}
	if resulter, ok := res.(interface {
		GetResult() *dashboard.Result
	}); ok {
		if result := resulter.GetResult(); result == nil || result.Errcode != 0 {
			return fmt.Errorf("http transaction failure: %v", string(respBody))
		}
	}
	return nil
}
