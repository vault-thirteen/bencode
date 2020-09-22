// Decoder.go.

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

// The 'Decoder' Class.

package bencode

import (
	"bufio"
	"bytes"
	"fmt"
)

//	1.	Parser's Settings.

//	1.1.	Integer Size (Number of ASCII Letters allowed).
// N.B.: Maximum Value of UInt64 is '18446744073709551615'.
const IntegerMaxLength = 20

//	1.2.	Byte String Size Header (Number of ASCII Letters allowed).
// N.B.: We are not going to read Byte Strings which have Length more than that.
const ByteStringMaxLength = IntegerMaxLength

// A 'bencode' Decoder.
type Decoder struct {
	reader *bufio.Reader
}

// Decoder's Constructor.
func NewDecoder(
	reader *bufio.Reader,
) (result *Decoder) {
	result = new(Decoder)

	result.reader = reader

	return
}

// Decodes a 'bencoded' Byte Stream into an Interface.
func (d Decoder) Decode() (result interface{}, err error) {
	return d.readBencodedValue()
}

// Reads a raw "bencoded" Value, including its Sub-values.
func (d Decoder) readBencodedValue() (result interface{}, err error) {

	// Get the first Byte from Stream to know its Type.
	var b byte
	b, err = d.reader.ReadByte()
	if err != nil {
		return
	}

	// Analyze the Type.
	if b == HeaderDictionary {
		return d.readDictionary()

	} else if b == HeaderList {
		return d.readList()

	} else if b == HeaderInteger {
		return d.readInteger(IntegerMaxLength)

	} else if isByteNonNegativeAsciiNumeric(b) {
		// It must be an ASCII Number indicating a Byte String.
		// => Byte String.

		// Rewind the Cursor back, as it does not have a Type-Prefix!
		// The 'bencode' Encoding is ugly...
		err = d.reader.UnreadByte()
		if err != nil {
			return nil, err
		}

		// Read the Byte String.
		return d.readByteString(ByteStringMaxLength)
	}

	// Otherwise, it is a Syntax Error.
	var errorArea []byte = []byte{b}
	err = fmt.Errorf(ErrFmtSyntaxErrorAt, errorArea)
	return nil, err
}

// Reads a Byte String from the Stream (Reader).
func (d Decoder) readByteString(
	byteStringMaxLen int,
) (ba []byte, err error) {

	// Read the Size Header and verify it.
	var byteStringLen uint
	var byteStringSizeHeaderMaxLen int = getByteStringSizeHeaderMaxLen(uint(byteStringMaxLen))
	byteStringLen, err = d.readByteStringSizeHeader(byteStringSizeHeaderMaxLen)
	if err != nil {
		return
	}

	// Now we should read the Byte String.
	var b byte
	var i uint = 0
	var bytesAccumulator *bytes.Buffer
	bytesAccumulator = bytes.NewBuffer([]byte{})
	for i < byteStringLen {

		// Read next Symbol.
		b, err = d.reader.ReadByte()
		if err != nil {
			return
		}

		// Save the Byte to the Accumulator.
		err = bytesAccumulator.WriteByte(b)
		if err != nil {
			return
		}

		i++
	}

	ba = bytesAccumulator.Bytes()
	return
}

// Reads the Size Header of a Byte String from the Stream (Reader) and
// converts its Value into an Integer.
func (d Decoder) readByteStringSizeHeader(
	byteStringSizeHeaderMaxLen int,
) (byteStringLen uint, err error) {

	// Read the first Byte.
	var b byte
	b, err = d.reader.ReadByte()
	if err != nil {
		return
	}

	var sizeHeader []byte
	for b != HeaderStringSizeValueDelimiter {

		// Syntax Check.
		if !isByteNonNegativeAsciiNumeric(b) {
			var errorArea []byte = append(sizeHeader, []byte{b}...)
			err = fmt.Errorf(ErrFmtSyntaxErrorAt, errorArea)
			return
		}

		// Save Byte to Size Header.
		if len(sizeHeader) < byteStringSizeHeaderMaxLen {
			sizeHeader = append(sizeHeader, b)
		} else {
			// The Length Header is too big!
			err = ErrHeaderLength
			return
		}

		// Read next Byte.
		b, err = d.reader.ReadByte()
		if err != nil {
			return
		}
	}

	// Check the Header's Length.
	if len(sizeHeader) == 0 {
		var errorArea []byte = append(sizeHeader, []byte{b}...)
		err = fmt.Errorf(ErrFmtSyntaxErrorAt, errorArea)
		return
	}

	// Convert the Size Header into normal integer Size Value.
	var byteStringLenUint64 uint64
	byteStringLenUint64, err = convertByteStringToNonNegativeInteger(sizeHeader)
	if err != nil {
		return
	}
	byteStringLen = uint(byteStringLenUint64)
	return
}

// Reads a Dictionary.
// We suppose that the Header of Dictionary ('d')
// has already been read from the Stream.
func (d Decoder) readDictionary() (result interface{}, err error) {

	// Prepare Data.
	var dictionary []DictionaryItem
	dictionary = make([]DictionaryItem, 0)

	// Probe the Next Byte to check the End of Dictionary.
	var b byte
	b, err = d.reader.ReadByte()
	if err != nil {
		return
	}
	for b != FooterCommon {

		// That single Byte (we probed) was not an End!
		// We must get back, rewind that Byte.
		err = d.reader.UnreadByte()
		if err != nil {
			return
		}

		// Get the Key.
		var dictKey []byte
		dictKey, err = d.readDictionaryKey()
		if err != nil {
			return
		}

		// Get the Value.
		var dictValue interface{}
		dictValue, err = d.readDictionaryValue()
		if err != nil {
			return
		}

		// Save Item into Dictionary.
		dictionary = append(
			dictionary,
			DictionaryItem{
				// System Fields.
				Key:   dictKey,
				Value: dictValue,

				// Additional Fields for special Purposes.
				KeyStr:   string(dictKey),
				ValueStr: convertInterfaceToString(dictValue),
			},
		)

		// Probe the Next Byte to check the End of Dictionary.
		b, err = d.reader.ReadByte()
		if err != nil {
			return
		}
	}

	result = dictionary
	return
}

// Reads a Dictionary's Key.
func (d Decoder) readDictionaryKey() ([]byte, error) {
	return d.readByteString(ByteStringMaxLength)
}

// Reads a Dictionary's Value.
func (d Decoder) readDictionaryValue() (interface{}, error) {
	return d.readBencodedValue()
}

// Reads an Integer from the Stream (Reader).
// We suppose that the Header of Integer ('i')
// has already been read from the Stream.
func (d Decoder) readInteger(
	integerMaxLen int,
) (value int64, err error) {

	// Prepare Data.
	var valueBA []byte = []byte{}

	// Read the first Byte.
	var b byte
	b, err = d.reader.ReadByte()
	if err != nil {
		return
	}

	for b != FooterCommon {

		// Syntax Check.
		if !isByteAsciiNumeric(b) {
			var errorArea []byte = append(valueBA, []byte{b}...)
			err = fmt.Errorf(ErrFmtSyntaxErrorAt, errorArea)
			return
		}

		// Save Byte to Value Byte Array.
		if len(valueBA) < integerMaxLen {
			valueBA = append(valueBA, b)
		} else {
			// The Integer is too big!
			err = ErrIntegerLength
			return
		}

		// Read next Byte.
		b, err = d.reader.ReadByte()
		if err != nil {
			return
		}
	}

	// We have read the Value.
	// Check that it is not empty.
	if len(valueBA) == 0 {
		var errorArea []byte = append(valueBA, []byte{b}...)
		err = fmt.Errorf(ErrFmtSyntaxErrorAt, errorArea)
		return
	}

	// Convert Value into normal integer Value.
	return convertByteStringToInteger(valueBA)
}

// Reads a List from the Stream (Reader).
// We suppose that the Header of List ('l')
// has already been read from the Stream.
func (d Decoder) readList() (list []interface{}, err error) {

	// Prepare Data.
	list = make([]interface{}, 0)

	// Probe the Next Byte to check the End of List.
	var b byte
	b, err = d.reader.ReadByte()
	if err != nil {
		return
	}
	for b != FooterCommon {

		// That single Byte (we probed) was not an End!
		// We must get back, rewind that Byte.
		err = d.reader.UnreadByte()
		if err != nil {
			return
		}

		// Get the Item.
		var listItem interface{}
		listItem, err = d.readBencodedValue()
		if err != nil {
			return
		}

		// Save Item into Dictionary.
		list = append(list, listItem)

		// Probe the Next Byte to check the End of List.
		b, err = d.reader.ReadByte()
		if err != nil {
			return
		}
	}

	return
}
