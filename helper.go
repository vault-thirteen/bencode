// helper.go.

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

// Some auxiliary Helper-Functions.

package bencode

import (
	"reflect"
	"strconv"
)

// Converts a Byte String into a signed 64-Bit Integer.
// Negative Numbers are possible.
func convertByteStringToInteger(
	ba []byte,
) (result int64, err error) {
	if len(ba) > ByteStringMaxLength {
		err = ErrByteStringToInt
		return
	}
	return strconv.ParseInt(string(ba), 10, 64)
}

// Converts a Byte String into an unsigned 64-Bit Integer.
// Negative Numbers are forbidden.
func convertByteStringToNonNegativeInteger(
	ba []byte,
) (result uint64, err error) {
	if len(ba) > ByteStringMaxLength {
		err = ErrByteStringToInt
		return
	}
	return strconv.ParseUint(string(ba), 10, 64)
}

// Tries to get a textual Data from an Interface.
func convertInterfaceToString(
	src interface{},
) (result string) {

	// Slice?
	var srcType reflect.Kind = reflect.TypeOf(src).Kind()
	if srcType == reflect.Slice {

		// Array Item's Type is Byte?
		var srcElementType reflect.Kind = reflect.TypeOf(src).Elem().Kind()
		if srcElementType == reflect.Uint8 {
			var bytes []byte
			var ok bool
			bytes, ok = src.([]byte)
			if !ok {
				return
			}
			result = string(bytes)
			return
		}
	}
	return
}

// Calculates how many numeric Symbols are required to write the Size Prefix.
func getByteStringSizeHeaderMaxLen(
	byteStringLen uint,
) (byteStringSizeHeaderMaxLen int) {
	s := strconv.FormatUint(uint64(byteStringLen), 10)
	byteStringSizeHeaderMaxLen = len([]rune(s))
	return
}

// Checks whether the Byte is ASCII numeric Symbol.
// Negative Numbers are possible.
func isByteAsciiNumeric(
	b byte,
) (result bool) {

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

// Checks whether the Byte is ASCII non-negative numeric Symbol.
// Negative Numbers are forbidden.
func isByteNonNegativeAsciiNumeric(
	b byte,
) (result bool) {

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
