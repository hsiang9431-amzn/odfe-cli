package security

import (
	"odfe-cli/client"
	"odfe-cli/entity"
	"odfe-cli/entity/security"
	"odfe-cli/gateway"
)

func NewAccount(c *client.Client, p *entity.Profile) (account, error) {
	return &accountClient{gateway: gateway.NewHTTPGateway(c, p)}, nil
}

func NewConfig(c *client.Client, p *entity.Profile) (config, error) {
	return &configClient{gateway: gateway.NewHTTPGateway(c, p)}, nil
}

func NewCache(c *client.Client, p *entity.Profile) (cache, error) {
	return &cacheClient{gateway: gateway.NewHTTPGateway(c, p)}, nil
}

func NewHealth(c *client.Client, p *entity.Profile) (health, error) {
	return &healthClient{gateway: gateway.NewHTTPGateway(c, p)}, nil
}

type accountClient struct {
	gateway *gateway.HTTPGateway
}

func (c *accountClient) GetDetails() (interface{}, error)                  { return nil, nil }
func (c *accountClient) ChangePassword(security.ChangePasswordQuery) error { return nil }

type configClient struct {
	gateway *gateway.HTTPGateway
}

func (c *configClient) Get() (interface{}, error)                 { return nil, nil }
func (c *configClient) Update(interface{}) error                  { return nil }
func (c *configClient) Patch(security.PatchQuery) error           { return nil }
func (c *configClient) PatchMultiple([]security.PatchQuery) error { return nil }

type cacheClient struct {
	gateway *gateway.HTTPGateway
}

func (c *cacheClient) Flush() error { return nil }

type healthClient struct {
	gateway *gateway.HTTPGateway
}

func (c *healthClient) Check() error { return nil }
