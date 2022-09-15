package dictpass

import "github.com/sfzhu93/fgg2go/internal/fg"

type UnusedVarRemover struct {
}

func (this UnusedVarRemover) VisitExpr(expr fg.FGExpr, defset map[string]struct{}, usedset map[string]struct{}) {
	switch e := expr.(type) {
	case fg.Call:
		this.VisitExpr(e.ExprRecv, defset, usedset)
		for _, arg := range e.GetArgs() {
			this.VisitExpr(arg, defset, usedset)
		}
	case fg.Assert:
		this.VisitExpr(e.GetExpr(), defset, usedset)
	case fg.Select:
		this.VisitExpr(e.GetExpr(), defset, usedset)
	case fg.StructLit:
		for _, field := range e.GetElems() {
			this.VisitExpr(field, defset, usedset)
		}
	case fg.MakeChan:case fg.FuncDecl:
		//do nothing
	case fg.Variable:
		usedset[e.String()] = struct{}{}
	case fg.ChanRecv:
		this.VisitExpr(e.Body, defset, usedset)
	default:
		panic("not implemented!")
	}
}

func (this UnusedVarRemover) VisitMethBody(body *fg.MethBody, defset map[string]struct{}, usedset map[string]struct{}) {
	def_new := make(map[string]struct{})
	for k, _ := range defset {
		def_new[k] = struct{}{}
	}
	for i, stmt_ := range body.StmtBody {
		switch stmt := stmt_.(type) {
		case fg.FGAssignmentStmt:
			this.VisitExpr(stmt.Body, def_new, usedset)
			def_new[stmt.VarName] = struct{}{}
		case fg.FGReturnStmt:
			this.VisitExpr(stmt.ExprBody, def_new, usedset)
		case fg.FGCaseSelectStmt:
			for j, c := range stmt.Cases {
				switch cg := c.CaseGuard.(type) {
				case fg.FGDispatchCaseGuard:
					this.VisitExpr(cg.RightExpr, def_new, usedset)
					this.VisitExpr(cg.LeftExpr, def_new, usedset)
				case fg.FGReceiveCaseGuard:
					this.VisitExpr(cg.RightExpr, def_new, usedset)
					def_new[cg.LeftName] = struct{}{}
				}
				this.VisitMethBody(&fg.MethBody{c.Body}, def_new, usedset)
				if recvcase, ok := c.CaseGuard.(fg.FGReceiveCaseGuard); ok {
					if _, ok2 := usedset[recvcase.LeftName]; !ok2 {
						println(recvcase.LeftName + " unused")
						recvcase.IsUnusedVar = true
						c.CaseGuard = recvcase
						stmt.Cases[j] = c
						body.StmtBody[i] = stmt
					}
				}
			}
			if stmt.DefaultCase != nil {
				this.VisitMethBody(&fg.MethBody{stmt.DefaultCase}, def_new, usedset)
			}
		case fg.FGChCloseStmt:
			this.VisitExpr(stmt.Body, def_new, usedset)
		case fg.FGChDispatchStmt:
			this.VisitExpr(stmt.LeftExpr, def_new, usedset)
			this.VisitExpr(stmt.RightExpr, def_new, usedset)
		case fg.GoroutineStmt:
			this.VisitExpr(stmt.CallExpr, def_new, usedset)
		default:
			panic("not implemented!")
		}
	}
	for i := range body.StmtBody {
		switch stmt := body.StmtBody[i].(type) {
		case fg.FGAssignmentStmt:
			if _, ok := usedset[stmt.VarName]; !ok {
				println(stmt.VarName, " unused")
				stmt.IsUnusedVar = true
				body.StmtBody[i] = stmt
			}
		}
	}
}
func (this UnusedVarRemover) VisitMethodDecl(decl fg.MethDecl) {
	defset := make(map[string]struct{})
	usedset := make(map[string]struct{})
	this.VisitMethBody(&decl.StmtBody, defset, usedset)
}
