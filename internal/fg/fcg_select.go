package fg

type FGCaseGuard interface {
	FGNode
	ToFGStmt() FGStmt
}

type FGDispatchCaseGuard struct {
	LeftExpr  FGExpr
	RightExpr FGExpr
}

func (F FGDispatchCaseGuard) ToFGStmt() FGStmt {
	return FGChDispatchStmt{LeftExpr: F.LeftExpr, RightExpr: F.RightExpr}
}

func (F FGDispatchCaseGuard) String() string {
	return F.LeftExpr.String() + "<-" + F.RightExpr.String()
}

type FGReceiveCaseGuard struct {
	LeftName    Name
	IsUnusedVar bool
	RightExpr   FGExpr
}

func (F FGReceiveCaseGuard) ToFGStmt() FGStmt {
	return FGAssignmentStmt{VarName: F.LeftName, Body: ChanRecv{F.RightExpr}}
}

func (F FGReceiveCaseGuard) String() string {
	if !F.IsUnusedVar {
		return F.LeftName + ":= " + F.RightExpr.String()
	} else {
		return F.RightExpr.String()
	}
}
