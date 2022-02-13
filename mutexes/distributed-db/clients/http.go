package clients

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

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

func (c *HTTPClient) Get(node string, key string) (models.CacheItem, error) {
	body := models.GetRequest{
		Keys: []string{key},
	}
	req, err := c.makeRequest(http.MethodGet, c.url(node, "get"), body)
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

func (c *HTTPClient) Gossip(node string, nodes []string, tokensChecksum string) error {
	body := models.GossipRequest{
		Nodes:          nodes,
		CreatedAt:      time.Now().UTC(),
		TokensChecksum: tokensChecksum,
	}
	req, err := c.makeRequest(http.MethodPost, c.url(node, "gossip"), body)
	if err != nil {
		return err
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *HTTPClient) Tokens(node string) (models.TokenMappings, error) {
	req, err := c.makeRequest(http.MethodGet, c.url(node, "tokens"), nil)
	if err != nil {
		return models.TokenMappings{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return models.TokenMappings{}, err
	}

	var tokensRes models.TokensResponse
	err = json.NewDecoder(res.Body).Decode(&tokensRes)
	if err != nil {
		return models.TokenMappings{}, err
	}

	return tokensRes.Tokens, nil
}

func (c *HTTPClient) url(node, path string) string {
	u := url.URL{
		Scheme: "http",
		Host:   node,
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
