package entityclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetAllEntities(authToken *string) (*[]Entity, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/entities", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	var entities []Entity
	err = json.Unmarshal(body, &entities)
	if err != nil {
		return nil, err
	}

	return &entities, nil
}

func (c *Client) GetEntity(entityID string, authToken *string) (*Entity, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/entities/%s", c.HostURL, entityID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	entity := Entity{}
	err = json.Unmarshal(body, &entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (c *Client) CreateEntity(entity Entity, authToken *string) (*Entity, error) {
	rb, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/entities", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newEntity := Entity{}
	err = json.Unmarshal(body, &newEntity)
	if err != nil {
		return nil, err
	}

	return &newEntity, nil
}

func (c *Client) UpdateEntity(entityID string, entity Entity, authToken *string) (*Entity, error) {
	rb, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/entities/%s", c.HostURL, entityID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	order := Entity{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) DeleteEntity(entityID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/entities/%s", c.HostURL, entityID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "Deleted order" {
		return errors.New(string(body))
	}

	return nil
}
