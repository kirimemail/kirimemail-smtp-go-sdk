package smtpsdk

import (
	"fmt"
	"strconv"
)

type DomainsService struct {
	client *Client
}

func (c *Client) Domains() *DomainsService {
	return &DomainsService{client: c}
}

func (s *DomainsService) List(page, limit *int, search *string) (*DomainListResponse, error) {
	query := make(map[string]string)
	if page != nil {
		query["page"] = strconv.Itoa(*page)
	}
	if limit != nil {
		query["limit"] = strconv.Itoa(*limit)
	}
	if search != nil {
		query["search"] = *search
	}

	resp, err := s.client.doGet("/api/domains", query)
	if err != nil {
		return nil, err
	}

	var result DomainListResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *DomainsService) Create(req DomainCreateRequest) (*DomainCreateResponse, error) {
	resp, err := s.client.doPost("/api/domains", req)
	if err != nil {
		return nil, err
	}

	var result DomainCreateResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *DomainsService) Get(domain string) (*Domain, error) {
	resp, err := s.client.doGet(fmt.Sprintf("/api/domains/%s", domain), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		APIResponse
		Data Domain `json:"data"`
	}
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func (s *DomainsService) Update(domain string, req DomainUpdateRequest) (*DomainUpdateResponse, error) {
	resp, err := s.client.doPut(fmt.Sprintf("/api/domains/%s", domain), req)
	if err != nil {
		return nil, err
	}

	var result DomainUpdateResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *DomainsService) Delete(domain string) (*APIResponse, error) {
	resp, err := s.client.doDelete(fmt.Sprintf("/api/domains/%s", domain))
	if err != nil {
		return nil, err
	}

	var result APIResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *DomainsService) SetupAuthDomain(domain string, req AuthDomainSetupRequest) (*AuthDomainSetupResponse, error) {
	resp, err := s.client.doPost(fmt.Sprintf("/api/domains/%s/setup-auth-domain", domain), req)
	if err != nil {
		return nil, err
	}

	var result AuthDomainSetupResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *DomainsService) VerifyMandatoryRecords(domain string) (*DNSVerificationResponse, error) {
	resp, err := s.client.doPost(fmt.Sprintf("/api/domains/%s/verify-mandatory", domain), nil)
	if err != nil {
		return nil, err
	}

	var result DNSVerificationResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *DomainsService) VerifyAuthDomain(domain string) (*AuthDomainVerificationResponse, error) {
	resp, err := s.client.doPost(fmt.Sprintf("/api/domains/%s/verify-auth-domain", domain), nil)
	if err != nil {
		return nil, err
	}

	var result AuthDomainVerificationResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *DomainsService) SetupTracklink(domain string, req TracklinkSetupRequest) (*TracklinkSetupResponse, error) {
	resp, err := s.client.doPost(fmt.Sprintf("/api/domains/%s/setup-tracklink", domain), req)
	if err != nil {
		return nil, err
	}

	var result TracklinkSetupResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *DomainsService) VerifyTracklink(domain string) (*TracklinkVerificationResponse, error) {
	resp, err := s.client.doPost(fmt.Sprintf("/api/domains/%s/verify-tracklink", domain), nil)
	if err != nil {
		return nil, err
	}

	var result TracklinkVerificationResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
