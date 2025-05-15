package adapters

import "github.com/SHshzik/homework_crud/services/user-server/entity"

type HTTPClient struct{}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{}
}

func (c *HTTPClient) Index() ([]*entity.User, error) {
	return nil, nil
}
