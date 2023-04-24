package acrolinx

type AppliedCheckOptions struct {
	GuidanceProfileID   string               `json:"guidanceProfileId"`
	GuidanceProfileName string               `json:"guidanceProfileName"`
	LanguageID          string               `json:"languageId"`
	ReportTypes         []string             `json:"reportTypes"`
	ContentFormat       string               `json:"contentFormat"`
	CheckType           string               `json:"checkType"`
	TermSets            []*TermSet           `json:"termSets"`
	PartialCheckRanges  []*PartialCheckRange `json:"partialCheckRanges"`
	Confidential        bool                 `json:"confidential"`
}

type CancelledCheck struct {
	ID string `json:"id"`
}

type Capabilities struct {
	DefaultGuidanceProfileID string             `json:"defaultGuidanceProfileId"`
	GuidanceProfiles         []*GuidanceProfile `json:"guidanceProfiles"`
	ContentFormats           []*ContentFormat   `json:"contentFormats"`
	ContentEncodings         []string           `json:"contentEncodings"`
	ReferencePattern         string             `json:"referencePattern"`
	CheckTypes               []string           `json:"checkTypes"`
	ReportTypes              []string           `json:"reportTypes"`
}

type Check struct {
	ID string `json:"id"`
}

type CheckOptions struct {
	GuidanceProfileID  string               `json:"guidanceProfileId"`
	ContentFormat      string               `json:"contentFormat"`
	ReportTypes        []string             `json:"reportTypes,omitempty"`
	CheckType          string               `json:"checkType"`
	PartialCheckRanges []*PartialCheckRange `json:"partialCheckRanges,omitempty"`
	BatchID            string               `json:"batchId"`
}

type CheckResult struct {
	ID                string               `json:"id"`
	CheckOptions      *AppliedCheckOptions `json:"checkOptions"`
	Document          *ResponseDocument    `json:"document"`
	Quality           *Quality             `json:"quality"`
	Counts            *Counts              `json:"counts"`
	Goals             []*Goal              `json:"goals"`
	Issues            []*Issue             `json:"issues"`
	Keywords          *Keywords            `json:"keywords"`
	Embed             []*EmbedItem         `json:"embed"`
	Reports           map[string]*Report   `json:"reports"`
	RuntimeStatistics *RuntimeStatistics   `json:"runtimeStatistics"`
	DictionaryScopes  []string             `json:"dictionaryScopes"`
	Progress          *Progress            `json:"progress"`
}

type ContentFormat struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type Counts struct {
	Sentences    int `json:"sentences"`
	Words        int `json:"words"`
	Issues       int `json:"issues"`
	ScoredIssues int `json:"scoredIssues"`
}

type CustomField struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Debug struct {
	Penalty float64 `json:"penalty"`
}

type DisplayInfo struct {
	Reference string `json:"reference"`
}

type Document struct {
	Reference    string         `json:"reference"`
	CustomFields []*CustomField `json:"customFields"`
}

type EmbedItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetCapabilitiesOptions struct {
	Locale string
}

type Goal struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Color       string `json:"color"`
	Scoring     string `json:"scoring"`
	Issues      int    `json:"issues,omitempty"`
}

type GuidanceProfile struct {
	ID          string     `json:"id"`
	DisplayName string     `json:"displayName"`
	Language    *Language  `json:"language"`
	Goals       []*Goal    `json:"goals"`
	TermSets    []*TermSet `json:"termSets"`
}

type Hashes struct {
	Issue       string `json:"issue"`
	Environment string `json:"environment"`
	Index       string `json:"index"`
}

type Issue struct {
	GoalID                string                 `json:"goalId"`
	InternalName          string                 `json:"internalName"`
	DisplayNameHTML       string                 `json:"displayNameHtml"`
	GuidanceHTML          string                 `json:"guidanceHtml"`
	DisplaySurface        string                 `json:"displaySurface"`
	IssueType             string                 `json:"issueType"`
	Scoring               string                 `json:"scoring"`
	PositionalInformation *PositionalInformation `json:"positionalInformation"`
	ReadOnly              bool                   `json:"readOnly"`
	IssueLocations        []*Location            `json:"issueLocations"`
	Suggestions           []*Suggestion          `json:"suggestions"`
	SubIssues             []*Issue               `json:"subIssues"`
	Debug                 *Debug                 `json:"debug"`
	CanAddToDictionary    bool                   `json:"canAddToDictionary"`
	Links                 Links                  `json:"links"`
}

type Keyword struct {
	Keyword     string                   `json:"keyword"`
	SortKey     string                   `json:"sortKey"`
	Density     float64                  `json:"density"`
	Count       int                      `json:"count"`
	Prominence  float64                  `json:"prominence"`
	Occurrences []*PositionalInformation `json:"occurrences"`
	Warnings    []*KeywordWarning        `json:"warnings"`
}

type Keywords struct {
	Discovered []*Keyword `json:"discovered"`
	Proposed   []*Keyword `json:"proposed"`
	Target     []*Keyword `json:"target"`
	Links      Links      `json:"links"`
}

type KeywordWarning struct {
	Type     string `json:"type"`
	Severity int    `json:"severity"`
}

type Language struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type Location struct {
	LocationID  string            `json:"locationId"`
	DisplayName string            `json:"displayName"`
	Values      map[string]string `json:"values"`
}

type Match struct {
	ExtractedPart  string `json:"extractedPart"`
	ExtractedBegin int    `json:"extractedBegin"`
	ExtractedEnd   int    `json:"extractedEnd"`
	OriginalPart   string `json:"originalPart"`
	OriginalBegin  int    `json:"originalBegin"`
	OriginalEnd    int    `json:"originalEnd"`
}

type PartialCheckRange struct {
	Begin int `json:"begin"`
	End   int `json:"end"`
}

type PositionalInformation struct {
	Hashes  *Hashes  `json:"hashes"`
	Matches []*Match `json:"matches"`
}

type Quality struct {
	Score            int
	Status           string
	ScoresByStrategy []*Score
	ScoresByGoal     []*Score
	Metrics          []*Score
}

type Report struct {
	DisplayName       string `json:"displayName"`
	Link              string `json:"link"`
	LinkAuthenticated string `json:"linkAuthenticated"`
}

type ResponseDocument struct {
	ID                   string         `json:"id"`
	DisplayInfo          *DisplayInfo   `json:"displayInfo"`
	CustomFields         []*CustomField `json:"customFields"`
	CustomFieldsComplete bool           `json:"customFieldsComplete"`
}

type RuntimeStatistics struct {
	StartedAt string `json:"startedAt"`
}

type Score struct {
	ID    string `json:"id"`
	Score int    `json:"score"`
}

type SubmitCheckOptions struct {
	Content      string        `json:"content"`
	CheckOptions *CheckOptions `json:"checkOptions"`
	Document     *Document     `json:"document"`
	Language     string        `json:"language"`
}

type Suggestion struct {
	Surface      string   `json:"surface"`
	GroupID      string   `json:"groupId"`
	Replacements []string `json:"replacements"`
	IconID       string   `json:"string"`
}

type TermSet struct {
	DisplayName string `json:"displayName"`
}
