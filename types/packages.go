package types

// Materials is the collection of package information for a specific instance
// of collection.
type Materials struct {
	AnnotationBase
	StructTypeBase

	BOMVersion string `json:"bomversion"`
	Packages   []Package
}

// StructTypeMATERIALS is for the Materials struct
const StructTypeMATERIALS StructType = "materials"

// Package is base metric of system state information to collect.
type Package struct {
	AnnotationBase

	Name    string        `json:"name"`
	Version string        `json:"version"`
	Release string        `json:"release"`
	Arch    string        `json:"arch"`
	URL     string        `json:"url,omitempty"`
	Format  PackageFormat `json:"format"`
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

	Format PackageFormat `json:"format"`
	Name   string        `json:"name"`
	URL    string        `json:"url,omitempty"`
	// like https://github.com/opencontainers/image-spec/blob/master/descriptor.md#digests
	Digest string `json:"digest"`
	// maybe also Version? Commit?
}
