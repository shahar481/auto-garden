package authentication

import (
	"auto-garden/crypto"
	"auto-garden/db"
	"auto-garden/db/queries"
	"auto-garden/web_server/api"
	"auto-garden/web_server/api/verification"
	"errors"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

var (
	authenticationParameters = map[string]func(param string) bool {
		"user": db.IsValidUser,
		"crypto": db.IsValidPassword,
	}
)

func Authenticate(ctx *atreugo.RequestCtx) error {
	if !verification.HasAllWantedVerifiedParameters(ctx, authenticationParameters) {
		return ctx.ErrorResponse(errors.New(api.BadRequestError), fasthttp.StatusBadRequest)
	}
	user := string(ctx.QueryArgs().Peek("user"))
	salt, err := queries.GetUserSalt(user)
	if err != nil {
		return ctx.ErrorResponse(errors.New(api.InvalidCredentialsError), fasthttp.StatusForbidden)
	}
	password := string(ctx.QueryArgs().Peek("password"))
	hashedPassword := crypto.HashPassword(password, salt)
	validCredentials, err := queries.IsValidUserCredentials(user, hashedPassword)
	if err != nil || !validCredentials {
		return ctx.ErrorResponse(errors.New(api.InvalidCredentialsError), fasthttp.StatusForbidden)
	}

	return nil
}