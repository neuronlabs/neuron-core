package controller

import "C"
import (
	"context"
	"time"

	"github.com/neuronlabs/neuron/mapping"
	"github.com/neuronlabs/neuron/repository"
)

// Options is the structure that contains, initialize and control the flow of the application.
// It contains repositories, model definitions.
type Controller struct {
	// Config is the configuration struct for the controller.
	Options *Options
	// ModelMap is a mapping for the model's structures.
	ModelMap *mapping.ModelMap
	// Repositories are the controller registered repositories.
	Repositories map[string]repository.Repository
	// ModelRepositories is the mapping between the model and related repositories.
	ModelRepositories map[*mapping.ModelStruct]repository.Repository
	// DefaultService is the default service or given controller.
	DefaultRepository repository.Repository
}

// NewController creates new controller for given config 'cfg'.
func New(options *Options) *Controller {
	return newController(options)
}

// NewDefault creates new default controller based on the default config.
func NewDefault() *Controller {
	return newController(defaultOptions())
}

// Now gets and returns current timestamp. If the configs specify the function might return UTC timestamp.
func (c *Controller) Now() time.Time {
	ts := time.Now()
	if c.Options.UTCTimestamps {
		ts = ts.UTC()
	}
	return ts
}

type ctxKeyS struct{}

var ctxKey = &ctxKeyS{}

// CtxStore stores the controller in the context.
func CtxStore(parent context.Context, c *Controller) context.Context {
	return context.WithValue(parent, ctxKey, c)
}

// CtxGet gets the controller from the context.
func CtxGet(ctx context.Context) (*Controller, bool) {
	c, ok := ctx.Value(ctxKey).(*Controller)
	return c, ok
}

func newController(options *Options) *Controller {
	c := &Controller{
		Repositories:      map[string]repository.Repository{},
		ModelRepositories: map[*mapping.ModelStruct]repository.Repository{},
	}
	if options == nil {
		options = &Options{}
		options.NamingConvention = mapping.SnakeCase
	}
	c.Options = options
	c.ModelMap = mapping.NewModelMap(c.Options.NamingConvention)
	return c
}
