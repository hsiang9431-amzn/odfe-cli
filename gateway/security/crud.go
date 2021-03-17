package security

import (
	"errors"
	"odfe-cli/entity/security"
)

var ErrInvalidCredentials = errors.New("Invalid credentials")

func NewActionGroupCRUD(u, p string) (crud, error) {
	if u == "" || p == "" {
		return nil, ErrInvalidCredentials
	}
	return &actionGroupCRUD{username: u, password: p}, nil
}

func NewUsersCRUD(u, p string) (crud, error) {
	if u == "" || p == "" {
		return nil, ErrInvalidCredentials
	}
	return &usersCRUD{username: u, password: p}, nil
}

func NewRolesCRUD(u, p string) (crud, error) {
	if u == "" || p == "" {
		return nil, ErrInvalidCredentials
	}
	return &rolesCRUD{username: u, password: p}, nil
}

func NewRoleMappingsCRUD(u, p string) (crud, error) {
	if u == "" || p == "" {
		return nil, ErrInvalidCredentials
	}
	return &roleMappingsCRUD{username: u, password: p}, nil
}

func NewTenantsCRUD(u, p string) (crud, error) {
	if u == "" || p == "" {
		return nil, ErrInvalidCredentials
	}
	return &tenantsCRUDCRUD{username: u, password: p}, nil
}

type actionGroupCRUD struct {
	username string
	password string
}

func (crud *actionGroupCRUD) Get() (interface{}, error)                 { return nil, nil }
func (crud *actionGroupCRUD) GetAll() (interface{}, error)              { return nil, nil }
func (crud *actionGroupCRUD) Delete(string) error                       { return nil }
func (crud *actionGroupCRUD) Create(string, interface{}) error          { return nil }
func (crud *actionGroupCRUD) Patch(string, security.PatchQuery) error   { return nil }
func (crud *actionGroupCRUD) PatchMultiple([]security.PatchQuery) error { return nil }

type usersCRUD struct {
	username string
	password string
}

func (crud *usersCRUD) Get() (interface{}, error)                 { return nil, nil }
func (crud *usersCRUD) GetAll() (interface{}, error)              { return nil, nil }
func (crud *usersCRUD) Delete(string) error                       { return nil }
func (crud *usersCRUD) Create(string, interface{}) error          { return nil }
func (crud *usersCRUD) Patch(string, security.PatchQuery) error   { return nil }
func (crud *usersCRUD) PatchMultiple([]security.PatchQuery) error { return nil }

type rolesCRUD struct {
	username string
	password string
}

func (crud *rolesCRUD) Get() (interface{}, error)                 { return nil, nil }
func (crud *rolesCRUD) GetAll() (interface{}, error)              { return nil, nil }
func (crud *rolesCRUD) Delete(string) error                       { return nil }
func (crud *rolesCRUD) Create(string, interface{}) error          { return nil }
func (crud *rolesCRUD) Patch(string, security.PatchQuery) error   { return nil }
func (crud *rolesCRUD) PatchMultiple([]security.PatchQuery) error { return nil }

type roleMappingsCRUD struct {
	username string
	password string
}

func (crud *roleMappingsCRUD) Get() (interface{}, error)                 { return nil, nil }
func (crud *roleMappingsCRUD) GetAll() (interface{}, error)              { return nil, nil }
func (crud *roleMappingsCRUD) Delete(string) error                       { return nil }
func (crud *roleMappingsCRUD) Create(string, interface{}) error          { return nil }
func (crud *roleMappingsCRUD) Patch(string, security.PatchQuery) error   { return nil }
func (crud *roleMappingsCRUD) PatchMultiple([]security.PatchQuery) error { return nil }

type tenantsCRUDCRUD struct {
	username string
	password string
}

func (crud *tenantsCRUDCRUD) Get() (interface{}, error)                 { return nil, nil }
func (crud *tenantsCRUDCRUD) GetAll() (interface{}, error)              { return nil, nil }
func (crud *tenantsCRUDCRUD) Delete(string) error                       { return nil }
func (crud *tenantsCRUDCRUD) Create(string, interface{}) error          { return nil }
func (crud *tenantsCRUDCRUD) Patch(string, security.PatchQuery) error   { return nil }
func (crud *tenantsCRUDCRUD) PatchMultiple([]security.PatchQuery) error { return nil }
