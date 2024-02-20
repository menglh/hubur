package asnmap

import (
	"net"

	"github.com/menglh/hubur/mapcidr"
)

func GetCIDR(output []*Response) ([]*net.IPNet, error) {
	var cidrs []*net.IPNet
	for _, res := range output {
		cidr, err := mapcidr.GetCIDRFromIPRange(net.ParseIP(res.FirstIp), net.ParseIP(res.LastIp))
		if err != nil {
			return nil, err
		}
		cidrs = append(cidrs, cidr...)
	}
	return cidrs, nil
}
