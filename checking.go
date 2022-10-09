package acrolinx

type CheckingService struct {
	client *Client
}

type Capabilities struct {
	GuidanceProfiles []*GuidanceProfile `json:"guidanceProfiles"`
	ContentFormats   []*ContentFormat   `json:"contentFormats"`
	ContentEncodings []string           `json:"contentEncodings"`
	ReferencePattern string             `json:"referencePattern"`
	CheckTypes       []string           `json:"checkTypes"`
	ReportTypes      []string           `json:"reportTypes"`
}

type GuidanceProfile struct {
	Id          string     `json:"id"`
	DisplayName string     `json:"displayName"`
	Language    *Language  `json:"language"`
	Goals       []*Goal    `json:"goals"`
	TermSets    []*TermSet `json:"termSets"`
}

type Language struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type Goal struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Color       string `json:"color"`
	Scoring     string `json:"scoring"`
}

type TermSet struct {
	DisplayName string `json:"displayName"`
}

type ContentFormat struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type ListCapabilitiesOptions struct {
	Locale string
}

func (s *CheckingService) ListCapabilities(opts *ListCapabilitiesOptions) (*Capabilities, Links, error) {
	req, err := s.client.newRequest("api/v1/checking/capabilities", nil)
	if err != nil {
		return nil, nil, err
	}

	if opts != nil && opts.Locale != "" {
		req.Header.Set(headerLocale, opts.Locale)
	}

	var caps Capabilities
	links := make(Links)
	resp := Response{
		Data:  caps,
		Links: links,
	}
	err = s.client.do(req, &resp)
	if err != nil {
		return nil, nil, err
	}

	return &caps, links, nil
}
