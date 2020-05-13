// Portions of this file are derived from JGraphT, a free Java graph-theory
// library.
//
// (C) Copyright 2006-2016, by France Telecom and Contributors.

package algo

import (
	"math"

	"github.com/puppetlabs/horsehead/v2/graph"
)

const (
	BellmanFordSupportedFeatures = graph.DeterministicIteration
)

type bellmanFordPathElement struct {
	vertex graph.Vertex
	cost   float64

	prevEdge        graph.Edge
	prevPathElement *bellmanFordPathElement
}

func (bfpe *bellmanFordPathElement) improve(prevPathElement *bellmanFordPathElement, prevEdge graph.Edge, cost float64) bool {
	if cost >= (bfpe.cost - 1e-7) {
		return false
	}

	bfpe.cost = cost

	bfpe.prevEdge = prevEdge
	bfpe.prevPathElement = prevPathElement

	return true
}

type bellmanFordExecutor struct {
	graph graph.Graph
	start graph.Vertex

	// Vertices whose shortest path costs have been improved during the
	// previous pass.
	prevImprovedVertices   []graph.Vertex
	prevVertexData         map[graph.Vertex]*bellmanFordPathElement
	startVertexEncountered bool

	// Vertices seen so far.
	vertexData map[graph.Vertex]*bellmanFordPathElement
}

func (e *bellmanFordExecutor) calculate() {
	for e.hasNext() {
		e.next()
	}
}

func (e *bellmanFordExecutor) hasNext() bool {
	e.encounterStartVertex()

	return len(e.prevImprovedVertices) > 0
}

func (e *bellmanFordExecutor) next() {
	e.encounterStartVertex()

	var improvedVertices []graph.Vertex
	for i := len(e.prevImprovedVertices) - 1; i >= 0; i-- {
		vertex := e.prevImprovedVertices[i]

		e.forEachEdgeOf(vertex, func(edge graph.Edge) error {
			otherVertex, _ := graph.OppositeVertexOf(e.graph, edge, vertex)

			if _, found := e.vertexData[otherVertex]; found {
				if e.relaxVertexAgain(otherVertex, edge) {
					improvedVertices = append(improvedVertices, otherVertex)
				}
			} else {
				e.relaxVertex(otherVertex, edge)
				improvedVertices = append(improvedVertices, otherVertex)
			}

			return nil
		})
	}

	e.savePassData(improvedVertices)
}

func (e *bellmanFordExecutor) calculatePathCost(vertex graph.Vertex, edge graph.Edge) float64 {
	other, _ := graph.OppositeVertexOf(e.graph, edge, vertex)
	prev := e.prevVertexData[other]

	cost, _ := e.graph.WeightOf(edge)
	if vertex != e.start {
		cost += prev.cost
	}

	return cost
}

func (e *bellmanFordExecutor) forEachEdgeOf(vertex graph.Vertex, fn graph.EdgeSetIterationFunc) {
	var edges graph.EdgeSet
	if dg, ok := e.graph.(graph.DirectedGraph); ok {
		edges, _ = dg.OutgoingEdgesOf(vertex)
	} else {
		edges, _ = e.graph.EdgesOf(vertex)
	}

	edges.ForEach(fn)
}

func (e *bellmanFordExecutor) createSeenData(vertex graph.Vertex, edge graph.Edge, cost float64) *bellmanFordPathElement {
	other, _ := graph.OppositeVertexOf(e.graph, edge, vertex)
	prev := e.prevVertexData[other]

	return &bellmanFordPathElement{
		vertex: vertex,
		cost:   cost,

		prevEdge:        edge,
		prevPathElement: prev,
	}
}

func (e *bellmanFordExecutor) encounterStartVertex() {
	if e.startVertexEncountered {
		return
	}

	el := &bellmanFordPathElement{vertex: e.start}

	e.prevImprovedVertices = append(e.prevImprovedVertices, e.start)
	e.vertexData[e.start] = el
	e.prevVertexData[e.start] = el

	e.startVertexEncountered = true
}

func (e *bellmanFordExecutor) relaxVertex(vertex graph.Vertex, edge graph.Edge) {
	cost := e.calculatePathCost(vertex, edge)

	e.vertexData[vertex] = e.createSeenData(vertex, edge, cost)
}

func (e *bellmanFordExecutor) relaxVertexAgain(vertex graph.Vertex, edge graph.Edge) bool {
	cost := e.calculatePathCost(vertex, edge)

	other, _ := graph.OppositeVertexOf(e.graph, edge, vertex)
	el := e.prevVertexData[other]
	return e.vertexData[vertex].improve(el, edge, cost)
}

func (e *bellmanFordExecutor) savePassData(improvedVertices []graph.Vertex) {
	for _, vertex := range improvedVertices {
		clone := &bellmanFordPathElement{}
		*clone = *e.vertexData[vertex]

		e.prevVertexData[vertex] = clone
	}

	e.prevImprovedVertices = improvedVertices
}

type BellmanFordShortestPaths struct {
	features graph.GraphFeature
	executor *bellmanFordExecutor
}

func (bfsp *BellmanFordShortestPaths) Features() graph.GraphFeature {
	return bfsp.features
}

func (bfsp *BellmanFordShortestPaths) EdgesTo(end graph.Vertex) ([]graph.Edge, error) {
	if !bfsp.executor.graph.ContainsVertex(end) {
		return nil, &graph.VertexNotFoundError{Vertex: end}
	}

	bfsp.executor.calculate()

	el, found := bfsp.executor.vertexData[end]
	if !found {
		return nil, &graph.NotConnectedError{Source: bfsp.executor.start, Target: end}
	}

	var edges []graph.Edge
	for el.prevEdge != nil {
		edges = append(edges, el.prevEdge)
		el = el.prevPathElement
	}

	// Reverse the list.
	for l, r := 0, len(edges)-1; l < r; l, r = l+1, r-1 {
		edges[l], edges[r] = edges[r], edges[l]
	}

	return edges, nil
}

func (bfsp *BellmanFordShortestPaths) CostTo(end graph.Vertex) (float64, error) {
	if !bfsp.executor.graph.ContainsVertex(end) {
		return math.Inf(1), &graph.VertexNotFoundError{Vertex: end}
	}

	bfsp.executor.calculate()

	el, found := bfsp.executor.vertexData[end]
	if !found {
		return math.Inf(1), nil
	}

	return el.cost, nil
}

func BellmanFordShortestPathsOf(g graph.Graph, start graph.Vertex) *BellmanFordShortestPaths {
	return &BellmanFordShortestPaths{
		features: g.Features() & BellmanFordSupportedFeatures,
		executor: &bellmanFordExecutor{
			graph: g,
			start: start,

			prevVertexData: make(map[graph.Vertex]*bellmanFordPathElement),
			vertexData:     make(map[graph.Vertex]*bellmanFordPathElement),
		},
	}
}
