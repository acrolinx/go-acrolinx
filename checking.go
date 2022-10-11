package acrolinx

import (
	"fmt"
	"net/http"
)

type CheckingService struct {
	client *Client
}

func (s *CheckingService) GetCapabilities(opts *GetCapabilitiesOptions) (*Capabilities, Links, error) {
	req, err := s.client.newRequest(http.MethodGet, "api/v1/checking/capabilities", nil)
	if err != nil {
		return nil, nil, err
	}

	if opts != nil && opts.Locale != "" {
		req.Header.Set(headerLocale, opts.Locale)
	}

	var caps Capabilities
	var reqError RequestError
	links := make(Links)
	resp := Response{
		Data:  &caps,
		Links: links,
		Error: &reqError,
	}
	err = s.client.do(req, &resp)
	if err != nil {
		return nil, nil, err
	}

	if reqError != (RequestError{}) {
		return nil, links, &reqError
	}

	return &caps, links, nil
}

func (s *CheckingService) SubmitCheck(opts *SubmitCheckOptions) (*Check, Links, error) {
	req, err := s.client.newRequest(http.MethodPost, "api/v1/checking/checks", opts)
	if err != nil {
		return nil, nil, fmt.Errorf("Error preparing check request: %w", err)
	}

	var check Check
	links := make(Links)
	var reqError RequestError
	resp := Response{Data: &check, Links: links, Error: &reqError}
	err = s.client.do(req, &resp)
	if err != nil {
		return nil, nil, fmt.Errorf("Error processing check request: %w", err)
	}

	if reqError != (RequestError{}) {
		return nil, nil, &reqError
	}

	return &check, links, nil
}

func (s *CheckingService) GetCheckResult(check *Check) (*CheckResult, Links, error) {
	path := fmt.Sprintf("api/v1/checking/checks/%s", check.ID)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("Error preparing check request: %w", err)
	}

	var result CheckResult
	var progress Progress
	links := make(Links)
	var reqError RequestError
	resp := Response{Data: &result, Links: links, Progress: &progress, Error: &reqError}
	err = s.client.do(req, &resp)
	if err != nil {
		return nil, nil, fmt.Errorf("Error processing check request: %w", err)
	}

	if reqError != (RequestError{}) {
		return nil, nil, &reqError
	}

	if progress != (Progress{}) {
		result.Progress = &progress
	}

	return &result, links, nil
}
