package acrolinx

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCapabilities(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/checking/capabilities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "get_capabilities.json")
	})

	opts := &GetCapabilitiesOptions{}
	caps, links, err := client.Checking.GetCapabilities(opts)
	if err != nil {
		t.Fatalf("Checking.ListCapabilities returned error: %v", err)
	}

	expectedCaps := &Capabilities{
		GuidanceProfiles: []*GuidanceProfile{
			{
				ID:          "710e1361-90b7-3867-a42e-35279b7f8aa2",
				DisplayName: "de - Blog Posts",
				Language:    &Language{"de", "German"},
				Goals: []*Goal{
					{
						ID:          "CLARITY",
						DisplayName: "Verständlichkeit",
						Color:       "#ec407a",
						Scoring:     "required",
					},
					{
						ID:          "CONSISTENCY",
						DisplayName: "Einheitlichkeit",
						Color:       "#ffd600",
						Scoring:     "required",
					},
				},
				TermSets: []*TermSet{},
			},
		},
		ContentFormats: []*ContentFormat{
			{"TEXT", "Plain Text"},
			{"XML", "XML"},
			{"MS_OFFICE", "Microsoft Office Container"},
		},
		ContentEncodings: []string{"none", "base64"},
		CheckTypes:       []string{"interactive", "batch", "baseline", "automated"},
		ReferencePattern: "\\.(xml|XML|xhtm|XHTM|xhtml|XHTML)$|\\.(svg|SVG|resx|RESX)$",
		ReportTypes:      []string{"scorecard", "contentAnalysisDashboard"},
	}

	expectedLinks := map[string]string{
		"submitCheck":          "https://example.com/api/v1/checking/checks",
		"checkingCapabilities": "https://example.com/api/v1/checking/capabilities",
	}

	assert.Equal(t, expectedCaps, caps)
	assert.Equal(t, expectedLinks, links)
}

func TestGetCapabilitiesWithError(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/checking/capabilities", func(w http.ResponseWriter, r *http.Request) {
		mustWriteHTTPResponse(t, w, "error.json")
	})

	_, _, err := client.Checking.GetCapabilities(&GetCapabilitiesOptions{})

	assert.EqualError(t, err, "Please provide a valid signature in the X-Acrolinx-Client header.")
}

func TestSubmitCheck(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/checking/checks", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		mustWriteHTTPResponse(t, w, "submit_check.json")
	})

	check, links, err := client.Checking.SubmitCheck(&SubmitCheckOptions{})
	assert.NoError(t, err)

	expectedCheck := &Check{"052929ee-be0c-46a7-87ce-eebd308fef6e"}
	expectedLinks := map[string]string{
		"cancel": "https://example.com/api/v1/checking/checks/052929ee-be0c-46a7-87ce-eebd308fef6e",
		"result": "https://example.com/api/v1/checking/checks/052929ee-be0c-46a7-87ce-eebd308fef6e",
	}

	assert.Equal(t, expectedCheck, check)
	assert.Equal(t, expectedLinks, links)
}

func TestGetCheckProgress(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/checking/checks/052929ee-be0c-46a7-87ce-eebd308fef6e",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			mustWriteHTTPResponse(t, w, "progress.json")
		})

	check := &Check{"052929ee-be0c-46a7-87ce-eebd308fef6e"}
	result, _, err := client.Checking.GetCheckResult(check)
	assert.NoError(t, err)

	assert.NotNil(t, result.Progress)

	expectedProgress := &Progress{
		Percent:    27,
		Message:    "Still processing in state ALLOCATED ...",
		RetryAfter: 1,
	}

	assert.Equal(t, expectedProgress, result.Progress)
}

func TestGetCheckResult(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/checking/checks/052929ee-be0c-46a7-87ce-eebd308fef6e",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			mustWriteHTTPResponse(t, w, "check_result.json")
		})

	check := &Check{"052929ee-be0c-46a7-87ce-eebd308fef6e"}
	result, _, err := client.Checking.GetCheckResult(check)
	assert.NoError(t, err)

	assert.Nil(t, result.Progress)

	expectedResult := &CheckResult{
		ID: "052929ee-be0c-46a7-87ce-eebd308fef6e",
		CheckOptions: AppliedCheckOptions{
			GuidanceProfileID:   "890b68c3-3fb2-369d-b86d-37151d236d9b",
			GuidanceProfileName: "Content and Creative (Prudential)",
			LanguageID:          "en",
			ReportTypes:         []string{"scorecard"},
			ContentFormat:       "TEXT",
			CheckType:           "interactive",
			TermSets: []*TermSet{
				{"ILI and Annuities"},
				{"Retirement PRT and MRT"},
				{"Content and Creative"},
				{"Group Insurance"},
			},
			PartialCheckRanges: []*PartialCheckRange{},
			Confidential:       false,
		},
		Document: &ResponseDocument{
			ID:                   "f8a6dad1-f3e0-4cc7-a5db-337e27abbb97",
			DisplayInfo:          &DisplayInfo{"6f5664bd3a2234f8a7768ff00d866989"},
			CustomFields:         []*CustomField{},
			CustomFieldsComplete: false,
		},
		Quality: &Quality{
			Score:  74,
			Status: "yellow",
			ScoresByStrategy: []*Score{
				{"average", 74},
				{"minimum", 29},
			},
			ScoresByGoal: []*Score{
				{"CONSISTENCY", 100},
				{"SCANNABILITY", 100},
			},
			Metrics: []*Score{
				{"Liveliness index", 85},
				{"Flesch Reading Ease", 71},
			},
		},
		Counts: &Counts{
			Sentences:    7,
			Words:        93,
			Issues:       6,
			ScoredIssues: 6,
		},
		Goals: []*Goal{
			{"CLARITY", "Clarity", "#ec407a", "required", 3},
			{"INCLUSIVE", "Inclusive Language", "#283593", "required", 0},
		},
		Issues: []*Issue{
			{
				GoalID:          "CLARITY",
				InternalName:    "use_comma_after_introductory_phrase",
				DisplayNameHTML: "<div lang=\"en\">Could you add a comma after the introductory phrase?</div>",
				GuidanceHTML:    "<div lang=\"en\">If you use a comma after a prepositional or introductory phrase, your content will be easier to read.</div>",
				DisplaySurface:  "In most cases",
				IssueType:       "actionable",
				Scoring:         "required",
				PositionalInformation: &PositionalInformation{
					Hashes: &Hashes{
						Issue:       "bpraLcyUF2mDHm2dMLp9jl9vBjyZyYu15ZBpbUFKPIk=",
						Environment: "aWN0BgvQacEHtAth/7D1eKlsk+E4+dBrdF1iRCwdQc4=",
						Index:       "1JTFUxBGybxfE4//VcukAkV0GXLWNyp8Tn83tnaNltk==1",
					},
					Matches: []*Match{
						{
							ExtractedPart:  "In",
							ExtractedBegin: 0,
							ExtractedEnd:   2,
							OriginalPart:   "In",
							OriginalBegin:  0,
							OriginalEnd:    2,
						},
					},
				},
				ReadOnly:       false,
				IssueLocations: []*Location{},
				Suggestions: []*Suggestion{
					{
						Surface:      "isn’t",
						GroupID:      "isn’t",
						Replacements: []string{"isn’t", "", ""},
					},
				},
				SubIssues: []*Issue{
					{
						GoalID:          "TONE",
						InternalName:    "phenomenon_RepetitiveStructureModalVerbCan_RepetitiveStructure",
						DisplayNameHTML: "<div lang=\"en\"><q>can just check</q>\n</div>",
						GuidanceHTML:    "",
						IssueType:       "actionable",
						PositionalInformation: &PositionalInformation{
							Hashes: &Hashes{
								Issue:       "KFME5q6POcqou27VfPUfIFa36QaCLwaO67yAYqcAYxI=",
								Environment: "aWN0BgvQacEHtAth/7D1eKlsk+E4+dBrdF1iRCwdQc4=",
								Index:       "ASZBzR4DQ7r2Na9Scnkj5FtEliGk4Js0b6wRvfxCPCc==1",
							},
							Matches: []*Match{
								{
									ExtractedPart:  "can",
									ExtractedBegin: 18,
									ExtractedEnd:   21,
									OriginalPart:   "can",
									OriginalBegin:  18,
									OriginalEnd:    21,
								},
							},
						},
						ReadOnly:           false,
						IssueLocations:     []*Location{},
						Suggestions:        []*Suggestion{},
						SubIssues:          []*Issue{},
						Debug:              &Debug{1000.},
						CanAddToDictionary: false,
						Links:              map[string]string{},
					},
				},
				Debug:              &Debug{},
				CanAddToDictionary: false,
				Links: Links{
					"help": "https://example.com/htmldata/en/rules/9b218748275029619c1f9cdfd66961e47a69e7c8.html",
				},
			},
		},
		Keywords: &Keywords{
			Links: Links{
				"getTargetKeywords": "https://example.com/iq/services/v1/rest/findability/targetKeywords?contextId=4b3b2a7326d29c9bf58a3d51cc0a59dc",
				"putTargetKeywords": "https://example.com/iq/services/v1/rest/findability/targetKeywords?contextId=4b3b2a7326d29c9bf58a3d51cc0a59dc",
			},
			Discovered: []*Keyword{
				{
					Keyword:    "error strings",
					SortKey:    "2",
					Density:    24.180959883021306,
					Count:      2,
					Prominence: 56.84981684981684,
					Occurrences: []*PositionalInformation{
						{
							Hashes: nil,
							Matches: []*Match{
								{
									ExtractedPart:  "error",
									ExtractedBegin: 86,
									ExtractedEnd:   91,
									OriginalPart:   "error",
									OriginalBegin:  86,
									OriginalEnd:    91,
								},
							},
						},
					},
					Warnings: []*KeywordWarning{},
				},
			},
			Target:   []*Keyword{},
			Proposed: []*Keyword{},
		},
		Embed: []*EmbedItem{},
		Reports: map[string]*Report{
			"scorecard": {
				DisplayName:       "Score Card",
				Link:              "https://example.com/api/v1/checking/scorecards/052929ee-be0c-46a7-87ce-eebd308fef6e",
				LinkAuthenticated: "https://example.com/api/v1/checking/scorecards/052929ee-be0c-46a7-87ce-eebd308fef6e",
			},
		},
		RuntimeStatistics: &RuntimeStatistics{
			StartedAt: "2022-10-10T11:55:19.603Z",
		},
		DictionaryScopes: []string{},
	}
	assert.Equal(t, expectedResult, result)
}
