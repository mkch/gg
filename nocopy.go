package gg

// NoCopy may be embedded into structs which must not be copied after the first use.
// The misuse of this struct will be detected by the go vet -copylocks tool,
// which will report an error if a struct containing NoCopy is copied.
type NoCopy struct{}

func (*NoCopy) Lock()   {}
func (*NoCopy) Unlock() {}
