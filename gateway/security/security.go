package security

import "odfe-cli/entity/security"

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
