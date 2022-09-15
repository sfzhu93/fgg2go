package fgg

type FGGStmt interface {
	FGGNode
	//TODO: should have memebers similar to FGExpr
	TSubs(subs map[TParam]Type) FGGStmt
	Subs(m map[Variable]FGGExpr) FGGExpr
	IsTerminateStmt() bool
	Typing(ds []Decl, delta Delta, gamma Gamma, sigma Type, allowStupid bool) Type
}

type ReturnStmt struct {
	ExprBody FGGExpr
}

func (F ReturnStmt) Subs(m map[Variable]FGGExpr) FGGExpr {
	panic("implement me")
}

func (F ReturnStmt) Typing(ds []Decl, delta Delta, gamma Gamma, sigma Type, allowStupid bool) Type {
	tau := F.ExprBody.Typing(ds, delta, gamma, allowStupid)
	if tau.ImplsDelta(ds, delta, sigma) {
		return sigma
	} else {
		panic("returned expression doesn't implement " + sigma.String())
	}
}

func (F ReturnStmt) TSubs(subs map[TParam]Type) FGGStmt {
	return ReturnStmt{ExprBody: F.ExprBody.TSubs(subs)}
}

func (F ReturnStmt) String() string {
	return "return " + F.ExprBody.String()
}

func (F ReturnStmt) IsTerminateStmt() bool {
	return true
}

type FGGCaseSelectStmt struct {
	Cases       []FGGCaseBody
	DefaultCase []FGGStmt
}

func (F FGGCaseSelectStmt) Subs(m map[Variable]FGGExpr) FGGExpr {
	panic("implement me")
}

func (F FGGCaseSelectStmt) Typing(ds []Decl, delta Delta, gamma Gamma, sigma Type, allowStupid bool) Type {
	var ret_ty Type = nil
	for _, casebody := range F.Cases {
		ty := casebody.Typing(ds, delta, gamma, sigma, allowStupid)
		if !ty.ImplsDelta(ds, delta, sigma) {
			panic("Type mismatch in select-case statement: previously was " + ret_ty.String() + ", but got " + ty.String() + " later on.")
		}
	}
	if F.DefaultCase != nil && len(F.DefaultCase) > 0 {
		ty := MethBody{StmtBody: F.DefaultCase}.Typing(ds, delta, gamma, sigma, allowStupid)
		if !ty.ImplsDelta(ds, delta, sigma) {
			panic("Type mismatch in select-case statement: previously was " + ret_ty.String() + ", but got " + ty.String() + " later on.")
		}
	}
	return sigma
}

func (F FGGCaseSelectStmt) String() string {
	ret := "select {\n"
	for _, c := range F.Cases {
		ret += "case " + c.CaseGuard.String() + ":\n"
		ret += getStringOfStmtArray(c.Body)
		ret += "\n"
	}
	if F.DefaultCase != nil && len(F.DefaultCase) > 0 {
		ret += "default:\n"
		ret += getStringOfStmtArray(F.DefaultCase)
	}
	ret += "}\n"
	return ret
}

func (F FGGCaseSelectStmt) IsTerminateStmt() bool {
	return true
}

func (F FGGCaseSelectStmt) TSubs(subs map[TParam]Type) FGGStmt {
	new_cases := []FGGCaseBody{}
	for _, x := range F.Cases {
		new_cases = append(new_cases, x.TSubs(subs))
	}
	defaultstmts := []FGGStmt{}
	for _, x := range F.DefaultCase {
		defaultstmts = append(defaultstmts, x.TSubs(subs))
	}
	return FGGCaseSelectStmt{Cases: new_cases, DefaultCase: defaultstmts}
}

type FGGCaseBody struct {
	CaseGuard FGGCaseGuard
	Body      []FGGStmt
}

func (F FGGCaseBody) String() string {
	panic("implement me")
}

func (F FGGCaseBody) Subs(m map[Variable]FGGExpr) FGGExpr {
	panic("implement me")
}

func (F FGGCaseBody) IsTerminateStmt() bool {
	panic("implement me")
}

func (F FGGCaseBody) TSubs(subs map[TParam]Type) FGGCaseBody {
	new_body := []FGGStmt{}
	for _, x := range F.Body {
		new_body = append(new_body, x.TSubs(subs))
	}
	return FGGCaseBody{
		CaseGuard: F.CaseGuard.TSubs(subs),
		Body:      new_body,
	}
}

func (F FGGCaseBody) Typing(ds []Decl, delta Delta, gamma Gamma, sigma Type, allowStupid bool) Type {
	new_stmts := []FGGStmt{F.CaseGuard.ToFGGStmt()}
	new_stmts = append(new_stmts, F.Body...)
	return MethBody{new_stmts}.Typing(ds, delta, gamma, sigma, allowStupid)
}

func getStringOfStmtArray(stmts []FGGStmt) string {
	ret := ""
	for i := 0; i < len(stmts)-1; i++ {
		ret += stmts[i].String() + ";"
	}
	if len(stmts) > 0 {
		ret += stmts[len(stmts)-1].String()
	}
	return ret
}

type FGGChCloseStmt struct {
	Body FGGExpr
}

func (F FGGChCloseStmt) Subs(m map[Variable]FGGExpr) FGGExpr {
	panic("implement me")
}

func (F FGGChCloseStmt) Typing(ds []Decl, delta Delta, gamma Gamma, sigma Type, allowStupid bool) Type {
	return nil //we use nil to represent statements with no types
}

func (F FGGChCloseStmt) String() string {
	return "close(" + F.Body.String() + ")"
}

func (F FGGChCloseStmt) IsTerminateStmt() bool {
	return false
}

func (F FGGChCloseStmt) TSubs(subs map[TParam]Type) FGGStmt {
	return FGGChCloseStmt{Body: F.Body}
}

type FGGChDispatchStmt struct {
	LeftExpr  FGGExpr
	RightExpr FGGExpr
}

func (F FGGChDispatchStmt) Subs(m map[Variable]FGGExpr) FGGExpr {
	panic("implement me")
}

func (F FGGChDispatchStmt) Typing(ds []Decl, delta Delta, gamma Gamma, sigma Type, allowStupid bool) Type {
	return nil
}

func (F FGGChDispatchStmt) String() string {
	return F.LeftExpr.String() + "<-" + F.RightExpr.String()
}

func (F FGGChDispatchStmt) IsTerminateStmt() bool {
	return false
}

func (F FGGChDispatchStmt) TSubs(subs map[TParam]Type) FGGStmt {
	return FGGChDispatchStmt{
		LeftExpr:  F.LeftExpr.TSubs(subs),
		RightExpr: F.RightExpr.TSubs(subs),
	}
}

type FGGAssignmentStmt struct {
	VarName Name
	Body    FGGExpr
}

func (F FGGAssignmentStmt) Subs(m map[Variable]FGGExpr) FGGExpr {
	panic("implement me")
}

func (F FGGAssignmentStmt) Typing(ds []Decl, delta Delta, gamma Gamma, sigma Type, allowStupid bool) Type {
	ty_e := F.Body.Typing(ds, delta, gamma, allowStupid)
	if _, ok := gamma[F.VarName]; ok {
		panic("Variable redefined: " + F.VarName) //Do we expect a panic here?
	} else {
		gamma[F.VarName] = ty_e
	}
	return ty_e
}

func (F FGGAssignmentStmt) TSubs(subs map[TParam]Type) FGGStmt {
	return FGGAssignmentStmt{Body: F.Body.TSubs(subs), VarName: F.VarName}
}

func (F FGGAssignmentStmt) String() string {
	return F.VarName + " := " + F.Body.String()
}

func (F FGGAssignmentStmt) IsTerminateStmt() bool {
	return false
}

type MethBody struct {
	StmtBody []FGGStmt
}

func (m2 MethBody) Subs(m map[Variable]FGGExpr) FGGExpr {
	panic("implement me")
}

func (m MethBody) TSubs(subs map[TParam]Type) FGGStmt {
	new_stmts := []FGGStmt{}
	for _, x := range m.StmtBody {
		new_stmts = append(new_stmts, x.TSubs(subs))
	}
	return MethBody{new_stmts}
}

func (m MethBody) IsTerminateStmt() bool {
	return true
}

func (m MethBody) String() string {
	ret := ""
	for _, stmt := range m.StmtBody {
		ret += stmt.String() + ";\n"
	}
	return ret
}

func (m MethBody) Typing(ds []Decl, delta Delta, gamma Gamma, sigma Type, stupid bool) Type {
	tmp_delta := make(map[TParam]Type)
	for k, v := range delta {
		tmp_delta[k] = v
	}
	tmp_gamma := make(map[Name]Type)
	for k, v := range gamma {
		tmp_gamma[k] = v
	}
	for _, stmt := range m.StmtBody {
		ty := stmt.Typing(ds, tmp_delta, tmp_gamma, sigma, stupid) //this may change delta and gamma.
		if stmt.IsTerminateStmt() && !ty.ImplsDelta(ds, delta, sigma){
			panic("Returned type doesn't match in method body.")
		}
	}
	return sigma
}

type GoroutineStmt struct {
	CallExpr Call
}

func (g GoroutineStmt) Subs(m map[Variable]FGGExpr) FGGExpr {
	panic("implement me")
}

func (g GoroutineStmt) Typing(ds []Decl, delta Delta, gamma Gamma, sigma Type, allowStupid bool) Type {
	return g.CallExpr.Typing(ds, delta, gamma, allowStupid)
}

func (g GoroutineStmt) TSubs(subs map[TParam]Type) FGGStmt {
	return GoroutineStmt{CallExpr: g.CallExpr.TSubs(subs).(Call)}
}

func (g GoroutineStmt) String() string {
	return "go " + g.CallExpr.String()
}

func (g GoroutineStmt) IsTerminateStmt() bool {
	return false
}
