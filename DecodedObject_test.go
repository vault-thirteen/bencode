package bencode

import (
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_DecodedObject_MakeSelfCheck(t *testing.T) {
	var aTest = tester.New(t)
	var object DecodedObject
	var ok bool

	// Test #1.
	{
		object = DecodedObject{
			RawObject: []DictionaryItem{
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

		ok = object.MakeSelfCheck()
		aTest.MustBeEqual(ok, true)
		aTest.MustBeEqual(object.IsSelfChecked, true)
	}

	// Test #2.
	{
		object = DecodedObject{
			RawObject: time.Time{},
		}

		ok = object.MakeSelfCheck()
		aTest.MustBeEqual(ok, false)
	}

	// Test #3.
	{
		object = DecodedObject{
			RawObject: []DictionaryItem{
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

		ok = object.MakeSelfCheck()
		aTest.MustBeEqual(ok, false)
	}
}
