package bencode

// Error messages and formats.
const (
	ErrByteStringToInt    = "byte string to integer conversion error"
	ErrDataType           = "unsupported type"
	ErrFileNotInitialized = "file is not initialized"
	ErrHeaderLength       = "the length header is too big: %v"
	ErrSelfCheck          = "self-check error"
	ErrTypeAssertion      = "type assertion error"
	ErrFIntegerLength     = "the integer is too big: %v"
	ErrFSyntaxErrorAt     = "syntax error at: '%v'"
)
