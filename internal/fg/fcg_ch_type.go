package fg

import "github.com/sfzhu93/fgg2go/internal/base"

type TypeBase interface {
	base.Type
}

const ChRecvOnly = 1
const ChSendOnly = 2
const ChBidirection = 3

var ChKindToString = map[int]string{ChRecvOnly: "chan<-", ChSendOnly: "<-chan", ChBidirection: "chan"}

type ChannelType struct {
	ElementType TypeBase
	ChType int
}

func (c ChannelType) Impls(ds []base.Decl, t base.Type) bool {
	u_ := t.(ChannelType)
	if !u_.ElementType.Equals(c.ElementType) {
		return false
	}
	if u_.ChType == c.ChType {
		return true
	}
	if c.ChType == ChBidirection && (u_.ChType == ChSendOnly || u_.ChType == ChRecvOnly) {
		return true
	}
	return false
}

func (c ChannelType) Equals(t base.Type) bool {
	panic("implement me")
}

func (c ChannelType) String() string {
	chLit := ChKindToString[c.ChType]
	return chLit + " " + c.ElementType.String()
}
