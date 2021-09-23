# Short Description
Library for encoding and decoding the `bencode` format.

# Full Description

This package provides a mechanism for data encoding to and data decoding from 
the `bencode` format.<br />
The `bencode` format was introduced with the appearance of the **BitTorrent** 
protocol.<br />
Apart from the encoding and decoding data with the `bencode` format, this 
packages also provides some additional functionality, such as:
  - Automatic self-check after file decoding;
  - Automatic calculation of the **BitTorrent Info Hash** (also known as
  **BTIH**) after the file decoding.

This package is focused on the safety and reliability rather than Speed.<br />
As opposed to many other existing `bencode` libraries, in this library, when 
decoding a stream, the decoder stops at syntax errors just as they appear. 
Moreover, the decoder is wise enough to stop when size fields are surprisingly 
long to prevent overflows in memory, so that the size-prefix overflow attack 
is not working on this decoder.<br />
Of course, this decoder is not the safest, it can only read files which can be 
placed into the system memory (RAM).<br />

# Installation

```
go get -u "github.com/neverwinter-nights/bencode"
```

# Usage

```
import "github.com/neverwinter-nights/bencode"
```
