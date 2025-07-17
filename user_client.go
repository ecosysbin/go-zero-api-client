package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Request and Response types
type GetUserRequest struct {
	Authorization string `header:"authorization"`
	Name          string `path:"name"`
	Delete        bool   `form:"delete,optional"`
}

type GetUserResponse struct {
	CreateTime string `json:"create_time"`
	Name       string `json:"name"`
	Age        string `json:"age"`
}

type AddUserRequest struct {
	Authorization string `header:"authorization"`
	Name          string `json:"name"`
	Age           string `json:"age"`
}

type AddUserResponse struct {
	Message string `json:"message"`
}

type DeleteUserRequest struct {
	Authorization string `header:"authorization"`
	Name          string `path:"name"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

type UpdateUserRequest struct {
	Authorization string `header:"authorization"`
	Name          string `path:"name"`
	Age           string `json:"age"`
}

type UpdateUserResponse struct {
	Message string `json:"message"`
}

// UserAPIClient represents the HTTP client for user API
type UserAPIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewUserAPIClient creates a new instance of UserAPIClient
func NewUserAPIClient(baseURL string) *UserAPIClient {
	return &UserAPIClient{
		BaseURL:    strings.TrimSuffix(baseURL, "/"),
		HTTPClient: &http.Client{},
	}
}

// AddUser creates a new user
func (c *UserAPIClient) AddUser(req AddUserRequest) (*AddUserResponse, error) {
	// Prepare request body
	body := struct {
		Name string `json:"name"`
		Age  string `json:"age"`
	}{
		Name: req.Name,
		Age:  req.Age,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequest("POST", c.BaseURL+"/v1/user", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", req.Authorization)

	// Execute request
	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse response
	var result AddUserResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, nil
}

// DeleteUser deletes a user by name
func (c *UserAPIClient) DeleteUser(req DeleteUserRequest) (*DeleteUserResponse, error) {
	// Create HTTP request
	url := fmt.Sprintf("%s/v1/user/%s", c.BaseURL, req.Name)
	httpReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Authorization", req.Authorization)

	// Execute request
	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse response
	var result DeleteUserResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, nil
}

// UpdateUser updates a user by name
func (c *UserAPIClient) UpdateUser(req UpdateUserRequest) (*UpdateUserResponse, error) {
	// Prepare request body
	body := struct {
		Age string `json:"age"`
	}{
		Age: req.Age,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	url := fmt.Sprintf("%s/v1/user/%s", c.BaseURL, req.Name)
	httpReq, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", req.Authorization)

	// Execute request
	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse response
	var result UpdateUserResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, nil
}

// GetUser retrieves a user by name
func (c *UserAPIClient) GetUser(req GetUserRequest) (*GetUserResponse, error) {
	// Create HTTP request
	baseURL := fmt.Sprintf("%s/v1/user/%s", c.BaseURL, req.Name)
	
	// Add query parameters if needed
	if req.Delete {
		u, err := url.Parse(baseURL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse URL: %w", err)
		}
		q := u.Query()
		q.Set("delete", "true")
		u.RawQuery = q.Encode()
		baseURL = u.String()
	}

	httpReq, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Authorization", req.Authorization)

	// Execute request
	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse response
	var result GetUserResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, nil
}