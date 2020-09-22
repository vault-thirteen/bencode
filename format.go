// format.go.

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

// Main Settings of the 'bencode' Format.

package bencode

//	1.	Special Symbols of 'bencode' Encoding.

//	1.1.	Headers.
const (
	HeaderDictionary               byte = 'd'
	HeaderInteger                  byte = 'i'
	HeaderList                     byte = 'l'
	HeaderStringSizeValueDelimiter byte = ':'
)

//	1.2.	Footers.
const (
	FooterCommon byte = 'e'
)

//	1.3.	Sections of a BitTorrent File.
const (
	FileSectionAnnounce     = "announce"
	FileSectionAnnounceList = "announce-list"
	FileSectionCreationDate = "creation date"
	FileSectionComment      = "comment"
	FileSectionCreatedBy    = "created by"
	FileSectionEncoding     = "encoding"
	FileSectionInfo         = "info"
)
