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

func (c *HTTPClient) Get(node string, keys []string) ([]models.CacheItem, error) {
	body := models.GetRequest{Keys: keys}
	req, err := c.makeRequest(http.MethodGet, c.url(node, "get"), body)
	if err != nil {
		return []models.CacheItem{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []models.CacheItem{}, err
	}

	var cacheItems []models.CacheItem
	err = json.NewDecoder(res.Body).Decode(&cacheItems)
	if err != nil {
		return []models.CacheItem{}, err
	}

	return cacheItems, nil
}

func (c *HTTPClient) Set(node string, key, value string) (models.CacheItem, error) {
	body := models.SetRequest{
		Key:   key,
		Value: value,
	}
	req, err := c.makeRequest(http.MethodPost, c.url(node, "set"), body)
	if err != nil {
		return models.CacheItem{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return models.CacheItem{}, err
	}

	var item models.CacheItem
	err = json.NewDecoder(res.Body).Decode(&item)
	if err != nil {
		return models.CacheItem{}, err
	}
	item.Node = node

	return item, nil
}

func (c *HTTPClient) SetBatch(node string, items map[int]models.CacheItem) ([]models.CacheItem, error) {
	body := models.SetBatchRequest{Items: items}
	req, err := c.makeRequest(http.MethodPost, c.url(node, "set/batch"), body)
	if err != nil {
		return []models.CacheItem{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []models.CacheItem{}, err
	}

	var cacheItems []models.CacheItem
	err = json.NewDecoder(res.Body).Decode(&items)
	if err != nil {
		return []models.CacheItem{}, err
	}

	return cacheItems, nil
}

func (c *HTTPClient) Gossip(node string, nodes models.NodesMap, tokensChecksum string) (models.NodesMap, error) {
	body := models.GossipRequest{
		Nodes:          nodes,
		TokensChecksum: tokensChecksum,
	}
	req, err := c.makeRequest(http.MethodPost, c.url(node, "gossip"), body)
	if err != nil {
		return models.NodesMap{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return models.NodesMap{}, err
	}

	var gossipRes models.GossipResponse
	err = json.NewDecoder(res.Body).Decode(&gossipRes)
	if err != nil {
		return models.NodesMap{}, err
	}

	return gossipRes.Nodes, nil
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
