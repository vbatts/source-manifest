package bom

type Packages struct {
	Name    string
	Version string
	Release string
	Arch    string
	Format  PackageFormat
	Source  SourceInfo
}

type PackageFormat string

const (
	PackageFormatRPM PackageFormat = "rpm"
	PackageFormatDEB PackageFormat = "deb"
)

type SourceInfo struct {
	Format PackageFormat
	Name   string
	URL    string

	// like https://github.com/opencontainers/image-spec/blob/master/descriptor.md#digests
	Digest string
}
