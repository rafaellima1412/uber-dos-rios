package usecase
import (
	"context"
	"github.com/riolivre/nautical_logistics/internal/application/ports/input"
	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)
type listCitiesUseCase struct {
	terminalRepository output.TerminalRepository
}
func NewListCitiesUseCase(terminalRepository output.TerminalRepository) input.ListCitiesUseCase {
	return &listCitiesUseCase{
		terminalRepository: terminalRepository,
	}
}
func (uc *listCitiesUseCase) Execute(ctx context.Context) ([]input.CityOutput, error) {
	cities, err := uc.terminalRepository.ListCities(ctx)
	if err != nil {
		return nil, err
	}
	var outputs []input.CityOutput
	for _, city := range cities {
		outputs = append(outputs, input.CityOutput{
			ID:    city.ID,
			Name:  city.Name,
			State: city.State,
		})
	}
	return outputs, nil
}
