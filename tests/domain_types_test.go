package tests

import (
	"encoding/json"
	"testing"

	smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"
)

func TestDomainStruct(t *testing.T) {
	data := `{"id":1,"domain":"example.com","tracklink_domain":"track.example.com","tracklink_domain_is_verified":true,"auth_domain_is_verified":true,"dns_selector":"default","dns_record":"v=DKIM1;k=rsa;p=publickey","click_track":true,"open_track":true,"unsub_track":true,"is_verified":true,"status":true,"created_at":1234567890,"modified_at":1234567890,"spf_record":"v=spf1 include:example.com ~all"}`
	var domain smtpsdk.Domain
	err := json.Unmarshal([]byte(data), &domain)
	if err != nil {
		t.Errorf("Failed to unmarshal: %v", err)
	}
	if domain.Domain != "example.com" {
		t.Errorf("Domain = %v, want example.com", domain.Domain)
	}
	if !domain.IsVerified {
		t.Errorf("IsVerified = false, want true")
	}
}

func TestDomainCreateRequest(t *testing.T) {
	req := smtpsdk.DomainCreateRequest{
		Domain:        "example.com",
		DKIMKeyLength: 2048,
	}
	if req.Domain != "example.com" {
		t.Errorf("Domain not set correctly")
	}
	if req.DKIMKeyLength != 2048 {
		t.Errorf("DKIMKeyLength not set correctly")
	}
}

func TestDomainUpdateRequest(t *testing.T) {
	openTrack := true
	clickTrack := true
	unsubTrack := false

	req := smtpsdk.DomainUpdateRequest{
		OpenTrack:  &openTrack,
		ClickTrack: &clickTrack,
		UnsubTrack: &unsubTrack,
	}

	if req.OpenTrack == nil || !*req.OpenTrack {
		t.Errorf("OpenTrack not set correctly")
	}
	if req.UnsubTrack == nil || *req.UnsubTrack {
		t.Errorf("UnsubTrack not set correctly")
	}
}

func TestAuthDomainSetupRequest(t *testing.T) {
	req := smtpsdk.AuthDomainSetupRequest{
		AuthDomain:    "auth.example.com",
		DKIMKeyLength: 2048,
	}
	if req.AuthDomain != "auth.example.com" {
		t.Errorf("AuthDomain not set correctly")
	}
}

func TestTracklinkSetupRequest(t *testing.T) {
	req := smtpsdk.TracklinkSetupRequest{
		TrackingDomain: "track.example.com",
	}
	if req.TrackingDomain != "track.example.com" {
		t.Errorf("TrackingDomain not set correctly")
	}
}

func TestDNSVerificationResponse(t *testing.T) {
	data := `{"records":{"dkim":true,"spf":true,"mx":true}}`
	var result smtpsdk.DNSVerificationResponse
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		t.Errorf("Failed to unmarshal: %v", err)
	}
	if !result.Records.DKIM || !result.Records.SPF || !result.Records.MX {
		t.Errorf("Records not all true")
	}
}

func TestPaginationStruct(t *testing.T) {
	data := `{"total":100,"page":2,"limit":10,"offset":10}`
	var pag smtpsdk.Pagination
	err := json.Unmarshal([]byte(data), &pag)
	if err != nil {
		t.Errorf("Failed to unmarshal: %v", err)
	}
	if pag.Total != 100 {
		t.Errorf("Total = %v, want 100", pag.Total)
	}
	if pag.Page != 2 {
		t.Errorf("Page = %v, want 2", pag.Page)
	}
}
