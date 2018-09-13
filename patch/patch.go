package patch

import (
	"compress/bzip2"
	"io"
	"io/ioutil"

	"github.com/teamrazr/go-bsdiff/internal"
	"github.com/teamrazr/go-bsdiff/internal/native"
)

func Patch(oldReader io.Reader, newWriter io.Writer, patchReader io.Reader) (err error) {
	oldBytes, err := ioutil.ReadAll(oldReader)
	if err != nil {
		return
	}

	newLen, err := internal.ReadHeader(patchReader)
	if err != nil {
		return
	}
	newBytes := make([]byte, newLen)

	// Decompression
	bz2Reader := bzip2.NewReader(patchReader)

	err = native.Patch(oldBytes, newBytes, bz2Reader)

	newWriter.Write(newBytes)
	return
}
