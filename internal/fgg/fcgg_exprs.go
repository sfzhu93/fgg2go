package fgg

import (
	"github.com/sfzhu93/fgg2go/internal/base"
)

type ChanRecv struct {
	Body FGGExpr
}

func (c ChanRecv) Subs(subs map[Variable]FGGExpr) FGGExpr {
	panic("implement me")
}

func (c ChanRecv) TSubs(subs map[TParam]Type) FGGExpr {
	return ChanRecv{Body: c.Body.TSubs(subs)}
}

func (c ChanRecv) Typing(ds []Decl, delta Delta, gamma Gamma, allowStupid bool) Type {
	body_ty := c.Body.Typing(ds, delta, gamma, allowStupid).(ChannelType)
	if body_ty.ChType == ChBidirection || body_ty.ChType == ChSendOnly {
		return body_ty.ElementType
	} else {
		panic("Cannot receive from a send-only channel.")
	}
}

func (c ChanRecv) TypingWithEnv(env *Env, delta Delta, gamma Gamma, allowStupid bool) Type {
	body_ty := c.Body.TypingWithEnv(env, delta, gamma, allowStupid).(ChannelType)
	if body_ty.ChType == ChBidirection || body_ty.ChType == ChSendOnly {
		return body_ty.ElementType
	} else {
		panic("Cannot receive from a send-only channel.")
	}
}

func (c ChanRecv) Eval(ds []Decl) (FGGExpr, string) {
	panic("implement me")
}

func (c ChanRecv) String() string {
	return "<-" + c.Body.String()
}

func (c ChanRecv) IsValue() bool {
	panic("implement me")
}

func (c ChanRecv) CanEval(ds []base.Decl) bool {
	panic("implement me")
}

func (c ChanRecv) ToGoString(ds []base.Decl) string {
	panic("implement me")
}

type MakeChan struct {
	ChType ChannelType
}

func (m MakeChan) Subs(subs map[Variable]FGGExpr) FGGExpr {
	panic("implement me")
}

func (m MakeChan) TSubs(subs map[TParam]Type) FGGExpr {
	return MakeChan{ChType: m.ChType.TSubs(subs).(ChannelType)}
}

func (m MakeChan) Typing(ds []Decl, delta Delta, gamma Gamma, allowStupid bool) Type {
	if m.ChType.ChType != ChBidirection {
		panic("Channel type is not bi-direction in `make`.")
	}
	m.ChType.ElementType.Ok(ds, delta)
	return m.ChType
}

func (m MakeChan) TypingWithEnv(env *Env, delta Delta, gamma Gamma, allowStupid bool) Type {
	if m.ChType.ChType != ChBidirection {
		panic("Channel type is not bi-direction in `make`.")
	}
	m.ChType.ElementType.Ok(env.fggDecls, delta)
	return m.ChType
}

func (m MakeChan) Eval(ds []Decl) (FGGExpr, string) {
	panic("implement me")
}

func (m MakeChan) String() string {
	return "make(" + m.ChType.String() + ")"
}

func (m MakeChan) IsValue() bool {
	panic("implement me")
}

func (m MakeChan) CanEval(ds []base.Decl) bool {
	panic("implement me")
}

func (m MakeChan) ToGoString(ds []base.Decl) string {
	panic("implement me")
}
