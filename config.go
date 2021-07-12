package datamesh

import "github.com/openziti/foundation/transport"

type Config struct {
	Listeners []*ListenerConfig
	Dialers   []*DialerConfig
}

type ListenerConfig struct {
	Id            string
	BindAddress   transport.Address
	Advertisement transport.Address
}

type DialerConfig struct {
	Id          string
	BindAddress transport.Address
}
