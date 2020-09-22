// DecodedObject.go.

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

// The 'DecodedObject' Class.

package bencode

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
)

// A decoded Object with some Meta-Data.
type DecodedObject struct {

	// Primary Information.
	FilePath        string
	SourceData      []byte
	DecodedObject   interface{}
	DecodeTimestamp int64

	// Secondary Information.
	IsSelfChecked bool
	BTIH          BtihData
}

// Calculates the BitTorrent Info Hash (BTIH) Check Sum.
func (do *DecodedObject) CalculateBtih() (err error) {

	// Get the 'info' Section from the decoded Object.
	var infoSection interface{}
	infoSection, err = do.GetInfoSection()
	if err != nil {
		return
	}

	// Encode the 'info' Section.
	var encoder *Encoder = NewEncoder()
	var infoSectionBA []byte
	infoSectionBA, err = encoder.EncodeAnInterface(infoSection)
	if err != nil {
		return
	}

	// Calculate BTIH.
	do.BTIH.Bytes, do.BTIH.Text = CalculateSha1(infoSectionBA)
	return
}

// Calculates SHA-1 Check Sum and
// returns it as a Hexadecimal Text and Byte Array.
func CalculateSha1(
	data []byte,
) (resultAsBytes Sha1Sum, resultAsText string) {
	resultAsBytes = sha1.Sum(data)
	resultAsText = hex.EncodeToString(resultAsBytes[:])
	return
}

// Gets an 'info' Section from the Object.
func (do DecodedObject) GetInfoSection() (result interface{}, err error) {

	// Get Dictionary.
	var dictionary []DictionaryItem
	var ok bool
	dictionary, ok = do.DecodedObject.([]DictionaryItem)
	if !ok {
		err = ErrTypeAssertion
		return
	}

	// Get the 'info' Section from the decoded Object.
	var dictItem DictionaryItem
	for _, dictItem = range dictionary {

		if string(dictItem.Key) == FileSectionInfo {
			result = dictItem.Value
			return
		}
	}

	err = ErrSectionDoesNotExist
	return
}

// A simple Self-Check.
// Encodes the decoded Data and compares with the Source.
func (do *DecodedObject) MakeSelfCheck() (success bool) {

	// Encode decoded Data.
	var encoder *Encoder = NewEncoder()
	var err error
	var baEncoded []byte
	baEncoded, err = encoder.EncodeAnInterface(do.DecodedObject)
	if err != nil {
		return
	}

	// Compare encoded decoded Data with original Data.
	var checkResult int
	checkResult = bytes.Compare(baEncoded, do.SourceData)
	if checkResult != 0 {
		return
	}

	do.IsSelfChecked = true
	success = true
	return
}
