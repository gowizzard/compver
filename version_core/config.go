package version_core

// Split is to save the version core as major, minor and patch
type Split struct {
	Major int
	Minor int
	Patch int
}

// Block is to save the version core blocks as a pair
type Block struct {
	Name    string
	Number1 int
	Number2 int
}
