// Encoder_test.go.

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
// Web Site:		'https://github.com/neverwinter-nights'.
// Author:			McArcher.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

package bencode

import (
	"fmt"
	"testing"
	"time"

	"github.com/vault-thirteen/tester"
)

func Test_addPostfixOfDictionary(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("abc")
	result := encoder.addPostfixOfDictionary(tmpResult)
	resultExpected := []byte("abce")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_addPostfixOfList(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("lt")
	result := encoder.addPostfixOfList(tmpResult)
	resultExpected := []byte("lte")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_addPrefixAndPostfixOfByteString(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("bst")
	result := encoder.addPrefixAndPostfixOfByteString(tmpResult)
	resultExpected := []byte("3:bst")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_addPrefixAndPostfixOfInteger(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("909")
	result := encoder.addPrefixAndPostfixOfInteger(tmpResult)
	resultExpected := []byte("i909e")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_addPrefixOfDictionary(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("xx")
	result := encoder.addPrefixOfDictionary(tmpResult)
	resultExpected := []byte("dxx")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_addPrefixOfList(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	encoder := NewEncoder()
	var tmpResult = []byte("nn")
	result := encoder.addPrefixOfList(tmpResult)
	resultExpected := []byte("lnn")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_createSizePrefix(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	encoder := NewEncoder()
	result := encoder.createSizePrefix(3)
	resultExpected := []byte("3:")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_createTextFromInteger(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	encoder := NewEncoder()
	result := encoder.createTextFromInteger(-56)
	resultExpected := []byte("-56")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_createTextFromUInteger(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	encoder := NewEncoder()
	result := encoder.createTextFromUInteger(56)
	resultExpected := []byte("56")
	aTest.MustBeEqual(result, resultExpected)
}

func Test_EncodeAnInterface(t *testing.T) {

	type TestData struct {
		dataToBeEncoded interface{}
		isErrorExpected bool
		expectedResult  interface{}
	}

	var aTest *tester.Test = tester.New(t)
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
		dataToBeEncoded: int(-124),
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

func Test_encodeDictionary(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1.
	encoder := NewEncoder()
	var dictionary []DictionaryItem = []DictionaryItem{
		{
			Key:      []byte("Aa"),
			Value:    int(123),
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
	var resultExpected = []byte("d2:Aai123e2:Bb6:QWERTYe")
	result, err := encoder.encodeDictionary(dictionary)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
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

func Test_encodeInterfaceOfInt(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = int(123)
	var resultExpected = []byte("i123e")
	result, err := encoder.encodeInterfaceOfInt(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = uint(123)
	result, err = encoder.encodeInterfaceOfInt(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfInt8(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = int8(127)
	var resultExpected = []byte("i127e")
	result, err := encoder.encodeInterfaceOfInt8(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = "qqq"
	result, err = encoder.encodeInterfaceOfInt8(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfInt16(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = int16(127)
	var resultExpected = []byte("i127e")
	result, err := encoder.encodeInterfaceOfInt16(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = "qqq"
	result, err = encoder.encodeInterfaceOfInt16(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfInt32(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = int32(127)
	var resultExpected = []byte("i127e")
	result, err := encoder.encodeInterfaceOfInt32(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = "qqq"
	result, err = encoder.encodeInterfaceOfInt32(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfInt64(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = int64(127)
	var resultExpected = []byte("i127e")
	result, err := encoder.encodeInterfaceOfInt64(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = "qqq"
	result, err = encoder.encodeInterfaceOfInt64(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfList(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data []interface{} = []interface{}{
		int8(123),
		"Qwe",
	}
	var resultExpected = []byte("li123e3:Qwee")
	result, err := encoder.encodeInterfaceOfList(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = []interface{}{
		int8(123),
		time.Time{},
	}
	result, err = encoder.encodeInterfaceOfList(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfSlice(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Slice of Bytes.
	encoder := NewEncoder()
	var data interface{} = []byte("ABC")
	var resultExpected = []byte("3:ABC")
	result, err := encoder.encodeInterfaceOfSlice(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Dictionary
	data = []DictionaryItem{
		{
			Key:      []byte("Aa"),
			Value:    int(123),
			KeyStr:   "Aa",
			ValueStr: "123",
		},
	}
	resultExpected = []byte("d2:Aai123ee")
	result, err = encoder.encodeInterfaceOfSlice(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #3. Slice of Interfaces.
	data = []interface{}{
		"Qwerty",
		uint16(6565),
	}
	resultExpected = []byte("l6:Qwertyi6565ee")
	result, err = encoder.encodeInterfaceOfSlice(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #4. unknown Type.
	data = []time.Time{}
	result, err = encoder.encodeInterfaceOfSlice(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfSliceOfBytes(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = []byte("Qwe")
	var resultExpected = []byte("3:Qwe")
	result, err := encoder.encodeInterfaceOfSliceOfBytes(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = 123
	result, err = encoder.encodeInterfaceOfSliceOfBytes(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfString(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = "Abc"
	var resultExpected = []byte("3:Abc")
	result, err := encoder.encodeInterfaceOfString(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = 123
	result, err = encoder.encodeInterfaceOfString(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfUint(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = uint(123)
	var resultExpected = []byte("i123e")
	result, err := encoder.encodeInterfaceOfUint(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = "x"
	result, err = encoder.encodeInterfaceOfUint(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfUint8(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = uint8(123)
	var resultExpected = []byte("i123e")
	result, err := encoder.encodeInterfaceOfUint8(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = "x"
	result, err = encoder.encodeInterfaceOfUint8(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfUint16(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = uint16(123)
	var resultExpected = []byte("i123e")
	result, err := encoder.encodeInterfaceOfUint16(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = "x"
	result, err = encoder.encodeInterfaceOfUint16(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfUint32(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = uint32(123)
	var resultExpected = []byte("i123e")
	result, err := encoder.encodeInterfaceOfUint32(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = "x"
	result, err = encoder.encodeInterfaceOfUint32(data)
	aTest.MustBeAnError(err)
}

func Test_encodeInterfaceOfUint64(t *testing.T) {

	var aTest *tester.Test = tester.New(t)

	// Test #1. Positive.
	encoder := NewEncoder()
	var data interface{} = uint64(123)
	var resultExpected = []byte("i123e")
	result, err := encoder.encodeInterfaceOfUint64(data)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Negative.
	data = "x"
	result, err = encoder.encodeInterfaceOfUint64(data)
	aTest.MustBeAnError(err)
}
