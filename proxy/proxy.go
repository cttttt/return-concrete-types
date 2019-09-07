package proxy

import "returnconcrete/resource"

// A Proxy provides a Go interface to a service
type Proxy struct {
}

// New returns a new proxy to a service
func New() *Proxy {
	return &Proxy{}
}

// GetResources returns a list of resources from the service
func (p *Proxy) GetResources() ([]*resource.Resource, error) {
	return []*resource.Resource{resource.New(1)}, nil
}
