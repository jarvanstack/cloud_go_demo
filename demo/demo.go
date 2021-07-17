package demo

import (
	"github.com/dengjiawen8955/go_utils/restfulu"
	"github.com/dengjiawen8955/goweb/goweb"
)

func Start() {
	web := goweb.NewWeb("/v1")
	web.Get("/ping", func(ctx *goweb.Context) {
		ctx.Json(restfulu.Ok("PONG"))
	})
	web.RunHTTP(9090)
}
