package fakehttp

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func ExampleFakeClient() {
	fc := &FakeClient{
		Error: errors.New("oh noes"),
		Response: &http.Response{
			StatusCode: 250,
			Body:       ioutil.NopCloser(strings.NewReader("oh wow, it works!")),
		},
	}
	client = httpClient(fc)

	st, body, err := test()
	fmt.Printf("%v %s %v\n", st, body, err)
	fmt.Printf("%s %s\n", fc.Requests[0].Method, fc.Requests[0].URL)
	// Output: 250 oh wow, it works! oh noes
	// GET https://example.com
}

// Use an interface.
type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

var client = httpClient(&http.Client{
	Timeout: time.Duration(5 * time.Second),
})

func test() (int, []byte, error) {
	resp, err := client.Get("https://example.com")
	_ = err
	defer resp.Body.Close() // nolint: errcheck

	b, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, b, err
}
