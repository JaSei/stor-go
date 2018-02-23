package smbutil

import (
	"testing"

	"github.com/JaSei/pathutil-go"
	"github.com/stretchr/testify/assert"
)

func TestMountCmd(t *testing.T) {
	mntDir, _ := pathutil.New("/mnt")

	t.Run("dom-user-pass", func(t *testing.T) {
		uri := &SmbUri{
			Domain:   "avast",
			Username: "seidl",
			Password: "heslo",
			Hostname: "server.avast.com",
			Share:    "share",
			Path:     "file",
		}
		cmd, err := uri.MountCmd(mntDir)
		assert.NoError(t, err)

		assert.Equal(t, "mount.cifs server.avast.com /mnt -o domain=avast,user=seidl,pass=heslo", cmd)
	})

	t.Run("user-pass", func(t *testing.T) {
		uri := &SmbUri{
			Domain:   "",
			Username: "seidl",
			Password: "heslo",
			Hostname: "server.avast.com",
			Share:    "share",
			Path:     "file",
		}
		cmd, err := uri.MountCmd(mntDir)
		assert.NoError(t, err)

		assert.Equal(t, "mount.cifs server.avast.com /mnt -o user=seidl,pass=heslo", cmd)
	})

	t.Run("user", func(t *testing.T) {
		uri := &SmbUri{
			Domain:   "",
			Username: "seidl",
			Password: "",
			Hostname: "server.avast.com",
			Share:    "share",
			Path:     "file",
		}
		cmd, err := uri.MountCmd(mntDir)
		assert.NoError(t, err)

		assert.Equal(t, "mount.cifs server.avast.com /mnt -o user=seidl", cmd)
	})

	t.Run("without options", func(t *testing.T) {
		uri := &SmbUri{
			Domain:   "",
			Username: "",
			Password: "",
			Hostname: "server.avast.com",
			Share:    "share",
			Path:     "file",
		}
		cmd, err := uri.MountCmd(mntDir)
		assert.NoError(t, err)

		assert.Equal(t, "mount.cifs server.avast.com /mnt", cmd)
	})
}
