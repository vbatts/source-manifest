package bom

type Materials struct {
	BOMVersion string `json:"bomversion"`
	Packages   []Package
}

type Package struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Release     string            `json:"release"`
	Arch        string            `json:"arch"`
	URL         string            `json:"url,omitempty"`
	Format      PackageFormat     `json:"format"`
	Source      PackageSource     `json:"source"`
	Annotations map[string]string `json:"annotations"`
}

type PackageFormat string

const (
	PackageFormatRPM PackageFormat = "rpm"
	PackageFormatDEB PackageFormat = "deb"
	PackageFormatTAR PackageFormat = "tar"
	// pypi? rubygem? cargo crate?
)

type PackageSource struct {
	Format      PackageFormat     `json:"format"`
	Name        string            `json:"name"`
	URL         string            `json:"url,omitempty"`
	Annotations map[string]string `json:"annotations"`
	// like https://github.com/opencontainers/image-spec/blob/master/descriptor.md#digests
	Digest string `json:"digest"`
}
