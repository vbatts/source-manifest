package types

// Host is the distribution information for the target root filesystem.
// Largely correlates with `/etc/os-release`
type Host struct {
	StructTypeBase
	AnnotationBase

	Name         string `json:"name"`
	PrettyName   string `json:"pretty_name"`
	ID           string `json:"id"`
	Version      string `json:"version"`
	VersionID    string `json:"version_id"`
	HomeURL      string `json:"home_url"`
	SupportURL   string `json:"support_url"`
	BugReportURL string `json:"bug_report_url"`
	StepUUID     string `json:"step_uuid,omitempty"` // to correlate this material list with a specific Step
}

// StructTypeHOST is for the Host struct
const StructTypeHOST StructType = "host"
