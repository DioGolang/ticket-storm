package resilience

import "net/http"

type Decorator func(http.RoundTripper) http.RoundTripper

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func Decorate(base http.RoundTripper, decorators ...Decorator) http.RoundTripper {
	transport := base

	for i := len(decorators) - 1; i >= 0; i-- {
		transport = decorators[i](transport)
	}
	return transport
}
