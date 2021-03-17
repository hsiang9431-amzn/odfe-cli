package security

import (
	"io"
	gw "odfe-cli/gateway/security"
	"os"
)

type resourceController struct {
	reader  io.Reader
	gateway gw.Gateway
	mapper  interface{}
	writer  io.Writer
}

func NewCRUDController(r io.Reader, gw gw.Gateway, mp interface{}, w io.Writer) (*resourceController, error) {
	if w == nil {
		w = os.Stdout
	}
	return &resourceController{reader: r, gateway: gw, mapper: mp, writer: w}, nil
}

func (rc *resourceController) Get(resource string) error    { return nil }
func (rc *resourceController) Delete(resource string) error { return nil }
func (rc *resourceController) Create(resource string) error { return nil }
func (rc *resourceController) Patch(resource string) error  { return nil }

type controller struct {
	reader  io.Reader
	gateway gw.Gateway
	mapper  interface{}
	writer  io.Writer
}

func NewController(r io.Reader, gw gw.Gateway, mp interface{}, w io.Writer) (*controller, error) {
	if w == nil {
		w = os.Stdout
	}
	return &controller{reader: r, gateway: gw, mapper: mp, writer: w}, nil
}

func (ctrl *controller) AccountGetDetails() error     { return nil }
func (ctrl *controller) AccountChangePassword() error { return nil }

func (ctrl *controller) GetConfig() error    { return nil }
func (ctrl *controller) UpdateConfig() error { return nil }
func (ctrl *controller) PatchConfig() error  { return nil }

func (ctrl *controller) CacheFlush() error  { return nil }
func (ctrl *controller) HealthCheck() error { return nil }
