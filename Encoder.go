// Encoder.go.

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
// Web Site:		'https://github.com/legacy-vault'.
// Author:			McArcher.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// The 'Encoder' Class.

package bencode

import (
	"reflect"
	"strconv"
)

// A 'bencode' Encoder.
type Encoder struct {

	// Cached Prefixes.
	commonPostfix    []byte
	dictionaryPrefix []byte
	integerPrefix    []byte
	listPrefix       []byte
}

// Encoder's Constructor.
func NewEncoder() (result *Encoder) {
	result = new(Encoder)

	result.commonPostfix = []byte{FooterCommon}
	result.dictionaryPrefix = []byte{HeaderDictionary}
	result.integerPrefix = []byte{HeaderInteger}
	result.listPrefix = []byte{HeaderList}

	return
}

// Adds a postfix to the 'bencode' Dictionary.
func (e Encoder) addPostfixOfDictionary(
	tmpResult []byte,
) (result []byte) {
	result = append(tmpResult, e.commonPostfix...)
	return
}

// Adds a Postfix to the 'bencode' List.
func (e Encoder) addPostfixOfList(
	tmpResult []byte,
) (result []byte) {
	result = append(tmpResult, e.commonPostfix...)
	return
}

// Adds a Prefix and a Postfix to the 'bencode' Byte String.
func (e Encoder) addPrefixAndPostfixOfByteString(
	tmpResult []byte,
) (result []byte) {
	result = append(
		e.createSizePrefix(
			uint64(len(tmpResult)),
		),
		result...,
	)
	return
}

// Adds a Prefix and a Postfix to the 'bencode' Integer.
func (e Encoder) addPrefixAndPostfixOfInteger(
	tmpResult []byte,
) (result []byte) {
	result = append(e.integerPrefix, tmpResult...)
	result = append(result, e.commonPostfix...)
	return
}

// Adds a Prefix to the 'bencode' Dictionary.
func (e Encoder) addPrefixOfDictionary(
	tmpResult []byte,
) (result []byte) {
	result = append(e.dictionaryPrefix, tmpResult...)
	return
}

// Adds a Prefix to the 'bencode' List.
func (e Encoder) addPrefixOfList(
	tmpResult []byte,
) (result []byte) {
	result = append(e.listPrefix, tmpResult...)
	return
}

// Creates a Size Prefix with a Delimiter.
func (e Encoder) createSizePrefix(
	size uint64,
) []byte {
	return append(
		[]byte(
			strconv.FormatUint(size, 10),
		),
		HeaderStringSizeValueDelimiter,
	)
}

// Creates an ASCII Text (Byte Array) of a signed Integer.
func (e Encoder) createTextFromInteger(
	value int64,
) []byte {
	return []byte(
		strconv.FormatInt(value, 10),
	)
}

// Creates an ASCII Text (Byte Array) of an unsigned Integer.
func (e Encoder) createTextFromUInteger(
	value uint64,
) []byte {
	return []byte(
		strconv.FormatUint(value, 10),
	)
}

// Encodes an Interface into an Array of Bytes.
func (e Encoder) EncodeAnInterface(
	ifc interface{},
) (result []byte, err error) {

	// Check an Interface's Type and encode it accordingly.
	var ifcType reflect.Kind = reflect.TypeOf(ifc).Kind()
	switch ifcType {

	case reflect.Slice:
		return e.encodeInterfaceOfSlice(ifc)

	case reflect.String:
		return e.encodeInterfaceOfString(ifc)

	case reflect.Uint:
		return e.encodeInterfaceOfUint(ifc)

	case reflect.Int:
		return e.encodeInterfaceOfInt(ifc)

	case reflect.Uint64:
		return e.encodeInterfaceOfUint64(ifc)

	case reflect.Int64:
		return e.encodeInterfaceOfInt64(ifc)

	case reflect.Uint32:
		return e.encodeInterfaceOfUint32(ifc)

	case reflect.Int32:
		return e.encodeInterfaceOfInt32(ifc)

	case reflect.Uint16:
		return e.encodeInterfaceOfUint16(ifc)

	case reflect.Int16:
		return e.encodeInterfaceOfInt16(ifc)

	case reflect.Uint8:
		return e.encodeInterfaceOfUint8(ifc)

	case reflect.Int8:
		return e.encodeInterfaceOfInt8(ifc)
	}

	// N.B.: Unfortunately, Go Language does not support generic Types,
	// so we must write a lot of similar Code doing the same Thing.

	// Unknown Type.
	return nil, ErrDataType
}

// Encodes a 'bencode' Dictionary.
func (e Encoder) encodeDictionary(
	dictionary []DictionaryItem,
) (result []byte, err error) {

	// Dictionary Prefix.
	result = e.addPrefixOfDictionary(result)

	// Add Keys and Values.
	var dictItem DictionaryItem
	for _, dictItem = range dictionary {

		var keyBA []byte
		var valueBA []byte

		// Add Key.
		keyBA, err = e.EncodeAnInterface(dictItem.Key)
		if err != nil {
			return
		}
		result = append(result, keyBA...)

		// Add Value.
		valueBA, err = e.EncodeAnInterface(dictItem.Value)
		if err != nil {
			return
		}
		result = append(result, valueBA...)
	}

	// Dictionary Postfix.
	result = e.addPostfixOfDictionary(result)
	return
}

// Encodes an int Interface as a 'bencode' Integer.
func (e Encoder) encodeInterfaceOfInt(
	intInterface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var intVar int
	var ok bool
	intVar, ok = intInterface.(int)
	if !ok {
		err = ErrTypeAssertion
		return
	}
	result = e.createTextFromInteger(int64(intVar))

	// Add Prefixes and Postfixes to the Integer.
	result = e.addPrefixAndPostfixOfInteger(result)
	return
}

// Encodes an int8 Interface as a 'bencode' Integer.
func (e Encoder) encodeInterfaceOfInt8(
	int8Interface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var int8var int8
	var ok bool
	int8var, ok = int8Interface.(int8)
	if !ok {
		err = ErrTypeAssertion
		return
	}
	result = e.createTextFromInteger(int64(int8var))

	// Add Prefixes and Postfixes to the Integer.
	result = e.addPrefixAndPostfixOfInteger(result)
	return
}

// Encodes an int16 Interface as a 'bencode' Integer.
func (e Encoder) encodeInterfaceOfInt16(
	int16Interface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var int16var int16
	var ok bool
	int16var, ok = int16Interface.(int16)
	if !ok {
		err = ErrTypeAssertion
		return
	}
	result = e.createTextFromInteger(int64(int16var))

	// Add Prefixes and Postfixes to the Integer.
	result = e.addPrefixAndPostfixOfInteger(result)
	return
}

// Encodes an int32 Interface as a 'bencode' Integer.
func (e Encoder) encodeInterfaceOfInt32(
	int32Interface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var int32var int32
	var ok bool
	int32var, ok = int32Interface.(int32)
	if !ok {
		err = ErrTypeAssertion
		return
	}
	result = e.createTextFromInteger(int64(int32var))

	// Add Prefixes and Postfixes to the Integer.
	result = e.addPrefixAndPostfixOfInteger(result)
	return
}

// Encodes an int64 Interface as a 'bencode' Integer.
func (e Encoder) encodeInterfaceOfInt64(
	int64Interface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var int64var int64
	var ok bool
	int64var, ok = int64Interface.(int64)
	if !ok {
		err = ErrTypeAssertion
		return
	}
	result = e.createTextFromInteger(int64var)

	// Add Prefixes and Postfixes to the Integer.
	result = e.addPrefixAndPostfixOfInteger(result)
	return
}

// Encodes an Interface as a 'bencode' List.
func (e Encoder) encodeInterfaceOfList(
	list []interface{},
) (result []byte, err error) {

	// List Prefix.
	result = e.addPrefixOfList(result)

	// Add Values.
	var (
		listItem interface{}
		valueBA  []byte
	)
	for _, listItem = range list {

		// Add Value.
		valueBA, err = e.EncodeAnInterface(listItem)
		if err != nil {
			return
		}
		result = append(result, valueBA...)
	}

	// List Postfix.
	result = e.addPostfixOfList(result)
	return
}

// Encodes a Slice Interface.
func (e Encoder) encodeInterfaceOfSlice(
	sliceInterface interface{},
) (result []byte, err error) {

	// Get Type of Sub-Elements.
	var ifcElementType reflect.Kind = reflect.TypeOf(sliceInterface).Elem().Kind()

	// Bytes Array ?
	if ifcElementType == reflect.Uint8 {
		return e.encodeInterfaceOfSliceOfBytes(sliceInterface)
	}

	// Try to change Type to Dictionary.
	var dictionary []DictionaryItem
	var ok bool
	dictionary, ok = sliceInterface.([]DictionaryItem)
	if ok {
		return e.encodeDictionary(dictionary)
	}

	// Try to change Type to List.
	var list []interface{}
	list, ok = sliceInterface.([]interface{})
	if ok {
		return e.encodeInterfaceOfList(list)
	}

	// Unknown Type.
	err = ErrDataType
	return
}

// Encodes a Bytes Slice Interface as a 'bencode' Byte String.
func (e Encoder) encodeInterfaceOfSliceOfBytes(
	sliceOfBytesInterface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var ok bool
	result, ok = sliceOfBytesInterface.([]byte)
	if !ok {
		err = ErrTypeAssertion
		return
	}

	// Add Prefixes and Postfixes to the Byte String.
	result = e.addPrefixAndPostfixOfByteString(result)
	return
}

// Encodes a String Interface as a 'bencode' Byte String.
func (e Encoder) encodeInterfaceOfString(
	stringInterface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var ok bool
	var stringVar string
	stringVar, ok = stringInterface.(string)
	if !ok {
		err = ErrTypeAssertion
		return
	}

	// String → Byte Array.
	result = []byte(stringVar)

	// Add Prefixes and Postfixes to the Byte String.
	result = e.addPrefixAndPostfixOfByteString(result)
	return
}

// Encodes an uint Interface as a 'bencode' Integer.
func (e Encoder) encodeInterfaceOfUint(
	uintInterface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var ok bool
	var uintVar uint
	uintVar, ok = uintInterface.(uint)
	if !ok {
		err = ErrTypeAssertion
		return
	}
	result = e.createTextFromUInteger(uint64(uintVar))

	// Add Prefixes and Postfixes to the Integer.
	result = e.addPrefixAndPostfixOfInteger(result)
	return
}

// Encodes an uint8 Interface as a 'bencode' Integer.
func (e Encoder) encodeInterfaceOfUint8(
	uint8Interface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var ok bool
	var uint8var uint8
	uint8var, ok = uint8Interface.(uint8)
	if !ok {
		err = ErrTypeAssertion
		return
	}
	result = e.createTextFromUInteger(uint64(uint8var))

	// Add Prefixes and Postfixes to the Integer.
	result = e.addPrefixAndPostfixOfInteger(result)
	return
}

// Encodes an uint16 Interface as a 'bencode' Integer.
func (e Encoder) encodeInterfaceOfUint16(
	uint16Interface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var ok bool
	var uint16var uint16
	uint16var, ok = uint16Interface.(uint16)
	if !ok {
		err = ErrTypeAssertion
		return
	}
	result = e.createTextFromUInteger(uint64(uint16var))

	// Add Prefixes and Postfixes to the Integer.
	result = e.addPrefixAndPostfixOfInteger(result)
	return
}

// Encodes an uint32 Interface as a 'bencode' Integer.
func (e Encoder) encodeInterfaceOfUint32(
	uint32Interface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var ok bool
	var uint32var uint32
	uint32var, ok = uint32Interface.(uint32)
	if !ok {
		err = ErrTypeAssertion
		return
	}
	result = e.createTextFromUInteger(uint64(uint32var))

	// Add Prefixes and Postfixes to the Integer.
	result = e.addPrefixAndPostfixOfInteger(result)
	return
}

// Encodes an uint64 Interface as a 'bencode' Integer.
func (e Encoder) encodeInterfaceOfUint64(
	uint64Interface interface{},
) (result []byte, err error) {

	// Convert the Type.
	var ok bool
	var uint64var uint64
	uint64var, ok = uint64Interface.(uint64)
	if !ok {
		err = ErrTypeAssertion
		return
	}
	result = e.createTextFromUInteger(uint64var)

	// Add Prefixes and Postfixes to the Integer.
	result = e.addPrefixAndPostfixOfInteger(result)
	return
}
