package nfs

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type FS string

// NewFS returns a new FS mounted under the given mountPoint. It will error
// if the mount point can't be read.
func NewFS(mountPoint string) (FS, error) {
	info, err := os.Stat(mountPoint)
	if err != nil {
		return "", fmt.Errorf("could not read %s: %s", mountPoint, err)
	}
	if !info.IsDir() {
		return "", fmt.Errorf("mount point %s is not a directory", mountPoint)
	}

	return FS(mountPoint), nil
}

// Path appends the given path elements to the filesystem path, adding separators
// as necessary.
func (fs FS) Path(p ...string) string {
	return filepath.Join(append([]string{string(fs)}, p...)...)
}

func regularExpress() *regexp.Regexp {
	return regexp.MustCompile(`^yyds://(loc|pub\.loc|priv\.loc)(/.*)$`)
}

func Parse(url string) *FS {
	reg := regularExpress()

	if reg.Match([]byte(url)) {
		sub := reg.FindStringSubmatch(url)
		store := sub[0]
		ext := sub[1]
		switch store {
		case "loc":
			fallthrough
		case "pub.loc":
			return NewPublicLoc(ext)
		case "priv.loc":
			return NewPrivateLoc(ext)
		}
	}
	return nil

}

func NewPublicLoc(ext string) *FS {
	return nil
}

func NewPrivateLoc(ext string) *FS {
	return nil
}
