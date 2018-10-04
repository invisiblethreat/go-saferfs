saferfs
======

## Safer Filesystem calls
This library is to make some filesystem calls safer than they currently are in
the standard Golang libraries.

In the age of containers and other interesting filesystems, some of the previous
"this will always be present" items have ceased to exist. This, in turn, exposes
design decisions that could not have anticipated these changes.

In particular, the Go libraries in the `os` package for interacting with `/tmp`
returns a hardcoded path if no environment variable is set. When using Docker,
and a `Dockerfile` that has `FROM scratch`, this path will be returned at `/tmp`
and cause an error.

This can't be fixed in Go 1.x due to it breaking the function signature.
Context: https://github.com/golang/go/issues/19695

## Functions

### `TempDir() (dir string, err error)`
This function reads the environment variable `TMPDIR`, which, if empty, defaults
to hardcoded paths. Before returning, the existence of the path is verified.
