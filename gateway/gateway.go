/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"odfe-cli/client"
	"odfe-cli/entity"
	"odfe-cli/entity/es"

	"github.com/hashicorp/go-retryablehttp"
)

//HTTPGateway type for gateway client
type HTTPGateway struct {
	Client  *client.Client
	Profile *entity.Profile
}

//GetDefaultHeaders returns common headers
func GetDefaultHeaders() map[string]string {
	return map[string]string{
		"content-type": "application/json",
	}
}

//NewHTTPGateway creates new HTTPGateway instance
func NewHTTPGateway(c *client.Client, p *entity.Profile) *HTTPGateway {
	return &HTTPGateway{
		Client:  c,
		Profile: p,
	}
}

//isValidResponse checks whether the response is valid or not by checking the status code
func (g *HTTPGateway) isValidResponse(response *http.Response) error {
	if response == nil {
		return errors.New("response is nil")
	}
	// client error if 400 <= status code < 500
	if response.StatusCode >= http.StatusBadRequest && response.StatusCode < http.StatusInternalServerError {

		return es.NewRequestError(
			response.StatusCode,
			response.Body,
			fmt.Errorf("%d Client Error: %s for url: %s", response.StatusCode, response.Status, response.Request.URL))
	}
	// server error if status code >= 500
	if response.StatusCode >= http.StatusInternalServerError {

		return es.NewRequestError(
			response.StatusCode,
			response.Body,
			fmt.Errorf("%d Server Error: %s for url: %s", response.StatusCode, response.Status, response.Request.URL))
	}
	return nil
}

//Execute calls request using http and check if status code is ok or not
func (g *HTTPGateway) Execute(req *retryablehttp.Request) ([]byte, error) {

	response, err := g.Client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			return
		}
	}()
	if err = g.isValidResponse(response); err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

//Call calls request using http and return error if status code is not expected
func (g *HTTPGateway) Call(req *retryablehttp.Request, statusCode int) ([]byte, error) {
	resBytes, err := g.Execute(req)
	if err == nil {
		return resBytes, nil
	}
	r, ok := err.(*es.RequestError)
	if !ok {
		return nil, err
	}
	if r.StatusCode() != statusCode {
		return nil, fmt.Errorf(r.GetResponse())
	}
	return nil, err

}

//BuildRequest builds request based on method and appends payload for given url with headers
// TODO: Deprecate this method by replace this with BuildCurlRequest
func (g *HTTPGateway) BuildRequest(ctx context.Context, method string, payload interface{}, url string, headers map[string]string) (*retryablehttp.Request, error) {
	reqBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return g.BuildCurlRequest(ctx, method, reqBytes, url, headers)
}

//BuildCurlRequest builds request based on method and add payload (in byte)
func (g *HTTPGateway) BuildCurlRequest(ctx context.Context, method string, payload []byte, url string, headers map[string]string) (*retryablehttp.Request, error) {
	r, err := retryablehttp.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req := r.WithContext(ctx)
	if len(g.Profile.UserName) != 0 {
		req.SetBasicAuth(g.Profile.UserName, g.Profile.Password)
	}
	if len(headers) == 0 {
		return req, nil
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req, nil
}

//GetValidEndpoint get url based on user config
func GetValidEndpoint(profile *entity.Profile) (*url.URL, error) {
	if len(profile.Endpoint) == 0 {
		return &url.URL{
			Scheme: "https",
			Host:   "localhost:9200",
		}, nil
	}
	u, err := url.Parse(profile.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint: %v due to %v", profile.Endpoint, err)
	}
	return u, nil
}
