package ecs

type spatialNode[T any] struct {
	data []T
}

func newSpatialNode[T any](size int) spatialNode[T] {
	return spatialNode[T]{
		data: make([]T, size),
	}
}

type spatialArray[T any, ChunkSize int] struct {
	nodes []*spatialNode[T]
}
