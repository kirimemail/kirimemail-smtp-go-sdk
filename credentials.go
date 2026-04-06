package smtpsdk

import (
	"fmt"
	"strconv"
)

type CredentialsService struct {
	client *Client
}

func (c *Client) Credentials() *CredentialsService {
	return &CredentialsService{client: c}
}

func (s *CredentialsService) List(domain string, page, limit *int) (*CredentialListResponse, error) {
	query := make(map[string]string)
	if page != nil {
		query["page"] = strconv.Itoa(*page)
	}
	if limit != nil {
		query["limit"] = strconv.Itoa(*limit)
	}

	resp, err := s.client.doGet(fmt.Sprintf("/api/domains/%s/credentials", domain), query)
	if err != nil {
		return nil, err
	}

	var result CredentialListResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CredentialsService) Create(domain string, req CredentialCreateRequest) (*CredentialCreateResponse, error) {
	resp, err := s.client.doPost(fmt.Sprintf("/api/domains/%s/credentials", domain), req)
	if err != nil {
		return nil, err
	}

	var result CredentialCreateResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CredentialsService) Get(domain, credentialGUID string) (*Credential, error) {
	resp, err := s.client.doGet(fmt.Sprintf("/api/domains/%s/credentials/%s", domain, credentialGUID), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		APIResponse
		Data Credential `json:"data"`
	}
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func (s *CredentialsService) Delete(domain, credentialGUID string) error {
	resp, err := s.client.doDelete(fmt.Sprintf("/api/domains/%s/credentials/%s", domain, credentialGUID))
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	return nil
}

func (s *CredentialsService) ResetPassword(domain, credentialGUID string) (*CredentialResetPasswordResponse, error) {
	resp, err := s.client.doPut(fmt.Sprintf("/api/domains/%s/credentials/%s/reset-password", domain, credentialGUID), nil)
	if err != nil {
		return nil, err
	}

	var result CredentialResetPasswordResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
