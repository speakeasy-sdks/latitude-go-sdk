package latitude

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/latitude-go-sdk/v2/pkg/utils"
	"net/http"
)

type sshKeys struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newSSHKeys(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *sshKeys {
	return &sshKeys{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// DeleteProjectSSHKey - Delete a Project SSH Key
// Allow you remove SSH Keys in a project. Remove a SSH Key from the project won't revoke the SSH Keys access for previously deploy and reinstall actions.
func (s *sshKeys) DeleteProjectSSHKey(ctx context.Context, request operations.DeleteProjectSSHKeyRequest) (*operations.DeleteProjectSSHKeyResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := utils.ConfigureSecurityClient(s.defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteProjectSSHKeyResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
	case httpRes.StatusCode == 404:
	}

	return res, nil
}

// GetProjectSSHKey - Retrieve a Project SSH Key
// List all SSH Keys in the project. These keys can be used to access servers after deploy and reinstall actions.
func (s *sshKeys) GetProjectSSHKey(ctx context.Context, request operations.GetProjectSSHKeyRequest) (*operations.GetProjectSSHKeyResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s.defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetProjectSSHKeyResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectSSHKey200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectSSHKey200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetProjectSSHKeys - List all Project SSH Keys
// List all SSH Keys in the project. These keys can be used to access servers after deploy and reinstall actions.
func (s *sshKeys) GetProjectSSHKeys(ctx context.Context, request operations.GetProjectSSHKeysRequest) (*operations.GetProjectSSHKeysResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/projects/{project_id_or_slug}/ssh_keys", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s.defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetProjectSSHKeysResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SSHKey
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.SSHKey = out
		}
	}

	return res, nil
}

// PostProjectSSHKey - Create a Project SSH Key
// Allow you create SSH Keys in a project. These keys can be used to access servers after deploy and reinstall actions.
func (s *sshKeys) PostProjectSSHKey(ctx context.Context, request operations.PostProjectSSHKeyRequest) (*operations.PostProjectSSHKeyResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/projects/{project_id_or_slug}/ssh_keys", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s.defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostProjectSSHKeyResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PostProjectSSHKey201ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PostProjectSSHKey201ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 400:
	case httpRes.StatusCode == 422:
	}

	return res, nil
}

// PutProjectSSHKey - Update a Project SSH Key
// Allow you update SSH Key in a project. These keys can be used to access servers after deploy and reinstall actions.
func (s *sshKeys) PutProjectSSHKey(ctx context.Context, request operations.PutProjectSSHKeyRequest) (*operations.PutProjectSSHKeyResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s.defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutProjectSSHKeyResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PutProjectSSHKey200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PutProjectSSHKey200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 400:
	case httpRes.StatusCode == 422:
	}

	return res, nil
}
