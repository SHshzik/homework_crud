package adapters

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SHshzik/homework_crud/pkg/logger"
	"github.com/SHshzik/homework_crud/services/user-client/config"
	"github.com/SHshzik/homework_crud/services/user-server/entity"
)

type HTTPClient struct {
	cfg *config.HTTP
	l   logger.Interface
}

func NewHTTPClient(cfg *config.HTTP, l logger.Interface) *HTTPClient {
	return &HTTPClient{cfg: cfg, l: l}
}

type usersIndex struct {
	Users []*entity.User `json:"users"`
}

func (c *HTTPClient) Index() ([]*entity.User, error) {
	url := fmt.Sprintf("http://localhost:%d/v1/users", c.cfg.PORT)

	// TODO: add check 5xx errors
	resp, err := http.Get(url)
	if err != nil {
		c.l.Error("fail to get users: %v", err)

		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.l.Error("fail to read body: %v", err)

		return nil, err
	}

	var usersIndex usersIndex

	err = json.Unmarshal(body, &usersIndex)
	if err != nil {
		c.l.Error("fail to unmarshal users: %v", err)

		return nil, err
	}

	return usersIndex.Users, nil
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
