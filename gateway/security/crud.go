package security

import (
	"errors"
	"odfe-cli/client"
	"odfe-cli/entity"
	"odfe-cli/entity/security"
	"odfe-cli/gateway"
)

var ErrInvalidCredentials = errors.New("Invalid credentials")

func NewActionGroupCRUD(c *client.Client, p *entity.Profile) (crud, error) {
	return &actionGroupCRUD{gateway: gateway.NewHTTPGateway(c, p)}, nil
}

func NewUsersCRUD(c *client.Client, p *entity.Profile) (crud, error) {
	return &usersCRUD{gateway: gateway.NewHTTPGateway(c, p)}, nil
}

func NewRolesCRUD(c *client.Client, p *entity.Profile) (crud, error) {
	return &rolesCRUD{gateway: gateway.NewHTTPGateway(c, p)}, nil
}

func NewRoleMappingsCRUD(c *client.Client, p *entity.Profile) (crud, error) {
	return &roleMappingsCRUD{gateway: gateway.NewHTTPGateway(c, p)}, nil
}

func NewTenantsCRUD(c *client.Client, p *entity.Profile) (crud, error) {
	return &tenantsCRUDCRUD{gateway: gateway.NewHTTPGateway(c, p)}, nil
}

type actionGroupCRUD struct {
	gateway *gateway.HTTPGateway
}

func (crud *actionGroupCRUD) Get() (interface{}, error)                 { return nil, nil }
func (crud *actionGroupCRUD) GetAll() (interface{}, error)              { return nil, nil }
func (crud *actionGroupCRUD) Delete(string) error                       { return nil }
func (crud *actionGroupCRUD) Create(string, interface{}) error          { return nil }
func (crud *actionGroupCRUD) Patch(string, security.PatchQuery) error   { return nil }
func (crud *actionGroupCRUD) PatchMultiple([]security.PatchQuery) error { return nil }

type usersCRUD struct {
	gateway *gateway.HTTPGateway
}

func (crud *usersCRUD) Get() (interface{}, error)                 { return nil, nil }
func (crud *usersCRUD) GetAll() (interface{}, error)              { return nil, nil }
func (crud *usersCRUD) Delete(string) error                       { return nil }
func (crud *usersCRUD) Create(string, interface{}) error          { return nil }
func (crud *usersCRUD) Patch(string, security.PatchQuery) error   { return nil }
func (crud *usersCRUD) PatchMultiple([]security.PatchQuery) error { return nil }

type rolesCRUD struct {
	gateway *gateway.HTTPGateway
}

func (crud *rolesCRUD) Get() (interface{}, error)                 { return nil, nil }
func (crud *rolesCRUD) GetAll() (interface{}, error)              { return nil, nil }
func (crud *rolesCRUD) Delete(string) error                       { return nil }
func (crud *rolesCRUD) Create(string, interface{}) error          { return nil }
func (crud *rolesCRUD) Patch(string, security.PatchQuery) error   { return nil }
func (crud *rolesCRUD) PatchMultiple([]security.PatchQuery) error { return nil }

type roleMappingsCRUD struct {
	gateway *gateway.HTTPGateway
}

func (crud *roleMappingsCRUD) Get() (interface{}, error)                 { return nil, nil }
func (crud *roleMappingsCRUD) GetAll() (interface{}, error)              { return nil, nil }
func (crud *roleMappingsCRUD) Delete(string) error                       { return nil }
func (crud *roleMappingsCRUD) Create(string, interface{}) error          { return nil }
func (crud *roleMappingsCRUD) Patch(string, security.PatchQuery) error   { return nil }
func (crud *roleMappingsCRUD) PatchMultiple([]security.PatchQuery) error { return nil }

type tenantsCRUDCRUD struct {
	gateway *gateway.HTTPGateway
}

func (crud *tenantsCRUDCRUD) Get() (interface{}, error)                 { return nil, nil }
func (crud *tenantsCRUDCRUD) GetAll() (interface{}, error)              { return nil, nil }
func (crud *tenantsCRUDCRUD) Delete(string) error                       { return nil }
func (crud *tenantsCRUDCRUD) Create(string, interface{}) error          { return nil }
func (crud *tenantsCRUDCRUD) Patch(string, security.PatchQuery) error   { return nil }
func (crud *tenantsCRUDCRUD) PatchMultiple([]security.PatchQuery) error { return nil }
