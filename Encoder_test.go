package bencode

import (
	"fmt"
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewEncoder(t *testing.T) {
	var aTest = tester.New(t)

	var encoder = NewEncoder()
	aTest.MustBeEqual(encoder.commonPostfix, []byte{'e'})
	aTest.MustBeEqual(encoder.dictionaryPrefix, []byte{'d'})
	aTest.MustBeEqual(encoder.integerPrefix, []byte{'i'})
	aTest.MustBeEqual(encoder.listPrefix, []byte{'l'})
}

func Test_Encoder_addPostfixOfDictionary(t *testing.T) {

	var aTest = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("abc")
	result := encoder.addPostfixOfDictionary(tmpResult)
	resultExpected := []byte("abce")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Encoder_addPostfixOfList(t *testing.T) {

	var aTest = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("lt")
	result := encoder.addPostfixOfList(tmpResult)
	resultExpected := []byte("lte")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Encoder_addPrefixAndPostfixOfByteString(t *testing.T) {

	var aTest = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("bst")
	result := encoder.addPrefixAndPostfixOfByteString(tmpResult)
	resultExpected := []byte("3:bst")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Encoder_addPrefixAndPostfixOfInteger(t *testing.T) {

	var aTest = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("909")
	result := encoder.addPrefixAndPostfixOfInteger(tmpResult)
	resultExpected := []byte("i909e")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Encoder_addPrefixOfDictionary(t *testing.T) {

	var aTest = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("xx")
	result := encoder.addPrefixOfDictionary(tmpResult)
	resultExpected := []byte("dxx")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Encoder_addPrefixOfList(t *testing.T) {

	var aTest = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("nn")
	result := encoder.addPrefixOfList(tmpResult)
	resultExpected := []byte("lnn")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Encoder_createSizePrefix(t *testing.T) {

	var aTest = tester.New(t)

	encoder := NewEncoder()
	result := encoder.createSizePrefix(3)
	resultExpected := []byte("3:")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Encoder_createTextFromInteger(t *testing.T) {

	var aTest = tester.New(t)

	encoder := NewEncoder()
	result := encoder.createTextFromInteger(-56)
	resultExpected := []byte("-56")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Encoder_createTextFromUInteger(t *testing.T) {

	var aTest = tester.New(t)

	encoder := NewEncoder()
	result := encoder.createTextFromUInteger(56)
	resultExpected := []byte("56")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Encoder_EncodeAnInterface(t *testing.T) {

	type TestData struct {
		dataToBeEncoded any
		isErrorExpected bool
		expectedResult  any
	}

	var aTest = tester.New(t)
	var tests []TestData

	// Test #1. Slice.
	tests = append(tests, TestData{
		dataToBeEncoded: []byte("ABC"),
		expectedResult:  []byte("3:ABC"),
	})

	// Test #2. String.
	tests = append(tests, TestData{
		dataToBeEncoded: "John",
		expectedResult:  []byte("4:John"),
	})

	// Test #3. Uint.
	tests = append(tests, TestData{
		dataToBeEncoded: uint(123),
		expectedResult:  []byte("i123e"),
	})

	// Test #4. Int.
	tests = append(tests, TestData{
		dataToBeEncoded: -124,
		expectedResult:  []byte("i-124e"),
	})

	// Test #5. Uint64.
	tests = append(tests, TestData{
		dataToBeEncoded: uint64(1064),
		expectedResult:  []byte("i1064e"),
	})

	// Test #6. Int64.
	tests = append(tests, TestData{
		dataToBeEncoded: int64(-1065),
		expectedResult:  []byte("i-1065e"),
	})

	// Test #7. Uint32.
	tests = append(tests, TestData{
		dataToBeEncoded: uint32(1066),
		expectedResult:  []byte("i1066e"),
	})

	// Test #8. Int32.
	tests = append(tests, TestData{
		dataToBeEncoded: int32(-1067),
		expectedResult:  []byte("i-1067e"),
	})

	// Test #8. Uint16.
	tests = append(tests, TestData{
		dataToBeEncoded: uint16(1068),
		expectedResult:  []byte("i1068e"),
	})

	// Test #9. Int16.
	tests = append(tests, TestData{
		dataToBeEncoded: int16(-1069),
		expectedResult:  []byte("i-1069e"),
	})

	// Test #10. Uint8.
	tests = append(tests, TestData{
		dataToBeEncoded: uint8(123),
		expectedResult:  []byte("i123e"),
	})

	// Test #11. Int8.
	tests = append(tests, TestData{
		dataToBeEncoded: int8(-124),
		expectedResult:  []byte("i-124e"),
	})

	// Test #12. Bad Type.
	tests = append(tests, TestData{
		dataToBeEncoded: time.Time{},
		isErrorExpected: true,
		expectedResult:  nil,
	})

	// Run the Tests.
	for i, test := range tests {
		fmt.Printf("Test #%v.\r\n", i+1)
		encoder := NewEncoder()
		result, err := encoder.EncodeAnInterface(test.dataToBeEncoded)
		if test.isErrorExpected {
			aTest.MustBeAnError(err)
			fmt.Println(err)
		} else {
			aTest.MustBeNoError(err)
			aTest.MustBeEqual(result, test.expectedResult)
			fmt.Println(string(result))
		}
	}
}

func Test_Encoder_encodeDictionary(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var dictionary []DictionaryItem
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1.
	{
		encoder = NewEncoder()
		dictionary = []DictionaryItem{
			{
				Key:      []byte("Aa"),
				Value:    123,
				KeyStr:   "Aa",
				ValueStr: "123",
			},
			{
				Key:      []byte("Bb"),
				Value:    "QWERTY",
				KeyStr:   "Bb",
				ValueStr: "QWERTY",
			},
		}
		resultExpected = []byte("d2:Aai123e2:Bb6:QWERTYe")
		result, err = encoder.encodeDictionary(dictionary)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		dictionary = []DictionaryItem{
			{
				Key:      []byte("Aa"),
				Value:    time.Time{},
				KeyStr:   "Aa",
				ValueStr: "123",
			},
		}
		result, err = encoder.encodeDictionary(dictionary)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfInt(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = 123
		resultExpected = []byte("i123e")
		result, err = encoder.encodeInterfaceOfInt(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = uint(123)
		result, err = encoder.encodeInterfaceOfInt(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfInt8(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = int8(127)
		resultExpected = []byte("i127e")
		result, err = encoder.encodeInterfaceOfInt8(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = "qqq"
		result, err = encoder.encodeInterfaceOfInt8(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfInt16(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = int16(127)
		resultExpected = []byte("i127e")
		result, err = encoder.encodeInterfaceOfInt16(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = "qqq"
		result, err = encoder.encodeInterfaceOfInt16(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfInt32(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = int32(127)
		resultExpected = []byte("i127e")
		result, err = encoder.encodeInterfaceOfInt32(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = "qqq"
		result, err = encoder.encodeInterfaceOfInt32(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfInt64(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = int64(127)
		resultExpected = []byte("i127e")
		result, err = encoder.encodeInterfaceOfInt64(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = "qqq"
		result, err = encoder.encodeInterfaceOfInt64(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfList(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data []any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = []any{
			int8(123),
			"Qwe",
		}
		resultExpected = []byte("li123e3:Qwee")
		result, err = encoder.encodeInterfaceOfList(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = []any{
			int8(123),
			time.Time{},
		}
		result, err = encoder.encodeInterfaceOfList(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfSlice(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Slice of Bytes.
	{
		encoder = NewEncoder()
		data = []byte("ABC")
		resultExpected = []byte("3:ABC")
		result, err = encoder.encodeInterfaceOfSlice(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Dictionary.
	{
		data = []DictionaryItem{
			{
				Key:      []byte("Aa"),
				Value:    123,
				KeyStr:   "Aa",
				ValueStr: "123",
			},
		}
		resultExpected = []byte("d2:Aai123ee")
		result, err = encoder.encodeInterfaceOfSlice(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #3. Slice of Interfaces.
	{
		data = []any{
			"Qwerty",
			uint16(6565),
		}
		resultExpected = []byte("l6:Qwertyi6565ee")
		result, err = encoder.encodeInterfaceOfSlice(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #4. unknown Type.
	{
		data = []time.Time{}
		result, err = encoder.encodeInterfaceOfSlice(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfSliceOfBytes(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = []byte("Qwe")
		resultExpected = []byte("3:Qwe")
		result, err = encoder.encodeInterfaceOfSliceOfBytes(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = 123
		result, err = encoder.encodeInterfaceOfSliceOfBytes(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfString(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = "Abc"
		resultExpected = []byte("3:Abc")
		result, err = encoder.encodeInterfaceOfString(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = 123
		result, err = encoder.encodeInterfaceOfString(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfUint(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = uint(123)
		resultExpected = []byte("i123e")
		result, err = encoder.encodeInterfaceOfUint(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = "x"
		result, err = encoder.encodeInterfaceOfUint(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfUint8(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = uint8(123)
		resultExpected = []byte("i123e")
		result, err = encoder.encodeInterfaceOfUint8(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = "x"
		result, err = encoder.encodeInterfaceOfUint8(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfUint16(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = uint16(123)
		resultExpected = []byte("i123e")
		result, err = encoder.encodeInterfaceOfUint16(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = "x"
		result, err = encoder.encodeInterfaceOfUint16(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfUint32(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = uint32(123)
		resultExpected = []byte("i123e")
		result, err = encoder.encodeInterfaceOfUint32(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = "x"
		result, err = encoder.encodeInterfaceOfUint32(data)
		aTest.MustBeAnError(err)
	}
}

func Test_Encoder_encodeInterfaceOfUint64(t *testing.T) {

	var aTest = tester.New(t)

	var encoder *Encoder
	var data any
	var result []byte
	var resultExpected []byte
	var err error

	// Test #1. Positive.
	{
		encoder = NewEncoder()
		data = uint64(123)
		resultExpected = []byte("i123e")
		result, err = encoder.encodeInterfaceOfUint64(data)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, resultExpected)
	}

	// Test #2. Negative.
	{
		data = "x"
		result, err = encoder.encodeInterfaceOfUint64(data)
		aTest.MustBeAnError(err)
	}
}
