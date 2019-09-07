package consumer

import "fmt"

type stringable interface {
	String() string
}

type resource interface {
	stringable
	ID() int
}

type proxy interface {
	GetResources() ([]resource, error)
}

// A Consumer uses a service through a proxy
type Consumer struct {
	service proxy
}

// New returns a new consumer
func New(service proxy) *Consumer {
	return &Consumer{
		service: service,
	}
}

// Do prints all of the resources from the service
func (c *Consumer) Do() error {
	resources, err := c.service.GetResources()

	if err != nil {
		return err
	}

	for _, resource := range resources {
		fmt.Println(resource.String())
	}

	return nil
}
