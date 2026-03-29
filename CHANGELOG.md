# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.3.0] - 2026-03-29

### Added
- `event_type` filter parameter for log retrieval (`Get()`, `Stream()`)
- `tags` filter parameter for log retrieval with partial match support
- `subject` filter parameter for log retrieval with partial match
- `GetByEventType()` helper method to filter logs by event type
- `GetByTags()` helper method to filter logs by tags
- Event type constants: `LogEventTypeQueued`, `LogEventTypeSend`, `LogEventTypeDelivered`, `LogEventTypeBounced`, `LogEventTypeFailed`, `LogEventTypeOpened`, `LogEventTypeClicked`, `LogEventTypeUnsubscribed`, `LogEventTypeTemporaryFail`, `LogEventTypePermanentFail`, `LogEventTypeDeferred`
- `IsValidLogEventType()` validation function

### Changed
- `Get()` method signature updated to include new filter parameters
- `Stream()` method signature updated to include new filter parameters

### Fixed
- None

### Deprecated
- None

## [1.0.0] - 2024-01-01

### Added
- Initial release with full API support
- Domain management (list, create, update, delete)
- Credential management (list, create, get, delete, reset password)
- Email sending (send, send with attachments, send with templates)
- Email validation (single, batch, strict mode)
- Logs retrieval and streaming
- Suppressions management (unsubscribe, bounce, whitelist)
- Webhooks management (create, list, get, update, delete, test)
- User quota information
