package meli

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ordersEndpoint struct {
	c Config
}

func newOrdersEndpoint(c Config) *ordersEndpoint {
	return &ordersEndpoint{
		c: c,
	}
}

// Get the order
// /orders/:id
func (e *ordersEndpoint) Get(ctx context.Context, accessToken string, id int) (*Order, error) {
	baseURL := fmt.Sprintf("%s/orders/%d", e.c.BaseURL, id)

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)

	if err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not create new request: %s", err),
			statusCode:  500,
		}
	}

	body, err := requestWithToken(ctx, accessToken, req, e.c.HTTP)

	if err != nil {
		return nil, err
	}

	jsonResponse := &Order{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not unmarshal json into struct: %s - raw body: %v", err, string(body)),
			statusCode:  500,
		}
	}

	return jsonResponse, nil
}
