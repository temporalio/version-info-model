package check

// Severity defines importance of the upgrade from the current version to the recommended version.
type Severity int

const (
	// SeverityUnspecified means that severity hasn't been set or is unknown.
	SeverityUnspecified Severity = iota
	// SeverityHigh means that there is an important update available and it should be applied ASAP.
	SeverityHigh
	// SeverityMedium means that there is an important update available however it is not urgent.
	SeverityMedium
	// SeverityLow means that no major changes are available and update is not required.
	SeverityLow
)

// ReleaseInfo contains information about a specific version of the product.
type ReleaseInfo struct {
	Version     string `json:"version"`
	ReleaseDate int64  `json:"release_date"`
	Notes       string `json:"notes"`
}

// Alert contains a message about given update and its importance.
type Alert struct {
	Message  string   `json:"message"`
	Severity Severity `json:"severity"`
}

// VersionCheckResponse contains recommendation about the best version of the product, any alerts for the current version as well as upgrade instructions.
type VersionCheckResponse struct {
	Current      ReleaseInfo `json:"current"`
	Recommended  ReleaseInfo `json:"recommended"`
	Instructions string      `json:"instructions"`
	Alerts       []Alert     `json:"alerts"`
}
