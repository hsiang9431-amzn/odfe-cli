package security

import (
	"odfe-cli/entity/security"
	"odfe-cli/gateway"
)

type accountClient struct {
	gateway *gateway.HTTPGateway
}

func (c *accountClient) GetDetails() (interface{}, error)                  { return nil, nil }
func (c *accountClient) ChangePassword(security.ChangePasswordQuery) error { return nil }

type configClient struct {
	gateway *gateway.HTTPGateway
}

func (c *configClient) Get() (interface{}, error)         { return nil, nil }
func (c *configClient) Update(interface{}) error          { return nil }
func (c *configClient) Patch([]security.PatchQuery) error { return nil }

type cacheClient struct {
	gateway *gateway.HTTPGateway
}

func (c *cacheClient) Flush() error { return nil }

type healthClient struct {
	gateway *gateway.HTTPGateway
}

func (c *healthClient) Check() error { return nil }
