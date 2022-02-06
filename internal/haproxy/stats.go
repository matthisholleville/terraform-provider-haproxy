package haproxy

import "net/http"

func (c *Client) TestApiCall() error {
	url := c.base_url + "/services/haproxy/stats/native"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if err := c.sendRequest(req, nil); err != nil {
		return err
	}

	return nil
}
