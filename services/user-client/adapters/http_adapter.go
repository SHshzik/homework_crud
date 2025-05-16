package adapters

import "github.com/SHshzik/homework_crud/services/user-server/entity"

type HTTPClient struct{}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{}
}

func (c *HTTPClient) Index() ([]*entity.User, error) {
	return nil, nil
}

func (c *HTTPClient) Create(name, email, phone string) (*entity.User, error) {
	return nil, nil
}

func (c *HTTPClient) Read(id int) (*entity.User, error) {
	return nil, nil
}

func (c *HTTPClient) Update(user *entity.User) (*entity.User, error) {
	return nil, nil
}
