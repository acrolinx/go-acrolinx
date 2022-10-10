package acrolinx

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCapabilities(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/checking/capabilities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "get_capabilities.json")
	})

	opts := &ListCapabilitiesOptions{}
	caps, _, err := client.Checking.ListCapabilities(opts)
	if err != nil {
		t.Fatalf("Checking.ListCapabilities returned error: %v", err)
	}

	expected := &Capabilities{
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

	assert.Equal(t, expected, caps)
}
