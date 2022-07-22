package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"terraform-provider-example/api/server"

	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) GetAll() (*map[string]server.Item, error) {
	body, err := c.httpRequest("item", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]server.Item{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) GetItem(name string) (*server.Item, error) {
	body, err := c.httpRequest(fmt.Sprintf("item/%v", name), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	item := &server.Item{}
	err = json.NewDecoder(body).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (c *Client) NewItem(item *server.Item) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return err
	}
	_, err = c.httpRequest("item", "POST", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateItem(item *server.Item) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("item/%s", item.Name), "PUT", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteItem(itemName string) error {
	_, err := c.httpRequest(fmt.Sprintf("item/%s", itemName), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s:%v/%s", "http://localhost", 3000, path)
}
