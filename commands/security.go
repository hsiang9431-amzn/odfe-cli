package commands

import (
	handler "odfe-cli/handler/security"

	"github.com/spf13/cobra"
)

// this file is strictly for defining cobra commands
// used in security plugin, all futher logic is implemented
// in handler/security and controller/security

const (
	securityCmdName = "security"

	securityCreateCmdName = "create"
	securityGetCmdName    = "get"
	securityPatchCmdName  = "patch"
	securityDeleteCmdName = "delete"

	securityAccountCmdName = "account"
	securityAccountAlias   = "ac"

	securityAGCmdName = "action-group"
	securityAGAlias   = "ag"

	securityUserCmdName = "user"
	securityUserAlias   = "u"

	securityRoleCmdName = "role"
	securityRoleAlias   = "r"

	securityRMCmdName = "role-mapping"
	securityRMAlias   = "rmp"

	securityTenantCmdName = "tenant"
	securityTenantAlias   = "tnt"

	securityConfigCmdName = "config"
	securityCacheCmdName  = "flush-cache"
	securityHealthCmdName = "health-check"

	securityInputFlag       = "file"
	securityInputFlagShort  = "f"
	securityOutputWideFlag  = "wide"
	securityOutputWideShort = "w"
	securityOutputJsonFlag  = "json"
	securityOutputYamlFlag  = "yaml"
)

var securityCreateCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

var securityAGCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	Run:   handler.Accounts,
}
