package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
)

type userServiceImpl struct {
	client        *http.Client
	tokenProvider AuthTokenProvider
}

func NewGetUsers(tokenProvider AuthTokenProvider) output.UserService {
	return &userServiceImpl{
		client:        &http.Client{Timeout: 10 * time.Second},
		tokenProvider: tokenProvider,
	}
}

// GetUsers fetches organization details from the external service.
func (o *userServiceImpl) GetUsers(ctx context.Context, id string) (*output.UsersDTO, error) {
	token, baseURL, err := o.tokenProvider.GetToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth token: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"/users/"+id, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := o.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
	var users output.UsersDTO
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return &users, nil

}
