package types

// Materials is the collection of package information for a specific instance
// of collection.
type Materials struct {
	StructTypeBase
	AnnotationBase

	Packages []Package `json:"packages"`
	StepUUID string    `json:"step_uuid,omitempty"` // to correlate this material list with a specific Step
}

// StructTypeMATERIALS is for the Materials struct
const StructTypeMATERIALS StructType = "materials"

// Package is base metric of system state information to collect.
type Package struct {
	AnnotationBase
	PackageURLBase
	PackageFormatBase

	Name    string        `json:"name"`
	Version string        `json:"version"`
	Release string        `json:"release"`
	Arch    string        `json:"arch"`
	Source  PackageSource `json:"source"`
}

// PackageFormat is the format or type of package. This is typed for seeing
// some defaults, but rather than locking it down, additional package formats
// can easily be provided without breaking this type.
type PackageFormat string

// initial packages. Additional could be pypi? rubygem? cargo crate?
const (
	PackageFormatRPM PackageFormat = "rpm"
	PackageFormatDEB PackageFormat = "deb"
	PackageFormatTAR PackageFormat = "tar"
	PackageFormatGIT PackageFormat = "git"
)

// PackageSource is the artifact or origin of source code that the Package is
// directly derived from (revision/commit or checksum of the source package).
type PackageSource struct {
	AnnotationBase
	PackageURLBase
	PackageFormatBase

	Name string `json:"name"`
	// like https://github.com/opencontainers/image-spec/blob/master/descriptor.md#digests
	Digest string `json:"digest"`
	// maybe also Version? Commit?
}

// PackageURLBase is generic for having a optional URL
type PackageURLBase struct {
	URL string `json:"url,omitempty"`
}

// PackageFormatBase is generic for having a Format
type PackageFormatBase struct {
	Format PackageFormat `json:"format"`
}
