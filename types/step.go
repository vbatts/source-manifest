package bom

// Step is recording of the target root filesystem at each step or iteration of
// it's build. There may only need to be a final "step", but if prior steps
// have removed or truncated the system to minimize or clean up build time
// dependencies, then you SHOULD record the state at those steps.
type Step struct {
}
