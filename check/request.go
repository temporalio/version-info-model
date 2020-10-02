package check

type VersionCheckRequest struct {
	Product string
	Version string
	Arch    string
	OS      string
	DB      string
	Token   string
}
