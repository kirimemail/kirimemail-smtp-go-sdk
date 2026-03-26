package tests

import (
	"encoding/json"
	"testing"

	smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"
)

func TestEmailValidationResult(t *testing.T) {
	data := `{"email":"test@example.com","is_valid":true,"error":null,"warnings":["Warning 1"],"cached":false,"validated_at":"2024-01-01T00:00:00Z","is_spamtrap":false,"spamtrap_score":0.0}`
	var result smtpsdk.EmailValidationResult
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		t.Errorf("Failed to unmarshal: %v", err)
	}
	if result.Email != "test@example.com" {
		t.Errorf("Email = %v, want test@example.com", result.Email)
	}
	if !result.IsValid {
		t.Errorf("IsValid = false, want true")
	}
	if len(result.Warnings) != 1 {
		t.Errorf("Expected 1 warning, got %d", len(result.Warnings))
	}
	if result.IsSpamtrap {
		t.Errorf("IsSpamtrap = true, want false")
	}
}

func TestEmailValidationResult_Invalid(t *testing.T) {
	data := `{"email":"invalid-email","is_valid":false,"error":"Invalid email format","warnings":[]}`
	var result smtpsdk.EmailValidationResult
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		t.Errorf("Failed to unmarshal: %v", err)
	}
	if result.IsValid {
		t.Errorf("IsValid = true, want false")
	}
	if result.Error == nil {
		t.Errorf("Error is nil, want error message")
	}
}

func TestEmailValidationResult_Spamtrap(t *testing.T) {
	data := `{"email":"trap@spamtrap.com","is_valid":true,"is_spamtrap":true,"spamtrap_score":0.95}`
	var result smtpsdk.EmailValidationResult
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		t.Errorf("Failed to unmarshal: %v", err)
	}
	if !result.IsSpamtrap {
		t.Errorf("IsSpamtrap = false, want true")
	}
	if result.SpamtrapScore != 0.95 {
		t.Errorf("SpamtrapScore = %v, want 0.95", result.SpamtrapScore)
	}
}

func TestEmailValidationBatchSummary(t *testing.T) {
	data := `{"total":10,"valid":8,"invalid":2,"cached":3,"validated":7}`
	var summary smtpsdk.EmailValidationBatchSummary
	err := json.Unmarshal([]byte(data), &summary)
	if err != nil {
		t.Errorf("Failed to unmarshal: %v", err)
	}
	if summary.Total != 10 {
		t.Errorf("Total = %v, want 10", summary.Total)
	}
	if summary.Valid != 8 {
		t.Errorf("Valid = %v, want 8", summary.Valid)
	}
	if summary.Invalid != 2 {
		t.Errorf("Invalid = %v, want 2", summary.Invalid)
	}
	if summary.Cached != 3 {
		t.Errorf("Cached = %v, want 3", summary.Cached)
	}
	if summary.Validated != 7 {
		t.Errorf("Validated = %v, want 7", summary.Validated)
	}
}

func TestEmailValidationRequest(t *testing.T) {
	req := smtpsdk.EmailValidationRequest{
		Email: "test@example.com",
	}

	if req.Email != "test@example.com" {
		t.Errorf("Email not set correctly")
	}
}

func TestEmailValidationBatchRequest(t *testing.T) {
	req := smtpsdk.EmailValidationBatchRequest{
		Emails: []string{
			"test1@example.com",
			"test2@example.com",
			"test3@example.com",
		},
	}

	if len(req.Emails) != 3 {
		t.Errorf("Expected 3 emails, got %d", len(req.Emails))
	}
}
