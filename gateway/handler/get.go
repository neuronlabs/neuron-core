package handler

import (
	"context"
	"github.com/neuronlabs/neuron/internal"
	ictrl "github.com/neuronlabs/neuron/internal/controller"
	"github.com/neuronlabs/neuron/internal/models"
	"github.com/neuronlabs/neuron/log"
	"github.com/neuronlabs/neuron/mapping"
	"github.com/neuronlabs/neuron/query/scope"
	"net/http"
)

func (h *Handler) HandleGet(m *mapping.ModelStruct) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Debugf("[GET] begins for model: '%s'", m.Type().String())
		defer func() { log.Debugf("[GET] finished for model: '%s'.", m.Type().String()) }()

		ic := (*ictrl.Controller)(h.c)

		s, errs, err := ic.QueryBuilder().BuildScopeSingle(req.Context(), (*models.ModelStruct)(m), req.URL, nil)
		if err != nil {
			log.Errorf("[GET] Building Scope for the request failed: %v", err)
			h.internalError(req, rw)
			return
		}
		if len(errs) > 0 {
			log.Debugf("Building Get Scope failed. ClientSide Error: %v", errs)
			h.marshalErrors(req, rw, unsetStatus, errs...)
			return
		}

		// set controller into scope's context
		ctx := context.WithValue(s.Context(), internal.ControllerIDCtxKey, h.c)
		s.WithContext(ctx)

		log.Debugf("[REQ-SCOPE-ID] %s", (*scope.Scope)(s).ID().String())

		/**

		TO DO:

		- rewrite language filters

		*/

		if err := (*scope.Scope)(s).Get(); err != nil {
			h.handleDBError(req, err, rw)
			return
		}

		h.marshalScope((*scope.Scope)(s), req, rw)
	})
}
