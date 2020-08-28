package meli

import (
	"fmt"
	"net/url"
)

type Client struct {
	OAuth      *oauthEndpoints
	Sites      *sitesEndpoints
	Categories *categoriesEndpoints
	Orders     *ordersEndpoint
}

// Config ...
type Config struct {
	HTTP        Doer
	ClientID    string
	Secret      string
	CallbackURL string
	BaseURL     string
}

func New(c Config) (*Client, error) {
	_, err := url.ParseRequestURI(c.CallbackURL)

	if err != nil {
		return nil, fmt.Errorf("invalid callback_url: %s", err)
	}

	if c.BaseURL == "" {
		c.BaseURL = "https://api.mercadolibre.com"
	}

	client := &Client{
		OAuth:      newOAuthEndpoints(c),
		Sites:      newSitesEndpoints(c),
		Categories: newCategoriesEndpoints(c),
		Orders:     newOrdersEndpoint(c),
	}

	return client, nil
}
