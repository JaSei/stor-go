package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

func GetFile(ctx *fasthttp.RequestCtx) {
	startTime := time.Now()
	defer func() {
		log.Println(ctx.Response.StatusCode())
		log.Println(ctx.Response.Header.ContentLength())
		log.Println(time.Since(startTime))
	}()

	sha256 := ctx.UserValue("sha256").(string)

	if sha256 == "status" {
		Status(ctx)
	} else {
		path := sha256[0:2] + "/" + sha256[2:4] + "/" + sha256[4:6] + "/" + sha256

		fasthttp.ServeFile(ctx, path)
	}

}

func About(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "This is <a href=\"github.com/avast/stor-go\">github.com/avast/stor-go</a> %s", version)
}

func Status(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(200)
}
