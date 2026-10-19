package bencode

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_convertByteStringToInteger(t *testing.T) {
	var aTest = tester.New(t)

	var bytes []byte
	var result int64
	var err error

	// Test #1. Negative.
	{
		bytes = []byte("123456789012345678912345")
		result, err = convertByteStringToInteger(bytes)
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(result, int64(0))
	}

	// Test #2. Positive.
	{
		bytes = []byte("-12345")
		result, err = convertByteStringToInteger(bytes)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, int64(-12345))
	}
}

func Test_convertByteStringToNonNegativeInteger(t *testing.T) {
	var aTest = tester.New(t)

	var bytes []byte
	var result uint64
	var err error

	// Test #1. Negative.
	{
		bytes = []byte("123456789012345678912345")
		result, err = convertByteStringToNonNegativeInteger(bytes)
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(result, uint64(0))
	}

	// Test #2. Positive.
	{
		bytes = []byte("12345")
		result, err = convertByteStringToNonNegativeInteger(bytes)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, uint64(12345))
	}
}

func Test_convertInterfaceToString(t *testing.T) {
	var aTest = tester.New(t)

	var result string

	// Test #1. Negative: Not a Slice.
	{
		result = convertInterfaceToString("Error")
		aTest.MustBeEqual(result, "")
	}

	// Test #2. Negative: Not a Slice of Bytes.
	{
		result = convertInterfaceToString([]rune("xyz"))
		aTest.MustBeEqual(result, "")
	}

	// Test #3. Positive.
	{
		result = convertInterfaceToString([]byte("Text"))
		aTest.MustBeEqual(result, "Text")
	}
}

func Test_isByteAsciiNumeric(t *testing.T) {
	var aTest = tester.New(t)

	aTest.MustBeEqual(isByteAsciiNumeric('0'), true)
	aTest.MustBeEqual(isByteAsciiNumeric('1'), true)
	aTest.MustBeEqual(isByteAsciiNumeric('2'), true)
	aTest.MustBeEqual(isByteAsciiNumeric('3'), true)
	aTest.MustBeEqual(isByteAsciiNumeric('4'), true)
	aTest.MustBeEqual(isByteAsciiNumeric('5'), true)
	aTest.MustBeEqual(isByteAsciiNumeric('6'), true)
	aTest.MustBeEqual(isByteAsciiNumeric('7'), true)
	aTest.MustBeEqual(isByteAsciiNumeric('8'), true)
	aTest.MustBeEqual(isByteAsciiNumeric('9'), true)
	aTest.MustBeEqual(isByteAsciiNumeric('-'), true)

	aTest.MustBeEqual(isByteAsciiNumeric('x'), false)
}

func Test_isByteNonNegativeAsciiNumeric(t *testing.T) {
	var aTest = tester.New(t)

	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('0'), true)
	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('1'), true)
	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('2'), true)
	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('3'), true)
	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('4'), true)
	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('5'), true)
	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('6'), true)
	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('7'), true)
	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('8'), true)
	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('9'), true)

	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('x'), false)
}
