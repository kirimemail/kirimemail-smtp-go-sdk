package tests

import (
	"encoding/json"
	"testing"

	smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"
)

func TestMessageSendRequest(t *testing.T) {
	req := smtpsdk.MessageSendRequest{
		From:     "noreply@example.com",
		FromName: "Company Name",
		To:       []string{"user@example.com"},
		Subject:  "Test Subject",
		Text:     "Test content",
		HTML:     "<h1>Test</h1>",
		ReplyTo:  "reply@example.com",
	}

	if req.From != "noreply@example.com" {
		t.Errorf("From not set correctly")
	}
	if req.Subject != "Test Subject" {
		t.Errorf("Subject not set correctly")
	}
}

func TestMessageSendRequest_SingleRecipient(t *testing.T) {
	req := smtpsdk.MessageSendRequest{
		From:    "noreply@example.com",
		To:      "user@example.com",
		Subject: "Test",
		Text:    "Content",
	}

	// Just verify the struct can be created with interface{} To field
	_ = req
}

func TestMessageSendRequest_WithHeaders(t *testing.T) {
	headers := smtpsdk.Headers{
		"X-Custom-Header": "custom-value",
		"X-Order-ID":      "12345",
	}

	req := smtpsdk.MessageSendRequest{
		From:    "noreply@example.com",
		Subject: "Test",
		Text:    "Content",
		Headers: headers,
	}

	_ = req
	if headers["X-Custom-Header"] != "custom-value" {
		t.Errorf("Custom header not set correctly")
	}
}

func TestMessageTemplateRequest(t *testing.T) {
	req := smtpsdk.MessageTemplateRequest{
		TemplateGUID: "template-123",
		To:           []string{"user@example.com"},
		From:         "noreply@example.com",
		FromName:     "Company",
		ReplyTo:      "reply@example.com",
	}

	if req.TemplateGUID != "template-123" {
		t.Errorf("TemplateGUID not set correctly")
	}

	if req.FromName != "Company" {
		t.Errorf("FromName not set correctly")
	}
}

func TestMessageTemplateRequest_WithVariables(t *testing.T) {
	variables := map[string]interface{}{
		"name":  "John Doe",
		"email": "john@example.com",
	}

	req := smtpsdk.MessageTemplateRequest{
		TemplateGUID: "template-123",
		To:           []string{"user@example.com"},
		Variables:    variables,
	}

	if vars, ok := req.Variables.(map[string]interface{}); ok {
		if vars["name"] != "John Doe" {
			t.Errorf("Variables not set correctly")
		}
	}
}

func TestMessageTemplateRequest_BulkSend(t *testing.T) {
	req := smtpsdk.MessageTemplateRequest{
		TemplateGUID: "template-123",
		To: []interface{}{
			"user1@example.com",
			"user2@example.com",
		},
		Variables: []map[string]interface{}{
			{"name": "John Doe"},
			{"name": "Jane Smith"},
		},
	}

	if recipients, ok := req.To.([]interface{}); ok {
		if len(recipients) != 2 {
			t.Errorf("Expected 2 recipients")
		}
	}
}

func TestHeadersMarshal(t *testing.T) {
	headers := smtpsdk.Headers{
		"X-Custom-Header":  "value1",
		"X-Another-Header": "value2",
	}

	data, err := json.Marshal(headers)
	if err != nil {
		t.Errorf("Headers.MarshalJSON() error = %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Expected non-empty JSON output")
	}
}

func TestHeadersUnmarshal(t *testing.T) {
	data := []byte(`{"X-Custom-Header":"value","X-Another":"another-value"}`)

	var headers smtpsdk.Headers
	err := json.Unmarshal(data, &headers)
	if err != nil {
		t.Errorf("Headers.UnmarshalJSON() error = %v", err)
	}

	if headers["X-Custom-Header"] != "value" {
		t.Errorf("X-Custom-Header = %v, want value", headers["X-Custom-Header"])
	}
	if headers["X-Another"] != "another-value" {
		t.Errorf("X-Another = %v, want another-value", headers["X-Another"])
	}
}

func TestMessageSendResponse_WithMessages(t *testing.T) {
	data := []byte(`{
		"success": true,
		"message": "Email send jobs dispatched for 2 recipient(s)",
		"messages": [
			{"to": "user1@example.com", "message_guid": "0193f3c4-aaaa-bbbb-cccc-000000000001"},
			{"to": "user2@example.com", "message_guid": "0193f3c4-aaaa-bbbb-cccc-000000000002"}
		]
	}`)

	var resp smtpsdk.MessageSendResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("MessageSendResponse unmarshal error = %v", err)
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if len(resp.Messages) != 2 {
		t.Fatalf("len(Messages) = %d, want 2", len(resp.Messages))
	}
	if resp.Messages[0].To != "user1@example.com" {
		t.Errorf("Messages[0].To = %v, want user1@example.com", resp.Messages[0].To)
	}
	if resp.Messages[0].MessageGUID != "0193f3c4-aaaa-bbbb-cccc-000000000001" {
		t.Errorf("Messages[0].MessageGUID = %v, want 0193f3c4-aaaa-bbbb-cccc-000000000001", resp.Messages[0].MessageGUID)
	}
}

func TestMessageSendResponse_WithoutMessages(t *testing.T) {
	data := []byte(`{"success": true, "message": "Email sent"}`)

	var resp smtpsdk.MessageSendResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("MessageSendResponse unmarshal error = %v", err)
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if resp.Messages != nil {
		t.Errorf("Messages = %v, want nil for backward compatibility", resp.Messages)
	}
}

func TestMessageTemplateResponse_WithMessages(t *testing.T) {
	data := []byte(`{
		"success": true,
		"message": "Template-based email jobs dispatched for 1 recipient(s)",
		"template_guid": "550e8400-e29b-41d4-a716-446655440000",
		"template_name": "Welcome Email",
		"messages": [
			{"to": "user@example.com", "message_guid": "0193f3c4-aaaa-bbbb-cccc-000000000003"}
		]
	}`)

	var resp smtpsdk.MessageTemplateResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("MessageTemplateResponse unmarshal error = %v", err)
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if resp.TemplateGUID != "550e8400-e29b-41d4-a716-446655440000" {
		t.Errorf("TemplateGUID = %v, want 550e8400-e29b-41d4-a716-446655440000", resp.TemplateGUID)
	}
	if len(resp.Messages) != 1 {
		t.Fatalf("len(Messages) = %d, want 1", len(resp.Messages))
	}
	if resp.Messages[0].To != "user@example.com" {
		t.Errorf("Messages[0].To = %v, want user@example.com", resp.Messages[0].To)
	}
	if resp.Messages[0].MessageGUID != "0193f3c4-aaaa-bbbb-cccc-000000000003" {
		t.Errorf("Messages[0].MessageGUID = %v, want 0193f3c4-aaaa-bbbb-cccc-000000000003", resp.Messages[0].MessageGUID)
	}
}
