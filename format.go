// format.go.

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

//	1. Special symbols of 'bencode' encoding.

//	1.1. Headers.
const (
	HeaderDictionary               byte = 'd'
	HeaderInteger                  byte = 'i'
	HeaderList                     byte = 'l'
	HeaderStringSizeValueDelimiter byte = ':'
)

//	1.2. Footers.
const (
	FooterCommon byte = 'e'
)

//	1.3. Sections of a BitTorrent file.
const (
	FileSectionAnnounce     = "announce"
	FileSectionAnnounceList = "announce-list"
	FileSectionCreationDate = "creation date"
	FileSectionComment      = "comment"
	FileSectionCreatedBy    = "created by"
	FileSectionEncoding     = "encoding"
	FileSectionInfo         = "info"
)
