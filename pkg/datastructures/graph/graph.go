package graph

import (
	"errors"
	"sync"
)

var (
	// ErrVertexNotFound is returned when an operation is requested on a
	// non-existent vertex.
	ErrVertexNotFound = errors.New("vertex not found")

	// ErrSelfLoop is returned when an operation tries to create a disallowed
	// self loop.
	ErrSelfLoop = errors.New("self loops not permitted")

	// ErrParallelEdge is returned when an operation tries to create a
	// disallowed parallel edge.
	ErrParallelEdge = errors.New("parallel edges are not permitted")
)

type Graph struct {
	mutex         sync.RWMutex
	adjacencyList map[interface{}]map[interface{}]struct{}
	v, e          int
}

type Grapher interface {
	V() int
	E() int
	AddEdge(v, w interface{}) error
	Adj(v interface{}) ([]interface{}, error)
	Degree(v interface{}) (int, error)
	addVertex(v interface{})
}

func (g *Graph) V() int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return g.v
}

func (g *Graph) E() int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return g.e
}

func (g *Graph) AddEdge(v, w interface{}) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if v == w {
		return ErrSelfLoop
	}

	g.addVertex(v)
	g.addVertex(w)

	if _, ok := g.adjacencyList[v][w]; ok {
		return ErrParallelEdge
	}

	g.adjacencyList[v][w] = struct{}{}
	g.adjacencyList[w][v] = struct{}{}

	g.e++
	return nil
}

func (g *Graph) addVertex(v interface{}) {
	_, ok := g.adjacencyList[v]

	if !ok {
		g.adjacencyList[v] = make(map[interface{}]struct{})
		g.v++
	}
}

func (g *Graph) Degree(v interface{}) (int, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	val, ok := g.adjacencyList[v]
	if !ok {
		return 0, ErrVertexNotFound
	}
	return len(val), nil
}

func (g *Graph) Adj(v interface{}) ([]interface{}, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	deg, err := g.Degree(v)
	if err != nil {
		return nil, ErrVertexNotFound
	}

	adj := make([]interface{}, deg)
	i := 0

	for key := range g.adjacencyList[v] {
		adj[i] = key
		i++
	}
	return adj, nil
}

func New() *Graph {
	return &Graph{
		adjacencyList: make(map[interface{}]map[interface{}]struct{}),
		v:             0,
		e:             0,
	}
}
