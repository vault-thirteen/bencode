// format_test.go.

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
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_format(t *testing.T) {
	var aTest = tester.New(t)

	aTest.MustBeEqual(FileSectionAnnounce, "announce")
	aTest.MustBeEqual(FileSectionAnnounceList, "announce-list")
	aTest.MustBeEqual(FileSectionCreationDate, "creation date")
	aTest.MustBeEqual(FileSectionComment, "comment")
	aTest.MustBeEqual(FileSectionCreatedBy, "created by")
	aTest.MustBeEqual(FileSectionEncoding, "encoding")
	aTest.MustBeEqual(FileSectionInfo, "info")
}
