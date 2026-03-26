package tests

import (
	"testing"

	smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"
)

func TestSuppression(t *testing.T) {
	sup := smtpsdk.Suppression{
		ID:            1,
		Type:          "unsubscribe",
		RecipientType: "email",
		Recipient:     "user@example.com",
		Tags:          stringPtr("test"),
		Description:   stringPtr("Unsubscribed user"),
		Source:        stringPtr("api"),
		CreatedAt:     1234567890,
	}

	if sup.Type != "unsubscribe" {
		t.Errorf("Type = %v, want unsubscribe", sup.Type)
	}
	if sup.Recipient != "user@example.com" {
		t.Errorf("Recipient = %v, want user@example.com", sup.Recipient)
	}
}

func TestSuppressionTypes(t *testing.T) {
	types := []string{"unsubscribe", "bounce", "whitelist"}

	for _, supType := range types {
		sup := smtpsdk.Suppression{
			ID:        1,
			Type:      supType,
			Recipient: "test@example.com",
		}

		if sup.Type != supType {
			t.Errorf("Type not set correctly: %v", supType)
		}
	}
}

func TestWhitelistCreateRequest(t *testing.T) {
	req := smtpsdk.WhitelistCreateRequest{
		Recipient:     "user@example.com",
		RecipientType: "email",
		Description:   "Trusted customer",
	}

	if req.Recipient != "user@example.com" {
		t.Errorf("Recipient not set correctly")
	}
	if req.RecipientType != "email" {
		t.Errorf("RecipientType = %v, want email", req.RecipientType)
	}
	if req.Description != "Trusted customer" {
		t.Errorf("Description not set correctly")
	}
}

func TestWhitelistCreateRequest_Domain(t *testing.T) {
	req := smtpsdk.WhitelistCreateRequest{
		Recipient:     "trusted-domain.com",
		RecipientType: "domain",
		Description:   "Trusted partner domain",
	}

	if req.RecipientType != "domain" {
		t.Errorf("RecipientType = %v, want domain", req.RecipientType)
	}
}

func TestSuppressionDeleteRequest(t *testing.T) {
	req := smtpsdk.SuppressionDeleteRequest{
		IDs: []int{1, 2, 3},
	}

	if len(req.IDs) != 3 {
		t.Errorf("Expected 3 IDs, got %d", len(req.IDs))
	}
}

func TestSuppressionDeleteResponse(t *testing.T) {
	resp := smtpsdk.SuppressionDeleteResponse{
		APIResponse: smtpsdk.APIResponse{
			Success: true,
			Message: "Deleted successfully",
		},
		Message:      "Deleted successfully",
		DeletedCount: 3,
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if resp.DeletedCount != 3 {
		t.Errorf("DeletedCount = %v, want 3", resp.DeletedCount)
	}
}

func TestSuppressionListResponse(t *testing.T) {
	resp := smtpsdk.SuppressionListResponse{
		APIResponse: smtpsdk.APIResponse{
			Success: true,
		},
		Data: []smtpsdk.Suppression{
			{
				ID:        1,
				Type:      "unsubscribe",
				Recipient: "user@example.com",
			},
		},
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 suppression")
	}
}

func TestWhitelistCreateResponse(t *testing.T) {
	resp := smtpsdk.WhitelistCreateResponse{
		APIResponse: smtpsdk.APIResponse{
			Success: true,
		},
		Data: smtpsdk.Suppression{
			ID:        1,
			Type:      "whitelist",
			Recipient: "user@example.com",
		},
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if resp.Data.Type != "whitelist" {
		t.Errorf("Type = %v, want whitelist", resp.Data.Type)
	}
}

func stringPtr(s string) *string {
	return &s
}
