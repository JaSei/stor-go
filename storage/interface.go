package storage

import (
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type Storages []Storage

type Storage interface {
	// check hash existense in storage
	Exist(hashutils.Hash) bool
	// send file from storage to
	ServeFile(*fasthttp.Request, hashutils.Hash) error
}

func (storages Storages) Exists(hash hashutils.Hash) bool {
}

func (storage Storages) ServeFile(ctx *fasthttp.Request, hash hashutils.Hash) error {

}
