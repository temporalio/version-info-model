package check

// Severity defines importance of the upgrade from the current version to the recommended version.
type Severity int

const (
	// SeverityUnspecified means that severity hasn't been set or is unknown.
	SeverityUnspecified Severity = iota
	// SeverityHigh means that there are important updates available for your version and they should be applied ASAP.
	SeverityHigh
	// SeverityMedium means that there are important updates available however neither of them is urgent.
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

// Alert contains message and severity of a given update.
type Alert struct {
	Message  string   `json:"message"`
	Severity Severity `json:"severity"`
}

// VersionCheckResponse contains recommendation about the best version of product and
type VersionCheckResponse struct {
	Current     ReleaseInfo `json:"current"`
	Recommended ReleaseInfo `json:"recommended"`
	Alerts      []Alert     `json:"alerts"`
}
