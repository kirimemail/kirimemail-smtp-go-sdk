package tests

import (
	"testing"

	smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"
)

func TestLogEntry(t *testing.T) {
	entry := smtpsdk.LogEntry{
		ID:             "log123",
		UserGUID:       "user123",
		UserDomainGUID: "domain123",
		EventType:      "delivered",
		MessageGUID:    "msg123",
		Timestamp:      1234567890,
	}

	if entry.EventType != "delivered" {
		t.Errorf("EventType = %v, want delivered", entry.EventType)
	}
	if entry.MessageGUID != "msg123" {
		t.Errorf("MessageGUID = %v, want msg123", entry.MessageGUID)
	}
}

func TestLogEntry_EventTypes(t *testing.T) {
	eventTypes := []string{
		"queued", "send", "delivered", "bounced", "failed",
		"permanent_fail", "opened", "clicked", "unsubscribed", "temporary_fail", "deferred",
	}

	for _, eventType := range eventTypes {
		entry := smtpsdk.LogEntry{
			ID:        "log1",
			EventType: eventType,
			Timestamp: 1234567890,
		}

		if entry.EventType != eventType {
			t.Errorf("EventType not set correctly: %v", eventType)
		}
	}
}

func TestLogListResponse(t *testing.T) {
	resp := smtpsdk.LogListResponse{
		Data: []smtpsdk.LogEntry{
			{
				ID:        "log1",
				EventType: "delivered",
			},
		},
		Count:  1,
		Offset: 0,
		Limit:  10,
		Pagination: struct {
			Total       int `json:"total"`
			PerPage     int `json:"per_page"`
			CurrentPage int `json:"current_page"`
			LastPage    int `json:"last_page"`
		}{
			Total:       100,
			PerPage:     10,
			CurrentPage: 1,
			LastPage:    10,
		},
	}

	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 log entry")
	}
	if resp.Count != 1 {
		t.Errorf("Count = %v, want 1", resp.Count)
	}
	if resp.Pagination.Total != 100 {
		t.Errorf("Pagination.Total = %v, want 100", resp.Pagination.Total)
	}
}

func TestLogMessageResponse(t *testing.T) {
	resp := smtpsdk.LogMessageResponse{
		Data: []smtpsdk.LogEntry{
			{
				ID:        "log1",
				EventType: "delivered",
			},
			{
				ID:        "log2",
				EventType: "opened",
			},
		},
	}

	if len(resp.Data) != 2 {
		t.Errorf("Expected 2 log entries, got %d", len(resp.Data))
	}
}
