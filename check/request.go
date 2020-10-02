package check

// VersionCheckRequest provides basic info about the client and is used to produce VersionCheckResponse.
type VersionCheckRequest struct {
	Product   string
	Version   string
	Arch      string
	OS        string
	DB        string
	ClusterID string
	Timestamp int64
}
