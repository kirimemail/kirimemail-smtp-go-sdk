package smtpsdk

import (
	"fmt"
	"net/http"
	"strconv"
)

type LogsService struct {
	client *Client
}

func (c *Client) Logs() *LogsService {
	return &LogsService{client: c}
}

func (s *LogsService) Get(domain string, start, end, sender, recipient, subject, eventType, tags *string, limit, offset *int) (*LogListResponse, error) {
	query := make(map[string]string)

	if start != nil {
		query["start"] = *start
	}
	if end != nil {
		query["end"] = *end
	}
	if sender != nil {
		query["sender"] = *sender
	}
	if recipient != nil {
		query["recipient"] = *recipient
	}
	if subject != nil {
		query["subject"] = *subject
	}
	if eventType != nil {
		if !IsValidLogEventType(*eventType) {
			return nil, fmt.Errorf("invalid event_type: %s", *eventType)
		}
		query["event_type"] = *eventType
	}
	if tags != nil {
		query["tags"] = *tags
	}
	if limit != nil {
		query["limit"] = strconv.Itoa(*limit)
	}
	if offset != nil {
		query["offset"] = strconv.Itoa(*offset)
	}

	resp, err := s.client.doGet(fmt.Sprintf("/api/domains/%s/log", domain), query)
	if err != nil {
		return nil, err
	}

	var result LogListResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *LogsService) GetByEventType(domain, eventType string, start, end, sender, recipient, subject, tags *string, limit, offset *int) (*LogListResponse, error) {
	return s.Get(domain, start, end, sender, recipient, subject, &eventType, tags, limit, offset)
}

func (s *LogsService) GetByTags(domain, tags string, start, end, sender, recipient, subject, eventType *string, limit, offset *int) (*LogListResponse, error) {
	return s.Get(domain, start, end, sender, recipient, subject, eventType, &tags, limit, offset)
}

func (s *LogsService) GetMessage(domain, messageGUID string) (*LogMessageResponse, error) {
	resp, err := s.client.doGet(fmt.Sprintf("/api/domains/%s/log/%s", domain, messageGUID), nil)
	if err != nil {
		return nil, err
	}

	var result LogMessageResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *LogsService) Stream(domain string, start, end, sender, recipient, subject, eventType, tags *string, limit *int) (*http.Response, error) {
	query := make(map[string]string)

	if start != nil {
		query["start"] = *start
	}
	if end != nil {
		query["end"] = *end
	}
	if sender != nil {
		query["sender"] = *sender
	}
	if recipient != nil {
		query["recipient"] = *recipient
	}
	if subject != nil {
		query["subject"] = *subject
	}
	if eventType != nil {
		if !IsValidLogEventType(*eventType) {
			return nil, fmt.Errorf("invalid event_type: %s", *eventType)
		}
		query["event_type"] = *eventType
	}
	if tags != nil {
		query["tags"] = *tags
	}
	if limit != nil {
		query["limit"] = strconv.Itoa(*limit)
	}

	return s.client.doGet(fmt.Sprintf("/api/domains/%s/log/stream", domain), query)
}
