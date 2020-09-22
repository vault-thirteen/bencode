// architecture.go.

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

// Architecture specific Methods.

package bencode

import (
	"math"
	"strconv"
)

const (
	ArchitectureIs64Bit bool = (strconv.IntSize == 64)
	ArchitectureIs32Bit bool = (strconv.IntSize == 32)
)

// Checks whether the unsigned Integer is able to be converted into 'int' Type.
func isUint64ConvertibleToInt(
	number uint64,
) bool {

	// 64-bit CPU Architecture.
	if ArchitectureIs64Bit {

		if number <= math.MaxInt64 {
			return true
		}
		return false
	}

	// 32-bit CPU Architecture.
	if ArchitectureIs32Bit {

		if number <= math.MaxInt32 {
			return true
		}
		return false
	}

	// UnKnown CPU Architecture.
	return false
}
