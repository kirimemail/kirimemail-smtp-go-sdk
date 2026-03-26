package tests

import (
	"testing"

	smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"
)

func TestWebhook(t *testing.T) {
	webhook := smtpsdk.Webhook{
		WebhookGUID:    "webhook123",
		UserGUID:       "user123",
		UserDomainGUID: "domain123",
		UserSMTPGUID:   "smtp123",
		Type:           "delivered",
		URL:            "https://example.com/webhook",
		IsDeleted:      false,
		CreatedAt:      1234567890,
		ModifiedAt:     1234567890,
	}

	if webhook.WebhookGUID != "webhook123" {
		t.Errorf("WebhookGUID = %v, want webhook123", webhook.WebhookGUID)
	}
	if webhook.Type != "delivered" {
		t.Errorf("Type = %v, want delivered", webhook.Type)
	}
	if webhook.URL != "https://example.com/webhook" {
		t.Errorf("URL = %v, want https://example.com/webhook", webhook.URL)
	}
}

func TestWebhookEventTypes(t *testing.T) {
	validTypes := []string{
		"queued", "send", "delivered", "bounced", "failed",
		"permanent_fail", "opened", "clicked", "unsubscribed", "temporary_fail", "deferred",
	}

	for _, eventType := range validTypes {
		webhook := smtpsdk.Webhook{
			WebhookGUID: "webhook1",
			Type:        eventType,
			URL:         "https://example.com/webhook",
		}

		if webhook.Type != eventType {
			t.Errorf("EventType not set correctly: %v", eventType)
		}
	}
}

func TestWebhookCreateRequest(t *testing.T) {
	req := smtpsdk.WebhookCreateRequest{
		Type: "delivered",
		URL:  "https://example.com/webhook",
	}

	if req.Type != "delivered" {
		t.Errorf("Type = %v, want delivered", req.Type)
	}
	if req.URL != "https://example.com/webhook" {
		t.Errorf("URL not set correctly")
	}
}

func TestWebhookCreateResponse(t *testing.T) {
	resp := smtpsdk.WebhookCreateResponse{
		APIResponse: smtpsdk.APIResponse{
			Success: true,
			Message: "Webhook created",
		},
		Message: "Webhook created",
		Data: smtpsdk.Webhook{
			WebhookGUID: "webhook123",
			Type:        "delivered",
			URL:         "https://example.com/webhook",
		},
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if resp.Data.WebhookGUID != "webhook123" {
		t.Errorf("WebhookGUID not set correctly")
	}
}

func TestWebhookUpdateRequest(t *testing.T) {
	newURL := "https://example.com/new-webhook"
	newType := "opened"

	req := smtpsdk.WebhookUpdateRequest{
		URL:  &newURL,
		Type: &newType,
	}

	if req.URL == nil || *req.URL != newURL {
		t.Errorf("URL pointer not set correctly")
	}
	if req.Type == nil || *req.Type != newType {
		t.Errorf("Type pointer not set correctly")
	}
}

func TestWebhookUpdateRequest_Partial(t *testing.T) {
	newURL := "https://example.com/updated"

	req := smtpsdk.WebhookUpdateRequest{
		URL: &newURL,
	}

	if req.URL == nil || *req.URL != newURL {
		t.Errorf("URL pointer not set correctly")
	}
	if req.Type != nil {
		t.Errorf("Type should be nil when not provided")
	}
}

func TestWebhookUpdateResponse(t *testing.T) {
	resp := smtpsdk.WebhookUpdateResponse{
		APIResponse: smtpsdk.APIResponse{
			Success: true,
			Message: "Webhook updated",
		},
		Message: "Webhook updated",
		Data: smtpsdk.Webhook{
			WebhookGUID: "webhook123",
			Type:        "opened",
			URL:         "https://example.com/updated",
		},
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if resp.Data.Type != "opened" {
		t.Errorf("Type = %v, want opened", resp.Data.Type)
	}
}

func TestWebhookListResponse(t *testing.T) {
	resp := smtpsdk.WebhookListResponse{
		APIResponse: smtpsdk.APIResponse{
			Success: true,
		},
		Data: []smtpsdk.Webhook{
			{
				WebhookGUID: "webhook1",
				Type:        "delivered",
				URL:         "https://example.com/webhook",
			},
		},
		Count: 1,
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 webhook")
	}
	if resp.Count != 1 {
		t.Errorf("Count = %v, want 1", resp.Count)
	}
}

func TestWebhookTestRequest(t *testing.T) {
	req := smtpsdk.WebhookTestRequest{
		URL:       "https://example.com/webhook",
		EventType: "delivered",
	}

	if req.URL != "https://example.com/webhook" {
		t.Errorf("URL not set correctly")
	}
	if req.EventType != "delivered" {
		t.Errorf("EventType = %v, want delivered", req.EventType)
	}
}

func TestWebhookTestResponse(t *testing.T) {
	resp := smtpsdk.WebhookTestResponse{
		APIResponse: smtpsdk.APIResponse{
			Success: true,
			Message: "Test successful",
		},
		Message: "Test successful",
		Data: struct {
			URL            string `json:"url"`
			EventType      string `json:"event_type"`
			ResponseStatus int    `json:"response_status"`
			ResponseTime   int    `json:"response_time"`
		}{
			URL:            "https://example.com/webhook",
			EventType:      "delivered",
			ResponseStatus: 200,
			ResponseTime:   123,
		},
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if resp.Data.ResponseStatus != 200 {
		t.Errorf("ResponseStatus = %v, want 200", resp.Data.ResponseStatus)
	}
	if resp.Data.ResponseTime != 123 {
		t.Errorf("ResponseTime = %v, want 123", resp.Data.ResponseTime)
	}
}
