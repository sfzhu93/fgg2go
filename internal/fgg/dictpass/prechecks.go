package dictpass

import (
	"github.com/sfzhu93/fgg2go/internal/fgg"
)

type PreCheck struct {
	hasAssert bool
}

func (this PreCheck) VisitExpr(expr fgg.FGGExpr) {
	switch e := expr.(type) {
	case fgg.Call:
		this.VisitExpr(e.E_recv)
		for _, arg := range e.GetArgs() {
			this.VisitExpr(arg)
		}
	case fgg.Assert:
		this.hasAssert = true
		this.VisitExpr(e.GetExpr())
	case fgg.Select:
		this.VisitExpr(e.GetExpr())
	case fgg.StructLit:
		for _, field := range e.GetElems() {
			this.VisitExpr(field)
		}
	case fgg.MakeChan:
		//do nothing
	case fgg.Variable:
	case fgg.ChanRecv:
		this.VisitExpr(e.Body)
	default:
		panic("not implemented!")
	}
}

func (this PreCheck) VisitMethBody(body *fgg.MethBody) {
	def_new := make(map[string]struct{})

	for _, stmt_ := range body.StmtBody {
		switch stmt := stmt_.(type) {
		case fgg.FGGAssignmentStmt:
			this.VisitExpr(stmt.Body)
			def_new[stmt.VarName] = struct{}{}
		case fgg.ReturnStmt:
			this.VisitExpr(stmt.ExprBody)
		case fgg.FGGCaseSelectStmt:
			for _, c := range stmt.Cases {
				switch cg := c.CaseGuard.(type) {
				case fgg.FGGDispatchCaseGuard:
					this.VisitExpr(cg.RightExpr)
					this.VisitExpr(cg.LeftExpr)
				case fgg.FGGReceiveCaseGuard:
					this.VisitExpr(cg.RightExpr)
					def_new[cg.LeftName] = struct{}{}
				}
				this.VisitMethBody(&fgg.MethBody{c.Body})
				if recvcase, ok := c.CaseGuard.(fgg.FGGReceiveCaseGuard); ok {
					this.VisitExpr(recvcase.RightExpr)
				}
			}
			if stmt.DefaultCase != nil {
				this.VisitMethBody(&fgg.MethBody{stmt.DefaultCase})
			}
		case fgg.FGGChCloseStmt:
			this.VisitExpr(stmt.Body)
		case fgg.FGGChDispatchStmt:
			this.VisitExpr(stmt.LeftExpr)
			this.VisitExpr(stmt.RightExpr)
		case fgg.GoroutineStmt:
			this.VisitExpr(stmt.CallExpr)
		default:
			panic("not implemented!")
		}
	}
}

func (this PreCheck) VisitMethodDecl(decl fgg.MethDecl) {
	this.VisitMethBody(&decl.BodyStmt)
}

func (this PreCheck) VisitProgram(program fgg.FGGProgram) {
	for _, decl := range program.GetDecls() {
		switch d := decl.(type) {
		case fgg.MethDecl:
			this.VisitMethodDecl(d)
		default:
			continue
		}
	}
	this.VisitExpr(program.GetMain().(fgg.FGGExpr))
}
