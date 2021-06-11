package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/steevehook/http/logging"
)

// routeParam fetches route param from context
func routeParam(r *http.Request, name string) string {
	ctx := r.Context()
	psCtx := ctx.Value(httprouter.ParamsKey)
	ps, ok := psCtx.(httprouter.Params)

	if !ok {
		logging.Logger().Error("could not extract params from context")
		return ""
	}
	return ps.ByName(name)
}
