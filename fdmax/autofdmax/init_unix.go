//go:build linux || openbsd || netbsd

package autofdmax

import (
	"github.com/menglh/hubur/fdmax"
)

func init() {
	_ = fdmax.Set(fdmax.UnixMax)
}
