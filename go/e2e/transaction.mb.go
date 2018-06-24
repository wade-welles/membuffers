package e2e

import (
	"github.com/orbs-network/membuffers/go"
)

/*
message Transaction {
	TransactionData data = 1;
	bytes signature = 2;
}
*/

// reader

type Transaction struct {
	message membuffers.Message
}

var m_Transaction_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes}
var m_Transaction_Unions = [][]membuffers.FieldType{{}}

func TransactionReader(buf []byte) *Transaction {
	x := &Transaction{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_Transaction_Scheme, m_Transaction_Unions)
	return x
}

func (x *Transaction) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *Transaction) Data() *TransactionData {
	b, s := x.message.GetMessage(0)
	return TransactionDataReader(b[:s])
}

func (x *Transaction) RawData() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *Transaction) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *Transaction) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type TransactionBuilder struct {
	builder   membuffers.Builder
	Data      *TransactionDataBuilder
	Signature []byte
}

func (w *TransactionBuilder) Write(buf []byte) {
	w.builder.Reset()
	w.builder.WriteMessage(buf, w.Data)
	w.builder.WriteBytes(buf, w.Signature)
}

func (w *TransactionBuilder) GetSize() membuffers.Offset {
	return w.builder.GetSize()
}

func (w *TransactionBuilder) CalcRequiredSize() membuffers.Offset {
	w.Write(nil)
	return w.builder.GetSize()
}

/*
message TransactionData {
	uint32 protocol_version = 1;
	uint64 virtual_chain = 2;
	repeated TransactionSender sender = 3;
	int64 time_stamp = 4;
}
*/

// reader

type TransactionData struct {
	message membuffers.Message
}

var m_TransactionData_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeMessageArray,membuffers.TypeUint64}
var m_TransactionData_Unions = [][]membuffers.FieldType{{}}

func TransactionDataReader(buf []byte) *TransactionData {
	x := &TransactionData{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TransactionData_Scheme, m_TransactionData_Unions)
	return x
}

func (x *TransactionData) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionData) ProtocolVersion() uint32 {
	return x.message.GetUint32(0)
}

func (x *TransactionData) RawProtocolVersion() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionData) VirtualChain() uint64 {
	return x.message.GetUint64(1)
}

func (x *TransactionData) RawVirtualChain() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *TransactionData) SenderIterator() *TransactionDataSenderIterator {
	return &TransactionDataSenderIterator{iterator: x.message.GetMessageArrayIterator(2)}
}

type TransactionDataSenderIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionDataSenderIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionDataSenderIterator) NextSender() *TransactionSender {
	b, s := i.iterator.NextMessage()
	return TransactionSenderReader(b[:s])
}

func (x *TransactionData) RawSenderArray() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *TransactionData) TimeStamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *TransactionData) RawTimeStamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

// builder

type TransactionDataBuilder struct {
	builder         membuffers.Builder
	ProtocolVersion uint32
	VirtualChain    uint64
	Sender          []*TransactionSenderBuilder
	TimeStamp       uint64
}

func (w *TransactionDataBuilder) sender() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.Sender))
	for i, v := range w.Sender {
		res[i] = v
	}
	return res
}

func (w *TransactionDataBuilder) Write(buf []byte) {
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.ProtocolVersion)
	w.builder.WriteUint64(buf, w.VirtualChain)
	w.builder.WriteMessageArray(buf, w.sender())
	w.builder.WriteUint64(buf, w.TimeStamp)
}

func (w *TransactionDataBuilder) GetSize() membuffers.Offset {
	return w.builder.GetSize()
}

func (w *TransactionDataBuilder) CalcRequiredSize() membuffers.Offset {
	w.Write(nil)
	return w.builder.GetSize()
}

/*
message TransactionSender {
	string name = 1;
	repeated string friend = 2;
}
*/

// reader

type TransactionSender struct {
	message membuffers.Message
}

var m_TransactionSender_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeStringArray}
var m_TransactionSender_Unions = [][]membuffers.FieldType{{}}

func TransactionSenderReader(buf []byte) *TransactionSender {
	x := &TransactionSender{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TransactionSender_Scheme, m_TransactionSender_Unions)
	return x
}

func (x *TransactionSender) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionSender) Name() string {
	return x.message.GetString(0)
}

func (x *TransactionSender) RawName() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionSender) FriendIterator() *TransactionSender_FriendIterator {
	return &TransactionSender_FriendIterator{iterator: x.message.GetStringArrayIterator(1)}
}

type TransactionSender_FriendIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionSender_FriendIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionSender_FriendIterator) NextFriend() string {
	return i.iterator.NextString()
}

func (x *TransactionSender) RawFriendArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type TransactionSenderBuilder struct {
	builder membuffers.Builder
	Name    string
	Friend  []string
}

func (w *TransactionSenderBuilder) Write(buf []byte) {
	w.builder.Reset()
	w.builder.WriteString(buf, w.Name)
	w.builder.WriteStringArray(buf, w.Friend)
}

func (w *TransactionSenderBuilder) GetSize() membuffers.Offset {
	return w.builder.GetSize()
}

func (w *TransactionSenderBuilder) CalcRequiredSize() membuffers.Offset {
	w.Write(nil)
	return w.builder.GetSize()
}