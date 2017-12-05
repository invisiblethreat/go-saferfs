package saferfs

import (
	"testing"

	"github.com/satori/go.uuid"
)

func TestTempDir(t *testing.T) {

	// Empty case, allowing for the OS environment variable to be used.
	_, err := TempDir("", "saferfs-testing")
	if err != nil {
		t.Errorf("Failed to create temporary directory: %s", err.Error())
	}

	dir := "/" + uuid.NewV4().String()
	_, err = TempDir(dir, "saferfs-testing")
	if err == nil {
		t.Errorf("Directory creation should have failed for: %s", dir)
	}

}

func TestTempFile(t *testing.T) {
	// Empty case, allowing for the OS environment variable to be used.
	defaultTempFile, err := TempFile("", "saferfs-testing")
	if err != nil {

		t.Errorf("Failed to create temporary file: %s, with error: %s",
			defaultTempFile.Name(),
			err.Error())
	}

	dir := "/" + uuid.NewV4().String()
	defaultTempFile, err = TempFile(dir, "saferfs-testing")
	if err == nil {
		t.Errorf("Directory creation should have failed for file: %s, with error: %s",
			dir,
			err.Error())
	}
}
