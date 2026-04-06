package smtpsdk

import (
	"fmt"
	"strconv"
)

type SuppressionsService struct {
	client *Client
}

func (c *Client) Suppressions() *SuppressionsService {
	return &SuppressionsService{client: c}
}

func (s *SuppressionsService) List(domain string, supressionType *string, search *string, page, perPage *int) (*SuppressionListResponse, error) {
	query := make(map[string]string)

	if supressionType != nil {
		query["type"] = *supressionType
	}
	if search != nil {
		query["search"] = *search
	}
	if page != nil {
		query["page"] = strconv.Itoa(*page)
	}
	if perPage != nil {
		query["per_page"] = strconv.Itoa(*perPage)
	}

	resp, err := s.client.doGet(fmt.Sprintf("/api/domains/%s/suppressions", domain), query)
	if err != nil {
		return nil, err
	}

	var result SuppressionListResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SuppressionsService) ListUnsubscribes(domain string, search *string, page, perPage *int) (*SuppressionListResponse, error) {
	query := make(map[string]string)

	if search != nil {
		query["search"] = *search
	}
	if page != nil {
		query["page"] = strconv.Itoa(*page)
	}
	if perPage != nil {
		query["per_page"] = strconv.Itoa(*perPage)
	}

	resp, err := s.client.doGet(fmt.Sprintf("/api/domains/%s/suppressions/unsubscribes", domain), query)
	if err != nil {
		return nil, err
	}

	var result SuppressionListResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SuppressionsService) ListBounces(domain string, search *string, page, perPage *int) (*SuppressionListResponse, error) {
	query := make(map[string]string)

	if search != nil {
		query["search"] = *search
	}
	if page != nil {
		query["page"] = strconv.Itoa(*page)
	}
	if perPage != nil {
		query["per_page"] = strconv.Itoa(*perPage)
	}

	resp, err := s.client.doGet(fmt.Sprintf("/api/domains/%s/suppressions/bounces", domain), query)
	if err != nil {
		return nil, err
	}

	var result SuppressionListResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SuppressionsService) ListWhitelists(domain string, search *string, page, perPage *int) (*SuppressionListResponse, error) {
	query := make(map[string]string)

	if search != nil {
		query["search"] = *search
	}
	if page != nil {
		query["page"] = strconv.Itoa(*page)
	}
	if perPage != nil {
		query["per_page"] = strconv.Itoa(*perPage)
	}

	resp, err := s.client.doGet(fmt.Sprintf("/api/domains/%s/suppressions/whitelist", domain), query)
	if err != nil {
		return nil, err
	}

	var result SuppressionListResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SuppressionsService) CreateWhitelist(domain string, req WhitelistCreateRequest) (*WhitelistCreateResponse, error) {
	resp, err := s.client.doPost(fmt.Sprintf("/api/domains/%s/suppressions/whitelist", domain), req)
	if err != nil {
		return nil, err
	}

	var result WhitelistCreateResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SuppressionsService) DeleteUnsubscribes(domain string, ids []int) error {
	req := SuppressionDeleteRequest{IDs: ids}

	resp, err := s.client.doDeleteWithBody(fmt.Sprintf("/api/domains/%s/suppressions/unsubscribes", domain), req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	return nil
}

func (s *SuppressionsService) DeleteBounces(domain string, ids []int) error {
	req := SuppressionDeleteRequest{IDs: ids}

	resp, err := s.client.doDeleteWithBody(fmt.Sprintf("/api/domains/%s/suppressions/bounces", domain), req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	return nil
}

func (s *SuppressionsService) DeleteWhitelists(domain string, ids []int) error {
	req := SuppressionDeleteRequest{IDs: ids}

	resp, err := s.client.doDeleteWithBody(fmt.Sprintf("/api/domains/%s/suppressions/whitelist", domain), req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	return nil
}
