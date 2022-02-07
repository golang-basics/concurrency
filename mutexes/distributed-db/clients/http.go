package clients

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"distributed-db/models"
)

func NewHTTP(host string) *HTTPClient {
	client := HTTPClient{
		host:       host,
		httpClient: &http.Client{},
	}
	return &client
}

type HTTPClient struct {
	host       string
	httpClient *http.Client
}

func (c *HTTPClient) Get(peer string, key string) (models.CacheItem, error) {
	body := models.GetRequest{
		Keys: []string{key},
	}
	req, err := c.makeRequest(http.MethodGet, c.url(peer, "get"), body)
	if err != nil {
		return models.CacheItem{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return models.CacheItem{}, err
	}

	var cacheItem []models.CacheItem
	err = json.NewDecoder(res.Body).Decode(&cacheItem)
	if err != nil {
		return models.CacheItem{}, err
	}

	return cacheItem[0], nil
}

func (c *HTTPClient) Gossip(peer string, summary models.Summary) error {
	body := models.GossipRequest{
		Summary: summary,
	}
	req, err := c.makeRequest(http.MethodPost, c.url(peer, "gossip"), body)
	if err != nil {
		return err
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *HTTPClient) url(peer, path string) string {
	u := url.URL{
		Scheme: "http",
		Host:   peer,
		Path:   path,
	}
	return u.String()
}

func (c *HTTPClient) makeRequest(method, url string, body interface{}) (*http.Request, error) {
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}
	req.Host = c.host

	return req, nil
}
