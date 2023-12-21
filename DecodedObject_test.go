package bencode

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_CalculateBtih(t *testing.T) {
	var aTest = tester.New(t)
	var err error

	// Test #1.
	var object = DecodedObject{
		DecodedObject: []DictionaryItem{
			{
				Key:   []byte("info"),
				Value: "Just a Test.",
			},
		},
	}
	var expectedBtihText = "6f1ef4ba8a877d657378dbbb78badfd2eaacf2a2"
	var expectedBtihBytes Sha1Sum
	var ba []byte
	// "Just a Test." -> "12:Just a Test." -> (SHA-1)
	ba, err = hex.DecodeString(expectedBtihText)
	aTest.MustBeNoError(err)
	copy(expectedBtihBytes[:], ba)
	//
	err = object.CalculateBtih()
	//
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(object.BTIH.Text, expectedBtihText)
	aTest.MustBeEqual(object.BTIH.Bytes, expectedBtihBytes)

	// Test #2.
	object = DecodedObject{
		DecodedObject: []DictionaryItem{
			{
				Key:   []byte("no-info"),
				Value: "",
			},
		},
	}
	//
	err = object.CalculateBtih()
	//
	aTest.MustBeAnError(err)
}

func Test_GetInfoSection(t *testing.T) {
	var aTest = tester.New(t)
	var output any
	var err error

	// Test #1.
	var input = DecodedObject{
		DecodedObject: []DictionaryItem{
			{
				Key:   []byte("aaa"),
				Value: nil,
			},
			{
				Key:   []byte("bbb"),
				Value: 123,
			},
			{
				Key:   []byte("INFO"),
				Value: uint8(255),
			},
			{
				Key:   []byte("info"),
				Value: int16(101),
			},
			{
				Key:   []byte("ccc"),
				Value: "John",
			},
		},
	}
	var outputExpected any = int16(101)
	//
	output, err = input.GetInfoSection()
	//
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(output, outputExpected)

	// Test #2.
	input = DecodedObject{
		DecodedObject: []DictionaryItem{
			{
				Key:   []byte("zzz"),
				Value: nil,
			},
		},
	}
	//
	output, err = input.GetInfoSection()
	//
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(output, nil)

	// Test #3.
	input = DecodedObject{
		DecodedObject: time.Time{},
	}
	//
	output, err = input.GetInfoSection()
	//
	aTest.MustBeAnError(err)
}

func Test_MakeSelfCheck(t *testing.T) {
	var aTest = tester.New(t)

	// Test #1.
	var object = DecodedObject{
		DecodedObject: []DictionaryItem{
			{
				Key:   []byte("test"),
				Value: "Just a Test.",
			},
			{
				Key:   []byte("aux"),
				Value: "Star",
			},
		},
		SourceData: []byte(
			"d4:test12:Just a Test.3:aux4:Stare",
		),
	}
	var ok bool
	//
	ok = object.MakeSelfCheck()
	//
	aTest.MustBeEqual(ok, true)
	aTest.MustBeEqual(object.IsSelfChecked, true)

	// Test #1.
	object = DecodedObject{
		DecodedObject: time.Time{},
	}
	//
	ok = object.MakeSelfCheck()
	//
	aTest.MustBeEqual(ok, false)

	// Test #3.
	object = DecodedObject{
		DecodedObject: []DictionaryItem{
			{
				Key:   []byte("test"),
				Value: "Just a Test.",
			},
			{
				Key:   []byte("aux"),
				Value: "Star",
			},
		},
		SourceData: []byte(
			"...Corrupted Data...",
		),
	}
	//
	ok = object.MakeSelfCheck()
	//
	aTest.MustBeEqual(ok, false)
}

func Test_CalculateSha1(t *testing.T) {
	const (
		Data        string = "Just a Test."
		HashSumText string = "7b708ef0a8efed41f005c67546a9467bf612a145"
	)

	var aTest = tester.New(t)

	// Test #1.
	var (
		ba                    []byte
		data                  []byte
		err                   error
		expectedResultAsBytes Sha1Sum
		expectedResultAsText  string
		resultAsBytes         Sha1Sum
		resultAsText          string
	)
	data = []byte(Data)
	expectedResultAsText = HashSumText
	ba, err = hex.DecodeString(HashSumText)
	aTest.MustBeNoError(err)
	copy(expectedResultAsBytes[:], ba)
	//
	resultAsBytes, resultAsText = CalculateSha1(data)
	//
	aTest.MustBeEqual(resultAsText, expectedResultAsText)
	aTest.MustBeEqual(resultAsBytes, expectedResultAsBytes)
}
