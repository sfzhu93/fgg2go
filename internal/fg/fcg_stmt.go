package fg

type FGStmt interface {
	FGNode
	//TODO: should have memebers similar to FGExpr
	IsTerminateStmt() bool
	Typing(ds []Decl, gamma map[Name]TypeBase, stupid bool) TypeBase
}

type FGReturnStmt struct {
	ExprBody FGExpr
}

func (F FGReturnStmt) Typing(ds []Decl, gamma map[Name]TypeBase, stupid bool) TypeBase {
	return F.ExprBody.Typing(ds, gamma, stupid)
}

func (F FGReturnStmt) String() string {
	return "return " + F.ExprBody.String()
}

func (F FGReturnStmt) IsTerminateStmt() bool {
	return true
}


type FGCaseSelectStmt struct {
	Cases       []FGCaseBody
	DefaultCase []FGStmt
}

func (F FGCaseSelectStmt) Typing(ds []Decl, gamma map[Name]TypeBase, allowStupid bool) TypeBase {
	var ret_ty TypeBase = nil
	for _, casebody := range F.Cases {
		ty := casebody.Typing(ds, gamma, allowStupid)
		if ret_ty == nil {
			ret_ty = ty
		}
		if !ret_ty.Equals(ty) {
			panic("Type mismatch in select-case statement: previously was " + ret_ty.String() + ", but got " + ty.String() + " later on.")
		}
	}
	if F.DefaultCase != nil && len(F.DefaultCase) > 0 {
		ty := MethBody{StmtBody: F.DefaultCase}.Typing(ds, gamma, allowStupid)
		if ret_ty == nil {
			ret_ty = ty
		}
		if !ret_ty.Equals(ty) {
			panic("Type mismatch in select-case statement: previously was " + ret_ty.String() + ", but got " + ty.String() + " later on.")
		}
	}
	return ret_ty
}

type FGCaseBody struct {
	CaseGuard FGCaseGuard
	Body      []FGStmt
}

func (F FGCaseBody) Typing(ds []Decl, gamma map[Name]TypeBase, allowStupid bool) TypeBase {
	new_stmts := []FGStmt{F.CaseGuard.ToFGStmt()}
	new_stmts = append(new_stmts, F.Body...)
	return MethBody{new_stmts}.Typing(ds, gamma, allowStupid)
}

func getStringOfStmtArray(stmts []FGStmt) string {
	ret := ""
	for i := 0; i < len(stmts)-1; i++ {
		ret += stmts[i].String() + ";"
	}
	if len(stmts) > 0 {
		ret += stmts[len(stmts)-1].String()
	}
	return ret
}

func (F FGCaseSelectStmt) String() string {
	ret := "select {\n"
	for _, c := range F.Cases {
		ret += "case " + c.CaseGuard.String() + ":\n"
		ret += getStringOfStmtArray(c.Body)
		ret += "\n"
	}
	if F.DefaultCase != nil {
		ret += "default:\n"
		ret += getStringOfStmtArray(F.DefaultCase)
	}
	ret += "}\n"
	return ret
}

func (F FGCaseSelectStmt) IsTerminateStmt() bool {
	return true
}

type FGChCloseStmt struct {
	Body FGExpr
}

func (F FGChCloseStmt) Typing(ds []Decl, gamma map[Name]TypeBase, stupid bool) TypeBase {
	return nil
}

func (F FGChCloseStmt) String() string {
	return "close(" + F.Body.String() + ")"
}

func (F FGChCloseStmt) IsTerminateStmt() bool {
	return false
}

type FGChDispatchStmt struct {
	LeftExpr  FGExpr
	RightExpr FGExpr
}

func (F FGChDispatchStmt) Typing(ds []Decl, gamma map[Name]TypeBase, stupid bool) TypeBase {
	return nil
}

func (F FGChDispatchStmt) String() string {
	return F.LeftExpr.String() + "<-" + F.RightExpr.String()
}

func (F FGChDispatchStmt) IsTerminateStmt() bool {
	return false
}

type FGAssignmentStmt struct {
	VarName     Name
	IsUnusedVar bool
	Body        FGExpr
}

func (F FGAssignmentStmt) Typing(ds []Decl, gamma map[Name]TypeBase, stupid bool) TypeBase {
	ty_e := F.Body.Typing(ds, gamma, stupid)
	if _, ok := gamma[F.VarName]; ok {
		panic("Variable redefined: " + F.VarName) //Do we expect a panic here?
	} else {
		gamma[F.VarName] = ty_e
	}
	return ty_e
}

func (F FGAssignmentStmt) String() string {
	if !F.IsUnusedVar {
		return F.VarName + " := " + F.Body.String()
	} else {
		return "_ = " + F.Body.String()
	}
}

func (F FGAssignmentStmt) IsTerminateStmt() bool {
	return false
}

type MethBody struct {
	StmtBody []FGStmt
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

func (m MethBody) Typing(ds []Decl, gamma Gamma, stupid bool) TypeBase {
	tmp_gamma := make(map[Name]TypeBase)
	for k, v := range gamma {
		tmp_gamma[k] = v
	}
	var ret_ty TypeBase = nil
	for _, stmt := range m.StmtBody {
		ty := stmt.Typing(ds, tmp_gamma, stupid) //this may change delta and gamma.
		if stmt.IsTerminateStmt() {
			if ret_ty == nil {
				ret_ty = ty
			} else if !ty.Equals(ret_ty) {
				panic("Returned type doesn't match in method body.")
			}
		}
	}
	return ret_ty
}

type GoroutineStmt struct {
	CallExpr Call
}

func (g GoroutineStmt) Typing(ds []Decl, gamma map[Name]TypeBase, stupid bool) TypeBase {
	return g.CallExpr.Typing(ds, gamma, stupid)
}

func (g GoroutineStmt) String() string {
	return "go " + g.CallExpr.String()
}

func (g GoroutineStmt) IsTerminateStmt() bool {
	return false
}

type OtherStmt struct {
	Body string
}

func (o OtherStmt) Typing(ds []Decl, gamma map[Name]TypeBase, stupid bool) TypeBase {
	panic("stmt does not exist in fcg.")
}

func (o OtherStmt) String() string {
	return o.Body
}

func (o OtherStmt) IsTerminateStmt() bool {
	panic("stmt does not exist in fcg.")
}
