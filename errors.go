package bencode

// Error messages and formats.
const (
	Err_ByteStringToInt    = "byte string to integer conversion error"
	Err_DataType           = "unsupported type"
	Err_FileNotInitialized = "file is not initialized"
	Err_SelfCheck          = "self-check error"
	Err_TypeAssertion      = "type assertion error"

	Errf_HeaderLength  = "the length header is too big: %v"
	Errf_IntegerLength = "the integer is too big: %v"
	Errf_SyntaxErrorAt = "syntax error at: '%v'"
)
