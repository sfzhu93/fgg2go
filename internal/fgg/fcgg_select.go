package fgg

type FGGCaseGuard interface {
	FGGNode
	TSubs(subs map[TParam]Type) FGGCaseGuard
	ToFGGStmt() FGGStmt
}

type FGGDispatchCaseGuard struct {
	LeftExpr  FGGExpr
	RightExpr FGGExpr
}

func (F FGGDispatchCaseGuard) ToFGGStmt() FGGStmt {
	return FGGChDispatchStmt{LeftExpr: F.LeftExpr, RightExpr: F.RightExpr}
}

func (F FGGDispatchCaseGuard) TSubs(subs map[TParam]Type) FGGCaseGuard {
	return FGGDispatchCaseGuard{F.LeftExpr.TSubs(subs), F.RightExpr.TSubs(subs)}
}

func (F FGGDispatchCaseGuard) String() string {
	return F.LeftExpr.String() + "<-" + F.RightExpr.String()
}

type FGGReceiveCaseGuard struct {
	LeftName  Name
	RightExpr FGGExpr
}

func (F FGGReceiveCaseGuard) ToFGGStmt() FGGStmt {
	return FGGAssignmentStmt{VarName: F.LeftName, Body: ChanRecv{F.RightExpr}}
}

func (F FGGReceiveCaseGuard) String() string {
	return F.LeftName + ":= " + F.RightExpr.String()
}

func (F FGGReceiveCaseGuard) TSubs(subs map[TParam]Type) FGGCaseGuard {
	return FGGReceiveCaseGuard{
		LeftName:  F.LeftName,
		RightExpr: F.RightExpr.TSubs(subs),
	}
}
