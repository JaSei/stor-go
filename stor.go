package main

import (
	"fmt"
	"log"
	//	"math/rand"
	"time"

	//	"github.com/alecthomas/kingpin"
	//	"github.com/alexcesaro/statsd"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var version = "master"

func main() {

	storList := storages{storage{FILESTORAGE, []string{
		"/nfs/prg24-004.srv.int.avast.com/data/storage/Samples",
		"/nfs/prg24-010.srv.int.avast.com/data/storage/Samples",
	}, 0}, storage{FILESTORAGE, []string{"b1", "b2"}, 0}}

	fmt.Println("%+v", storList)

	router := fasthttprouter.New()
	router.GET("/:sha256", getFileHandler)
	router.HEAD("/:sha256", getFileHandler)
	router.GET("/", aboutHandler)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

func getFileHandler(ctx *fasthttp.RequestCtx) {
	startTime := time.Now()
	defer func() {
		log.Println(ctx.Response.StatusCode())
		log.Println(ctx.Response.Header.ContentLength())
		log.Println(time.Since(startTime))
	}()

	sha256 := ctx.UserValue("sha256").(string)

	if sha256 == "status" {
		statusHandler(ctx)
	} else {
		path := sha256[0:2] + "/" + sha256[2:4] + "/" + sha256[4:6] + "/" + sha256

		fasthttp.ServeFile(ctx, path)
	}

}

func aboutHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "This is <a href=\"github.com/avast/stor-go\">github.com/avast/stor-go</a> %s", version)
}

func statusHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(200)
}
