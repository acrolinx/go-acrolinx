package acrolinx

import (
	"fmt"
	"net/http"
)

type CheckingService struct {
	client *Client
}

type Capabilities struct {
	DefaultGuidanceProfileID string             `json:"defaultGuidanceProfileID"`
	GuidanceProfiles         []*GuidanceProfile `json:"guidanceProfiles"`
	ContentFormats           []*ContentFormat   `json:"contentFormats"`
	ContentEncodings         []string           `json:"contentEncodings"`
	ReferencePattern         string             `json:"referencePattern"`
	CheckTypes               []string           `json:"checkTypes"`
	ReportTypes              []string           `json:"reportTypes"`
}

type GuidanceProfile struct {
	ID          string     `json:"id"`
	DisplayName string     `json:"displayName"`
	Language    *Language  `json:"language"`
	Goals       []*Goal    `json:"goals"`
	TermSets    []*TermSet `json:"termSets"`
}

type Language struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type Goal struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Color       string `json:"color"`
	Scoring     string `json:"scoring"`
}

type TermSet struct {
	DisplayName string `json:"displayName"`
}

type ContentFormat struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type GetCapabilitiesOptions struct {
	Locale string
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

type SubmitCheckOptions struct {
	Content      string        `json:"content"`
	CheckOptions *CheckOptions `json:"checkOptions"`
	Document     *Document     `json:"document"`
	Language     string        `json:"language"`
}

type CheckOptions struct {
	GuidanceProfileID  string               `json:"guidanceProfileId"`
	ContentFormat      string               `json:"contentFormat"`
	ReportTypes        []string             `json:"reportTypes"`
	CheckType          string               `json:"checkType"`
	PartialCheckRanges []*PartialCheckRange `json:"partialCheckRanges"`
	BatchID            string               `json:"batchId"`
}

type Document struct {
	Reference    string         `json:"reference"`
	CustomFields []*CustomField `json:"customFields"`
}

type PartialCheckRange struct {
	Begin int `json:"begin"`
	End   int `json:"end"`
}

type CustomField struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Check struct {
	ID string `json:"id"`
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
