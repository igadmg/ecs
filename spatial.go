package ecs

import (
	"fmt"
)

var (
	errNotFound = fmt.Errorf("item not found")
)

type spatialNode[T any] struct {
	mask int64
	data []T
}

func newSpatialNode[T any](size int) *spatialNode[T] {
	return &spatialNode[T]{
		mask: 0,
		data: make([]T, size),
	}
}

type SpatialArray[T any] struct {
	chunkSize int
	nodes     []*spatialNode[T]
}

func (a *SpatialArray[T]) Get(index int) (item T, err error) {
	ni := index / a.chunkSize
	if ni >= len(a.nodes) || a.nodes[ni] == nil {
		err = errNotFound
		return
	}

	n := a.nodes[ni]
	ci := index % a.chunkSize
	if n.mask&(1<<ci) == 0 {
		err = errNotFound
		return
	}

	return n.data[ci], nil
}

func (a *SpatialArray[T]) Put(index int, item T) {
	ni := index / a.chunkSize
	if ni >= len(a.nodes) {
		a.nodes = append(a.nodes, make([]*spatialNode[T], ni-len(a.nodes))...)
	}

	n := a.nodes[ni]
	if n == nil {
		n = newSpatialNode[T](a.chunkSize)
		a.nodes[ni] = n
	}

	ci := index % a.chunkSize
	n.data[ci] = item
	n.mask |= 1 << ci
}
