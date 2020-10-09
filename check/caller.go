package check

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Caller interface {
	Call(r *VersionCheckRequest) (*VersionCheckResponse, error)
}

type callerImpl struct {
	scheme string
	host   string
}

func NewCaller() Caller {
	return &callerImpl{"https", "version-info.temporal.io"}
}

func (vc *callerImpl) Call(r *VersionCheckRequest) (*VersionCheckResponse, error) {
	err := validate(r)
	if err != nil {
		return nil, err
	}
	u := vc.getUrl(r)
	tr := &http.Transport{
		DisableKeepAlives:   true,
		MaxIdleConnsPerHost: -1,
	}
	if vc.scheme == "https" {
		tr.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	versionCheck := &VersionCheckResponse{}
	err = json.Unmarshal(body, versionCheck)
	if err != nil {
		return nil, err
	}
	return versionCheck, nil
}

func validate(r *VersionCheckRequest) error {
	if r.Product == "" || r.Version == "" || r.ClusterID == "" || r.DB == "" || r.OS == "" || r.Arch == "" || r.Timestamp == 0 {
		return errors.New("invalid request: missing required fields")
	}
	return nil
}

func (vc *callerImpl) getUrl(r *VersionCheckRequest) *url.URL {
	var u url.URL
	v := u.Query()
	v.Set("product", r.Product)
	v.Set("version", r.Version)
	v.Set("arch", r.Arch)
	v.Set("os", r.OS)
	v.Set("db", r.DB)
	v.Set("cluster", r.ClusterID)
	v.Set("timestamp", strconv.FormatInt(r.Timestamp, 10))
	u.Scheme = vc.scheme
	u.Host = vc.host
	u.Path = fmt.Sprintf("check")
	u.RawQuery = v.Encode()
	return &u
}
