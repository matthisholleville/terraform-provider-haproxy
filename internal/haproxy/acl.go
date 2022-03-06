package haproxy

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy/models"
)

func (c *Client) GetAcl(transactionId string, acl *models.ACL, parentName string, parentType string) (*models.ACL, error) {
	url := c.base_url + "/services/haproxy/configuration/acls/" + strconv.Itoa(acl.Index) + "?parent_name=" + parentName + "&transaction_id=" + transactionId + "&parent_type" + parentType
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.GetAcl{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (c *Client) CreateAcl(transactionId string, acl *models.ACL, parentName string, parentType string) (*models.ACL, error) {
	url := c.base_url + "/services/haproxy/configuration/acls?parent_name=" + parentName + "&transaction_id=" + transactionId + "&parent_type" + parentType
	bodyStr, _ := json.Marshal(acl)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res := models.ACL{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateAcl(transactionId string, acl *models.ACL, parentName string, parentType string) (*models.ACL, error) {
	url := c.base_url + "/services/haproxy/configuration/acls/" + strconv.Itoa(acl.Index) + "?parent_name=" + parentName + "&transaction_id=" + transactionId + "&parent_type" + parentType
	bodyStr, _ := json.Marshal(acl)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res := models.ACL{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteAcl(transactionId string, acl *models.ACL, parentName string, parentType string) error {
	url := c.base_url + "/services/haproxy/configuration/acls/" + strconv.Itoa(acl.Index) + "?parent_name=" + parentName + "&transaction_id=" + transactionId + "&parent_type" + parentType
	bodyStr, _ := json.Marshal(acl)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(bodyStr))
	if err != nil {
		return err
	}
	if err := c.sendRequest(req, nil); err != nil {
		return err
	}

	return nil
}
