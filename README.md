saferfs
======

## Safer Filesystem calls
This library is to make some filesystem calls safer than they currently are in
the standard Golang libraries.

In the age of containers and other interesting filesystems, some of the previous
"this will always be present" items have ceased to exist. This, in turn, exposes
design decisions that could not have anticipated these changes.

## Functions

### `TempDir() (s string, err error)`
This function reads the environment variable `TMPDIR`, which, if empty, defaults
to hardcoded paths. Before returning, the existence of the path is verified.