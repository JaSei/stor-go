/*
smbutil ...

	uri, err := ParseUri("smb://domain;user:pass@host/share/path")
	if err != nil {
		panic(err)
	}

	err := exec.Run(uri.MountCmd())
	if err != nil {
		panic(err)
	}

*/
package smbutil

import (
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type SmbUri struct {
	Hostname string
	Domain   string
	Username string
	Password string
	Share    string
	Path     string
}

/*
ParseUri parse samba / cifs URI
inspired by https://lists.samba.org/archive/samba-technical/2001-April/012999.html

URI structure

	smb://[[domain;]user[:pass]@]host/share[/path]

*/
func ParseUri(uriString string) (*SmbUri, error) {
	u, err := url.Parse(uriString)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "smb" {
		return nil, errors.Errorf("Unsupported scheme %s", u.Scheme)
	}

	uri := &SmbUri{
		Hostname: u.Hostname(),
	}

	if u.User != nil {
		if err := parseUserUriPart(u.User.String(), uri); err != nil {
			return nil, err
		}
	}

	if err := parsePathPart(u.Path, uri); err != nil {
		return nil, err
	}

	return uri, nil
}

func parseUserUriPart(userPart string, uri *SmbUri) error {
	if userPart == "" {
		return errors.New("username must be defined")
	}

	splitDomain := strings.SplitN(userPart, ";", 2)
	if len(splitDomain) == 2 {
		uri.Domain = splitDomain[0]
		userPass := strings.SplitN(splitDomain[1], ":", 2)

		if len(userPass) == 2 {
			uri.Username = userPass[0]
			uri.Password = userPass[1]
		} else {
			uri.Username = userPass[0]
		}
	} else {
		uri.Username = splitDomain[0]
	}

	return nil
}

func parsePathPart(path string, uri *SmbUri) error {
	p := strings.SplitN(path, "/", 3)

	switch len(p) {
	case 3:
		uri.Share = p[1]
		uri.Path = p[2]
	case 2:
		uri.Share = p[1]
	default:
		return errors.New("Share not defined")
	}

	return nil
}
