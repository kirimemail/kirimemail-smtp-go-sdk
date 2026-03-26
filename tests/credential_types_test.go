package tests

import (
	"testing"

	smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"
)

func TestCredentialStruct(t *testing.T) {
	cred := smtpsdk.Credential{
		ID:           1,
		UserSMTPGUID: "guid123",
		Username:     "testuser",
		CreatedAt:    1234567890,
		ModifiedAt:   1234567890,
	}

	if cred.Username != "testuser" {
		t.Errorf("Username = %v, want testuser", cred.Username)
	}
	if cred.ID != 1 {
		t.Errorf("ID = %v, want 1", cred.ID)
	}
}

func TestCredentialCreateRequest(t *testing.T) {
	req := smtpsdk.CredentialCreateRequest{
		Username: "testuser",
	}

	if req.Username != "testuser" {
		t.Errorf("Username not set correctly")
	}
}

func TestCredentialCreateResponse(t *testing.T) {
	resp := smtpsdk.CredentialCreateResponse{
		APIResponse: smtpsdk.APIResponse{
			Success: true,
			Message: "Credential created",
		},
		Data: struct {
			Credential   smtpsdk.Credential `json:"credential"`
			Password     string             `json:"password"`
			RemoteSynced bool               `json:"remote_synced"`
			StrengthInfo interface{}        `json:"strength_info"`
		}{
			Password:     "secret123",
			RemoteSynced: true,
		},
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if resp.Data.Password != "secret123" {
		t.Errorf("Password not set correctly")
	}
	if !resp.Data.RemoteSynced {
		t.Errorf("RemoteSynced = false, want true")
	}
}
