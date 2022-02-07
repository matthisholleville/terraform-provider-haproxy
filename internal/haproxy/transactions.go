package haproxy

import (
	"net/http"
	"strconv"
)

type Transaction struct {
	Version int    `json:"_version"`
	Id      string `json:"id"`
	Status  string `json:"status"`
}

func (c *Client) CreateTransaction(version int) (*Transaction, error) {
	url := c.base_url + "/services/haproxy/transaction?version=" + strconv.Itoa(version)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	res := Transaction{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CommitTransaction(transactionId string) (*Transaction, error) {
	url := c.base_url + "/services/haproxy/transaction/" + transactionId
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	res := Transaction{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
