package main

import (
	"fmt"
	"log"
	//	"math/rand"
	"time"

	"github.com/alecthomas/kingpin"
	//	"github.com/alexcesaro/statsd"
	"github.com/avast/stor-go/handler"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var version = "master"

var (
	port     = kingpin.Flag("port", "server port").Default("8080").Uint16()
	storages = kingpin.Flag("storage", "storage definition in format 'STORAGETYPE;PATH[,PATH,...][;PRIORITY]'").Required().Strings()
)

func main() {
	kingpin.Parse()

	storList := storages{storage{FILESTORAGE, []string{
		"/nfs/prg24-004.srv.int.avast.com/data/storage/Samples",
		"/nfs/prg24-010.srv.int.avast.com/data/storage/Samples",
	}, 0}, storage{FILESTORAGE, []string{"b1", "b2"}, 0}}

	fmt.Println("%+v", storList)

	router := fasthttprouter.New()
	router.GET("/:sha256", handler.GetFile)
	router.HEAD("/:sha256", handler.GetFile)
	router.GET("/", handler.About)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
