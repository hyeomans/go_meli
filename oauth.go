package meli

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type oauthEndpoints struct {
	c Config
}

func newOAuthEndpoints(c Config) *oauthEndpoints {
	return &oauthEndpoints{
		c: c,
	}
}

func (e *oauthEndpoints) Token(ctx context.Context) (*OauthTokenResponse, error) {
	baseURL := fmt.Sprintf("%s/oauth/token", e.c.BaseURL)
	data := url.Values{}
	data.Add("grant_type", "authorization_code")
	data.Add("client_id", e.c.ClientID)
	data.Add("client_secret", e.c.Secret)
	data.Add("code", e.c.UserCode)
	data.Add("redirect_uri", e.c.CallbackURL)
	values := strings.NewReader(data.Encode())

	req, err := http.NewRequest(http.MethodPost, baseURL, values)

	if err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not create new request: %s", err),
			statusCode:  500,
		}
	}

	body, err := request(ctx, req, e.c.HTTP)
	if err != nil {
		return nil, err
	}

	jsonResponse := OauthTokenResponse{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not unmarshal json into struct: %s - raw body: %v", err, string(body)),
			statusCode:  500,
		}
	}

	return &jsonResponse, nil
}

func (e *oauthEndpoints) RefreshToken(ctx context.Context, refreshToken string) (*OauthTokenResponse, error) {
	baseURL := fmt.Sprintf("%s/oauth/token", e.c.BaseURL)
	data := url.Values{}
	data.Add("grant_type", "refresh_token")
	data.Add("client_id", e.c.ClientID)
	data.Add("client_secret", e.c.Secret)
	data.Add("refresh_token", refreshToken)

	values := strings.NewReader(data.Encode())

	req, err := http.NewRequest(http.MethodPost, baseURL, values)

	if err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not create new request: %s", err),
			statusCode:  500,
		}
	}

	body, err := request(ctx, req, e.c.HTTP)
	if err != nil {
		return nil, err
	}

	jsonResponse := OauthTokenResponse{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not unmarshal json into struct: %s - raw body: %v", err, string(body)),
			statusCode:  500,
		}
	}

	return &jsonResponse, nil
}
