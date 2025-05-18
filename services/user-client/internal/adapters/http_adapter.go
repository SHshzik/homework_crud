package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/SHshzik/homework_crud/pkg/logger"
	"github.com/SHshzik/homework_crud/services/user-client/config"
	"github.com/SHshzik/homework_crud/services/user-client/internal/entity"
)

var ErrServerError = errors.New("server error")

type HTTPClient struct {
	cfg    *config.HTTP
	l      logger.Interface
	client *http.Client
}

func NewHTTPClient(cfg *config.HTTP, l logger.Interface) *HTTPClient {
	return &HTTPClient{cfg: cfg, l: l, client: &http.Client{}}
}

type usersIndex struct {
	Users []*entity.User `json:"users"`
}

func (c *HTTPClient) Index() ([]*entity.User, error) {
	urlStr := fmt.Sprintf("http://localhost:%d/v1/users", c.cfg.PORT)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, urlStr, http.NoBody)
	if err != nil {
		c.l.Error("fail to create request: %v", err)

		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		c.l.Error("fail to get users: %v", err)

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusInternalServerError {
		c.l.Error("server error: %d", resp.StatusCode)

		return nil, fmt.Errorf("%w: %d", ErrServerError, resp.StatusCode)
	}

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

func (c *HTTPClient) Create(name, email string) (*entity.User, error) {
	urlStr := fmt.Sprintf("http://localhost:%d/v1/users", c.cfg.PORT)

	data, err := json.Marshal(entity.User{Name: name, Email: email})
	if err != nil {
		c.l.Error("fail to marshal user: %v", err)

		return nil, err
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, urlStr, bytes.NewBuffer(data))
	if err != nil {
		c.l.Error("fail to create request: %v", err)

		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
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
	urlStr := fmt.Sprintf("http://localhost:%d/v1/users/%d", c.cfg.PORT, id)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, urlStr, http.NoBody)
	if err != nil {
		c.l.Error("fail to create request: %v", err)

		return nil, err
	}

	resp, err := c.client.Do(req)
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
	url := fmt.Sprintf("http://localhost:%d/v1/users/%d", c.cfg.PORT, user.ID)

	data, err := json.Marshal(user)
	if err != nil {
		c.l.Error("fail to marshal user: %v", err)

		return nil, err
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		c.l.Error("fail to create request: %v", err)

		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		c.l.Error("fail to do request: %v", err)

		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.l.Error("fail to read body: %v", err)

		return nil, err
	}

	var updatedUser entity.User

	err = json.Unmarshal(body, &updatedUser)
	if err != nil {
		c.l.Error("fail to unmarshal user: %v", err)

		return nil, err
	}

	return &updatedUser, nil
}

func (c *HTTPClient) Delete(id int) error {
	url := fmt.Sprintf("http://localhost:%d/v1/users/%d", c.cfg.PORT, id)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, url, http.NoBody)
	if err != nil {
		c.l.Error("fail to create request: %v", err)

		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		c.l.Error("fail to delete user: %v", err)

		return err
	}
	defer resp.Body.Close()

	return nil
}
