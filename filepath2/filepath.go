// Package filepath implements utility routines for manipulating filename paths.
package filepath2

import (
	"path/filepath"
	"strings"
)

// Name returns the filename of path without extension.
func Name(path string) string {
	base := filepath.Base(path)
	if dot := strings.LastIndexByte(base, '.'); dot != -1 {
		return base[:dot]
	}
	return base
}

// ChangeExt returns a copy of path but with extension
// changed to newExt. If path has no extension, newExt
// is added as the extension.
func ChangeExt(path, newExt string) string {
	if newExt != "" && !strings.HasPrefix(newExt, ".") {
		newExt = "." + newExt
	}
	return filepath.Join(filepath.Dir(path), Name(path)+newExt)
}
