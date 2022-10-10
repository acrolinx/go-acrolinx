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
					{"CLARITY", "Verständlichkeit", "#ec407a", "required"},
					{"CONSISTENCY", "Einheitlichkeit", "#ffd600", "required"},
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
