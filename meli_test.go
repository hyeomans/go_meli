package meli

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	ctx := context.Background()
	mockServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		t.Logf("Sending fake request to %v", req.RequestURI)
		rw.WriteHeader(200)
		rw.Write([]byte(`
		{
			"access_token": "reallylongaccesstoken", 
			"token_type": "fake", 
			"expires_in": 100, 
			"scope": "all", 
			"refresh_token": "refreshtoken"
		}
		`))
	}))
	defer mockServer.Close()

	c := Config{
		HTTP:        mockServer.Client(),
		BaseURL:     mockServer.URL,
		CallbackURL: "https://localhost:3000/callback/mercadolibre?something=token",
		ClientID:    "clientID",
		Secret:      "clientSecret",
	}
	client, _ := New(c)

	r, err := client.OAuth.Token(ctx, "token")

	if err != nil {
		t.Errorf("OAuth.Token returned an error: %v", err)
	}

	if r.AccessToken != "reallylongaccesstoken" {
		t.Errorf("Expecting %v but got %v", "reallylongaccesstoken", r.AccessToken)
	}
}
