// Decoder_test.go.

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

// The 'Decoder' Class.

package bencode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_Decode(t *testing.T) {
	// See Test_readBencodedValue.
}

func Test_readBencodedValue(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  interface{}
	}

	var aTest *tester.Test = tester.New(t)
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
		expectedResult: []interface{}{
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

func Test_readByteString(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  interface{}
	}

	var aTest *tester.Test = tester.New(t)
	var tests []TestData
	var byteStringMaxLen int = 100

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

	// Run the Tests.
	for i, test := range tests {
		fmt.Printf("Test #%v.\r\n", i+1)
		decoder := NewDecoder(test.reader)
		result, err := decoder.readByteString(byteStringMaxLen)
		if test.isErrorExpected {
			aTest.MustBeAnError(err)
			fmt.Println(err)
		} else {
			aTest.MustBeNoError(err)
			aTest.MustBeEqual(result, test.expectedResult)
		}
	}
}

func Test_readByteStringSizeHeader(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  interface{}
	}

	var aTest *tester.Test = tester.New(t)
	var tests []TestData
	var byteStringSizeHeaderMaxLen int = 2

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

	// Test #2. Negative: bad Symbol.
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
				"12345:qwerty",
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

	// Run the Tests.
	for i, test := range tests {
		fmt.Printf("Test #%v.\r\n", i+1)
		decoder := NewDecoder(test.reader)
		result, err := decoder.readByteStringSizeHeader(byteStringSizeHeaderMaxLen)
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

func Test_readDictionary(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  interface{}
	}

	var aTest *tester.Test = tester.New(t)
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

func Test_readDictionaryKey(t *testing.T) {
	// See Test_readByteString.
}

func Test_readDictionaryValue(t *testing.T) {
	// See Test_readBencodedValue.
}

func Test_readInteger(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  interface{}
	}

	var aTest *tester.Test = tester.New(t)
	var tests []TestData
	var integerMaxLen int = 4

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
				"12345e", // Without 'i' Prefix !
			),
		),
		isErrorExpected: true,
		expectedResult:  int64(12345),
	})

	// Run the Tests.
	for i, test := range tests {
		fmt.Printf("Test #%v.\r\n", i+1)
		decoder := NewDecoder(test.reader)
		result, err := decoder.readInteger(integerMaxLen)
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

func Test_readList(t *testing.T) {

	type TestData struct {
		reader          *bufio.Reader
		isErrorExpected bool
		expectedResult  interface{}
	}

	var aTest *tester.Test = tester.New(t)
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
		expectedResult: []interface{}{
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
		expectedResult: []interface{}{
			[]byte("Cat"),
			[]byte("Us"),
		},
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
