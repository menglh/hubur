package ipranger

import (
	"github.com/menglh/hubur/mapcidr"
)

// Ips of a cidr
func Ips(cidr string) ([]string, error) {
	return mapcidr.IPAddresses(cidr)
}
