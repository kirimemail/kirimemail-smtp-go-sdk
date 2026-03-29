# AGENTS.md - SDK Update Guide

This document provides guidance for agents updating the Kirim.Email SMTP Go SDK.

## SDK Structure

```
golang/
├── client.go          # Core client, HTTP handling, auth
├── types.go          # All shared types (requests, responses, models)
├── domains.go        # Domain management service
├── credentials.go    # SMTP credentials service
├── messages.go       # Email sending service
├── validation.go     # Email validation service
├── logs.go           # Email logs service
├── supprressions.go  # Suppression management service
├── webhooks.go       # Webhook management service
├── user.go           # User quota service
├── tests/            # Unit tests
└── example/          # Usage examples
```

## Adding a New Service

1. Create a new file `newservice.go` following the service pattern:

```go
package smtpsdk

type NewService struct {
    client *Client
}

func (c *Client) NewService() *NewService {
    return &NewService{client: c}
}
```

2. Add accessor method to `client.go`:

```go
func (c *Client) NewService() *NewService {
    return &NewService{client: c}
}
```

## Adding New Types

Add request/response types to `types.go`:

```go
type NewTypeRequest struct {
    Field1 string `json:"field1"`
    Field2 int    `json:"field2"`
}

type NewTypeResponse struct {
    Field1 string `json:"field1"`
    Data   []Item `json:"data"`
}
```

## Adding New Methods

1. Add method to appropriate service file:

```go
func (s *DomainsService) NewMethod(req DomainRequest) (*DomainResponse, error) {
    path := "/v1/domains/new-method"
    reqBody, _ := json.Marshal(req)
    
    httpResp, err := s.client.request("POST", path, reqBody, nil)
    if err != nil {
        return nil, err
    }
    
    var resp DomainResponse
    if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
        return nil, err
    }
    return &resp, nil
}
```

2. Return `(*Response, error)` for methods that only return status, or specific response types for methods with data.

## Request Helper Pattern

Use the client's built-in request method:

```go
httpResp, err := s.client.request(method, path, body, queryParams)
```

Query params should be built from pointer values (nil = omit):

```go
func buildQuery(params ...*string) map[string]string {
    q := make(map[string]string)
    for _, p := range params {
        if p != nil {
            q["key"] = *p
        }
    }
    return q
}
```

## Response Patterns

**List response:**
```go
type DomainListResponse struct {
    Data       []Domain `json:"data"`
    Pagination `json:"pagination"`
}
```

**Single item response:**
```go
type DomainResponse struct {
    Data Domain `json:"data"`
}
```

**Simple API response:**
```go
type APIResponse struct {
    Message string `json:"message"`
    Code    int    `json:"code"`
}
```

## Required Updates When Adding API Endpoints

1. **types.go** - Add request/response structs
2. **client.go** - Add service accessor method
3. **newservice.go** - Add service struct and methods
4. **example/main.go** - Add usage example
5. **tests/** - Add unit tests
6. **SKILLS.md** - Update usage documentation

## Code Style

- Use pointer types (`*string`, `*int`) for optional parameters
- Provide helper functions `StringPtr`, `BoolPtr`, `IntPtr`
- Return `(*Response, error)` for operations without data
- Use meaningful struct tags: `json:"field_name"`
- Follow existing naming conventions (PascalCase for exports)

## Testing

Add tests in `tests/` directory following the existing test pattern:

```go
func TestDomainsList(t *testing.T) {
    // Setup
    client := NewClient("user", "token")
    
    // Test
    resp, err := client.Domains().List(nil, nil, nil)
    
    // Assert
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    // ... more assertions
}
```

## Running Tests

```bash
go test ./tests/...          # All tests
go test -cover ./tests/...   # With coverage
go test -v ./tests/...       # Verbose
```

## Lint and Typecheck

Run before committing changes:

```bash
go fmt ./...
go vet ./...
go build ./...
```
