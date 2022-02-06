package haproxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	username   string
	password   string
	base_url   string
	HTTPClient *http.Client
}

func NewClient(username string, password string, server_url string) *Client {
	return &Client{
		username: username,
		password: password,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		base_url: "http://" + server_url + ":5555/v2",
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Basic "+basicAuth(c.username, c.password))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Try to unmarshall into errorResponse
	if res.StatusCode >= 300 {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if res.StatusCode == http.StatusNoContent {
		return nil
	}

	// bodyBytes, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bodyString := string(bodyBytes)
	// fmt.Println("bodystring", bodyString)
	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil

}
