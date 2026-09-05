package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/riolivre/nautical_logistics/internal/application/ports/output"
	"github.com/riolivre/nautical_logistics/internal/logger"
	"go.uber.org/zap"
)

type AuthTokenProvider interface {
	GetToken(ctx context.Context) (string, string, error)
}

type organizationServiceImpl struct {
	client        *http.Client
	tokenProvider AuthTokenProvider
}

type TokenManager struct {
	Token        string
	ExpiresAt    time.Time
	ClientID     string
	ClientSecret string
	AuthURL      string
	mu           sync.Mutex
}

func (tm *TokenManager) GetToken(ctx context.Context) (string, string, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if time.Now().After(tm.ExpiresAt.Add(-5 * time.Minute)) { // Renova 5min antes de expirar
		reqBody := map[string]string{
			"client_id":     tm.ClientID,
			"client_secret": tm.ClientSecret,
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequestWithContext(ctx, "POST", tm.AuthURL+"/auth/service-token", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		logger.Error("failed: ", zap.Error(err))
		if err != nil {
			return "", "", err
		}
		defer resp.Body.Close()
		var respData struct {
			AccessToken string `json:"access_token"`
			ExpiresIn   int    `json:"expires_in"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			return "", "", err
		}
		tm.Token = respData.AccessToken
		tm.ExpiresAt = time.Now().Add(time.Duration(respData.ExpiresIn) * time.Second)
	}
	return tm.Token, tm.AuthURL, nil
}

// NewOrganizationService creates a new instance of OrganizationService.
func NewOrganizationService(tokenProvider AuthTokenProvider) output.OrganizationService {
	return &organizationServiceImpl{
		client:        &http.Client{Timeout: 10 * time.Second},
		tokenProvider: tokenProvider,
	}
}

// GetOrganization fetches organization details from the external service.
func (o *organizationServiceImpl) GetOrganization(ctx context.Context, id string) (*output.OrganizationDTO, error) {
	token, baseURL, err := o.tokenProvider.GetToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth token: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"/organizations/"+id, nil)
	if err != nil {
		return nil, err
	}
	logger.Error("failed: ", zap.Error(err))
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := o.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var org output.OrganizationDTO
	if err := json.NewDecoder(resp.Body).Decode(&org); err != nil {
		return nil, err
	}

	return &org, nil
}

