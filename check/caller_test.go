package check

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestPostInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Method != POST (%s)", r.Method)
		}
		if r.URL.Path != "/check" {
			t.Errorf("URL.Path != /check (%s)", r.URL.Path)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type != application/json (%s)", r.Header.Get("Content-Type"))
		}
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body %s", err)
		}
		versionCheckRequest := &VersionCheckRequest{}
		err = json.Unmarshal(body, versionCheckRequest)
		if err != nil {
			t.Fatalf("Failed to unmarshal request body %s", err)
		}
		// Unmarshalling works
		res, err := json.Marshal(VersionCheckResponse{
			Products: []ProductVersionReport{
				ProductVersionReport{
					Product: "server",
					Current: ReleaseInfo{
						Version:     "0.1",
						ReleaseTime: time.Now().UnixNano(),
						Notes:       "",
					},
					Recommended: ReleaseInfo{
						Version:     "0.1",
						ReleaseTime: time.Now().UnixNano(),
						Notes:       "",
					},
					Instructions: "instructions",
					Alerts:       []Alert{},
				},
			},
		})
		if err != nil {
			t.Fatalf("Failed to marshal response %s", err)
		}
		w.Write(res)
	}))
	url, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("Request failed: %s", err)
	}
	caller := &callerImpl{url.Scheme, url.Host}
	sdkInfo := []SDKInfo{{
		Name:    "sdk-java",
		Version: "3.11",
	}}
	_, err = caller.Call(&VersionCheckRequest{
		Product:   "server",
		Version:   "0.1",
		ClusterID: "foo",
		DB:        "cassandra",
		OS:        "linux",
		Arch:      "arm64",
		Timestamp: time.Now().UnixNano(),
		SDKInfo:   sdkInfo,
	})
	if err != nil {
		t.Fatalf("Request failed: %s", err)
	}
}
