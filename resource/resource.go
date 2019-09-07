package resource

import "fmt"

// A Resource is an object being persisted by a service
type Resource struct {
	id int
}

// ID returns the id of the resource
func (r *Resource) ID() int {
	return r.id
}

// String returns a string represetnation of the resource
func (r *Resource) String() string {
	return fmt.Sprintf("%d", r.id)
}

// New returns a new Resource
func New(id int) *Resource {
	return &Resource{id: id}
}
