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
