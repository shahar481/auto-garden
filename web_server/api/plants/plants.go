package plants

import (
	"auto-garden/db"
	"errors"
	"fmt"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"strconv"
)

var (
	shouldWaterParameters = map[string]func(param string) bool {
		"plant_id":         db.IsInt8,
		"currently_on":     db.IsBool,
		"current_humidity": db.IsInt2}
)


func hasAllWantedVerifiedParameters(ctx *atreugo.RequestCtx, parameters map[string]func(param string) bool) bool {
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

func ShouldWaterRequest(ctx *atreugo.RequestCtx) error {
	if !hasAllWantedVerifiedParameters(ctx, shouldWaterParameters) {
		return ctx.ErrorResponse(errors.New("missing parameters\n"), fasthttp.StatusBadRequest)
	}
	id, _ := strconv.ParseInt(string(ctx.QueryArgs().Peek("plant_id")), 10, 64)
	currentlyOn := string(ctx.QueryArgs().Peek("currently_on"))
	currentHumidity, _ := strconv.ParseInt(string(ctx.QueryArgs().Peek("current_humidity")), 10, 16)

	water, err := db.ShouldWater(id, currentlyOn == "0", currentHumidity)
	if err != nil {
		return ctx.ErrorResponse(errors.New("Bad request\n"), fasthttp.StatusBadRequest)
	}

	response := fmt.Sprintf("%+v", water)
	return ctx.HTTPResponse(response, fasthttp.StatusOK)
}

