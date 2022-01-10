package check

// VersionCheckRequest provides basic info about the client and is used to produce VersionCheckResponse.
type SDKInfo struct {
	Name      string `json:"sdk_name"`
	Version   string `json:"sdk_version"`
	TimesSeen int64  `json:"times_seen"`
}

// VersionCheckRequest provides basic info about the client and is used to produce VersionCheckResponse.
type VersionCheckRequest struct {
	Product   string `json:"product"`
	Version   string `json:"version"`
	Arch      string `json:"arch"`
	OS        string `json:"os"`
	DB        string `json:"db"`
	ClusterID string `json:"cluster"`
	// Encode as string for JS compatibility
	Timestamp int64     `json:"timestamp,string"`
	SDKInfo   []SDKInfo `json:"sdk_info"`
}
