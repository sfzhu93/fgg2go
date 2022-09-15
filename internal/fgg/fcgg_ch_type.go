package fgg

import (
	"github.com/sfzhu93/fgg2go/internal/base"
)

type TypeBase interface {
	base.Type
}

const ChRecvOnly = 1
const ChSendOnly = 2
const ChBidirection = 3

var ChKindToString = map[int]string{ChRecvOnly: "chan<-", ChSendOnly: "<-chan", ChBidirection: "chan"}

type ChannelType struct {
	ElementType Type
	ChType      int
}

func IsChannelType(ty Type) bool {
	_, ok := ty.(ChannelType)
	return ok
}
func (c ChannelType) Impls(ds []base.Decl, t base.Type) bool {
	panic("implement me")
}

func (c ChannelType) Equals(t base.Type) bool {
	ch_ty := t.(ChannelType)
	return c.ElementType.Equals(ch_ty.ElementType) && c.ChType == ch_ty.ChType
}

func (c ChannelType) String() string {
	chLit := ChKindToString[c.ChType]
	return chLit + " " + c.ElementType.String()
}

//c <:u
func (c ChannelType) ImplsDelta(ds []Decl, delta Delta, u Type) bool {
	u_ := u.(ChannelType)
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

func (c ChannelType) ImplsDeltaWithEnv(env *Env, delta Delta, u Type) bool {
	u_ := u.(ChannelType)
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

func (c ChannelType) TSubs(subs map[TParam]Type) Type {
	return ChannelType{
		ElementType: c.ElementType.TSubs(subs),
		ChType:      c.ChType,
	}
}

func (c ChannelType) SubsEta(eta Eta) Type {
	return ChannelType{
		ElementType: c.ElementType.SubsEta(eta),
		ChType:      c.ChType,
	}
}

func (c ChannelType) SubsEtaOpen(eta EtaOpen) Type {
	return ChannelType{
		ElementType: c.ElementType.SubsEtaOpen(eta),
		ChType:      c.ChType,
	}
}

func (c ChannelType) Ok(ds []Decl, delta Delta) {
	c.ElementType.Ok(ds, delta)
}

func (c ChannelType) ToGoString(ds []Decl) string {
	panic("implement me")
}
