package smtpsdk

import (
	"fmt"
)

type WebhooksService struct {
	client *Client
}

func (c *Client) Webhooks() *WebhooksService {
	return &WebhooksService{client: c}
}

func (s *WebhooksService) List(domain string, webhookType *string) (*WebhookListResponse, error) {
	query := make(map[string]string)
	if webhookType != nil {
		query["type"] = *webhookType
	}

	resp, err := s.client.doGet(fmt.Sprintf("/api/domains/%s/webhooks", domain), query)
	if err != nil {
		return nil, err
	}

	var result WebhookListResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *WebhooksService) Create(domain string, req WebhookCreateRequest) (*WebhookCreateResponse, error) {
	resp, err := s.client.doPost(fmt.Sprintf("/api/domains/%s/webhooks", domain), req)
	if err != nil {
		return nil, err
	}

	var result WebhookCreateResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *WebhooksService) Get(domain, webhookGUID string) (*Webhook, error) {
	resp, err := s.client.doGet(fmt.Sprintf("/api/domains/%s/webhooks/%s", domain, webhookGUID), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		APIResponse
		Data Webhook `json:"data"`
	}
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func (s *WebhooksService) Update(domain, webhookGUID string, req WebhookUpdateRequest) (*WebhookUpdateResponse, error) {
	resp, err := s.client.doPut(fmt.Sprintf("/api/domains/%s/webhooks/%s", domain, webhookGUID), req)
	if err != nil {
		return nil, err
	}

	var result WebhookUpdateResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *WebhooksService) Delete(domain, webhookGUID string) (*APIResponse, error) {
	resp, err := s.client.doDelete(fmt.Sprintf("/api/domains/%s/webhooks/%s", domain, webhookGUID))
	if err != nil {
		return nil, err
	}

	var result APIResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *WebhooksService) Test(domain string, req WebhookTestRequest) (*WebhookTestResponse, error) {
	resp, err := s.client.doPost(fmt.Sprintf("/api/domains/%s/webhooks/test", domain), req)
	if err != nil {
		return nil, err
	}

	var result WebhookTestResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
