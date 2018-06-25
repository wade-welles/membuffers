# MemBuffers

**MemBuffers** is an efficient cross platform serialization library. It allows you to directly access serialized data without unpacking/parsing it first. The wire format is deterministic and canonical, perfect for cryptographic fields and hashes. It uses **protobuf** schema (ver 3) to define messages but naturally encodes them in a different format.

This standalone library was originally created as part of the [ORBS](https://www.orbs.com) blockchain project. Its features are designed to address the requirements of a highly performant system dealing with cryptography, hashes and signatures. 

## Features

* *Canonical encoding* - Two parties encoding the same message with the same values will always receive the exact same encoding byte by byte.

* *Direct access* - There's no packing/unpacking process for messages, serialized fields are accessed directly. Messages are mapped to memory.

* *Familiar schema* - Messages are defined using Google [Protobuf](https://developers.google.com/protocol-buffers/docs/proto3) schema (ver 3) for easy migration and familiar syntax.

* *Zero copies* - Without packing/unpacking, excessive data copying is avoided. Data can even be mutated and changed in-place.

* *Lazy parsing* - Fields are parsed lazily and never recursively. Accessing a single field in a complexly nested message will not parse the entire message.

* *Self-contained fields* - Every field is serialized sequentially with all its children, this permits direct byte operations over fields such as hashing or signing.

* *Almost human readable* - Encoding format is simple enough to be decoded by hand and even be generated manually without a fancy library.

## Usage

1. Write schema file `transaction.proto`:

    ```proto
    syntax = "proto3";
    package types;
    
    message Transaction {
       TransactionData data = 1;
       bytes hash = 2;
    }
    
    message TransactionData {
       uint32 protocol_version = 1;
       uint64 sender_account = 2;
       string contract_method = 3;
    }
    ```
    
2. Compile schema to code file `transaction.mb.go`:

    ```sh
    membufc --go transaction.proto
    ```

3. Encode a transaction message for sending over the wire:

    ```go
    builder := &types.TransactionBuilder{
      Data: &types.TransactionDataBuilder{
        ProtocolVersion: 0x01,
        SenderAccount: 0x11223344,
        ContractMethod: "ZincToken.Transfer",
      },
      Hash: []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,},
    }
    buf := make([]byte, builder.CalcRequiredSize())
    err := builder.Write(buf)
    if err != nil {
      fmt.Println("serialization error")
    }
    ```

4. Decode a transaction message received over the wire:

    ```go
    transaction := types.TransactionReader(buf)
    // validate format
    if !transaction.IsValid() {
      fmt.Println("serialized data is invalid")
    }
    // check hash
    calculated := md5.Sum(transaction.RawData())
    if !bytes.Equal(calculated, transaction.Hash()) {
      fmt.Println("hash mismatch")
    }
    // access fields
    ver := transaction.Data().ProtocolVersion()
    sender := transaction.Data().SenderAccount()
    contract := transaction.Data().ContractMethod
    ```
    
5. Mutate fields as needed:

    ```go
    transaction := types.TransactionReader(buf)
    calculated := md5.Sum(transaction.RawData())
    err := transaction.MutateHash(calculated)
    if err != nil {
      fmt.Println("mutation error")
    }
    ```

## Installation

#### Prerequisites

1. Make sure [Go](https://golang.org/doc/install) is installed (version 1.10 or later).
  
    > Verify with `go version`

2. Make sure [Go workspace bin](https://stackoverflow.com/questions/42965673/cant-run-go-bin-in-terminal) is in your path.
  
    > Install with ``export PATH=$PATH:`go env GOPATH`/bin``
  
    > Verify with `echo $PATH`

#### Get and build

1. Get the library into your Go workspace:
 
     ```sh
     go get github.com/orbs-network/membuffers/go
     cd `go env GOPATH`/src/github.com/orbs-network/membuffers/go
     ```

* Install the `membufc` compiler:

    ```sh
    cd ./membufc
    go install
    ```
    > Verify with `membufc --version`
    
## Test

1. Test the library (unit tests and end to end tests):

    ```sh
    ./test.sh
    ```
    
2. Test the `membufc` compiler (end to end test):

    ```sh
    cd ./membufc
    ./test.sh
    ```
    
## Comparison to other libraries

#### [Google Protobuf](https://developers.google.com/protocol-buffers/)

###### Strengths of MemBuffers
* *Canonical encoding* - Protobuf isn't canonical. Different encoders may get different results over the same source data.
* *Direct access* - Protobuf messages are sent packed over the wire and must be unpacked before access.  
* *Familiar schema* - Protobuf has a great schema, we use it too.
* *Zero copies* - Protobuf unpacking process involves copying the data to unpacked form.
* *Lazy parsing* - Protobuf unpacking process is recursive and must decode the entire message before any field can be accessed.
* *Self-contained fields* - Protobuf fields must be re-packed separately to a byte steam so they can be hashed.
* *Almost human readable* - Protobuf encoding format is complex and not easily readable by humans or quick and dirty scripts.

###### Strengths of Protobuf
* Protobuf currently has wider language support.
* Protobuf packed wire format is more efficient in space (optimizes for bandwidth instead of access performance).
* Protobuf is better optimized to deal with deprecated or optional fields (they don't take any space on the wire).
* Protobuf is backed by Google.

#### [Google Flatbuffers](https://google.github.io/flatbuffers/)

###### Strengths of MemBuffers
* *Canonical encoding* - Flatbuffers isn't canonical. Different encoders may get different results over the same source data.
* *Direct access* - Flatbuffers messages also provide direct access and do not need to be unpacked before access.  
* *Familiar schema* - Flatbuffers has its own proprietary schema.
* *Zero copies* - Flatbuffers avoids unpacking so it also has very little data copying.
* *Lazy parsing* - Flatbuffers decode process is recursive and must parse the entire message before any field can be accessed.
* *Self-contained fields* - Flatbuffers fields are not self-contained so they cannot be hashed directly from wire format.
* *Almost human readable* - Flatbuffers encoding format is complex and not easily readable by humans or quick and dirty scripts.

###### Strengths of Flatbuffers
* Flatbuffers currently has wider language support.
* Flatbuffers messages are memory mapped for fast direct access (although the Go implementation doesn't take advantage of that).
* Flatbuffers is backed by Google.

#### [Cap'n Proto](https://capnproto.org/)

###### Strengths of MemBuffers
* *Canonical encoding* - Cap'n Proto isn't canonical. Different encoders may get different results over the same source data.
* *Direct access* - Cap'n Proto messages also provide direct access and do not need to be unpacked before access.  
* *Familiar schema* - Cap'n Proto has its own proprietary schema.
* *Zero copies* - Cap'n Proto avoids unpacking so it also has very little data copying.
* *Lazy parsing* - Cap'n Proto decode process is recursive and must parse the entire message before any field can be accessed.
* *Self-contained fields* - Cap'n Proto fields are not self-contained so they cannot be hashed directly from wire format.
* *Almost human readable* - Cap'n Proto encoding format is complex and not easily readable by humans or quick and dirty scripts.

###### Strengths of Flatbuffers
* Cap'n Proto currently has wider language support.
* Cap'n Proto messages are memory mapped for fast direct access (although the Go implementation doesn't take advantage of that).
* Cap'n Proto is a relatively popular project with a large community.

## Wire format

#### Alignment

* Large fields (like `uint32`, `uint64`) are aligned to 4 bytes.
* Smaller fields (like `uint8`, `uint16`) are aligned to their own size.

#### Primitives

* Such as `uint8`, `uint16`, `uint32`, `uint64`.
* All numbers (and size fields) are encoded in little-endian format.
* Primitives are encoded directly without any additional meta data (take exactly 1/2/4/64 bytes of space).

#### Arrays of primitives

* Such as `bytes`, `string`, `uint8[]`, `uint16[]`, `uint32[]`, `uint64[]`.
* Total content size in bytes is encoded first as a `uint32` (a 4 byte field).
* The actual content is encoded immediately after.
* Strings are not null terminated, they are encoded exactly like `bytes` or `uint8[]`.

#### Enums

* Always encoded as a `uint16` (a 2 byte field).

#### Messages

* Messages include other fields which may be primitives or other messages. The root object is always a message.
* Total content size in bytes (including all children content) is encoded first as a `uint32` (a 4 byte field). 
* The actual content is encoded immediately after (all children, which also include their own children).
* All fields in the schema are always encoded. Fields that were not given are encoded as zero value.
* Fields are encoded immediately one after the other without special meta data (zero padding may be added for alignment).
* Fields are encoded according to the tag order defined in the schema.
    > TIP: To reduce the wire size, order fields in the schema so that no space is lost due to alignment. 

#### Unions (oneof)

* Indication of the type (which one of the OneOf it is) is encoded first as a `uint16` (a 4 byte field).
* Encoding of the field itself appears immediately after (zero padding may be added for alignment).

#### The decode process

* Whenever a message field is first accessed, this message's fields are parsed (lazily).
* Only this message is parsed, this is not recursive. Child messages or arrays are not parsed until they are accessed.
* Parsing requires knowledge of the schema because the schema defines the order of fields and their types and sizes.
* Parsing is not random access, all previous fields (in the same level) must be parsed in order to know where a field starts.
* Parsing creates an in-memory offset table when tells in what offset (of the parent message content) each field starts.
* The offset table is generated by skipping over each field. Dynamic sized fields require reading their size from the wire. 

#### Example

Assuming the following schema:
```proto
syntax = "proto3";
package types;

message Transaction {
    TransactionData data = 1;
    bytes hash = 2;
}

message TransactionData {
    uint16 protocol_version = 1;
    uint32 sender_account = 2;
    string contract_method = 3;
}
```

Initialized to the following values:
```go
&types.TransactionBuilder{
  Data: &types.TransactionDataBuilder{
    ProtocolVersion: 0x01,
    SenderAccount: 0x11223344,
    ContractMethod: "abc",
  },
  Hash: []byte{0x55,0x56,0x57,0x58,0x59,},
}
```

Will produce the following encoding over the wire (total of 29 bytes):
```
0F 00 00 00                 // size of TransactionData (15 bytes)
01 00 00 00                 // TransactionData.protocol_version with 2 byte padding after
44 33 22 11                 // TransactionData.sender_account
03 00 00 00 61 62 63 00     // TransactionData.contract_method starting with size (3 bytes) and then the content and padding
05 00 00 00 55 56 57 58 59  // Transaction.hash starting with size (5 bytes) and then the content
```

## Evolving your schema

* The following rules will help you avoid making breaking changes and maintain backwards and forwards compatibility.
* Never delete schema fields. If you deprecate a field, stop using it and it will be encoded as zero value.
* Always add new fields at the end of the schema.
* Keep the numeric tag of fields sequential.

## License

MIT