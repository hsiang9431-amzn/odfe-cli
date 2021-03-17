package security

import (
	"errors"
	"odfe-cli/entity/security"
	"odfe-cli/gateway"
)

var ErrInvalidCredentials = errors.New("Invalid credentials")

type actionGroupCRUD struct {
	gateway *gateway.HTTPGateway
}

func (crud *actionGroupCRUD) Get(string) (interface{}, error)           { return nil, nil }
func (crud *actionGroupCRUD) GetAll() (interface{}, error)              { return nil, nil }
func (crud *actionGroupCRUD) Delete(string) error                       { return nil }
func (crud *actionGroupCRUD) Create(string, interface{}) error          { return nil }
func (crud *actionGroupCRUD) Patch(string, []security.PatchQuery) error { return nil }

type usersCRUD struct {
	gateway *gateway.HTTPGateway
}

func (crud *usersCRUD) Get(string) (interface{}, error)           { return nil, nil }
func (crud *usersCRUD) Delete(string) error                       { return nil }
func (crud *usersCRUD) Create(string, interface{}) error          { return nil }
func (crud *usersCRUD) Patch(string, []security.PatchQuery) error { return nil }

type rolesCRUD struct {
	gateway *gateway.HTTPGateway
}

func (crud *rolesCRUD) Get(string) (interface{}, error)           { return nil, nil }
func (crud *rolesCRUD) Delete(string) error                       { return nil }
func (crud *rolesCRUD) Create(string, interface{}) error          { return nil }
func (crud *rolesCRUD) Patch(string, []security.PatchQuery) error { return nil }

type roleMappingsCRUD struct {
	gateway *gateway.HTTPGateway
}

func (crud *roleMappingsCRUD) Get(string) (interface{}, error)           { return nil, nil }
func (crud *roleMappingsCRUD) Delete(string) error                       { return nil }
func (crud *roleMappingsCRUD) Create(string, interface{}) error          { return nil }
func (crud *roleMappingsCRUD) Patch(string, []security.PatchQuery) error { return nil }

type tenantsCRUD struct {
	gateway *gateway.HTTPGateway
}

func (crud *tenantsCRUD) Get(string) (interface{}, error)           { return nil, nil }
func (crud *tenantsCRUD) Delete(string) error                       { return nil }
func (crud *tenantsCRUD) Create(string, interface{}) error          { return nil }
func (crud *tenantsCRUD) Patch(string, []security.PatchQuery) error { return nil }
