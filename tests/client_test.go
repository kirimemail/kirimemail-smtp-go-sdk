package tests

import (
	"testing"

	smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"
)

func TestNewClient(t *testing.T) {
	client := smtpsdk.NewClient("testuser", "testtoken")
	if client.Username != "testuser" {
		t.Errorf("Username = %v, want testuser", client.Username)
	}
	if client.Token != "testtoken" {
		t.Errorf("Token = %v, want testtoken", client.Token)
	}
	if client.BaseURL != "https://smtp-app.kirim.email" {
		t.Errorf("BaseURL = %v, want https://smtp-app.kirim.email", client.BaseURL)
	}
}

func TestNewClientWithBaseURL(t *testing.T) {
	customURL := "https://custom.example.com"
	client := smtpsdk.NewClient("testuser", "testtoken", smtpsdk.WithBaseURL(customURL))
	if client.BaseURL != customURL {
		t.Errorf("BaseURL = %v, want %v", client.BaseURL, customURL)
	}
}

func TestIntPtr(t *testing.T) {
	value := 42
	ptr := smtpsdk.IntPtr(value)
	if *ptr != value {
		t.Errorf("IntPtr() = %v, want %v", *ptr, value)
	}
}

func TestBoolPtr(t *testing.T) {
	value := true
	ptr := smtpsdk.BoolPtr(value)
	if *ptr != value {
		t.Errorf("BoolPtr() = %v, want %v", *ptr, value)
	}
}

func TestStringPtr(t *testing.T) {
	value := "test"
	ptr := smtpsdk.StringPtr(value)
	if *ptr != value {
		t.Errorf("StringPtr() = %v, want %v", *ptr, value)
	}
}
