/*
storagefactory ...

	smb://domain;user:pass@host/share/path
	nfs://host:port/path
	s3://aswKeyId:aswSecretAccessKey@hostname/bucket


	storageFactory := storagefactory.New(s3storage.New, filestorage.New)
	storages, err := storageFactory.ParseFromStrings([]string{
		"file;/tmp/dir,/tmp/other",
		"s3;http://s3-us-west-2.amazonaws.com/mybucket;1"
	})

	if storages.Exists(sha256) {
		if err := storages.ServeFile(sha256, ctx); err != nil {
			log.Error(err)
		}
	}
*/
package storagefactory

import (
	"github.com/avast/stor-go/storage"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type SupportedStorages []SupportedStorage

type SupportedStorage func(string, uint16) storage.Storage

// New create new Storage factory initialized by SupportedStorage
//
// storagefactory.New(s3storage.New, filestorage.New, storstorage.New)
func New(supports ...SupportedStorage) SupportedStorages {
	registered := make([]Storage, len(supports))
}

func (ss SupportedStorages) ParseFromStrings(strStorages []string) (Storages, error) {
	var storages Storages
	var multiErr multierror.Error

	for _, str := range strStorages {
		storage, err := splitStorageString(str)
		if err != nil {
			multierror.Append(multiErr, err)
		} else {
			storages = append(storages, storage)
		}
	}

	return storages, multiErr.ErrorOrNil()
}

func (ss SupportedStorage) splitStorageString(str string) (Storage, error) {
	storageParts = strings.SplitN(str, ";", 3)

	if len(storageParts) < 2 || len(storageParts) > 3 {
		return errors.New("Invalid storage format")
	}

	paths := strings.Split(storageParts[1], ",")

	var storage Storage
	switch storageParts[0] {
	case "file":
		storage = filestorage.New(paths, priority)
	default:
		return storage, errors.Errorf("Storage type '%s' isn't supported", storageParts[0])
	}

	return storage, nil
}
