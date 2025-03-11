package client

import (
	"context"
	"fmt"
	"net/http"
)

type TokenAuth struct {
	token string
}

func NewTokenAuth(token string) (*TokenAuth, error) {
	return &TokenAuth{
		token: token,
	}, nil
}

func (s *TokenAuth) Intercept(ctx context.Context, req *http.Request) error {
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", s.token))
	return nil
}
