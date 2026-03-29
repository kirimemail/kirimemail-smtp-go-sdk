# SKILLS.md - Kirim.Email SMTP Go SDK Usage Guide

## Overview

This is the official Go SDK for Kirim.Email SMTP API. It provides a clean, service-oriented interface for managing domains, SMTP credentials, sending emails, validation, logs, suppressions, webhooks, and user quotas.

**Module:** `github.com/kirimemail/kirimemail-smtp-go-sdk`

---

## Installation

```bash
go get github.com/kirimemail/kirimemail-smtp-go-sdk
```

---

## Quick Start

```go
import smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"

client := smtpsdk.NewClient("username", "api-token")

domains, err := client.Domains().List(nil, nil, nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found %d domains\n", len(domains.Data))
```

---

## Client Configuration

```go
client := smtpsdk.NewClient("username", "api-token",
    smtpsdk.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
    smtpsdk.WithBaseURL("https://custom.smtp.example.com"),
)
```

---

## Domain Management

```go
domains := client.Domains()

list, err := domains.List(nil, nil, nil)
create, err := domains.Create(smtpsdk.DomainCreateRequest{Domain: "example.com"})
get, err := domains.Get("example.com")
err := domains.VerifyMandatoryRecords("example.com")
err := domains.SetupAuthDomain("example.com", smtpsdk.AuthDomainSetupRequest{...})
err := domains.SetupTracklink("example.com", smtpsdk.TracklinkSetupRequest{...})
```

---

## SMTP Credentials

```go
creds := client.Credentials()

list, err := creds.List("example.com", nil, nil)
create, err := creds.Create("example.com", smtpsdk.CredentialCreateRequest{
    Email: "sender@example.com",
})
reset, err := creds.ResetPassword("example.com", "credential-guid")
err := creds.Delete("example.com", "credential-guid")
```

---

## Sending Emails

**Simple email:**
```go
resp, err := client.Messages().Send("example.com", smtpsdk.MessageSendRequest{
    From:    "noreply@example.com",
    To:      []string{"recipient@example.com"},
    Subject: "Hello",
    Text:    "Message body",
}, nil)
```

**With attachments:**
```go
fileData, _ := os.ReadFile("document.pdf")
resp, err := client.Messages().Send("example.com", smtpsdk.MessageSendRequest{
    From:    "noreply@example.com",
    To:      []string{"recipient@example.com"},
    Subject: "Document",
    Text:    "Please find attached",
}, map[string][]byte{"document.pdf": fileData})
```

**Using templates:**
```go
resp, err := client.Messages().SendTemplate("example.com", smtpsdk.MessageTemplateRequest{
    From:      "noreply@example.com",
    To:        []string{"recipient@example.com"},
    Template:  "welcome-template",
    Variables: map[string]string{"name": "John"},
}, nil)
```

---

## Email Validation

```go
val := client.Validation()

result, err := val.ValidateEmail("user@example.com")
result, err := val.ValidateEmailStrict("user@example.com")
batch, err := val.ValidateEmailsBatch([]string{"a@example.com", "b@example.com"})
batch, err := val.ValidateEmailsBatchStrict([]string{"a@example.com", "b@example.com"})
```

---

## Email Logs

```go
logs := client.Logs()

list, err := logs.Get("example.com", nil, nil, nil, nil, nil, nil)
msg, err := logs.GetMessage("example.com", "message-guid")

streamResp, err := logs.Stream("example.com", nil, nil, nil, nil, nil)
// streamResp is an http.Response with SSE stream
```

---

## Suppressions

```go
supp := client.Suppressions()

list, err := supp.List("example.com", nil, nil, nil, nil)
unsubs, err := supp.ListUnsubscribes("example.com", nil, nil, nil)
bounces, err := supp.ListBounces("example.com", nil, nil, nil)
whitelists, err := supp.ListWhitelists("example.com", nil, nil, nil)

err = supp.CreateWhitelist("example.com", smtpsdk.WhitelistCreateRequest{
    Email: "user@example.com",
})
err = supp.DeleteBounces("example.com", []int{1, 2, 3})
```

---

## Webhooks

```go
hooks := client.Webhooks()

list, err := hooks.List("example.com", nil)
create, err := hooks.Create("example.com", smtpsdk.WebhookCreateRequest{
    URL:    "https://example.com/webhook",
    Events: []string{"bounce", "unsubscribe"},
})
test, err := hooks.Test("example.com", smtpsdk.WebhookTestRequest{
    URL: "https://example.com/webhook",
})
err = hooks.Delete("example.com", "webhook-guid")
```

---

## User Quota

```go
quota, err := client.User().GetQuota()
fmt.Printf("Used: %d, Limit: %d\n", quota.Usage, quota.Limit)
```

---

## Helper Functions

```go
strPtr := smtpsdk.StringPtr("value")   // *string
boolPtr := smtpsdk.BoolPtr(true)       // *bool
intPtr := smtpsdk.IntPtr(10)           // *int
```

---

## Testing

```bash
go test ./tests/...
go test -cover ./tests/...
go test -v ./tests/...
```

---

## Error Handling

All methods return `error` as the last return value. Always check for errors:

```go
result, err := client.Domains().List(nil, nil, nil)
if err != nil {
    log.Fatalf("API Error: %v", err)
}
```
