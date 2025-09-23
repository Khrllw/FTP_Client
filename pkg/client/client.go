package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"ftp_client/pkg/swagger"
	"net/http"

	"ftp_client/internal/domain/models"
)

// ClientAPI определяет интерфейс для взаимодействия с сервисом
type ClientAPI interface {
	GetFileWithAuth(ctx context.Context, req *models.GetFileWithAuthRequest) (*swagger.GetResponse, *http.Response, error)
	GetFileAnonymous(ctx context.Context, req *models.GetFileAnonymousRequest) (*swagger.GetResponse, *http.Response, error)
	SendFileWithAuth(ctx context.Context, req *models.SendFileRequest) (*swagger.SendResponse, *http.Response, error)
	SendFileAnonymous(ctx context.Context, req *models.SendFileRequest) (*swagger.SendResponse, *http.Response, error)
}

// Client реализует интерфейс ClientAPI
type Client struct {
	service *ClientService
}

// NewClient создает нового клиента
func NewClient(host string) ClientAPI {
	return &Client{
		service: NewClientService(host),
	}
}

// GetFileWithAuth скачивает файл с FTP с авторизацией
func (c *Client) GetFileWithAuth(ctx context.Context, req *models.GetFileWithAuthRequest) (*swagger.GetResponse, *http.Response, error) {
	const endpoint = "/api/v1/get/auth"

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := c.service.createRequestJSONWithContext(ctx, http.MethodPost, endpoint, nil, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, nil, err
	}

	body, httpResp, err := c.service.doRequest(httpReq)
	if err != nil {
		return nil, httpResp, err
	}

	var resp swagger.GetResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, httpResp, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, httpResp, nil
}

// GetFileAnonymous скачивает файл с FTP анонимно
func (c *Client) GetFileAnonymous(ctx context.Context, req *models.GetFileAnonymousRequest) (*swagger.GetResponse, *http.Response, error) {
	const endpoint = "/api/v1/get/anon"

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := c.service.createRequestJSONWithContext(ctx, http.MethodPost, endpoint, nil, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, nil, err
	}

	body, httpResp, err := c.service.doRequest(httpReq)
	if err != nil {
		return nil, httpResp, err
	}

	var resp swagger.GetResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, httpResp, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, httpResp, nil
}

// SendFileWithAuth загружает файл на FTP с авторизацией
func (c *Client) SendFileWithAuth(ctx context.Context, req *models.SendFileRequest) (*swagger.SendResponse, *http.Response, error) {
	const endpoint = "/api/v1/send/auth"

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := c.service.createRequestJSONWithContext(ctx, http.MethodPost, endpoint, nil, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, nil, err
	}

	body, httpResp, err := c.service.doRequest(httpReq)
	if err != nil {
		return nil, httpResp, err
	}

	var resp swagger.SendResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, httpResp, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, httpResp, nil
}

// SendFileAnonymous загружает файл на FTP анонимно
func (c *Client) SendFileAnonymous(ctx context.Context, req *models.SendFileRequest) (*swagger.SendResponse, *http.Response, error) {
	const endpoint = "/api/v1/send/anon"

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := c.service.createRequestJSONWithContext(ctx, http.MethodPost, endpoint, nil, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, nil, err
	}

	body, httpResp, err := c.service.doRequest(httpReq)
	if err != nil {
		return nil, httpResp, err
	}

	var resp swagger.SendResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, httpResp, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, httpResp, nil
}
