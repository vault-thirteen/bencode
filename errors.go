// errors.go.

//============================================================================//
//
// Copyright © 2018..2021 by McArcher.
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

import "errors"

// Error messages and formats.
const (
	ErrByteStringToIntError     = "byte string to integer conversion error"
	ErrDataTypeError            = "unsupported type"
	ErrFileNotInitializedError  = "file is not initialized"
	ErrHeaderLengthError        = "the length header is too big: %v"
	ErrSectionDoesNotExistError = "section does not exist"
	ErrSelfCheckError           = "self-check error"
	ErrTypeAssertionError       = "type assertion error"

	ErrFIntegerLengthError = "the integer is too big: %v"
)

// Cached errors.
var (
	ErrByteStringToInt     = errors.New(ErrByteStringToIntError)
	ErrDataType            = errors.New(ErrDataTypeError)
	ErrFileNotInitialized  = errors.New(ErrFileNotInitializedError)
	ErrSectionDoesNotExist = errors.New(ErrSectionDoesNotExistError)
	ErrSelfCheck           = errors.New(ErrSelfCheckError)
	ErrTypeAssertion       = errors.New(ErrTypeAssertionError)
)

// Formats of error messages.
const (
	ErrFSyntaxErrorAt = "syntax error at: '%v'"
	ErrCombinator     = ": "
)

// combineErrors combines two errors.
func combineErrors(
	error1 error,
	error2 error,
) (result error) {
	if error1 == nil {
		if error2 == nil {
			return nil
		} else {
			return error2
		}
	} else {
		if error2 == nil {
			return error1
		} else {
			return errors.New(error1.Error() + ErrCombinator + error2.Error())
		}
	}
}
