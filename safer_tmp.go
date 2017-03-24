package saferfs

import (
	"os"
	"runtime"
)

// This is in response to https://github.com/golang/go/issues/19695
// TempDir ensures that a path exists and is accessable before returning. This
// is important for things like Docker containers that are derived from the
// scratch base image.
func TempDir() (s string, err error) {
	dir := os.Getenv("TMPDIR")
	if dir == "" {
		if runtime.GOOS == "android" {
			dir = "/data/local/tmp"
		} else {
			dir = "/tmp"
		}
	}
	_, err = os.Stat(dir)
	if err != nil {
		return "", err
	}
	return dir, nil
}
