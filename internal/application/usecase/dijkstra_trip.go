package usecase

import (
	"container/heap"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
)

type DijkstraTripUseCase struct{}

func NewDijkstraTripUseCase() *DijkstraTripUseCase {
	return &DijkstraTripUseCase{}
}

var _ input.DijkstraTripInput = (*DijkstraTripUseCase)(nil)

type Node struct {
	City     int
	Priority int
	Index    int
}

type PrevNode struct {
	FromCity          int
	Cost              int
	OrganizationID    uuid.UUID
	RouteID           uuid.UUID
	RouteName         string
	DateDepartureTime time.Time
	DateArrivalTime   time.Time
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	node := x.(*Node)
	node.Index = n
	*pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func (uc *DijkstraTripUseCase) Execute(graph input.Graph, origin, destination int,
) (input.TripSearchResponse, bool) {

	const INF = int(^uint(0) >> 1)

	dist := make(map[int]int)
	prev := make(map[int]PrevNode)

	for from, edges := range graph {
		if _, ok := dist[from]; !ok {
			dist[from] = INF
		}
		for _, edge := range edges {
			if _, ok := dist[edge.To]; !ok {
				dist[edge.To] = INF
			}
		}
	}
	dist[origin] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Node{City: origin, Priority: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Node)
		if curr.Priority > dist[curr.City] {
			continue
		}
		if curr.City == destination {
			break
		}

		for _, edge := range graph[curr.City] {
			newCost := dist[curr.City] + edge.Cost

			if newCost < dist[edge.To] {
				dist[edge.To] = newCost
				prev[edge.To] = PrevNode{
					FromCity:          curr.City,
					Cost:              edge.Cost,
					OrganizationID:    edge.OrganizationID,
					RouteID:           edge.RouteID,
					RouteName:         edge.RouteName,
					DateDepartureTime: edge.DateDepartureTime,
					DateArrivalTime:   edge.DateArrivalTime,
				}
				heap.Push(pq, &Node{
					City:     edge.To,
					Priority: newCost})
			}
		}
	}

	if dist[destination] == INF {
		return input.TripSearchResponse{}, false
	}

	// reconstrói caminho
	path := []int{}
	for at := destination; at != origin; {
		path = append(path, at)
		prevNode, ok := prev[at]
		if !ok {
			return input.TripSearchResponse{}, false
		}
		at = prevNode.FromCity
	}
	path = append(path, origin)
	slices.Reverse(path)

	steps := make([]input.PathStep, 0, len(path))
	for i := 1; i < len(path); i++ {
		step := input.PathStep{
			CityID: path[i],
			Cost:   0,
		}

		if i > 0 {
			p := prev[path[i]]
			step.OrganizationID = p.OrganizationID
			step.Cost = p.Cost
			step.RouteID = p.RouteID
			step.RouteName = p.RouteName
			step.DateDepartureTime = p.DateDepartureTime
			step.DateArrivalTime = p.DateArrivalTime
		}

		steps = append(steps, step)
	}

	return input.TripSearchResponse{
		TotalCost: dist[destination],
		Path:      steps,
	}, true
}
