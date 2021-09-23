// errors_test.go.

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

import (
	"errors"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_combineErrors(t *testing.T) {

	type TestData struct {
		e1                    error
		e2                    error
		expectedCombinedError error
	}

	var aTest = tester.New(t)
	var tests []TestData

	// Test #1.
	tests = append(tests, TestData{
		e1:                    nil,
		e2:                    nil,
		expectedCombinedError: nil,
	})

	// Test #2.
	tests = append(tests, TestData{
		e1:                    nil,
		e2:                    errors.New("qwe"),
		expectedCombinedError: errors.New("qwe"),
	})

	// Test #3.
	tests = append(tests, TestData{
		e1:                    errors.New("qwe"),
		e2:                    nil,
		expectedCombinedError: errors.New("qwe"),
	})

	// Test #4.
	tests = append(tests, TestData{
		e1:                    errors.New("aaa"),
		e2:                    errors.New("bbb"),
		expectedCombinedError: errors.New("aaa: bbb"),
	})

	// Run the Tests.
	for _, test := range tests {
		result := combineErrors(test.e1, test.e2)
		aTest.MustBeEqual(result, test.expectedCombinedError)
	}
}
