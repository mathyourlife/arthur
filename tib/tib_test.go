package tib

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func endpointContents(urlPath string) ([]byte, error) {
	mux := http.NewServeMux()
	SetupMux("api/v1/", mux)

	ts := httptest.NewServer(mux)
	defer ts.Close()
	client := ts.Client()

	u, err := url.Parse(ts.URL)
	if err != nil {
		return nil, err
	}

	u, err = u.Parse(urlPath)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u.String(), nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func TestFraction(t *testing.T) {
	contents, err := endpointContents("/api/v1/fraction?type=unit&format=latex")
	if err != nil {
		t.Error(err)
	}
	assert.Regexp(t, regexp.MustCompile(`\\frac{[-]?[\d]+}{[-]?[\d]+}`), string(contents))

	contents, err = endpointContents("/api/v1/fraction?type=unit&format=json")
	if err != nil {
		t.Error(err)
	}
	assert.Regexp(t, regexp.MustCompile(`{"numerator":[\d]+,"denominator":[\d]+}`), string(contents))

	contents, err = endpointContents("/api/v1/fraction?type=proper")
	if err != nil {
		t.Error(err)
	}
	assert.Regexp(t, regexp.MustCompile(`\\frac{[-]?[\d]+}{[-]?[\d]+}`), string(contents))
}

func TestInteger(t *testing.T) {
	contents, err := endpointContents("/api/v1/integer")
	if err != nil {
		t.Error(err)
	}
	assert.Regexp(t, regexp.MustCompile(`[\d]{1,2}`), string(contents))

	contents, err = endpointContents("/api/v1/integer?size=medium&format=json")
	if err != nil {
		t.Error(err)
	}
	assert.Regexp(t, regexp.MustCompile(`{"value":[\d]{2,3}}`), string(contents))

	contents, err = endpointContents("/api/v1/integer?size=large&format=json")
	if err != nil {
		t.Error(err)
	}
	assert.Regexp(t, regexp.MustCompile(`{"value":[\d]{5}}`), string(contents))

	contents, err = endpointContents("/api/v1/integer?format=json")
	if err != nil {
		t.Error(err)
	}
	assert.Regexp(t, regexp.MustCompile(`{"value":[\d]+}`), string(contents))
}
