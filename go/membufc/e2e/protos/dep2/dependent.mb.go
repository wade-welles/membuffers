// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package dep2

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/dep1"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/dep1/dep11"
)

/////////////////////////////////////////////////////////////////////////////
// message Dependent

// reader

type Dependent struct {
	message membuffers.Message
}

var _Dependent_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeUint16,membuffers.TypeMessage,}
var _Dependent_Unions = [][]membuffers.FieldType{}

func DependentReader(buf []byte) *Dependent {
	x := &Dependent{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _Dependent_Scheme, _Dependent_Unions)
	return x
}

func (x *Dependent) IsValid() bool {
	return x.message.IsValid()
}

func (x *Dependent) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *Dependent) A() *dep1.DependencyMessage {
	b, s := x.message.GetMessage(0)
	return dep1.DependencyMessageReader(b[:s])
}

func (x *Dependent) RawA() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *Dependent) B() dep11.DependencyEnum {
	return dep11.DependencyEnum(x.message.GetUint16(1))
}

func (x *Dependent) RawB() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *Dependent) MutateB(v dep11.DependencyEnum) error {
	return x.message.SetUint16(1, uint16(v))
}

func (x *Dependent) C() *SamePackageDependencyMessage {
	b, s := x.message.GetMessage(2)
	return SamePackageDependencyMessageReader(b[:s])
}

func (x *Dependent) RawC() []byte {
	return x.message.RawBufferForField(2, 0)
}

// builder

type DependentBuilder struct {
	builder membuffers.Builder
	A *dep1.DependencyMessageBuilder
	B dep11.DependencyEnum
	C *SamePackageDependencyMessageBuilder
}

func (w *DependentBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.A)
	if err != nil {
		return
	}
	w.builder.WriteUint16(buf, uint16(w.B))
	err = w.builder.WriteMessage(buf, w.C)
	if err != nil {
		return
	}
	return nil
}

func (w *DependentBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *DependentBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *DependentBuilder) Build() *Dependent {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return DependentReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

