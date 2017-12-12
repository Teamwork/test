// Package fakehttp provides a "fake" http.Client implementation.
package fakehttp

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// FakeClient is a fake http.Client implementation.
//
// It will always return the provided data from the Response and Error field,
// and all requests are recorded in the Requests field.
type FakeClient struct {
	Response *http.Response  // Response to return from all Do() calls.
	Error    error           // Error to return from all Do() calls.
	Requests []*http.Request // Records all requests made with Do().
}

// Do mocks client.Do.
func (c *FakeClient) Do(req *http.Request) (*http.Response, error) {
	c.Requests = append(c.Requests, req)
	return c.Response, c.Error
}

// Get mocks client.Get.
func (c *FakeClient) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// Head mocks client.Head.
func (c *FakeClient) Head(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// Post mocks client.Post.
func (c *FakeClient) Post(
	url string, contentType string, body io.Reader,
) (resp *http.Response, err error) {

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	return c.Do(req)
}

// PostForm mocks client.PostForm.
func (c *FakeClient) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return c.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()))
}
