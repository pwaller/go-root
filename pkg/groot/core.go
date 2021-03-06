package groot

import (
	"fmt"
	"os"
)

var g_verbose = false

func printf(format string, args ...interface{}) {
	if g_verbose {
		fmt.Fprintf(os.Stderr, format, args...)
	}
}

func dprintf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	//printf(format, args...)
}

const (
	sz_int16  = 2
	sz_int32  = 4
	sz_int64  = 8
	sz_uint16 = 2
	sz_uint32 = 4
	sz_uint64 = 8

	g_START_BIG_FILE = 2000000000
)

// EOF
