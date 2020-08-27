package meli

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func request(ctx context.Context, req *http.Request, doer Doer) ([]byte, error) {
	req.Header.Add("Content-Type", "application/json")
	res, err := doer.Do(req.WithContext(ctx))

	if err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not get response from request: %s", err),
			statusCode:  500,
		}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not read body into memory: %s", err),
			statusCode:  500,
		}
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var isTemporary bool
		if http.StatusRequestTimeout == res.StatusCode ||
			http.StatusTooManyRequests == res.StatusCode ||
			http.StatusInternalServerError == res.StatusCode ||
			http.StatusServiceUnavailable == res.StatusCode {
			isTemporary = true
		}

		return nil, ClientError{
			isTemporary: isTemporary,
			message:     fmt.Sprintf("could not get a valid response code: %v - raw response: %v", res.StatusCode, string(body)),
			statusCode:  res.StatusCode,
		}
	}

	return body, nil
}
