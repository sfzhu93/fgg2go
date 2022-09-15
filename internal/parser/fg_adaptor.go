/*
 * TODO: fix many magic numbers and other sloppy hacks
 */

package parser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/sfzhu93/fgg2go/internal/base"
	"github.com/sfzhu93/fgg2go/internal/base/testutils"
	"github.com/sfzhu93/fgg2go/internal/fg"
	"github.com/sfzhu93/fgg2go/internal/parser/util"
	"github.com/sfzhu93/fgg2go/parser/fg/parser"
)

var _ = fmt.Errorf
var _ = reflect.Append

// Convert ANTLR generated CST to an fg.FGNode AST
type FGAdaptor struct {
	*parser.BaseFGListener
	stack []fg.FGNode // Because Listener methods don't return...
}

var _ base.Adaptor = &FGAdaptor{}

func (a *FGAdaptor) push(n fg.FGNode) {
	a.stack = append(a.stack, n)
}

func (a *FGAdaptor) pop() fg.FGNode {
	if len(a.stack) < 1 {
		panic(testutils.PARSER_PANIC_PREFIX + "Stack is empty")
	}
	res := a.stack[len(a.stack)-1]
	a.stack = a.stack[:len(a.stack)-1]
	return res
}

// strictParse means panic upon any parsing error -- o/w error recovery is attempted
func (a *FGAdaptor) Parse(strictParse bool, input string) base.Program {
	is := antlr.NewInputStream(input)
	var lexer antlr.Lexer
	if strictParse { // https://stackoverflow.com/questions/51683104/how-to-catch-minor-errors
		lexer = FGBailLexer{parser.NewFGLexer(is)} // FIXME: not working -- e.g., incr{1}, bad token
	} else {
		lexer = parser.NewFGLexer(is)
	}
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewFGParser(stream)
	if strictParse {
		p.RemoveErrorListeners()
		p.SetErrorHandler(&util.StrictErrorStrategy{})
	}
	antlr.ParseTreeWalkerDefault.Walk(a, p.Program())
	return a.pop().(fg.FGProgram)
}

/* "program" */

func (a *FGAdaptor) ExitProgram(ctx *parser.ProgramContext) {
	body := a.pop().(fg.FGExpr)
	ds := []fg.Decl{}
	offset := 0 // TODO: refactor
	printf := false
	c3 := ctx.GetChild(3)                                     // Check if this child is "import"
	foo := ctx.GetChild(ctx.GetChildCount() - 4).GetPayload() // Check if this child is the "=" in "_ = ..."
	if c3_cast, ok := c3.GetPayload().(*antlr.CommonToken); ok &&
		c3_cast.GetText() == "import" {
		if pkg := ctx.GetChild(4).GetPayload().(*antlr.CommonToken).GetText(); pkg != "\"fmt\"" { // TODO: refactor
			panic(testutils.PARSER_PANIC_PREFIX + "The only allowed import is \"fmt\"; found: " + pkg)
		}
		offset = 3
		if cast, ok := foo.(*antlr.CommonToken); !ok || cast.GetText() != "=" { // Looking for: _ = ...
			printf = true
		}
	} else if cast, ok := foo.(*antlr.CommonToken); !ok || cast.GetText() != "=" {
		panic(testutils.PARSER_PANIC_PREFIX + "Missing \"import fmt;\".")
	}
	bar := ctx.GetChild(offset + 3)                                   // Check if this child is "func", i.e., no decls
	if _, ok := bar.GetPayload().(*antlr.BaseParserRuleContext); ok { // If "func", then *antlr.CommonToken
		nds := ctx.GetChild(offset+3).GetChildCount() / 2 // (decl ';')+ -- i.e, includes trailing ';'
		ds = make([]fg.Decl, nds)
		for i := nds - 1; i >= 0; i-- {
			ds[i] = a.pop().(fg.Decl) // Adding backwards
		}
	}
	a.push(fg.NewFGProgram(ds, body, printf))
}

/* "typeDecl" */

// Children: 1=NAME, 2=typeLit
func (a *FGAdaptor) ExitTypeDecl(ctx *parser.TypeDeclContext) {
	t := fg.Type(ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText())
	td := a.pop().(fg.TDecl)
	if s, ok := td.(fg.STypeLit); ok { // N.B. s is a *copy* of td
		/*s.t_S = t
		a.push()*/
		a.push(fg.NewSTypeLit(t, s.GetFieldDecls()))
	} else if c, ok := td.(fg.ITypeLit); ok {
		/*c.t_I = t
		a.push(c)*/
		a.push(fg.NewITypeLit(t, c.GetSpecs()))
	} else {
		panic(testutils.PARSER_PANIC_PREFIX + "Unknown type decl: " + reflect.TypeOf(td).String())
	}
}

/*func (a *FGAdaptor) ExitChannelTypeLit(ctx *parser.InterfaceTypeLitContext) {
	if ctx.GetChildCount() == 1 {
		println("should be bidirectional channel")
		a.push(fg.NewChTypeLit("^", fg.ChBidirection))
	} else if ctx.GetChildCount() == 2{
		println(ctx.GetChild(0))
		println(ctx.GetChild(1))
	}
}*/

/* #StructTypeLit ("typeLit"), "fieldDecls", "fieldDecl" */

// Children: 2=fieldDecls
func (a *FGAdaptor) ExitStructTypeLit(ctx *parser.StructTypeLitContext) {
	fds := []fg.FieldDecl{}
	if ctx.GetChildCount() > 3 {
		nfds := (ctx.GetChild(2).GetChildCount() + 1) / 2 // fieldDecl (';' fieldDecl)*
		fds = make([]fg.FieldDecl, nfds)
		for i := nfds - 1; i >= 0; i-- {
			fd := a.pop().(fg.FieldDecl)
			fds[i] = fd // Adding backwards
		}
	}
	a.push(fg.NewSTypeLit("^", fds)) // "^" to be overwritten in ExitTypeDecl
}

func (a *FGAdaptor) ExitFieldDecl(ctx *parser.FieldDeclContext) {
	f := fg.Name(ctx.GetField().GetText())
	u := a.pop().(fg.TypeBase)
	a.push(fg.NewFieldDecl(f, u))
}

/* "methDecl", "paramDecl" */

func (a *FGAdaptor) ExitMethDecl(ctx *parser.MethDeclContext) {
	// Reverse order
	top := a.pop()

	stmt := makeMethBody(top)
	g := a.pop().(fg.Sig)
	recv := a.pop().(fg.ParamDecl)
	a.push(fg.NewMDecl(recv, g.GetMethod(), g.GetParamDecls(), g.GetReturn(), nil, stmt.StmtBody))
}

func makeMethBody(top fg.FGNode) fg.MethBody {
	stmt := fg.MethBody{}
	if body, ok := top.(fg.MethBody); ok {
		stmt = body
	} else {
		stmt.StmtBody = []fg.FGStmt{top.(fg.FGStmt)}
	}
	return stmt
}

//methBody   : RETURN expr                            # MethReturn
func (a *FGAdaptor) ExitMethReturn(ctx *parser.MethReturnContext) {
	e := a.pop().(fg.FGExpr)
	a.push(fg.FGReturnStmt{
		e,
	})
}

//           | stmt ';' methBody                      # Sequence
func (a *FGAdaptor) ExitSequence(ctx *parser.SequenceContext) {
	StmtOrSeq := a.pop()
	if methbody, ok := StmtOrSeq.(fg.MethBody); ok {
		methbody.StmtBody = append([]fg.FGStmt{a.pop().(fg.FGStmt)}, methbody.StmtBody...)
		a.push(methbody)
	} else {
		methbody = fg.MethBody{StmtBody: []fg.FGStmt{a.pop().(fg.FGStmt), StmtOrSeq.(fg.FGStmt)}}
		a.push(methbody)
	}
}

//           | SELECT '{' (CASE caseGuard ':' methBody)* (DEFAULT ':' methBody)? '}' #CaseSelect
func (a *FGAdaptor) ExitCaseSelect(ctx *parser.CaseSelectContext) {
	cnt := ctx.GetChildCount()
	println(cnt)
	for i := 0; i < cnt; i++ {
		print(strconv.FormatInt(int64(i), 10) + ": ")
		if child, ok := ctx.GetChild(i).(*antlr.TerminalNodeImpl); ok {
			println(child.GetText())
		} else {
			println(ctx.GetChild(i))
		}
	}

	ret := fg.FGCaseSelectStmt{
		Cases:       []fg.FGCaseBody{},
		DefaultCase: nil,
	}
	if cnt > 3 { //select not empty
		if child, ok := ctx.GetChild(cnt - 4).(*antlr.TerminalNodeImpl); ok && child.GetText() == "default" {
			ret.DefaultCase = makeMethBody(a.pop()).StmtBody
			cnt -= 5
		} else {
			cnt -= 2
		}
		//println(cnt)
		for cnt > 2 {
			//print(strconv.FormatInt(int64(cnt), 10) + ": ")
			//println(ctx.GetChild(cnt))
			methbody := makeMethBody(a.pop())
			cnt -= 2
			//print(strconv.FormatInt(int64(cnt), 10) + ": ")
			//println(ctx.GetChild(cnt))
			caseguard := a.pop().(fg.FGCaseGuard)
			cnt -= 2
			ret.Cases = append([]fg.FGCaseBody{fg.FGCaseBody{
				CaseGuard: caseguard,
				Body:      methbody.StmtBody,
			}}, ret.Cases...)
		}
	}
	a.push(ret)
}

func (a *FGAdaptor) ExitMethCallPrime(ctx *parser.MethCallPrimeContext) {
	args := []fg.FGExpr{}
	if ctx.GetChildCount() > 3 { // TODO: refactor as ctx.GetArgs() != nil -- and child-count-checks elsewhere
		nargs := (ctx.GetChild(3).GetChildCount() + 1) / 2 // e.g., e ',' e ',' e
		args = make([]fg.FGExpr, nargs)
		for i := nargs - 1; i >= 0; i-- {
			args[i] = a.pop().(fg.FGExpr) // Adding backwards
		}
	}
	m := fg.Name(ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText())
	a.push(fg.NewCall(nil, m, args))
}

func (a *FGAdaptor) ExitChClose(ctx *parser.ChCloseContext) {
	a.push(fg.FGChCloseStmt{Body: a.pop().(fg.FGExpr)})
}

func (a *FGAdaptor) ExitChDispatch(ctx *parser.ChDispatchContext) {
	right := a.pop().(fg.FGExpr)
	left := a.pop().(fg.FGExpr)
	a.push(fg.FGChDispatchStmt{
		LeftExpr:  left,
		RightExpr: right,
	})
}

func (a *FGAdaptor) ExitGoroutine(ctx *parser.GoroutineContext) {
	right := a.pop().(fg.Call)
	left := a.pop().(fg.FGExpr)
	right.ExprRecv = left
	a.push(fg.GoroutineStmt{CallExpr: right})
}

func (a *FGAdaptor) ExitAssignment(ctx *parser.AssignmentContext) {
	lvalue := ctx.GetLvalue().GetText()
	expr := a.pop().(fg.FGExpr)
	a.push(fg.FGAssignmentStmt{
		VarName: lvalue,
		Body:    expr,
	})
}

func (a *FGAdaptor) ExitGuardDispatch(ctx *parser.GuardDispatchContext) {
	right := a.pop().(fg.FGExpr)
	left := a.pop().(fg.FGExpr)
	a.push(fg.FGDispatchCaseGuard{
		LeftExpr:  left,
		RightExpr: right,
	})
}

func (a *FGAdaptor) ExitGuardReceive(ctx *parser.GuardReceiveContext) {
	right := a.pop().(fg.FGExpr)
	left := ctx.GetLvalue().GetText()
	a.push(fg.FGReceiveCaseGuard{
		LeftName:  left,
		RightExpr: right,
	})
}

// Cf. ExitFieldDecl
func (a *FGAdaptor) ExitParamDecl(ctx *parser.ParamDeclContext) {
	x := ctx.GetVari().GetText()
	t := a.pop().(fg.TypeBase) //fg.Type(ctx.GetTyp().GetText())
	a.push(fg.NewParamDecl(x, t))
}

/* #InterfaceTypeLit ("typeLit"), "specs", #SigSpec ("spec"), #InterfaceSpec ("spec"), "sig" */

// Cf. ExitStructTypeLit
func (a *FGAdaptor) ExitInterfaceTypeLit(ctx *parser.InterfaceTypeLitContext) {
	ss := []fg.Spec{}
	if ctx.GetChildCount() > 3 {
		nss := (ctx.GetChild(2).GetChildCount() + 1) / 2 // e.g., s ';' s ';' s
		ss = make([]fg.Spec, nss)
		for i := nss - 1; i >= 0; i-- {
			s := a.pop().(fg.Spec)
			ss[i] = s // Adding backwards
		}
	}
	a.push(fg.NewITypeLit("^", ss)) // "^" to be overwritten in ExitTypeDecl
}

func (a *FGAdaptor) ExitSigSpec(ctx *parser.SigSpecContext) {
	// No action -- Sig is at a.stack[len(a.stack)-1]
}

func (a *FGAdaptor) ExitInterfaceSpec(ctx *parser.InterfaceSpecContext) {
	t := fg.Type(ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText())
	a.push(t)
}

func (a *FGAdaptor) ExitSig(ctx *parser.SigContext) {
	m := ctx.GetMeth().GetText()
	// Reverse order
	t := fg.Type(ctx.GetRet().GetText())
	pds := []fg.ParamDecl{}
	if ctx.GetChildCount() > 4 {
		npds := (ctx.GetChild(2).GetChildCount() + 1) / 2 // e.g., pd ',' pd ',' pd
		pds = make([]fg.ParamDecl, npds)
		for i := npds - 1; i >= 0; i-- {
			pds[i] = a.pop().(fg.ParamDecl) // Adding backwards
		}
	}
	a.push(fg.NewSig(m, pds, t))
}

/* "expr": #Variable, #StructLit, #Select, #Call, #Assert, #Sprintf */

func (a *FGAdaptor) ExitVariable(ctx *parser.VariableContext) {
	id := fg.Name(ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText())
	a.push(fg.NewVariable(id))
}

// Children: 0=typ (*antlr.TerminalNodeImpl), 1='{', 2=exprs (*parser.ExprsContext), 3='}'
// N.B. ExprsContext is a "helper" Context, actual exprs are its children
func (a *FGAdaptor) ExitStructLit(ctx *parser.StructLitContext) {
	t := fg.Type(ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText())
	es := []fg.FGExpr{}
	if ctx.GetChildCount() > 3 {
		nes := (ctx.GetChild(2).GetChildCount() + 1) / 2 // e.g., 'x' ',' 'y' ',' 'z'
		es = make([]fg.FGExpr, nes)
		for i := nes - 1; i >= 0; i-- {
			es[i] = a.pop().(fg.FGExpr) // Adding backwards
		}
	}
	a.push(fg.NewStructLit(t, es))
}

func (a *FGAdaptor) ExitSelect(ctx *parser.SelectContext) {
	e := a.pop().(fg.FGExpr)
	f := fg.Name(ctx.GetChild(2).(*antlr.TerminalNodeImpl).GetText())
	a.push(fg.NewSelect(e, f))
}

//copied from goroutine
func (a *FGAdaptor) ExitCall(ctx *parser.CallContext) {
	right := a.pop().(fg.Call)
	left := a.pop().(fg.FGExpr)
	right.ExprRecv = left
	a.push(right)
}

func (a *FGAdaptor) ExitAssert(ctx *parser.AssertContext) {
	t := fg.Type(ctx.GetChild(3).(*antlr.TerminalNodeImpl).GetText())
	e := a.pop().(fg.FGExpr)
	a.push(fg.NewAssert(e, t))
}

// TODO: check for import "fmt"
func (a *FGAdaptor) ExitSprintf(ctx *parser.SprintfContext) { //TODO: changed in fcgg
	panic("haven't implemented for fcg")
	var format string = ctx.GetChild(4).(*antlr.TerminalNodeImpl).GetText()
	nargs := (ctx.GetChildCount() - 6) / 2 // Because of the comma
	args := make([]fg.FGExpr, nargs)
	for i := nargs - 1; i >= 0; i-- {
		tmp := a.pop()
		args[i] = tmp.(fg.FGExpr)
	}
	a.push(fg.NewSprintf(format, args))
}

func (a *FGAdaptor) ExitMakeChan(ctx *parser.MakeChanContext) {
	ChTy := a.pop().(fg.ChannelType)
	a.push(fg.MakeChan{ChTy})
}

func (a *FGAdaptor) ExitChanRecv(ctx *parser.ChanRecvContext) {
	a.push(fg.ChanRecv{Body: a.pop().(fg.FGExpr)})
}

/* For "strict" parsing, *lexer* errors */

func (a *FGAdaptor) ExitType_(ctx *parser.Type_Context) {
	chdcnt := ctx.GetChildCount()
	if chdcnt == 1 {
		tyName := ctx.GetTyName().GetText()
		a.push(fg.Type(tyName))
	} else {
		elemTy := a.pop().(fg.TypeBase)
		chTy := a.pop().(fg.ChannelType)
		chTy.ElementType = elemTy
		a.push(chTy)
	}
}

func (a *FGAdaptor) ExitChannelTypeLit(ctx *parser.ChannelTypeLitContext) {
	txt := ctx.GetText()
	switch txt {
	case "<-chan":
		a.push(fg.ChannelType{
			ElementType: nil,
			ChType:      fg.ChSendOnly,
		})
	case "chan<-":
		a.push(fg.ChannelType{
			ElementType: nil,
			ChType:      fg.ChRecvOnly,
		})
	case "chan":
		a.push(fg.ChannelType{
			ElementType: nil,
			ChType:      fg.ChBidirection,
		})
	}

}

type FGBailLexer struct {
	*parser.FGLexer
}

// FIXME: not working -- e.g., incr{1}, bad token
// Want to "override" *BaseLexer.Recover -- XXX that's not how Go works (because BaseLexer is a struct, not interface)
func (b *FGBailLexer) Recover(re antlr.RecognitionException) {
	message := "lex error after token " + re.GetOffendingToken().GetText() +
		" at position " + strconv.Itoa(re.GetOffendingToken().GetStart())
	panic(message)
}

/*public FGBailLexer(ICharStream input) : base(input) { }

public override void Recover(LexerNoViableAltException e)
{
	string message = string.Format("lex error after token {0} at position {1}", _lasttoken.Text, e.StartIndex);
	BasicEnvironment.SyntaxError = message;
	BasicEnvironment.ErrorStartIndex = e.StartIndex;
	throw new ParseCanceledException(BasicEnvironment.SyntaxError);
}*/
