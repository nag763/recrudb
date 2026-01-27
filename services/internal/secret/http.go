package secret

import "net/http"

type transport struct {
	base   http.RoundTripper
	secret string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	newReq := req.Clone(req.Context())
	newReq.Header.Set("X-Api-Token", t.secret)

	return t.base.RoundTrip(newReq)
}
