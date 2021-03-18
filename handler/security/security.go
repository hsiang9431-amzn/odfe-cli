package security

import (
	"odfe-cli/entity"

	"github.com/spf13/cobra"
)

// handlers are cobra functions defined with signature
// func(cmd *cobra.Command, args []string)
// these functions can be used directly when declaring cobra.Command structures

type handler struct {
	prof *entity.Profile
}

func NewHandler(p *entity.Profile) *handler {
	return &handler{prof: p}
}

func (h *handler) Handle(cmd *cobra.Command, args []string) {
	// route to correct function based on the command
}

func (h *handler) accounts(cmd *cobra.Command, args []string) {

}

func (h *handler) actionGroups(cmd *cobra.Command, args []string) {

}

func (h *handler) users(cmd *cobra.Command, args []string) {

}

func (h *handler) roles(cmd *cobra.Command, args []string) {

}

func (h *handler) roleMappings(cmd *cobra.Command, args []string) {

}

func (h *handler) tenants(cmd *cobra.Command, args []string) {

}

func (h *handler) config(cmd *cobra.Command, args []string) {

}

func (h *handler) cache(cmd *cobra.Command, args []string) {

}

func (h *handler) health(cmd *cobra.Command, args []string) {

}
