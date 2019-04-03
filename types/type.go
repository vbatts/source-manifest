package types

// AnnotationBase is the base struct to include for the annotations field
type AnnotationBase struct {
	Annotations map[string]string `json:"annotations"`
}

// StructTypeBase is the base struct to include for a field to have a struct inform of its type.
//
// Any data collector that runs on a host MUST set a "struct_type" on the items it exports.
type StructTypeBase struct {
	StructType StructType `json:"struct_type"`
}

// StructType is the type that any data exported by a collector MUST set
type StructType string
