package commands

// this file is strictly for defining cobra commands
// used in security plugin, all futher logic is implemented
// in handler/security and controller/security

const (
	securityCmd = "security"

	securityCreateCmd = "create"
	securityGetCmd    = "get"
	securityPatchCmd  = "patch"
	securityDeleteCmd = "delete"

	securityAccountCmd   = "account"
	securityAccountAlias = "ac"

	securityAGCmd   = "action-groups"
	securityAGAlias = "ag"

	securityUserCmd   = "user"
	securityUserAlias = "u"

	securityRoleCmd   = "role"
	securityRoleAlias = "r"

	securityRMCmd   = "role-mapping"
	securityRMAlias = "rmp"

	securityTenantCmd   = "tenant"
	securityTenantAlias = "tnt"

	securityConfigCmd = "config"
	securityCacheCmd  = "flush-cache"
	securityHealthCmd = "health-check"

	securityInputFlag       = "file"
	securityInputFlagShort  = "f"
	securityOutputWideFlag  = "wide"
	securityOutputWideShort = "w"
	securityOutputJsonFlag  = "json"
	securityOutputYamlFlag  = "yaml"
)
