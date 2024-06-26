package bencode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewDecoder(t *testing.T) {

	var aTest = tester.New(t)

	var reader = bufio.NewReader(strings.NewReader(""))
	var decoder = NewDecoder(reader)
	aTest.MustBeEqual(decoder.reader, reader)
}

func Test_Decoder_Decode(t *testing.T) {
	// See Test_readBencodedValue.
}

func Test_Decoder_readBencodedValue(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  any
	}

	var aTest = tester.New(t)
	var tests []TestData

	// Test #1. Negative: No Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"",
			),
		),
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Test #2. Negative: Bad Prefix.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"x",
			),
		),
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Test #3. Positive: A Dictionary.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"d4:info3:abce",
			),
		),
		isErrorExpected: false,
		expectedResult: []DictionaryItem{
			{
				Key:      []byte("info"),
				Value:    []byte("abc"),
				KeyStr:   "info",
				ValueStr: "abc",
			},
		},
	})

	// Test #4. Positive: A List.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"l4:John3:Cate",
			),
		),
		isErrorExpected: false,
		expectedResult: []any{
			[]byte("John"),
			[]byte("Cat"),
		},
	})

	// Test #5. Positive: An Integer.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"i123e",
			),
		),
		isErrorExpected: false,
		expectedResult:  int64(123),
	})

	// Test #6. Positive: A Byte String.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"4:Rain",
			),
		),
		isErrorExpected: false,
		expectedResult:  []byte("Rain"),
	})

	// Run the Tests.
	for i, test := range tests {
		fmt.Printf("Test #%v.\r\n", i+1)
		decoder := NewDecoder(test.reader)
		result, err := decoder.Decode()
		if test.isErrorExpected {
			aTest.MustBeAnError(err)
			fmt.Println(err)
		} else {
			aTest.MustBeNoError(err)
			aTest.MustBeEqual(result, test.expectedResult)
		}
	}
}

func Test_Decoder_readByteString(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  any
	}

	var aTest = tester.New(t)
	var tests []TestData

	// Test #1. Negative: Bad Size Header.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"-2",
			),
		),
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Test #2. Positive.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"3:Sun",
			),
		),
		isErrorExpected: false,
		expectedResult:  []byte("Sun"),
	})

	// Test #3. Negative: Not enough Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"4:X",
			),
		),
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Run the Tests.
	for i, test := range tests {
		fmt.Printf("Test #%v.\r\n", i+1)
		decoder := NewDecoder(test.reader)
		result, err := decoder.readByteString()
		if test.isErrorExpected {
			aTest.MustBeAnError(err)
			fmt.Println(err)
		} else {
			aTest.MustBeNoError(err)
			aTest.MustBeEqual(result, test.expectedResult)
		}
	}
}

func Test_Decoder_readByteStringSizeHeader(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  any
	}

	var aTest = tester.New(t)
	var tests []TestData

	// Test #1. Negative: No Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"",
			),
		),
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Test #2. Negative: Bad Symbol.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"x",
			),
		),
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Test #3. Negative: Size Header Overflow.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"123456789012345678901:qwerty", // 21 Symbol in Size Header.
			),
		),
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Test #4. Positive: Normal Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"16:Abcdefghijklmnop",
			),
		),
		isErrorExpected: false,
		expectedResult:  uint(16),
	})

	// Test #5. Negative: Empty Header.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				":Qwe",
			),
		),
		isErrorExpected: true,
	})

	// Test #6. Negative: Not enough Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"2",
			),
		),
		isErrorExpected: true,
	})

	// Run the Tests.
	for i, test := range tests {
		fmt.Printf("Test #%v.\r\n", i+1)
		decoder := NewDecoder(test.reader)
		result, err := decoder.readByteStringSizeHeader()
		if test.isErrorExpected {
			aTest.MustBeAnError(err)
			fmt.Println(err)
		} else {
			aTest.MustBeNoError(err)
			aTest.MustBeEqual(result, test.expectedResult)
			fmt.Println(result)
		}
	}
}

func Test_Decoder_readDictionary(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  any
	}

	var aTest = tester.New(t)
	var tests []TestData

	// Test #1. Negative: No Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"",
			),
		),
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Test #2. Positive: Normal Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"4:info3:abc2:rx3:cate", // Without 'd' Prefix !
			),
		),
		isErrorExpected: false,
		expectedResult: []DictionaryItem{
			{
				Key:      []byte("info"),
				Value:    []byte("abc"),
				KeyStr:   "info",
				ValueStr: "abc",
			},
			{
				Key:      []byte("rx"),
				Value:    []byte("cat"),
				KeyStr:   "rx",
				ValueStr: "cat",
			},
		},
	})

	// Test #3. Negative: Bad Value.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"4:info3:ab", // Without 'd' Prefix !
			),
		),
		isErrorExpected: true,
	})

	// Test #4. Negative: Bad Key.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"4:ix", // Without 'd' Prefix !
			),
		),
		isErrorExpected: true,
	})

	// Test #5. Negative: No Ending.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"4:info3:abc", // Without 'd' Prefix !
			),
		),
		isErrorExpected: true,
	})

	// Run the Tests.
	for i, test := range tests {
		fmt.Printf("Test #%v.\r\n", i+1)
		decoder := NewDecoder(test.reader)
		result, err := decoder.readDictionary()
		if test.isErrorExpected {
			aTest.MustBeAnError(err)
			fmt.Println(err)
		} else {
			aTest.MustBeNoError(err)
			aTest.MustBeEqual(result, test.expectedResult)
			fmt.Println(result)
		}
	}
}

func Test_Decoder_readDictionaryKey(t *testing.T) {
	// See Test_readByteString.
}

func Test_Decoder_readDictionaryValue(t *testing.T) {
	// See Test_readBencodedValue.
}

func Test_Decoder_readInteger(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  any
	}

	var aTest = tester.New(t)
	var tests []TestData

	// Test #1. Negative: No Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"",
			),
		),
		isErrorExpected: true,
	})

	// Test #2. Negative: Bad Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"q",
			),
		),
		isErrorExpected: true,
	})

	// Test #3. Positive: Normal Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"1234e", // Without 'i' Prefix !
			),
		),
		isErrorExpected: false,
		expectedResult:  int64(1234),
	})

	// Test #4. Negative: Overflow.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"1234567890123456789012345e", // Without 'i' Prefix !
			),
		),
		isErrorExpected: true,
	})

	// Test #5. Negative: Unexpected End.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"123", // Without 'i' Prefix !
			),
		),
		isErrorExpected: true,
	})

	// Test #6. Negative: No Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"e", // Without 'i' Prefix !
			),
		),
		isErrorExpected: true,
	})

	// Test #7. Positive: Maximum Size.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"-2345678901234567890e", // Without 'i' Prefix !
			),
		),
		isErrorExpected: false,
		expectedResult:  int64(-2345678901234567890),
	})

	// Run the Tests.
	for i, test := range tests {
		fmt.Printf("Test #%v.\r\n", i+1)
		decoder := NewDecoder(test.reader)
		result, err := decoder.readInteger()
		if test.isErrorExpected {
			aTest.MustBeAnError(err)
			fmt.Println(err)
		} else {
			aTest.MustBeNoError(err)
			aTest.MustBeEqual(result, test.expectedResult)
			fmt.Println(result)
		}
	}
}

func Test_Decoder_readList(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  any
	}

	var aTest = tester.New(t)
	var tests []TestData

	// Test #1. Negative: No Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"",
			),
		),
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Test #2. Negative: Bad Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"q",
			),
		),
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Test #3. Positive: Normal Data.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"2:Wee", // Without 'l' Prefix !
			),
		),
		isErrorExpected: false,
		expectedResult: []any{
			[]byte("We"),
		},
	})

	// Test #4. Positive: Overflow.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"3:Cat2:Use", // Without 'l' Prefix !
			),
		),
		isErrorExpected: false,
		expectedResult: []any{
			[]byte("Cat"),
			[]byte("Us"),
		},
	})

	// Test #5. Negative: Unexpected End.
	tests = append(tests, TestData{
		reader: bufio.NewReader(
			strings.NewReader(
				"2:We",
			),
		),
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Run the Tests.
	for i, test := range tests {
		fmt.Printf("Test #%v.\r\n", i+1)
		decoder := NewDecoder(test.reader)
		result, err := decoder.readList()
		if test.isErrorExpected {
			aTest.MustBeAnError(err)
			fmt.Println(err)
		} else {
			aTest.MustBeNoError(err)
			aTest.MustBeEqual(result, test.expectedResult)
			fmt.Println(result)
		}
	}
}
