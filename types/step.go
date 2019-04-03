package types

// Step is recording of the target root filesystem at each step or iteration of
// it's build. There may only need to be a final "step", but if prior steps
// have removed or truncated the system to minimize or clean up build time
// dependencies, then you SHOULD record the state at those steps.
type Step struct {
	AnnotationBase
	StructTypeBase

	Time      string `json:"time"`      // SHOULD be RFC 3339
	Operation string `json:"operation"` // like RUN, FROM, COPY, ADD
	Action    string `json:"action"`    // 'fedora:latest@sha256:aabbccdd', 'dnf install -y redis'
}

// StructTypeSTEP is for the Step struct
const StructTypeSTEP StructType = "step"
