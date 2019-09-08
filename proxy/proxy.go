package proxy

import (
	"returnconcrete/resource"
	"returnconcrete/resourcelist"
)

// A Proxy provides a Go interface to a service
type Proxy struct {
}

// New returns a new proxy to a service
func New() *Proxy {
	return &Proxy{}
}

// GetResources returns a list of resources from the service
func (p *Proxy) GetResources() (*resourcelist.ResourceList, error) {
	resources := []*resource.Resource{resource.New(1)}
	return resourcelist.New(resources...), nil
}
