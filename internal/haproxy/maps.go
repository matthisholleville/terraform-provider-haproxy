package haproxy

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy/models"
)

func (c *Client) GetMapEntrie(entrieName string, mapName string) (*models.MapEntrie, error) {
	url := c.base_url + "/services/haproxy/runtime/maps_entries/" + encodeUrl(entrieName) + "?map=" + mapName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.MapEntrie{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateMapEntrie(entrie *models.MapEntrie, mapName string, forceSync bool) (*models.MapEntrie, error) {
	url := c.base_url + "/services/haproxy/runtime/maps_entries/?map=" + mapName + "&force_sync=" + strconv.FormatBool(forceSync)
	bodyStr, _ := json.Marshal(entrie)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res := models.MapEntrie{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateMapEntrie(entrie *models.MapEntrie, mapName string, forceSync bool) (*models.MapEntrie, error) {
	url := c.base_url + "/services/haproxy/runtime/maps_entries/" + encodeUrl(entrie.Key) + "?map=" + mapName + "&force_sync=" + strconv.FormatBool(forceSync)
	entrieValue := &models.MapEntrie{
		Value: entrie.Value,
	}
	bodyStr, _ := json.Marshal(entrieValue)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res := models.MapEntrie{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteMapEntrie(entrieName string, mapName string, forceSync bool) error {
	url := c.base_url + "/services/haproxy/runtime/maps_entries/" + encodeUrl(entrieName) + "?map=" + mapName + "&force_sync=" + strconv.FormatBool(forceSync)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	res := models.MapEntrie{}
	if err := c.sendRequest(req, res); err != nil {
		return err
	}

	return nil

}
