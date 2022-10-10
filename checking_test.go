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
