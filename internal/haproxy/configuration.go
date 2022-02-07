package haproxy

import (
	"net/http"

	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy/models"
)

func (c *Client) GetConfiguration(transactionId string) (*models.Configuration, error) {
	url := c.base_url + "/services/haproxy/configuration/raw?transaction_id" + transactionId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.Configuration{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
