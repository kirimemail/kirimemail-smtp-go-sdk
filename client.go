package smtpsdk

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Username   string
	Token      string
}

type ClientOption func(*Client)

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) { c.HTTPClient = client }
}

func WithBaseURL(url string) ClientOption {
	return func(c *Client) { c.BaseURL = url }
}

func NewClient(username, token string, opts ...ClientOption) *Client {
	c := &Client{
		BaseURL:    "https://smtp-app.kirim.email",
		Username:   username,
		Token:      token,
		HTTPClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) buildURL(path string) string {
	return c.BaseURL + path
}

func (c *Client) getAuth() string {
	auth := c.Username + ":" + c.Token
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c *Client) doRequest(method, path string, body io.Reader, contentType string) (*http.Response, error) {
	req, err := http.NewRequest(method, c.buildURL(path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.getAuth())
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	return c.HTTPClient.Do(req)
}

func (c *Client) doGet(path string, query map[string]string) (*http.Response, error) {
	if len(query) > 0 {
		values := url.Values{}
		for k, v := range query {
			values.Add(k, v)
		}
		path += "?" + values.Encode()
	}
	return c.doRequest("GET", path, nil, "")
}

func (c *Client) doPost(path string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return c.doRequest("POST", path, bytes.NewReader(jsonBody), "application/json")
}

func (c *Client) doPut(path string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return c.doRequest("PUT", path, bytes.NewReader(jsonBody), "application/json")
}

func (c *Client) doDelete(path string) (*http.Response, error) {
	return c.doRequest("DELETE", path, nil, "")
}

func (c *Client) doMultipartRequest(path string, fields map[string]interface{}, files map[string][]byte) (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, value := range fields {
		switch v := value.(type) {
		case string:
			_ = writer.WriteField(key, v)
		case int, int64, int32, float64, bool:
			_ = writer.WriteField(key, fmt.Sprintf("%v", v))
		case Headers:
			jsonData, err := json.Marshal(v)
			if err == nil {
				_ = writer.WriteField(key, string(jsonData))
			}
		case []string:
			for i, item := range v {
				_ = writer.WriteField(fmt.Sprintf("%s[%d]", key, i), item)
			}
		case []interface{}:
			jsonData, err := json.Marshal(v)
			if err == nil {
				_ = writer.WriteField(key, string(jsonData))
			}
		default:
			jsonData, err := json.Marshal(v)
			if err == nil {
				_ = writer.WriteField(key, string(jsonData))
			}
		}
	}

	for filename, data := range files {
		part, err := writer.CreateFormFile("attachments", filename)
		if err != nil {
			return nil, err
		}
		if _, err := part.Write(data); err != nil {
			return nil, err
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.buildURL(path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.getAuth())
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return c.HTTPClient.Do(req)
}

func (c *Client) decodeResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		var errResp struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
			Error   string `json:"error"`
		}
		if err := json.Unmarshal(body, &errResp); err == nil {
			if errResp.Message != "" {
				return fmt.Errorf("%s (status: %d)", errResp.Message, resp.StatusCode)
			}
			if errResp.Error != "" {
				return fmt.Errorf("%s (status: %d)", errResp.Error, resp.StatusCode)
			}
		}
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return json.Unmarshal(body, target)
}

func buildQuery(query map[string]string) string {
	if len(query) == 0 {
		return ""
	}
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}
	return values.Encode()
}

func parseIntQueryParam(value *int) string {
	if value == nil {
		return ""
	}
	return strconv.Itoa(*value)
}

func parseStringQueryParam(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func parseBoolQueryParam(value *bool) string {
	if value == nil {
		return ""
	}
	return strconv.FormatBool(*value)
}

func StringPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}

func IntPtr(i int) *int {
	return &i
}

func splitRecipients(recipients interface{}) ([]string, error) {
	switch r := recipients.(type) {
	case string:
		return []string{r}, nil
	case []string:
		return r, nil
	case []interface{}:
		result := make([]string, len(r))
		for i, v := range r {
			if s, ok := v.(string); ok {
				result[i] = s
			} else {
				return nil, fmt.Errorf("invalid recipient type at index %d", i)
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("invalid recipients type: %T", recipients)
	}
}
