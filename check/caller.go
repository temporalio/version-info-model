package check

import (
	"crypto/tls"
	"encoding/json"
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

func NewCaller() *callerImpl {
	return &callerImpl{"https", "version-info.temporal.io"}
}

func (vc *callerImpl) Call(r *VersionCheckRequest) (*VersionCheckResponse, error) {
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
