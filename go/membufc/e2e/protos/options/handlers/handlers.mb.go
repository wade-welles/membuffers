// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package handlers

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// service ServicesINeedFromOthersHandler

type ServicesINeedFromOthersHandler interface {
	SomeMethodINeedFromOthers(input *SomeMessage) (*SomeMessage, error)
}

/////////////////////////////////////////////////////////////////////////////
// service ServicesIProvideToOthersHandler

type ServicesIProvideToOthersHandler interface {
	SomeMethodIProvideToOthers(input *SomeMessage) (*SomeMessage, error)
}

/////////////////////////////////////////////////////////////////////////////
// message SomeMessage

// reader

type SomeMessage struct {
	message membuffers.Message
}

var _SomeMessage_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var _SomeMessage_Unions = [][]membuffers.FieldType{}

func SomeMessageReader(buf []byte) *SomeMessage {
	x := &SomeMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _SomeMessage_Scheme, _SomeMessage_Unions)
	return x
}

func (x *SomeMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *SomeMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SomeMessage) Str() string {
	return x.message.GetString(0)
}

func (x *SomeMessage) RawStr() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SomeMessage) MutateStr(v string) error {
	return x.message.SetString(0, v)
}

// builder

type SomeMessageBuilder struct {
	builder membuffers.Builder
	Str string
}

func (w *SomeMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteString(buf, w.Str)
	return nil
}

func (w *SomeMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SomeMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SomeMessageBuilder) Build() *SomeMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SomeMessageReader(buf)
}

