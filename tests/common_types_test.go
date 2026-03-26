package tests

import (
	"testing"

	smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"
)

func TestQuotaResponse(t *testing.T) {
	resp := smtpsdk.QuotaResponse{
		APIResponse: smtpsdk.APIResponse{
			Success: true,
		},
		Data: struct {
			CurrentQuota    int     `json:"current_quota"`
			MaxQuota        int     `json:"max_quota"`
			UsagePercentage float64 `json:"usage_percentage"`
		}{
			CurrentQuota:    9500,
			MaxQuota:        10000,
			UsagePercentage: 95.0,
		},
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if resp.Data.CurrentQuota != 9500 {
		t.Errorf("CurrentQuota = %v, want 9500", resp.Data.CurrentQuota)
	}
	if resp.Data.MaxQuota != 10000 {
		t.Errorf("MaxQuota = %v, want 10000", resp.Data.MaxQuota)
	}
	if resp.Data.UsagePercentage != 95.0 {
		t.Errorf("UsagePercentage = %v, want 95.0", resp.Data.UsagePercentage)
	}
}

func TestQuotaResponse_LowUsage(t *testing.T) {
	resp := smtpsdk.QuotaResponse{
		APIResponse: smtpsdk.APIResponse{
			Success: true,
		},
		Data: struct {
			CurrentQuota    int     `json:"current_quota"`
			MaxQuota        int     `json:"max_quota"`
			UsagePercentage float64 `json:"usage_percentage"`
		}{
			CurrentQuota:    1000,
			MaxQuota:        10000,
			UsagePercentage: 10.0,
		},
	}

	if resp.Data.UsagePercentage != 10.0 {
		t.Errorf("UsagePercentage = %v, want 10.0", resp.Data.UsagePercentage)
	}
}

func TestQuotaResponse_HighUsage(t *testing.T) {
	resp := smtpsdk.QuotaResponse{
		APIResponse: smtpsdk.APIResponse{
			Success: true,
		},
		Data: struct {
			CurrentQuota    int     `json:"current_quota"`
			MaxQuota        int     `json:"max_quota"`
			UsagePercentage float64 `json:"usage_percentage"`
		}{
			CurrentQuota:    500,
			MaxQuota:        10000,
			UsagePercentage: 95.0,
		},
	}

	if resp.Data.CurrentQuota != 500 {
		t.Errorf("CurrentQuota = %v, want 500", resp.Data.CurrentQuota)
	}
	if resp.Data.UsagePercentage < 90.0 {
		t.Errorf("UsagePercentage should be high, got %v", resp.Data.UsagePercentage)
	}
}

func TestAPIResponse(t *testing.T) {
	resp := smtpsdk.APIResponse{
		Success: true,
		Message: "Operation successful",
	}

	if !resp.Success {
		t.Errorf("Success = false, want true")
	}
	if resp.Message != "Operation successful" {
		t.Errorf("Message = %v, want 'Operation successful'", resp.Message)
	}
}

func TestAPIResponse_Error(t *testing.T) {
	resp := smtpsdk.APIResponse{
		Success: false,
		Error:   "Invalid request",
	}

	if resp.Success {
		t.Errorf("Success = true, want false")
	}
	if resp.Error != "Invalid request" {
		t.Errorf("Error = %v, want 'Invalid request'", resp.Error)
	}
}
