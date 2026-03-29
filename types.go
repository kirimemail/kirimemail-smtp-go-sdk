package smtpsdk

import "encoding/json"

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Credential struct {
	ID           int64  `json:"id"`
	UserSMTPGUID string `json:"user_smtp_guid"`
	Username     string `json:"username"`
	CreatedAt    int64  `json:"created_at"`
	ModifiedAt   int64  `json:"modified_at"`
}

type CredentialCreateRequest struct {
	Username string `json:"username"`
}

type CredentialCreateResponse struct {
	APIResponse
	Data struct {
		Credential   Credential  `json:"credential"`
		Password     string      `json:"password"`
		RemoteSynced bool        `json:"remote_synced"`
		StrengthInfo interface{} `json:"strength_info"`
	} `json:"data"`
}

type CredentialListResponse struct {
	APIResponse
	Data   []Credential `json:"data"`
	Domain string       `json:"domain"`
}

type CredentialResetPasswordResponse struct {
	APIResponse
	Data struct {
		Credential   Credential  `json:"credential"`
		NewPassword  string      `json:"new_password"`
		StrengthInfo interface{} `json:"strength_info"`
		RemoteSynced bool        `json:"remote_synced"`
	} `json:"data"`
}

type Domain struct {
	ID                int64  `json:"id"`
	Domain            string `json:"domain"`
	TracklinkDomain   string `json:"tracklink_domain"`
	TracklinkVerified bool   `json:"tracklink_domain_is_verified"`
	AuthVerified      bool   `json:"auth_domain_is_verified"`
	DNSSelector       string `json:"dns_selector"`
	DNSRecord         string `json:"dns_record"`
	ClickTrack        bool   `json:"click_track"`
	OpenTrack         bool   `json:"open_track"`
	UnsubTrack        bool   `json:"unsub_track"`
	IsVerified        bool   `json:"is_verified"`
	Status            bool   `json:"status"`
	CreatedAt         int64  `json:"created_at"`
	ModifiedAt        int64  `json:"modified_at"`
	SPFRecord         string `json:"spf_record"`
}

type DomainCreateRequest struct {
	Domain        string `json:"domain"`
	DKIMKeyLength int    `json:"dkim_key_length"`
}

type DomainCreateResponse struct {
	APIResponse
	Data struct {
		Domain string `json:"domain"`
	} `json:"data"`
}

type DomainListResponse struct {
	APIResponse
	Data       []Domain   `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type DomainUpdateRequest struct {
	OpenTrack  *bool `json:"open_track,omitempty"`
	ClickTrack *bool `json:"click_track,omitempty"`
	UnsubTrack *bool `json:"unsub_track,omitempty"`
}

type DomainUpdateResponse struct {
	APIResponse
	Data struct {
		OpenTrack  bool `json:"open_track"`
		ClickTrack bool `json:"click_track"`
		UnsubTrack bool `json:"unsub_track"`
	} `json:"data"`
}

type AuthDomainSetupRequest struct {
	AuthDomain    string `json:"auth_domain,omitempty"`
	DKIMKeyLength int    `json:"dkim_key_length"`
}

type AuthDomainSetupResponse struct {
	APIResponse
	Data struct {
		AuthDomain string `json:"auth_domain"`
	} `json:"data"`
}

type DNSVerificationResponse struct {
	Records struct {
		DKIM bool `json:"dkim"`
		SPF  bool `json:"spf"`
		MX   bool `json:"mx"`
	} `json:"records"`
}

type AuthDomainVerificationResponse struct {
	Records struct {
		AuthDKIM bool `json:"auth_dkim"`
		AuthSPF  bool `json:"auth_spf"`
		AuthMX   bool `json:"auth_mx"`
	} `json:"records"`
}

type TracklinkSetupRequest struct {
	TrackingDomain string `json:"tracking_domain"`
}

type TracklinkSetupResponse struct {
	APIResponse
	Data struct {
		TrackingDomain string `json:"tracking_domain"`
	} `json:"data"`
}

type TracklinkVerificationResponse struct {
	Records struct {
		CNAME          string `json:"cname"`
		TrackingDomain string `json:"tracking_domain"`
	} `json:"records"`
}

type Pagination struct {
	Total  int `json:"total"`
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type EmailValidationRequest struct {
	Email string `json:"email"`
}

type EmailValidationResult struct {
	Email         string   `json:"email"`
	IsValid       bool     `json:"is_valid"`
	Error         *string  `json:"error"`
	Warnings      []string `json:"warnings"`
	Cached        bool     `json:"cached"`
	ValidatedAt   string   `json:"validated_at"`
	IsSpamtrap    bool     `json:"is_spamtrap"`
	SpamtrapScore float64  `json:"spamtrap_score"`
}

type EmailValidationResponse struct {
	APIResponse
	Data EmailValidationResult `json:"data"`
}

type EmailValidationBatchRequest struct {
	Emails []string `json:"emails"`
}

type EmailValidationBatchSummary struct {
	Total     int `json:"total"`
	Valid     int `json:"valid"`
	Invalid   int `json:"invalid"`
	Cached    int `json:"cached"`
	Validated int `json:"validated"`
}

type EmailValidationBatchResponse struct {
	APIResponse
	Data struct {
		Results []EmailValidationResult     `json:"results"`
		Summary EmailValidationBatchSummary `json:"summary"`
	} `json:"data"`
}

type MessageSendRequest struct {
	From              string      `json:"from"`
	FromName          string      `json:"from_name,omitempty"`
	To                interface{} `json:"to"`
	Subject           string      `json:"subject"`
	Text              string      `json:"text"`
	HTML              string      `json:"html,omitempty"`
	Headers           interface{} `json:"headers,omitempty"`
	ReplyTo           string      `json:"reply_to,omitempty"`
	Attachments       interface{} `json:"attachments,omitempty"`
	AttachmentOptions string      `json:"attachment_options,omitempty"`
}

type MessageSendResponse struct {
	APIResponse
	Message string `json:"message"`
}

type MessageTemplateRequest struct {
	TemplateGUID      string      `json:"template_guid"`
	To                interface{} `json:"to"`
	From              string      `json:"from,omitempty"`
	FromName          string      `json:"from_name,omitempty"`
	Variables         interface{} `json:"variables,omitempty"`
	Headers           interface{} `json:"headers,omitempty"`
	ReplyTo           string      `json:"reply_to,omitempty"`
	Attachments       interface{} `json:"attachments,omitempty"`
	AttachmentOptions string      `json:"attachment_options,omitempty"`
}

type MessageTemplateResponse struct {
	APIResponse
	Message      string `json:"message"`
	TemplateGUID string `json:"template_guid,omitempty"`
	TemplateName string `json:"template_name,omitempty"`
}

type LogEntry struct {
	ID             string `json:"id"`
	UserGUID       string `json:"user_guid"`
	UserDomainGUID string `json:"user_domain_guid"`
	EventType      string `json:"event_type"`
	MessageGUID    string `json:"message_guid"`
	Timestamp      int64  `json:"timestamp"`
}

const (
	LogEventTypeQueued        = "queued"
	LogEventTypeSend          = "send"
	LogEventTypeDelivered     = "delivered"
	LogEventTypeBounced       = "bounced"
	LogEventTypeFailed        = "failed"
	LogEventTypeOpened        = "opened"
	LogEventTypeClicked       = "clicked"
	LogEventTypeUnsubscribed  = "unsubscribed"
	LogEventTypeTemporaryFail = "temporary_fail"
	LogEventTypePermanentFail = "permanent_fail"
	LogEventTypeDeferred      = "deferred"
)

var LogEventTypes = []string{
	LogEventTypeQueued,
	LogEventTypeSend,
	LogEventTypeDelivered,
	LogEventTypeBounced,
	LogEventTypeFailed,
	LogEventTypeOpened,
	LogEventTypeClicked,
	LogEventTypeUnsubscribed,
	LogEventTypeTemporaryFail,
	LogEventTypePermanentFail,
	LogEventTypeDeferred,
}

func IsValidLogEventType(eventType string) bool {
	for _, et := range LogEventTypes {
		if et == eventType {
			return true
		}
	}
	return false
}

type LogListResponse struct {
	Data       []LogEntry `json:"data"`
	Count      int        `json:"count"`
	Offset     int        `json:"offset"`
	Limit      int        `json:"limit"`
	Pagination struct {
		Total       int `json:"total"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		LastPage    int `json:"last_page"`
	} `json:"pagination"`
}

type LogMessageResponse struct {
	Data []LogEntry `json:"data"`
}

type QuotaResponse struct {
	APIResponse
	Data struct {
		CurrentQuota    int     `json:"current_quota"`
		MaxQuota        int     `json:"max_quota"`
		UsagePercentage float64 `json:"usage_percentage"`
	} `json:"data"`
}

type Suppression struct {
	ID            int64   `json:"id"`
	Type          string  `json:"type"`
	RecipientType string  `json:"recipient_type"`
	Recipient     string  `json:"recipient"`
	Tags          *string `json:"tags"`
	Description   *string `json:"description"`
	Source        *string `json:"source"`
	CreatedAt     int64   `json:"created_at"`
}

type SuppressionListResponse struct {
	APIResponse
	Data       []Suppression `json:"data"`
	Pagination interface{}   `json:"pagination,omitempty"`
	Filters    interface{}   `json:"filters,omitempty"`
}

type WhitelistCreateRequest struct {
	Recipient     string `json:"recipient"`
	RecipientType string `json:"recipient_type"`
	Description   string `json:"description,omitempty"`
}

type WhitelistCreateResponse struct {
	APIResponse
	Data Suppression `json:"data"`
}

type SuppressionDeleteRequest struct {
	IDs []int `json:"ids"`
}

type SuppressionDeleteResponse struct {
	APIResponse
	Message      string `json:"message"`
	DeletedCount int    `json:"deleted_count"`
}

type Webhook struct {
	WebhookGUID    string `json:"webhook_guid"`
	UserGUID       string `json:"user_guid"`
	UserDomainGUID string `json:"user_domain_guid"`
	UserSMTPGUID   string `json:"user_smtp_guid"`
	Type           string `json:"type"`
	URL            string `json:"url"`
	IsDeleted      bool   `json:"is_deleted"`
	CreatedAt      int64  `json:"created_at"`
	ModifiedAt     int64  `json:"modified_at"`
}

type WebhookCreateRequest struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type WebhookCreateResponse struct {
	APIResponse
	Message string  `json:"message"`
	Data    Webhook `json:"data"`
}

type WebhookUpdateRequest struct {
	Type *string `json:"type,omitempty"`
	URL  *string `json:"url,omitempty"`
}

type WebhookUpdateResponse struct {
	APIResponse
	Message string  `json:"message"`
	Data    Webhook `json:"data"`
}

type WebhookListResponse struct {
	APIResponse
	Data  []Webhook `json:"data"`
	Count int       `json:"count"`
}

type WebhookTestRequest struct {
	URL       string `json:"url"`
	EventType string `json:"event_type"`
}

type WebhookTestResponse struct {
	APIResponse
	Message string `json:"message"`
	Data    struct {
		URL            string `json:"url"`
		EventType      string `json:"event_type"`
		ResponseStatus int    `json:"response_status"`
		ResponseTime   int    `json:"response_time"`
	} `json:"data"`
}

type Headers map[string]string

func (h Headers) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string(h))
}

func (h *Headers) UnmarshalJSON(data []byte) error {
	var m map[string]string
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	*h = Headers(m)
	return nil
}
