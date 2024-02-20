package port

import (
	"fmt"
	"projectdiscovery/hubur/protocol"
)

type Port struct {
	Port     int               `json:"port"`
	Protocol protocol.Protocol `json:"protocol"`
	TLS      bool              `json:"tls"`
}

func (p *Port) String() string {
	return fmt.Sprintf("%d-%d-%v", p.Port, p.Protocol, p.TLS)
}
