package acrolinx

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientWithErr(t *testing.T) {
	_, err := NewClient("signature", "not a URL")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "Error parsing platform URL")
}

func TestSignIn(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/dashboard/api/signin/authenticate", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		mustWriteHTTPResponse(t, w, "sign_in.json")
	})

	err := client.SignIn("username", "password")
	assert.NoError(t, err)

	assert.Equal(t, "VGhlIGZhbGNvbiBoZWFycyB0aGUgZmFsY29uZXIK", client.accessToken)
}

func setup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)

	client, err := NewClient("testsignature", server.URL)
	if err != nil {
		server.Close()
		t.Fatalf("Failed to create client: %v", err)
	}

	return mux, server, client
}

func teardown(server *httptest.Server) {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if actual := r.Method; actual != expected {
		t.Errorf("Request method: %s, expected %s", actual, expected)
	}
}

func mustWriteHTTPResponse(t *testing.T, w io.Writer, fixtureFile string) {
	fixturePath := "testdata/" + fixtureFile
	f, err := os.Open(fixturePath)
	if err != nil {
		t.Fatalf("Error opening fixture at %s: %v", fixturePath, err)
	}

	if _, err = io.Copy(w, f); err != nil {
		t.Fatalf("Error writing HTTP response: %v", err)
	}
}
