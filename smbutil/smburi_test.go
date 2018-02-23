package smbutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUri(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		uri, err := ParseUri("smb://avast;seidl:heslo@server.avast.com/share/file")
		assert.NoError(t, err)
		assert.Equal(t, &SmbUri{
			Domain:   "avast",
			Username: "seidl",
			Password: "heslo",
			Hostname: "server.avast.com",
			Share:    "share",
			Path:     "file",
		}, uri)
	})

	t.Run("username-hostname-share-path", func(t *testing.T) {
		uri, err := ParseUri("smb://seidl@server.avast.com/share/file")
		assert.NoError(t, err)
		assert.Equal(t, &SmbUri{
			Domain:   "",
			Username: "seidl",
			Password: "",
			Hostname: "server.avast.com",
			Share:    "share",
			Path:     "file",
		}, uri)
	})

	t.Run("hostname-share-path", func(t *testing.T) {
		uri, err := ParseUri("smb://server.avast.com/share/file")
		assert.NoError(t, err)
		assert.Equal(t, &SmbUri{
			Domain:   "",
			Username: "",
			Password: "",
			Hostname: "server.avast.com",
			Share:    "share",
			Path:     "file",
		}, uri)
	})

	t.Run("hostname-share-path", func(t *testing.T) {
		uri, err := ParseUri("smb://server.avast.com/share/file")
		assert.NoError(t, err)
		assert.Equal(t, &SmbUri{
			Domain:   "",
			Username: "",
			Password: "",
			Hostname: "server.avast.com",
			Share:    "share",
			Path:     "file",
		}, uri)
	})

	t.Run("hostname-share", func(t *testing.T) {
		uri, err := ParseUri("smb://server.avast.com/share")
		assert.NoError(t, err)
		assert.Equal(t, &SmbUri{
			Domain:   "",
			Username: "",
			Password: "",
			Hostname: "server.avast.com",
			Share:    "share",
			Path:     "",
		}, uri)
	})

	t.Run("hostname-share", func(t *testing.T) {
		uri, err := ParseUri("smb://server.avast.com")
		assert.Error(t, err)
		assert.Nil(t, uri)
	})
}
