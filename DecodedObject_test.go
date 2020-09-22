// DecodedObject_test.go.

//============================================================================//
//
// Copyright © 2018..2020 by McArcher.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
//============================================================================//
//
// Web Sites:		'https://github.com/neverwinter-nights',
//					'https://github.com/vault-thirteen',
//					'https://github.com/legacy-vault'.
// Author:			McArcher.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

package bencode

import (
	"encoding/hex"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_CalculateBtih(t *testing.T) {

	var aTest *tester.Test = tester.New(t)
	var err error

	// Test #1.
	var object DecodedObject = DecodedObject{
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
}

func Test_CalculateSha1(t *testing.T) {

	const (
		Data        string = "Just a Test."
		HashSumText string = "7b708ef0a8efed41f005c67546a9467bf612a145"
	)

	var aTest *tester.Test = tester.New(t)

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

func Test_GetInfoSection(t *testing.T) {

	var aTest *tester.Test = tester.New(t)
	var output interface{}
	var err error

	// Test #1.
	var input DecodedObject = DecodedObject{
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
	var outputExpected interface{} = int16(101)
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
}

func Test_MakeSelfCheck(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1.
	var object DecodedObject = DecodedObject{
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
}
