package secret

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"net/http"
)

type SecretService interface {
	VerifySecret(value string) bool
	VerifySecretFromReq(req *http.Request) bool
	VerifySecretFromHeader(header *http.Header) bool
	NewAuthenticatedClient() *http.Client
}

type secretServiceInternal struct {
	internalSecret string
}

func newSecretServiceInternal() *secretServiceInternal {
	return &secretServiceInternal{
		internalSecret: fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprint(rand.Text())))),
	}
}

func (sp *secretServiceInternal) VerifySecret(value string) bool {
	return sp.internalSecret != "" && value == sp.internalSecret
}

func (sp *secretServiceInternal) VerifySecretFromHeader(header *http.Header) bool {
	return sp.VerifySecret(header.Get("X-Api-Token"))
}

func (sp *secretServiceInternal) VerifySecretFromReq(req *http.Request) bool {
	return sp.VerifySecretFromHeader(&req.Header)
}

func (sp *secretServiceInternal) NewAuthenticatedClient() *http.Client {
	return &http.Client{
		Transport: &transport{
			base:   http.DefaultTransport,
			secret: sp.internalSecret,
		},
	}
}
