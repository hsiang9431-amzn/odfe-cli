package security

import (
	"odfe-cli/client"
	"odfe-cli/entity"
	"odfe-cli/entity/security"
	"odfe-cli/gateway"
)

const (
	accountURL = baseURL + "/account"

	baseURL         = "_opendistro/_security/api"
	actiongroupsURL = baseURL + "/actiongroups"
	usersURL        = baseURL + "/internalusers"
	rolesURL        = baseURL + "/roles"
	rolesmappingURL = baseURL + "/rolesmapping"
	tenantsURL      = baseURL + "/tenants"

	configURL = baseURL + "/securityconfig"
	cacheURL  = baseURL + "/cache"
	healthURL = baseURL + "/health"
)

type Gateway struct {
	Account account
	Config  config
	Cache   cache
	Health  health

	ActionGroup  crud
	Users        crud
	Roles        crud
	RoleMappings crud
	Tenants      crud
}

func NewGateway(c *client.Client, p *entity.Profile) (*Gateway, error) {
	gw := gateway.NewHTTPGateway(c, p)
	return &Gateway{
		Account: &accountClient{gateway: gw},
		Config:  &configClient{gateway: gw},
		Cache:   &cacheClient{gateway: gw},
		Health:  &healthClient{gateway: gw},

		ActionGroup:  &actionGroupCRUD{gateway: gw},
		Users:        &usersCRUD{gateway: gw},
		Roles:        &rolesCRUD{gateway: gw},
		RoleMappings: &roleMappingsCRUD{gateway: gw},
		Tenants:      &tenantsCRUD{gateway: gw},
	}, nil
}

type account interface {
	GetDetails() (interface{}, error)
	ChangePassword(security.ChangePasswordQuery) error
}

type crud interface {
	Get() (interface{}, error)
	GetAll() (interface{}, error)
	Delete(string) error
	Create(string, interface{}) error
	Patch(string, security.PatchQuery) error
	PatchMultiple([]security.PatchQuery) error
}

type config interface {
	Get() (interface{}, error)
	Update(interface{}) error
	Patch(security.PatchQuery) error
	PatchMultiple([]security.PatchQuery) error
}

type cache interface {
	Flush() error
}

type health interface {
	Check() error
}
