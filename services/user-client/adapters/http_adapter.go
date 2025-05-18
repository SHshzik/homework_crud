package adapters

import (
	"bytes"
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
	url := fmt.Sprintf("http://localhost:%d/v1/users", c.cfg.PORT)
	data, err := json.Marshal(entity.User{Name: name, Email: email, Phone: phone})
	if err != nil {
		c.l.Error("fail to marshal user: %v", err)

		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		c.l.Error("fail to create user: %v", err)

		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.l.Error("fail to read body: %v", err)

		return nil, err
	}

	var user entity.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		c.l.Info("body: %v", string(body))
		c.l.Error("fail to unmarshal user: %v", err)

		return nil, err
	}

	return &user, nil
}

func (c *HTTPClient) Read(id int) (*entity.User, error) {
	url := fmt.Sprintf("http://localhost:%d/v1/users/%d", c.cfg.PORT, id)

	resp, err := http.Get(url)
	if err != nil {
		c.l.Error("fail to get user: %v", err)

		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.l.Error("fail to read body: %v", err)

		return nil, err
	}

	var user entity.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		c.l.Error("fail to unmarshal user: %v", err)

		return nil, err
	}

	return &user, nil
}

func (c *HTTPClient) Update(user *entity.User) (*entity.User, error) {
	return nil, nil
}
