package smtpsdk

type UserService struct {
	client *Client
}

func (c *Client) User() *UserService {
	return &UserService{client: c}
}

func (s *UserService) GetQuota() (*QuotaResponse, error) {
	resp, err := s.client.doGet("/api/quota", nil)
	if err != nil {
		return nil, err
	}

	var result QuotaResponse
	if err := s.client.decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
