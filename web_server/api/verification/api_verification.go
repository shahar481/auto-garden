package verification

import "github.com/savsgio/atreugo/v11"

func HasAllWantedVerifiedParameters(ctx *atreugo.RequestCtx, parameters map[string]func(param string) bool) bool {
	q := ctx.QueryArgs()
	for index, fun := range parameters {
		if !q.Has(index) {
			return false
		}
		if !fun(string(q.Peek(index))) {
			return false
		}
	}
	return true
}
