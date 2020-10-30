package demo

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RewriteTransport struct {
	Transport http.RoundTripper
}

func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func HTTPTestFacilities(t *testing.T) (*http.Client, *http.ServeMux, *httptest.Server) {
	t.Helper()

	mux := http.NewServeMux()

	server := httptest.NewServer(mux)

	transport := &RewriteTransport{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse(server.URL) // httptest server.URL
			},
		}}
	client := &http.Client{Transport: transport}

	return client, mux, server
}

func TestClient(t *testing.T) {
	httpClient, mux, server := HTTPTestFacilities(t)
	defer server.Close()

	mux.HandleFunc("/pay/unifiedorder", func(w http.ResponseWriter, r *http.Request) {
		assert.EqualValues(t, "POST", r.Method)

		// TODO: assert.EqualValues(t, &Request{}, xml.Unmarshal(r.Body))

		w.Header().Set("Content-Type", "application/xml")

		f, _ := os.Open("testdata/response.xml")
		io.Copy(w, f)
	})

	expected := &Response{
		XMLName:    xml.Name{"", "xml"},
		ReturnCode: "SUCCESS",
		ReturnMsg:  "OK",
		AppID:      "wx2421b1c4370ec43b",
		MchID:      "10000100",
		NonceStr:   "IITRi8Iabbblz1Jc",
		Sign:       "7921E432F65EB8ED0CE9755F0E86D72F",
		ResultCode: "SUCCESS",
		PrepayID:   "wx201411101639507cbf6ffd8b0779950874",
		TradeType:  "JSAPI",
	}

	request := &Request{}
	client := NewClient(httpClient, URL) // NOTE: or server.URL without RewriteTransport
	response, err := client.UnifiedOrder(request)

	assert.NoError(t, err)
	assert.EqualValues(t, expected, response)
}
