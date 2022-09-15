package fg

import "github.com/sfzhu93/fgg2go/internal/base"

type ChanRecv struct {
	Body FGExpr
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

func (c ChanRecv) Subs(subs map[Variable]FGExpr) FGExpr {
	panic("implement me")
}

func (c ChanRecv) Typing(ds []Decl, gamma Gamma, allowStupid bool) TypeBase {
	body_ty := c.Body.Typing(ds, gamma, allowStupid).(ChannelType)
	if body_ty.ChType == ChBidirection || body_ty.ChType == ChSendOnly {
		return body_ty.ElementType
	} else {
		panic("Cannot receive from a send-only channel.")
	}
}

func (c ChanRecv) Eval(ds []Decl) (FGExpr, string) {
	panic("implement me")
}

type MakeChan struct {
	ChType ChannelType
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

func (m MakeChan) Subs(subs map[Variable]FGExpr) FGExpr {
	panic("implement me")
}

func (m MakeChan) Typing(ds []Decl, gamma Gamma, allowStupid bool) TypeBase {
	if m.ChType.ChType != ChBidirection {
		panic("Channel type is not bi-direction in `make`.")
	}
	if !isTypeOk(ds, m.ChType.ElementType) {
		panic(m.ChType.String() + " type check fails!")
	}
	return m.ChType
}

func (m MakeChan) Eval(ds []Decl) (FGExpr, string) {
	panic("implement me")
}
