package api

import (
	"auto-garden/db"
	"auto-garden/web_server/api/verification"
	"github.com/savsgio/atreugo/v11"
	"strconv"
)

var (
	runningSessionParameters = map[string]func(param string) bool {
		"session-id": db.IsInt8,
	}
)


func CheckAuthenticated(ctx *atreugo.RequestCtx) error {
	if !verification.HasAllWantedVerifiedParameters(ctx, runningSessionParameters) {
		return ctx.RedirectResponse("/", 304)
	}
	id, _ := strconv.ParseInt(string(ctx.QueryArgs().Peek("session-id")), 10, 64)
	valid, err := db.IsSessionValid(id)
	if err != nil || !valid {
		return ctx.RedirectResponse("/", 304)
	}
	return ctx.Next()
}

func CheckWantedParameters(ctx *atreugo.RequestCtx) error {

	return ctx.Next()
}