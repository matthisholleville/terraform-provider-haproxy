package haproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type MapEntrie struct {
	Id    string `json:"id,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value"`
}

func (c *Client) GetMapEntrie(entrieName string, mapName string) (*MapEntrie, error) {
	url := c.base_url + "/services/haproxy/runtime/maps_entries/" + replaceSlashInString(entrieName) + "?map=" + mapName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := MapEntrie{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateMapEntrie(entrie *MapEntrie, mapName string, forceSync bool) (*MapEntrie, error) {
	url := c.base_url + "/services/haproxy/runtime/maps_entries/?map=" + mapName + "&force_sync=" + strconv.FormatBool(forceSync)
	bodyStr, _ := json.Marshal(entrie)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res := MapEntrie{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateMapEntrie(entrie *MapEntrie, mapName string, forceSync bool) (*MapEntrie, error) {
	url := c.base_url + "/services/haproxy/runtime/maps_entries/" + replaceSlashInString(entrie.Key) + "?map=" + mapName + "&force_sync=" + strconv.FormatBool(forceSync)
	entrieValue := &MapEntrie{
		Value: entrie.Value,
	}
	bodyStr, _ := json.Marshal(entrieValue)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res := MapEntrie{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteMapEntrie(entrieName string, mapName string, forceSync bool) error {
	url := c.base_url + "/services/haproxy/runtime/maps_entries/" + replaceSlashInString(entrieName) + "?map=" + mapName + "&force_sync=" + strconv.FormatBool(forceSync)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	res := MapEntrie{}
	if err := c.sendRequest(req, res); err != nil {
		fmt.Println("toto")
		return err
	}

	return nil

}
