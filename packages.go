package bom

type Materials struct {
	Packages []Package
}

type Package struct {
	Name    string
	Version string
	Release string
	Arch    string
	Format  PackageFormat
	Source  PackageSource
}

type PackageFormat string

const (
	PackageFormatRPM PackageFormat = "rpm"
	PackageFormatDEB PackageFormat = "deb"
)

type PackageSource struct {
	Format PackageFormat
	Name   string
	URL    string

	// like https://github.com/opencontainers/image-spec/blob/master/descriptor.md#digests
	Digest string
}
