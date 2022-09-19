package version_core

// Core is to save the version core as major, minor and patch.
type Core struct {
	Major int
	Minor int
	Patch int
}

// Block is to save the version core blocks as a pair.
type Block struct {
	Name    string
	Number1 int
	Number2 int
}
