package web_server

import "github.com/savsgio/atreugo/v11"

type apiFunction struct {
	Func func(ctx *atreugo.RequestCtx) error
	Parameters map[string]func(param string) bool
	Path string
	Server *atreugo.Router
	BeforeMiddleware func(ctx *atreugo.RequestCtx) error
}

type apiGroup struct {
	Functions map[string]apiFunction
	Server *atreugo.Router
	Middleware func(ctx *atreugo.RequestCtx) error
	subRouter *atreugo.Router
}

type apiInterface interface {
	SetChilds()
}

func (f apiFunction) SetChilds() {
	get := f.Server.GET(f.Path, f.Func)
	if f.BeforeMiddleware != nil {
		get.UseBefore(f.BeforeMiddleware)
	}
}

func (g apiGroup) SetChilds() {

}