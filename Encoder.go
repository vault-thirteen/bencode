package bencode

import (
	"errors"
	"reflect"
	"strconv"
)

// Encoder is a 'bencode' encoder.
type Encoder struct {

	// Cached prefixes.
	commonPostfix    []byte
	dictionaryPrefix []byte
	integerPrefix    []byte
	listPrefix       []byte
}

// NewEncoder is an encoder's constructor.
func NewEncoder() (e *Encoder) {
	e = &Encoder{
		commonPostfix:    []byte{FooterCommon},
		dictionaryPrefix: []byte{HeaderDictionary},
		integerPrefix:    []byte{HeaderInteger},
		listPrefix:       []byte{HeaderList},
	}

	return e
}

// addPostfixOfDictionary adds a postfix to the 'bencode' dictionary.
func (e Encoder) addPostfixOfDictionary(tmpResult []byte) (result []byte) {
	return append(tmpResult, e.commonPostfix...)
}

// addPostfixOfList adds a Postfix to the 'bencode' list.
func (e Encoder) addPostfixOfList(tmpResult []byte) (result []byte) {
	return append(tmpResult, e.commonPostfix...)
}

// addPrefixAndPostfixOfByteString adds a prefix and a postfix to the 'bencode'
// byte string.
func (e Encoder) addPrefixAndPostfixOfByteString(tmpResult []byte) (result []byte) {
	return append(
		e.createSizePrefix(
			uint64(len(tmpResult)),
		),
		tmpResult...,
	)
}

// addPrefixAndPostfixOfInteger adds a prefix and a postfix to the 'bencode'
// integer.
func (e Encoder) addPrefixAndPostfixOfInteger(tmpResult []byte) (result []byte) {
	return append(
		append(e.integerPrefix, tmpResult...),
		e.commonPostfix...,
	)
}

// addPrefixOfDictionary adds a prefix to the 'bencode' dictionary.
func (e Encoder) addPrefixOfDictionary(tmpResult []byte) (result []byte) {
	return append(e.dictionaryPrefix, tmpResult...)
}

// addPrefixOfList adds a prefix to the 'bencode' list.
func (e Encoder) addPrefixOfList(tmpResult []byte) (result []byte) {
	return append(e.listPrefix, tmpResult...)
}

// createSizePrefix creates a size prefix with a delimiter.
func (e Encoder) createSizePrefix(size uint64) []byte {
	return append(
		[]byte(
			strconv.FormatUint(size, 10),
		),
		HeaderStringSizeValueDelimiter,
	)
}

// createTextFromInteger creates an ASCII text (byte array) of a signed integer.
func (e Encoder) createTextFromInteger(value int64) []byte {
	return []byte(
		strconv.FormatInt(value, 10),
	)
}

// createTextFromUInteger creates an ASCII text (byte array) of an unsigned
// integer.
func (e Encoder) createTextFromUInteger(value uint64) []byte {
	return []byte(
		strconv.FormatUint(value, 10),
	)
}

// EncodeAnInterface encodes an interface into an array of bytes.
func (e Encoder) EncodeAnInterface(ifc any) (result []byte, err error) {

	// Check the interface's type and encode it accordingly.
	var ifcType = reflect.TypeOf(ifc).Kind()
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

	// N.B.: Unfortunately, Go language does not support generic types,
	// so we must write a lot of similar code doing the same thing.

	// Unknown type.
	return nil, errors.New(ErrDataType)
}

// encodeDictionary encodes a 'bencode' dictionary.
func (e Encoder) encodeDictionary(dictionary []DictionaryItem) (result []byte, err error) {

	// Dictionary prefix.
	result = e.addPrefixOfDictionary(result)

	// Add keys and values.
	var dictItem DictionaryItem
	for _, dictItem = range dictionary {

		var keyBA []byte
		var valueBA []byte

		// Add the key.
		keyBA, err = e.EncodeAnInterface(dictItem.Key)
		if err != nil {
			return nil, err
		}

		result = append(result, keyBA...)

		// Add the value.
		valueBA, err = e.EncodeAnInterface(dictItem.Value)
		if err != nil {
			return nil, err
		}

		result = append(result, valueBA...)
	}

	// Dictionary postfix.
	return e.addPostfixOfDictionary(result), nil
}

// encodeInterfaceOfInt encodes an int interface as a 'bencode' integer.
func (e Encoder) encodeInterfaceOfInt(intInterface any) (result []byte, err error) {

	// Convert the type.
	var intVar int
	var ok bool
	intVar, ok = intInterface.(int)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	result = e.createTextFromInteger(int64(intVar))

	// Add a prefix and a postfix to the integer.
	return e.addPrefixAndPostfixOfInteger(result), nil
}

// encodeInterfaceOfInt8 encodes an int8 interface as a 'bencode' integer.
func (e Encoder) encodeInterfaceOfInt8(int8Interface any) (result []byte, err error) {

	// Convert the type.
	var int8var int8
	var ok bool
	int8var, ok = int8Interface.(int8)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	result = e.createTextFromInteger(int64(int8var))

	// Add a prefix and a postfix to the integer.
	return e.addPrefixAndPostfixOfInteger(result), nil
}

// encodeInterfaceOfInt16 encodes an int16 interface as a 'bencode' integer.
func (e Encoder) encodeInterfaceOfInt16(int16Interface any) (result []byte, err error) {

	// Convert the type.
	var int16var int16
	var ok bool
	int16var, ok = int16Interface.(int16)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	result = e.createTextFromInteger(int64(int16var))

	// Add a prefix and a postfix to the integer.
	return e.addPrefixAndPostfixOfInteger(result), nil
}

// encodeInterfaceOfInt32 encodes an int32 interface as a 'bencode' integer.
func (e Encoder) encodeInterfaceOfInt32(int32Interface any) (result []byte, err error) {

	// Convert the type.
	var int32var int32
	var ok bool
	int32var, ok = int32Interface.(int32)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	result = e.createTextFromInteger(int64(int32var))

	// Add a prefix and a postfix to the integer.
	return e.addPrefixAndPostfixOfInteger(result), nil
}

// encodeInterfaceOfInt64 encodes an int64 interface as a 'bencode' integer.
func (e Encoder) encodeInterfaceOfInt64(int64Interface any) (result []byte, err error) {

	// Convert the type.
	var int64var int64
	var ok bool
	int64var, ok = int64Interface.(int64)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	result = e.createTextFromInteger(int64var)

	// Add a prefix and a postfix to the integer.
	return e.addPrefixAndPostfixOfInteger(result), nil
}

// encodeInterfaceOfList encodes an interface as a 'bencode' list.
func (e Encoder) encodeInterfaceOfList(list []any) (result []byte, err error) {

	// List prefix.
	result = e.addPrefixOfList(result)

	// Add the values.
	var (
		listItem any
		valueBA  []byte
	)
	for _, listItem = range list {

		// Add the value.
		valueBA, err = e.EncodeAnInterface(listItem)
		if err != nil {
			return nil, err
		}

		result = append(result, valueBA...)
	}

	// List postfix.
	return e.addPostfixOfList(result), nil
}

// encodeInterfaceOfSlice encodes a slice interface.
func (e Encoder) encodeInterfaceOfSlice(sliceInterface any) (result []byte, err error) {

	// Get the type of sub-elements.
	var ifcElementType = reflect.TypeOf(sliceInterface).Elem().Kind()

	// Bytes array ?
	if ifcElementType == reflect.Uint8 {
		return e.encodeInterfaceOfSliceOfBytes(sliceInterface)
	}

	// Try to change the type to dictionary.
	var dictionary []DictionaryItem
	var ok bool
	dictionary, ok = sliceInterface.([]DictionaryItem)
	if ok {
		return e.encodeDictionary(dictionary)
	}

	// Try to change the type to list.
	var list []any
	list, ok = sliceInterface.([]any)
	if ok {
		return e.encodeInterfaceOfList(list)
	}

	// Unknown type.
	return nil, errors.New(ErrDataType)
}

// encodeInterfaceOfSliceOfBytes encodes a bytes slice interface as a 'bencode'
// byte string.
func (e Encoder) encodeInterfaceOfSliceOfBytes(sliceOfBytesInterface any) (result []byte, err error) {

	// Convert the type.
	var ok bool
	result, ok = sliceOfBytesInterface.([]byte)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	// Add a prefix and a postfix to the byte string.
	return e.addPrefixAndPostfixOfByteString(result), nil
}

// encodeInterfaceOfString encodes a string interface as a 'bencode' byte
// string.
func (e Encoder) encodeInterfaceOfString(stringInterface any) (result []byte, err error) {

	// Convert the type.
	var ok bool
	var stringVar string
	stringVar, ok = stringInterface.(string)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	// String → Byte array.
	result = []byte(stringVar)

	// Add a prefix and a postfix to the byte string.
	return e.addPrefixAndPostfixOfByteString(result), nil
}

// encodeInterfaceOfUint encodes an uint interface as a 'bencode' integer.
func (e Encoder) encodeInterfaceOfUint(uintInterface any) (result []byte, err error) {

	// Convert the type.
	var ok bool
	var uintVar uint
	uintVar, ok = uintInterface.(uint)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	result = e.createTextFromUInteger(uint64(uintVar))

	// Add a prefix and a postfix to the integer.
	return e.addPrefixAndPostfixOfInteger(result), nil
}

// encodeInterfaceOfUint8 encodes an uint8 interface as a 'bencode' integer.
func (e Encoder) encodeInterfaceOfUint8(uint8Interface any) (result []byte, err error) {

	// Convert the type.
	var ok bool
	var uint8var uint8
	uint8var, ok = uint8Interface.(uint8)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	result = e.createTextFromUInteger(uint64(uint8var))

	// Add a prefix and a postfix to the integer.
	return e.addPrefixAndPostfixOfInteger(result), nil
}

// encodeInterfaceOfUint16 encodes an uint16 interface as a 'bencode' integer.
func (e Encoder) encodeInterfaceOfUint16(uint16Interface any) (result []byte, err error) {

	// Convert the type.
	var ok bool
	var uint16var uint16
	uint16var, ok = uint16Interface.(uint16)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	result = e.createTextFromUInteger(uint64(uint16var))

	// Add a prefix and a postfix to the integer.
	return e.addPrefixAndPostfixOfInteger(result), nil
}

// encodeInterfaceOfUint32 encodes an uint32 interface as a 'bencode' integer.
func (e Encoder) encodeInterfaceOfUint32(uint32Interface any) (result []byte, err error) {

	// Convert the type.
	var ok bool
	var uint32var uint32
	uint32var, ok = uint32Interface.(uint32)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	result = e.createTextFromUInteger(uint64(uint32var))

	// Add a prefix and a postfix to the integer.
	return e.addPrefixAndPostfixOfInteger(result), nil
}

// encodeInterfaceOfUint64 encodes an uint64 interface as a 'bencode' integer.
func (e Encoder) encodeInterfaceOfUint64(uint64Interface any) (result []byte, err error) {

	// Convert the type.
	var ok bool
	var uint64var uint64
	uint64var, ok = uint64Interface.(uint64)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	result = e.createTextFromUInteger(uint64var)

	// Add a prefix and a postfix to the integer.
	return e.addPrefixAndPostfixOfInteger(result), nil
}
