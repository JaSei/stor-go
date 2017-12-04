package storages

type storageType int

const (
	FILESTORAGE = iota
)

type storage struct {
	storageType storageType
	paths       []string
	priority    uint
}

type storages []storage

type Storage interface {
	Exists(hashutils.Hash) bool
	Get(hashutils.Hash) io.ReaderCloser, error
}
