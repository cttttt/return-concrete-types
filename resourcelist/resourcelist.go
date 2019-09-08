package resourcelist

import "returnconcrete/resource"

// A ResourceList encapsulates a slice of resources
type ResourceList struct {
	resources []*resource.Resource
}

// New returns a new ResourceList
func New(resources ...*resource.Resource) *ResourceList {
	return &ResourceList{resources: resources}
}

// Get fetches the i's element in this resource list overrunning the list will
// cause panic
func (r *ResourceList) Get(i int) *resource.Resource {
	return r.resources[i]
}

// Len returns the length of the resource list
func (r *ResourceList) Len() int {
	return len(r.resources)
}
