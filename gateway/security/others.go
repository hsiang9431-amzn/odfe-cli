package security

import "odfe-cli/entity/security"

func NewAccount(u, p string) (account, error) {
	if u == "" || p == "" {
		return nil, ErrInvalidCredentials
	}
	return &accountClient{username: u, password: p}, nil
}

func NewConfig(u, p string) (config, error) {
	if u == "" || p == "" {
		return nil, ErrInvalidCredentials
	}
	return &configClient{username: u, password: p}, nil
}

func NewCache(u, p string) (cache, error) {
	if u == "" || p == "" {
		return nil, ErrInvalidCredentials
	}
	return &cacheClient{username: u, password: p}, nil
}

func NewHealth(u, p string) (health, error) {
	if u == "" || p == "" {
		return nil, ErrInvalidCredentials
	}
	return &healthClient{username: u, password: p}, nil
}

type accountClient struct {
	username string
	password string
}

func (c *accountClient) GetDetails() (interface{}, error)                  { return nil, nil }
func (c *accountClient) ChangePassword(security.ChangePasswordQuery) error { return nil }

type configClient struct {
	username string
	password string
}

func (c *configClient) Get() (interface{}, error)                 { return nil, nil }
func (c *configClient) Update(interface{}) error                  { return nil }
func (c *configClient) Patch(security.PatchQuery) error           { return nil }
func (c *configClient) PatchMultiple([]security.PatchQuery) error { return nil }

type cacheClient struct {
	username string
	password string
}

func (c *cacheClient) Flush() error { return nil }

type healthClient struct {
	username string
	password string
}

func (c *healthClient) Check() error { return nil }
