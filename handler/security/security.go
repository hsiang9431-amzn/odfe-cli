package security

import (
	"odfe-cli/entity"

	"github.com/spf13/cobra"
)

// handlers are cobra functions defined with signature
// func(cmd *cobra.Command, args []string)
// these functions can be used directly when declaring cobra.Command structures

func Accounts(p *entity.Profile) func(cmd *cobra.Command, args []string) {

	return nil
}

func ActionGroups(p *entity.Profile) func(cmd *cobra.Command, args []string) {

	return nil
}

func Users(p *entity.Profile) func(cmd *cobra.Command, args []string) {

	return nil
}

func Roles(p *entity.Profile) func(cmd *cobra.Command, args []string) {

	return nil
}

func RoleMappings(p *entity.Profile) func(cmd *cobra.Command, args []string) {

	return nil
}

func Tenants(p *entity.Profile) func(cmd *cobra.Command, args []string) {

	return nil
}

func Config(p *entity.Profile) func(cmd *cobra.Command, args []string) {

	return nil
}

func Cache(p *entity.Profile) func(cmd *cobra.Command, args []string) {

	return nil
}

func Health(p *entity.Profile) func(cmd *cobra.Command, args []string) {

	return nil
}
