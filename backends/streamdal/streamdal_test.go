package streamdal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"

	"github.com/batchcorp/plumber/config"
	"github.com/batchcorp/plumber/options"
	"github.com/batchcorp/plumber/server/types"
)

var logger = &logrus.Logger{Out: ioutil.Discard}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// newHttpClient returns *http.Client with Transport replaced to avoid making real calls
func newHttpClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func StreamdalWithMockResponse(httpCode int, responseBody string) *Streamdal {
	client := newHttpClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: httpCode,
			Body:       ioutil.NopCloser(bytes.NewBufferString(responseBody)),
		}
	})

	return &Streamdal{
		Log:    logrus.NewEntry(logger),
		Client: client,
		Opts:   options.NewCLIOptions(),
		PersistentConfig: &config.Config{
			Connections: make(map[string]*types.Connection),
		},
	}
}

var _ = Describe("Streamdal", func() {
	Context("New", func() {
		It("returns a new instance of Streamdal struct", func() {
			b := New(&opts.CLIOptions{Streamdal: &opts.StreamdalOptions{}}, &config.Config{})
			Expect(b).To(BeAssignableToTypeOf(&Streamdal{}))
		})
	})
	Context("getCookieJar", func() {
		It("returns cookie jar with auth_token", func() {
			const Token = "testin123"

			b := &Streamdal{
				PersistentConfig: &config.Config{Token: Token},
				ApiUrl:           "https://api.streamdal.com",
			}

			jar := b.getCookieJar("/v1/collection")

			u, _ := url.Parse(b.ApiUrl + "/v1/collection")
			cookies := jar.Cookies(u)

			Expect(len(cookies)).To(Equal(1))
			Expect(cookies[0].Name).To(Equal("auth_token"))
			Expect(cookies[0].Value).To(Equal(Token))
		})
	})
	Context("Post", func() {
		It("returns unauthenticated error", func() {
			b := StreamdalWithMockResponse(http.StatusUnauthorized, "")

			_, _, err := b.Post("/", nil)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(errNotAuthenticated))
		})
	})
})
