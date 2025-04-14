package os

import (
	"io"
	"os"

	"github.com/mkch/gg"
)

// CopyFile copies the content from src to dest, and sets dest with the same file mode as src.
// If overwrite is true, it will overwrite the existing content of dest.
func CopyFile(src, dest string, overwrite bool) (err error) {
	r, err := os.Open(src)
	if err != nil {
		return
	}
	defer gg.ChainError(r.Close, &err)
	srcInfo, err := r.Stat()
	if err != nil {
		return
	}
	w, err := os.OpenFile(dest,
		os.O_CREATE|os.O_WRONLY|gg.If(overwrite, os.O_TRUNC, os.O_EXCL),
		srcInfo.Mode())
	if err != nil {
		return
	}
	defer gg.ChainError(w.Close, &err)
	_, err = io.Copy(w, r)
	return
}
