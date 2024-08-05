package archiving

import (
	"io"

	"github.com/docker/docker/pkg/archive"
)

// Tar creates an archive from the directory at `path`, and returns it as a stream of bytes.
//
// Arguments:
//   - path: The path to the directory to archive.
//
// Returns:
//   - A stream of bytes representing the archive.
//   - An error if the archive could not be created.
func Tar(path string) (io.Reader, error) {
	return archive.Tar(path, archive.Uncompressed)
}
