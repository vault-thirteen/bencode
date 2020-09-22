# Short Description.
Library for encoding and decoding the 'bencode' Format.

# Full Description.

This Package provides Data Encoding to and Data Decoding from the 'bencode' Format.<br />
The 'bencode' Format was introduced with the Appearance of the BitTorrent Protocol.<br />
Apart from the encoding and decoding Data with the 'bencode' Format, this Packages also provides some additional Functionality, such as:
  - Automatic Self-Check after File Decoding;
  - Automatic Calculation of the BitTorrent Info Hash (also known as 'BTIH') after the File Decoding.

This Package is focused on the Safety and Reliability rather than Speed.<br />
As opposed to many other existing 'bencode' Libraries, in this Library, when decoding a Stream, the Decoder stops at Syntax Errors just as they appear. Moreover, the Decoder is wise enough to stop when Size Fields are surprisingly long to prevent Overflows in Memory, so that Size-Prefix Overflow Attack is not working on this Decoder.<br />
Of course, this Decoder is not the safest, it can only read Files which can be placed into RAM.<br />

# Installation.

```
go get -u "github.com/neverwinter-nights/bencode"
```

# Usage.

```
import ...
```
