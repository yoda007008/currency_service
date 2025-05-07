package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL string
}

func NewClient(baseUrl string) *Client {
	return &Client{BaseURL: baseUrl}
}

func (c *Client) GetCurrencyRate(code string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/currency/%s", c.BaseURL, code)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
