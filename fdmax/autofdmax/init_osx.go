//go:build darwin

package autofdmax

import "github.com/menglh/hubur/fdmax"

func init() {
	_ = fdmax.Set(fdmax.OSXMax)
}
