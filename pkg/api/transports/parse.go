package transports

import (
	"github.com/SimonRichardson/coherence/pkg/api"
	"github.com/SimonRichardson/coherence/pkg/api/client"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/pkg/errors"
)

// Strategy wraps a Transport protocol layer for querying requests.
type Strategy struct {
	fn func(string) api.Transport
}

// Apply a host to a Transport
func (s Strategy) Apply(host string) api.Transport {
	return s.fn(host)
}

// Parse a protocol transport config and return a Strategy for creating a
// Transport on demand.
func Parse(protocol string) (Strategy, error) {
	switch protocol {
	case "http":
		pooledClient := cleanhttp.DefaultPooledClient()
		return Strategy{
			fn: func(host string) api.Transport {
				return NewHTTPTransport(client.New(pooledClient, protocol, host))
			},
		}, nil
	case "nop":
		return Strategy{
			fn: func(host string) api.Transport {
				return Nop{}
			},
		}, nil
	default:
		return Strategy{}, errors.Errorf("unexpected protocol: %q", protocol)
	}
}
