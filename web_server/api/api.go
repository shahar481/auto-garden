package api

import (
	"auto-garden/web_server/api/plants"
	"github.com/savsgio/atreugo/v11"
)

var (
	apiFunctions = map[string]func(ctx *atreugo.RequestCtx) error {
		"/should-water": plants.ShouldWaterRequest,
	}
)

func SetHTTPFunctions(server *atreugo.Atreugo) {
	apiGroup := server.NewGroupPath(apiWebPath)
	for path, handlingFunction := range apiFunctions {
		apiGroup.GET(path, handlingFunction)
	}
}
