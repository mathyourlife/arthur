package tib

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArthurHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ArthurFractionHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	contents, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	assert.Regexp(t, regexp.MustCompile("\\\\frac{[-]?[\\d]+}{[-]?[\\d]+}"), string(contents))
}
