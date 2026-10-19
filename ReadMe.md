# Bencode

A library for encoding and decoding the _Bencode_ format.

## Description

This package provides a mechanism for data encoding to and data decoding from
the _Bencode_ format.

The _Bencode_ format was introduced with the appearance of the _BitTorrent_
protocol.

Apart from the encoding and decoding data with the _Bencode_ format, this
package also provides some additional functionality, such as:
- Automatic self-check after file decoding.

This package is focused on safety and reliability rather than speed.

As opposed to many other existing _Bencode_ format libraries, this library
follows the principle that during the decoding process of a stream the decoder
stops at syntax errors just as they appear. Moreover, the decoder is wise
enough to stop when size fields are surprisingly long to prevent overflows in
memory, so that the size-prefix overflow attack is not working on this decoder.
Of course, this does not make the decoder the safest one, while it can only
read those files which can be fully placed into the system memory (RAM).

## Importing

```
import "github.com/vault-thirteen/bencode"
```
