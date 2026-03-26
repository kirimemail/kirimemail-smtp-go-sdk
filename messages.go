package smtpsdk

import (
	"fmt"
)

type MessagesService struct {
	client *Client
}

func (c *Client) Messages() *MessagesService {
	return &MessagesService{client: c}
}

func (s *MessagesService) Send(domain string, req MessageSendRequest, attachments map[string][]byte) (*MessageSendResponse, error) {
	fields := map[string]interface{}{
		"from":    req.From,
		"subject": req.Subject,
		"text":    req.Text,
	}

	if req.FromName != "" {
		fields["from_name"] = req.FromName
	}
	if req.To != nil {
		fields["to"] = req.To
	}
	if req.HTML != "" {
		fields["html"] = req.HTML
	}
	if req.Headers != nil {
		fields["headers"] = req.Headers
	}
	if req.ReplyTo != "" {
		fields["reply_to"] = req.ReplyTo
	}
	if req.AttachmentOptions != "" {
		fields["attachment_options"] = req.AttachmentOptions
	}

	resp, err := s.client.doMultipartRequest(fmt.Sprintf("/api/domains/%s/message", domain), fields, attachments)
	if err != nil {
		return nil, err
	}

	var result MessageSendResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *MessagesService) SendTemplate(domain string, req MessageTemplateRequest, attachments map[string][]byte) (*MessageTemplateResponse, error) {
	fields := map[string]interface{}{
		"template_guid": req.TemplateGUID,
		"to":            req.To,
	}

	if req.From != "" {
		fields["from"] = req.From
	}
	if req.FromName != "" {
		fields["from_name"] = req.FromName
	}
	if req.Variables != nil {
		fields["variables"] = req.Variables
	}
	if req.Headers != nil {
		fields["headers"] = req.Headers
	}
	if req.ReplyTo != "" {
		fields["reply_to"] = req.ReplyTo
	}
	if req.AttachmentOptions != "" {
		fields["attachment_options"] = req.AttachmentOptions
	}

	resp, err := s.client.doMultipartRequest(fmt.Sprintf("/api/domains/%s/message/template", domain), fields, attachments)
	if err != nil {
		return nil, err
	}

	var result MessageTemplateResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
