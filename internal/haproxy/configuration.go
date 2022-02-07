package haproxy

import "net/http"

type Configuration struct {
	Version int    `json:"_version"`
	Data    string `json:"data"`
}

func (c *Client) GetConfiguration(transactionId string) (*Configuration, error) {
	url := c.base_url + "/services/haproxy/configuration/raw?transaction_id" + transactionId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := Configuration{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
