package check

const (
	Unspecified Severity = iota
	High
	Medium
	Low
)

type Severity int

type ReleaseInfo struct {
	Version     string `json:"version"`
	ReleaseDate int64  `json:"release_date"`
	Notes       string `json:"notes"`
}

type VersionCheckResponse struct {
	Current     ReleaseInfo `json:"current"`
	Recommended ReleaseInfo `json:"recommended"`
	Severity    Severity    `json:"important"`
	Message     string      `json:"message"`
}
