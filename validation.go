package smtpsdk

type ValidationService struct {
	client *Client
}

func (c *Client) Validation() *ValidationService {
	return &ValidationService{client: c}
}

func (s *ValidationService) ValidateEmail(email string) (*EmailValidationResponse, error) {
	req := EmailValidationRequest{Email: email}

	resp, err := s.client.doPost("/api/email/validate", req)
	if err != nil {
		return nil, err
	}

	var result EmailValidationResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *ValidationService) ValidateEmailStrict(email string) (*EmailValidationResponse, error) {
	req := EmailValidationRequest{Email: email}

	resp, err := s.client.doPost("/api/email/validate/strict", req)
	if err != nil {
		return nil, err
	}

	var result EmailValidationResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *ValidationService) ValidateEmailsBatch(emails []string) (*EmailValidationBatchResponse, error) {
	req := EmailValidationBatchRequest{Emails: emails}

	resp, err := s.client.doPost("/api/email/validate/bulk", req)
	if err != nil {
		return nil, err
	}

	var result EmailValidationBatchResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *ValidationService) ValidateEmailsBatchStrict(emails []string) (*EmailValidationBatchResponse, error) {
	req := EmailValidationBatchRequest{Emails: emails}

	resp, err := s.client.doPost("/api/email/validate/bulk/strict", req)
	if err != nil {
		return nil, err
	}

	var result EmailValidationBatchResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
