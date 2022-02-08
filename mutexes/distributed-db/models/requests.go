package models

type GetRequest struct {
	Keys []string `json:"keys"`
}

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
