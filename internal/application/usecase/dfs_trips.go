package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/rafaellima1412/uber-dos-rios/internal/application/ports/input"
)

type DFSTripsUseCase struct{}

func NewDFSTripsUseCase() input.DFSTripsUseCase {
	return &DFSTripsUseCase{}
}
func canBoard(
	search input.TripSearchInput,
	last input.Path,
	next *input.Edge,
	connection time.Duration,
) bool {

	// edge sem data válida
	if next.DateDepartureTime.IsZero() || next.DateArrivalTime.IsZero() {
		return false
	}

	dep := next.DateDepartureTime
	arr := next.DateArrivalTime

	// chegou depois do limite da busca
	if !search.DateArrivalTime.IsZero() && arr.After(search.DateArrivalTime) {
		return false
	}

	// PRIMEIRO EMBARQUE
	if last.TripConfigurationID == uuid.Nil {

		if dep.Before(search.DateDepartureTime) {
			return false
		}

		return true
	}

	// CONEXÃO
	if last.DateArrivalTime.IsZero() {
		return false
	}

	minNextDeparture := last.DateArrivalTime.Add(connection)

	if dep.Before(minNextDeparture) {
		return false
	}

	return true
}

// input.DfsTripsUseCase
func (uc *DFSTripsUseCase) Execute(
	graph input.Graph,
	search input.TripSearchInput,
) ([]input.TripSearchOutput, error) {

	results := []input.TripSearchOutput{}
	visited := make(map[uuid.UUID]bool)

	start := input.Path{
		CityID:            search.CityOriginID,
		DateDepartureTime: search.DateDepartureTime,
		DateArrivalTime:   search.DateDepartureTime,
		Cost:              0,
	}

	limits := input.TripSearchInput{
		MaxStops:   search.MaxStops,
		MaxCost:    search.MaxCost,
		MaxResults: search.MaxResults,
	}

	uc.dfs(
		graph,
		search.CityOriginID,
		search.CityDestID,
		visited,
		[]input.Path{start},
		0,
		limits,
		0,
		&results,
	)

	return results, nil
}

func (uc *DFSTripsUseCase) dfs(
	graph input.Graph,
	current int,
	destination int,
	visited map[uuid.UUID]bool,
	path []input.Path,
	totalCost int,
	search input.TripSearchInput,
	stops int,
	results *[]input.TripSearchOutput,
) {
	// regra de negócio cliente quer parar
	if search.MaxStops > 0 && stops > search.MaxStops {
		return
	}
	// regra de negócio cliente pode gastar
	if search.MaxCost > 0 && totalCost > search.MaxCost {
		return
	}
	// limite de quantidade de resultados
	if search.MaxResults > 0 && len(*results) >= search.MaxResults {
		return
	}
	// CHEGOU NO DESTINO
	if current == destination {
		finalPath := make([]input.Path, len(path))
		copy(finalPath, path)

		*results = append(*results, input.TripSearchOutput{
			TotalCost: totalCost,
			Path:      finalPath,
		})

		return
	}
	
	last := path[len(path)-1]

	for _, edge := range graph[current] {

		if visited[edge.TripConfigurationID] {
			continue
		}
		last = path[len(path)-1]
		if !canBoard(search, last, edge, 30*time.Minute) {
			continue
		}

		step := input.Path{
			OrganizationID:      edge.OrganizationID,
			CityID:              edge.To,
			Cost:                edge.Cost,
			RouteID:             edge.RouteID,
			RouteName:           edge.RouteName,
			DateDepartureTime:   edge.DateDepartureTime,
			DateArrivalTime:     edge.DateArrivalTime,
			ShipID:              edge.ShipID,
			ShipName:            edge.ShipName,
			ShipURL:             edge.ShipImageURL,
			TripConfigurationID: edge.TripConfigurationID,
		}

		// PUSH (sem alocar slice novo) melhoria
		visited[edge.TripConfigurationID] = true
		path = append(path, step)

		uc.dfs(
			graph,
			edge.To,
			destination,
			visited,
			path,
			totalCost+edge.Cost,
			search,
			stops+1,
			results,
		)

		path = path[:len(path)-1]
		delete(visited, edge.TripConfigurationID)
	}
}
