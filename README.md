# SMTP SDK for Go

Go SDK for Kirim.Email SMTP API - A comprehensive SDK for managing domains, credentials, email sending, validation, logs, suppressions, and webhooks.

## Installation

```bash
go get github.com/kirimemail/kirimemail-smtp-go-sdk
```

## Getting Started

```go
package main

import (
    "fmt"
    "log"
    
    smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"
)

func main() {
    client := smtpsdk.NewClient("your-username", "your-api-token")
    
    domains, err := client.Domains().List(nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d domains\n", len(domains.Data))
}
```

## Features

- **Domain Management**: List, create, update, delete domains; setup and verify DNS records
- **Credential Management**: List, create, get, delete SMTP credentials; reset passwords
- **Email Sending**: Send transactional emails with or without templates; support for attachments
- **Email Validation**: Single and batch email validation with strict mode options
- **Logs**: Retrieve email logs, including streaming via Server-Sent Events
- **Suppressions**: Manage unsubscribe, bounce, and whitelist suppressions
- **Webhooks**: Full webhook lifecycle management (create, list, get, update, delete, test)
- **User**: Retrieve quota information

## Domain Management

```go
client := smtpsdk.NewClient("username", "token")

// List all domains
domains, err := client.Domains().List(nil, nil, nil)

// Create a new domain
domain, err := client.Domains().Create(smtpsdk.DomainCreateRequest{
    Domain:        "example.com",
    DKIMKeyLength: 2048,
})

// Get a specific domain
domain, err := client.Domains().Get("example.com")

// Update domain tracking settings
domain, err := client.Domains().Update("example.com", smtpsdk.DomainUpdateRequest{
    OpenTrack:  smtpsdk.BoolPtr(true),
    ClickTrack: smtpsdk.BoolPtr(true),
})

// Delete a domain
err := client.Domains().Delete("example.com")

// Setup authentication domain
authDomain, err := client.Domains().SetupAuthDomain("example.com", smtpsdk.AuthDomainSetupRequest{
    AuthDomain:    "auth.example.com",
    DKIMKeyLength: 2048,
})

// Verify mandatory DNS records (DKIM, SPF, MX)
verification, err := client.Domains().VerifyMandatoryRecords("example.com")
fmt.Printf("DKIM: %v, SPF: %v, MX: %v\n", 
    verification.Records.DKIM, verification.Records.SPF, verification.Records.MX)

// Verify authentication domain records
authVerification, err := client.Domains().VerifyAuthDomain("example.com")

// Setup tracking domain
tracklink, err := client.Domains().SetupTracklink("example.com", smtpsdk.TracklinkSetupRequest{
    TrackingDomain: "track.example.com",
})

// Verify tracking domain
verification, err := client.Domains().VerifyTracklink("example.com")
```

## Credential Management

```go
// List credentials for a domain
credentials, err := client.Credentials().List("example.com", nil, nil)

// Create a new credential
credential, err := client.Credentials().Create("example.com", smtpsdk.CredentialCreateRequest{
    Username: "myuser",
})
fmt.Printf("Generated password: %s\n", credential.Data.Password)

// Get a specific credential
credential, err := client.Credentials().Get("example.com", "credential-guid")

// Delete a credential
err := client.Credentials().Delete("example.com", "credential-guid")

// Reset credential password
result, err := client.Credentials().ResetPassword("example.com", "credential-guid")
fmt.Printf("New password: %s\n", result.Data.NewPassword)
```

## Email Sending

### Send Email

```go
response, err := client.Messages().Send("example.com", smtpsdk.MessageSendRequest{
    From:     "noreply@example.com",
    FromName: "Company Name",
    To:       []string{"customer@example.com"},
    Subject:  "Welcome",
    Text:     "Hello World",
    HTML:     "<h1>Hello World</h1>",
    Headers:  smtpsdk.Headers{"X-Custom-Header": "value"},
}, nil)
```

### Send with Attachments

```go
// Read file
fileData, _ := os.ReadFile("document.pdf")

response, err := client.Messages().Send("example.com", smtpsdk.MessageSendRequest{
    From:    "noreply@example.com",
    To:      []string{"customer@example.com"},
    Subject: "Document attached",
    Text:    "Please find attached document",
}, map[string][]byte{
    "document.pdf": fileData,
})

// Multiple attachments
response, err := client.Messages().Send("example.com", smtpsdk.MessageSendRequest{
    From:   "noreply@example.com",
    To:      []string{"customer@example.com"},
    Subject: "Multiple files",
    Text:    "Attachments included",
}, map[string][]byte{
    "file1.pdf": data1,
    "file2.txt": data2,
})
```

### Send Using Template

```go
// Send with template
response, err := client.Messages().SendTemplate("example.com", smtpsdk.MessageTemplateRequest{
    TemplateGUID: "template-guid",
    To:           []string{"customer@example.com"},
    Variables: map[string]interface{}{
        "name": "John Doe",
    },
}, nil)

// Bulk send with personalized variables
response, err := client.Messages().SendTemplate("example.com", smtpsdk.MessageTemplateRequest{
    TemplateGUID: "template-guid",
    To: []interface{}{
        "user1@example.com",
        "user2@example.com",
    },
    Variables: []map[string]interface{}{
        {"name": "John Doe"},
        {"name": "Jane Smith"},
    },
}, nil)
```

### Advanced Email Features

```go
// Custom headers
headers := smtpsdk.Headers{
    "X-Priority":      "1",
    "X-Order-ID":      "12345",
    "List-Unsubscribe": "<mailto:unsubscribe@example.com>",
}

// Attachment options for processing
attachmentOptions := `{
    "compress": true,
    "password": "SecureDoc2024",
    "watermark": {
        "enabled": true,
        "text": "CONFIDENTIAL",
        "position": "center"
    }
}`

response, err := client.Messages().Send("example.com", smtpsdk.MessageSendRequest{
    From:             "noreply@example.com",
    To:               []string{"customer@example.com"},
    Subject:          "Secure Document",
    Text:             "Attached secure document",
    Headers:          headers,
    AttachmentOptions: attachmentOptions,
}, map[string][]byte{
    "document.pdf": fileData,
})
```

## Email Validation

```go
// Validate single email
result, err := client.Validation().ValidateEmail("user@example.com")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Valid: %v\n", result.Data.IsValid)
fmt.Printf("Is spamtrap: %v\n", result.Data.IsSpamtrap)
fmt.Printf("Spamtrap score: %.2f\n", result.Data.SpamtrapScore)

// Strict validation (no warnings)
result, err := client.Validation().ValidateEmailStrict("user@example.com")

// Batch validation (max 100 emails)
result, err := client.Validation().ValidateEmailsBatch([]string{
    "user1@example.com",
    "user2@example.com",
    "invalid-email",
})
fmt.Printf("Total: %d, Valid: %d, Invalid: %d\n", 
    result.Data.Summary.Total, 
    result.Data.Summary.Valid, 
    result.Data.Summary.Invalid)

// Batch with strict mode
result, err := client.Validation().ValidateEmailsBatchStrict(emails)
```

## Logs

```go
import (
    "bufio"
    "net/http"
    "time"
)

// Get logs with filters
limit := 100
offset := 0
logs, err := client.Logs().Get("example.com", nil, nil, nil, nil, nil, nil, nil, &limit, &offset)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found %d log entries\n", len(logs.Data))

// Filter by date range
start := "2024-01-01T00:00:00Z"
end := "2024-12-31T23:59:59Z"
logs, err := client.Logs().Get("example.com", &start, &end, nil, nil, nil, nil, nil, nil, nil)

// Filter by sender/recipient
logs, err := client.Logs().Get("example.com", nil, nil, 
    smtpsdk.StringPtr("sender@example.com"), 
    smtpsdk.StringPtr("recipient@example.com"), 
    nil, nil, nil, nil)

// Filter by event type
logs, err = client.Logs().GetByEventType("example.com", smtpsdk.LogEventTypeDelivered, nil, nil, nil, nil, nil, nil, nil, nil)

// Filter by tags
logs, err = client.Logs().GetByTags("example.com", "newsletter", nil, nil, nil, nil, nil, nil, nil, nil)

// Filter by subject (partial match)
logs, err = client.Logs().Get("example.com", nil, nil, nil, nil, smtpsdk.StringPtr("Welcome"), nil, nil, nil, nil)

// Combined filters
eventType := smtpsdk.LogEventTypeOpened
tags := "campaign-2024"
logs, err = client.Logs().Get("example.com", &start, &end, 
    smtpsdk.StringPtr("sender@example.com"), 
    smtpsdk.StringPtr("recipient@example.com"),
    smtpsdk.StringPtr("Welcome"),
    &eventType,
    &tags,
    &limit, &offset)

// Get logs for specific message
logEntry, err := client.Logs().GetMessage("example.com", "message-guid")
if len(logEntry.Data) > 0 {
    fmt.Printf("Event: %s at %d\n", 
        logEntry.Data[0].EventType, 
        logEntry.Data[0].Timestamp)
}

// Stream logs (Server-Sent Events)
resp, err := client.Logs().Stream("example.com", nil, nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

if resp.StatusCode == http.StatusOK {
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        event := scanner.Text()
        fmt.Printf("Event: %s\n", event)
    }
}
```

## Suppressions

```go
// List all suppressions
suppressions, err := client.Suppressions().List(
    "example.com", 
    nil, nil, nil, nil)

// Filter by type
unsubscribes, err := client.Suppressions().List(
    "example.com", 
    smtpsdk.StringPtr("unsubscribe"), nil, nil, nil)

bounces, err := client.Suppressions().List(
    "example.com", 
    smtpsdk.StringPtr("bounce"), nil, nil, nil)

// List by type-specific endpoints
unsubscribes, err := client.Suppressions().ListUnsubscribes(
    "example.com", nil, nil, nil)

bounces, err := client.Suppressions().ListBounces(
    "example.com", nil, nil, nil)

whitelists, err := client.Suppressions().ListWhitelists(
    "example.com", nil, nil, nil)

// Search suppressions
unsubscribes, err := client.Suppressions().ListUnsubscribes(
    "example.com", 
    smtpsdk.StringPtr("user@example.com"), 
    nil, nil)

// Pagination
page := 1
perPage := 20
unsubscribes, err := client.Suppressions().ListUnsubscribes(
    "example.com", nil, &page, &perPage)

// Create whitelist entry
result, err := client.Suppressions().CreateWhitelist("example.com", smtpsdk.WhitelistCreateRequest{
    Recipient:     "user@example.com",
    RecipientType: "email",
    Description:   "Trusted customer",
})

// Whitelist by domain
result, err := client.Suppressions().CreateWhitelist("example.com", smtpsdk.WhitelistCreateRequest{
    Recipient:     "trusted.com",
    RecipientType: "domain",
    Description:   "Trusted partner domain",
})

// Delete suppressions
result, err := client.Suppressions().DeleteUnsubscribes(
    "example.com", []int{1, 2, 3})

result, err := client.Suppressions().DeleteBounces(
    "example.com", []int{1, 2, 3})

result, err := client.Suppressions().DeleteWhitelists(
    "example.com", []int{1, 2, 3})
fmt.Printf("Deleted %d suppressions\n", result.Data.DeletedCount)
```

## Webhooks

```go
// List webhooks
webhooks, err := client.Webhooks().List("example.com", nil)

// Filter by event type
webhooks, err := client.Webhooks().List(
    "example.com", smtpsdk.StringPtr("delivered"))

// Available event types:
// "queued", "send", "delivered", "bounced", "failed", 
// "permanent_fail", "opened", "clicked", "unsubscribed", "temporary_fail", "deferred"

// Create webhook
webhook, err := client.Webhooks().Create("example.com", smtpsdk.WebhookCreateRequest{
    Type: "delivered",
    URL:  "https://example.com/webhook",
})
fmt.Printf("Webhook GUID: %s\n", webhook.Data.WebhookGUID)

// Get webhook details
webhook, err := client.Webhooks().Get("example.com", "webhook-guid")

// Update webhook
webhook, err := client.Webhooks().Update("example.com", "webhook-guid", smtpsdk.WebhookUpdateRequest{
    URL:  smtpsdk.StringPtr("https://example.com/new-webhook"),
    Type: smtpsdk.StringPtr("opened"),
})

// Delete webhook
err := client.Webhooks().Delete("example.com", "webhook-guid")

// Test webhook
result, err := client.Webhooks().Test("example.com", smtpsdk.WebhookTestRequest{
    URL:       "https://example.com/webhook",
    EventType: "delivered",
})
if result.Data.ResponseStatus == 200 {
    fmt.Printf("Test successful! Response time: %dms\n", result.Data.ResponseTime)
} else {
    fmt.Printf("Test failed with status: %d\n", result.Data.ResponseStatus)
}
```

## User

```go
// Get quota information
quota, err := client.User().GetQuota()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Current quota: %d\n", quota.Data.CurrentQuota)
fmt.Printf("Max quota: %d\n", quota.Data.MaxQuota)
fmt.Printf("Usage: %.1f%%\n", quota.Data.UsagePercentage)
fmt.Printf("Remaining: %d\n", quota.Data.CurrentQuota)

// Check if near limit
if quota.Data.UsagePercentage > 90.0 {
    log.Println("Warning: Quota nearly exhausted!")
}
```

## Client Configuration

### Custom HTTP Client

```go
import (
    "net/http"
    "time"
)

httpClient := &http.Client{
    Timeout: 30 * time.Second,
}

client := smtpsdk.NewClient("username", "token",
    smtpsdk.WithHTTPClient(httpClient),
)
```

### Custom Base URL

```go
client := smtpsdk.NewClient("username", "token",
    smtpsdk.WithBaseURL("https://custom.smtp.example.com"),
)
```

### Multiple Options

```go
client := smtpsdk.NewClient("username", "token",
    smtpsdk.WithHTTPClient(httpClient),
    smtpsdk.WithBaseURL("https://api.example.com"),
)
```

## Error Handling

All methods return errors when API requests fail:

```go
domains, err := client.Domains().List(nil, nil, nil)
if err != nil {
    log.Fatalf("Failed to fetch domains: %v", err)
}

// API errors include status and message
if err != nil {
    // Error contains: "message (status: XXX)"
    log.Fatal(err)
}
```

### Common Error Responses

- **401 Unauthorized**: Invalid username or token
- **404 Not Found**: Domain, credential, or webhook not found
- **422 Validation Error**: Invalid request parameters
- **429 Rate Limit**: Too many requests
- **500 Server Error**: Internal server error

## Advanced Usage

### Working with Pointers

The SDK uses pointer types for optional parameters:

```go
// Optional parameter
page := 1
domains, err := client.Domains().List(&page, nil, nil)

// Omit parameter
domains, err := client.Domains().List(nil, nil, nil)

// Helper functions
limit := smtpsdk.IntPtr(100)
enabled := smtpsdk.BoolPtr(true)
email := smtpsdk.StringPtr("test@example.com")
```

### Handling Recipients

The `To` field accepts different types:

```go
// Single string
req := smtpsdk.MessageSendRequest{
    To: "recipient@example.com",
}

// Array of strings
req := smtpsdk.MessageSendRequest{
    To: []string{"user1@example.com", "user2@example.com"},
}

// Array of interfaces (for template bulk)
req := smtpsdk.MessageTemplateRequest{
    To: []interface{}{"user1@example.com", "user2@example.com"},
}
```

### Custom Headers

```go
headers := smtpsdk.Headers{
    "X-Campaign-ID":   "welcome-series",
    "X-Priority":      "1",
    "X-Customer-Tier": "premium",
}
```

## Testing

Run tests:

```bash
go test ./tests/...
```

Run tests with coverage:

```bash
go test -cover ./tests/...
```

Run tests with verbose output:

```bash
go test -v ./tests/...
```

Run specific test:

```bash
go test -v ./tests/... -run TestDomainsService_List
```

## Examples

See the `example/` directory for a complete working example:

```bash
cd example
go run main.go
```

## License

See LICENSE file for details.

## Support

For API documentation, visit: https://smtp-app.kirim.email

For issues, please open a GitHub issue.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Version History

- **v1.3.0** - Added event_type and tags filter support for log retrieval
- **v1.0.0** - Initial release with full API support
