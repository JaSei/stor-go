package smbutil

import (
	"fmt"
	"strings"

	"github.com/JaSei/pathutil-go"
	"github.com/pkg/errors"
)

/*
MountCmd prepare mount.cifs command for run
*/

func (uri *SmbUri) MountCmd(mntDir pathutil.Path) (string, error) {
	if !mntDir.IsDir() {
		return "", errors.Errorf("Local mount dir %s is not directory", mntDir)
	}

	cmd := fmt.Sprintf("mount.cifs %s %s", uri.Hostname, mntDir)

	options := make([]string, 0)
	if uri.Domain != "" {
		options = append(options, "domain="+uri.Domain)
	}
	if uri.Username != "" {
		options = append(options, "user="+uri.Username)
	}
	if uri.Password != "" {
		options = append(options, "pass="+uri.Password)
	}

	if len(options) != 0 {
		cmd += " -o " + strings.Join(options, ",")
	}

	return cmd, nil
}
