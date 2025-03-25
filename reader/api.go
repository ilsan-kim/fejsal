package reader

import (
	"io"
	"time"
)

// StreamReader defined methods to handle reading data line-by-line from a stream source
type StreamReader interface {
	LoadNextLine() bool
	InputStream(input io.Reader)
	StringGetter(key any) func() (string, bool)
	IntGetter(key any) func() (int, bool)
	TimeGetter(key any, layout string) func() (time.Time, bool)

	read(key any) (string, bool)
}
