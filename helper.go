package bencode

import (
	"errors"
	"reflect"
	"strconv"
)

// convertByteStringToInteger converts a byte string into a signed 64-bit
// integer. Negative numbers are possible.
func convertByteStringToInteger(ba []byte) (result int64, err error) {
	if len(ba) > IntegerMaxLength {
		return 0, errors.New(ErrByteStringToInt)
	}

	return strconv.ParseInt(string(ba), 10, 64)
}

// convertByteStringToNonNegativeInteger converts a byte string into an
// unsigned 64-bit integer. Negative numbers are forbidden.
func convertByteStringToNonNegativeInteger(ba []byte) (result uint64, err error) {
	if len(ba) > IntegerMaxLength {
		return 0, errors.New(ErrByteStringToInt)
	}

	return strconv.ParseUint(string(ba), 10, 64)
}

// convertInterfaceToString tries to get a textual data from an interface.
func convertInterfaceToString(src any) (result string) {

	// Not a slice ?
	if reflect.TypeOf(src).Kind() != reflect.Slice {
		return ""
	}

	// Is array item's type not a byte ?
	if reflect.TypeOf(src).Elem().Kind() != reflect.Uint8 {
		return ""
	}

	var bytes []byte
	var ok bool
	bytes, ok = src.([]byte)
	if !ok {
		return ""
	}

	return string(bytes)
}

// isByteAsciiNumeric checks whether the byte is ASCII numeric symbol. Negative
// numbers are possible.
func isByteAsciiNumeric(b byte) (result bool) {

	if (b == '0') ||
		(b == '1') ||
		(b == '2') ||
		(b == '3') ||
		(b == '4') ||
		(b == '5') ||
		(b == '6') ||
		(b == '7') ||
		(b == '8') ||
		(b == '9') ||
		(b == '-') {
		return true
	}

	return false
}

// isByteNonNegativeAsciiNumeric checks whether the byte is ASCII non-negative
// numeric symbol. Negative numbers are forbidden.
func isByteNonNegativeAsciiNumeric(b byte) (result bool) {

	if (b == '0') ||
		(b == '1') ||
		(b == '2') ||
		(b == '3') ||
		(b == '4') ||
		(b == '5') ||
		(b == '6') ||
		(b == '7') ||
		(b == '8') ||
		(b == '9') {
		return true
	}

	return false
}
