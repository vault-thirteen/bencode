// helper_test.go.

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
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_convertByteStringToInteger(t *testing.T) {
	var aTest *tester.Test = tester.New(t)

	// Test #1. Negative.
	var bytes []byte = []byte("123456789012345678912345")
	var result int64
	var err error
	result, err = convertByteStringToInteger(bytes)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, int64(0))

	// Test #2. Positive.
	bytes = []byte("-12345")
	result, err = convertByteStringToInteger(bytes)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, int64(-12345))
}

func Test_convertByteStringToNonNegativeInteger(t *testing.T) {
	var aTest *tester.Test = tester.New(t)

	// Test #1. Negative.
	var bytes []byte = []byte("123456789012345678912345")
	var result uint64
	var err error
	result, err = convertByteStringToNonNegativeInteger(bytes)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, uint64(0))

	// Test #2. Positive.
	bytes = []byte("12345")
	result, err = convertByteStringToNonNegativeInteger(bytes)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, uint64(12345))
}

func Test_convertInterfaceToString(t *testing.T) {
	var aTest *tester.Test = tester.New(t)

	// Test #1. Negative: Not a Slice.
	var result string
	result = convertInterfaceToString("Error")
	aTest.MustBeEqual(result, "")

	// Test #2. Negative: Not a Slice of Bytes.
	result = convertInterfaceToString([]rune("xyz"))
	aTest.MustBeEqual(result, "")

	// Test #3. Positive.
	result = convertInterfaceToString([]byte("Text"))
	aTest.MustBeEqual(result, "Text")
}

func Test_isByteAsciiNumeric(t *testing.T) {
	var aTest *tester.Test = tester.New(t)

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
	//
	aTest.MustBeEqual(isByteAsciiNumeric('x'), false)
}

func Test_isByteNonNegativeAsciiNumeric(t *testing.T) {
	var aTest *tester.Test = tester.New(t)

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
	//
	aTest.MustBeEqual(isByteNonNegativeAsciiNumeric('x'), false)
}
