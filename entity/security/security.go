package security

type ChangePasswordQuery struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"password"`
}

type PatchQuery struct {
	OP   string `json:"op"`
	Path string `json:"path"`
	// value can be either []string or map[string]string
	Value interface{} `json:"value"`
}

type CreateActionGroupQuery struct {
	AllowedActions []string `json:"allowed_actions"`
}

type CreateUserQuery struct {
	Password     string            `json:"password"`
	BackendRoles []string          `json:"backend_roles"`
	Attributes   map[string]string `json:"attributes"`
}

type CreateRoleQuery struct {
	ClusterPermissions []string           `json:"cluster_permissions"`
	IndexPermissions   []IndexPermission  `json:"index_permissions"`
	TenantPermissions  []TenantPermission `json:"tenant_permissions"`
}

type IndexPermission struct {
	IndexPatterns  []string `json:"index_patterns"`
	DLS            string   `json:"dls"`
	FLS            []string `json:"fls"`
	MaskedFields   []string `json:"masked_fields"`
	AllowedActions []string `json:"allowed_actions"`
}

type TenantPermission struct {
	TenantPatterns []string `json:"tenant_patterns"`
	AllowedActions []string `json:"allowed_actions"`
}

type CreateRoleMappingQuery struct {
	BackendRoles []string `json:"backend_roles"`
	Hosts        []string `json:"hosts"`
	Users        []string `json:"users"`
}

type CreateTenantQuery struct {
	Description string `json:"description"`
}

type Config struct {
	Dynamic interface{} `json:"dynamic"`
}
