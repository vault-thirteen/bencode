package bencode

import "errors"

// Error messages and formats.
const (
	ErrByteStringToIntError     = "byte string to integer conversion error"
	ErrDataTypeError            = "unsupported type"
	ErrFileNotInitializedError  = "file is not initialized"
	ErrHeaderLengthError        = "the length header is too big: %v"
	ErrSectionDoesNotExistError = "section does not exist"
	ErrSelfCheckError           = "self-check error"
	ErrTypeAssertionError       = "type assertion error"

	ErrFIntegerLengthError = "the integer is too big: %v"
)

// Cached errors.
var (
	ErrByteStringToInt     = errors.New(ErrByteStringToIntError)
	ErrDataType            = errors.New(ErrDataTypeError)
	ErrFileNotInitialized  = errors.New(ErrFileNotInitializedError)
	ErrSectionDoesNotExist = errors.New(ErrSectionDoesNotExistError)
	ErrSelfCheck           = errors.New(ErrSelfCheckError)
	ErrTypeAssertion       = errors.New(ErrTypeAssertionError)
)

// Formats of error messages.
const (
	ErrFSyntaxErrorAt = "syntax error at: '%v'"
	ErrCombinator     = ": "
)

// combineErrors combines two errors.
func combineErrors(
	error1 error,
	error2 error,
) (result error) {
	if error1 == nil {
		if error2 == nil {
			return nil
		} else {
			return error2
		}
	} else {
		if error2 == nil {
			return error1
		} else {
			return errors.New(error1.Error() + ErrCombinator + error2.Error())
		}
	}
}
