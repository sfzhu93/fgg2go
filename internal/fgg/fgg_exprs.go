/*
 * This file contains defs for "concrete" syntax w.r.t. exprs.
 * Base ("abstract") types, interfaces, etc. are in fg.go.
 */

package fgg

import (
	"fmt"
	"reflect"
	"strings"
)

var _ = fmt.Errorf
var _ = reflect.Append
var _ = strings.Compare

/* Public constructors */

func NewVariable(id Name) Variable                            { return Variable{id} }
func NewStructLit(u_S TNamed, es []FGGExpr) StructLit         { return StructLit{u_S, es} }
func NewSelect(e FGGExpr, f Name) Select                      { return Select{e, f} }
func NewCall(e FGGExpr, m Name, us []Type, es []FGGExpr) Call { return Call{e, m, us, es} }
func NewAssert(e FGGExpr, t Type) Assert                      { return Assert{e, t} }
func NewString(v string) StringLit                            { return StringLit{v} }
func NewSprintf(format string, args []FGGExpr) Sprintf        { return Sprintf{format, args} }

/* Variable */

type Variable struct {
	name Name
}

var _ FGGExpr = Variable{}

func (x Variable) GetName() Name { return x.name }

func (x Variable) Subs(m map[Variable]FGGExpr) FGGExpr {
	res, ok := m[x]
	if !ok {
		panic("Unknown var: " + x.String())
	}
	return res
}

func (x Variable) TSubs(subs map[TParam]Type) FGGExpr {
	return x
}

func (x Variable) Eval(ds []Decl) (FGGExpr, string) {
	panic("Cannot evaluate free variable: " + x.name)
}

// TODO: refactor as Typing and StupidTyping? (clearer than bool param)
func (x Variable) Typing(ds []Decl, delta Delta, gamma Gamma,
	allowStupid bool) Type {
	res, ok := gamma[x.name]
	if !ok {
		panic("Var not in env: " + x.String())
	}
	return res
}

func (x Variable) TypingWithEnv(env *Env, delta Delta, gamma Gamma,
	allowStupid bool) Type {
	res, ok := gamma[x.name]
	if !ok {
		panic("Var not in env: " + x.String())
	}
	return res
}

// From base.Expr
func (x Variable) IsValue() bool {
	return false
}

func (x Variable) CanEval(ds []Decl) bool {
	return false
}

func (x Variable) String() string {
	return x.name
}

func (x Variable) ToGoString(ds []Decl) string {
	return x.name
}

/* StructLit */

type StructLit struct {
	U_S   TNamed // u.t is a t_S
	Elems []FGGExpr
}

var _ FGGExpr = StructLit{}

func (s StructLit) GetNamedType() TNamed { return s.U_S }
func (s StructLit) GetElems() []FGGExpr  { return s.Elems }

func (s StructLit) Subs(subs map[Variable]FGGExpr) FGGExpr {
	es := make([]FGGExpr, len(s.Elems))
	for i := 0; i < len(s.Elems); i++ {
		es[i] = s.Elems[i].Subs(subs)
	}
	return StructLit{s.U_S, es}
}

func (s StructLit) TSubs(subs map[TParam]Type) FGGExpr {
	es := make([]FGGExpr, len(s.Elems))
	for i := 0; i < len(s.Elems); i++ {
		es[i] = s.Elems[i].TSubs(subs)
	}
	return StructLit{s.U_S.TSubs(subs).(TNamed), es}
}

func (s StructLit) Eval(ds []Decl) (FGGExpr, string) {
	es := make([]FGGExpr, len(s.Elems))
	done := false
	var rule string
	for i := 0; i < len(s.Elems); i++ {
		v := s.Elems[i]
		if !done && !v.IsValue() {
			v, rule = v.Eval(ds)
			done = true
		}
		es[i] = v
	}
	if done {
		return StructLit{s.U_S, es}, rule
	} else {
		panic("Cannot reduce: " + s.String())
	}
}

func (s StructLit) Typing(ds []Decl, delta Delta, gamma Gamma,
	allowStupid bool) Type {
	s.U_S.Ok(ds, delta)
	fs := fields(ds, s.U_S)
	if len(s.Elems) != len(fs) {
		var b strings.Builder
		b.WriteString("Arity mismatch: args=[")
		writeExprs(&b, s.Elems)
		b.WriteString("], fields=[")
		writeFieldDecls(&b, fs)
		b.WriteString("]\n\t")
		b.WriteString(s.String())
		panic(b.String())
	}
	for i := 0; i < len(s.Elems); i++ {
		u := s.Elems[i].Typing(ds, delta, gamma, allowStupid)
		r := fs[i].u
		if !u.ImplsDelta(ds, delta, r) {
			panic("Arg expr must implement field type: arg=" + u.String() +
				", field=" + r.String() + "\n\t" + s.String())
		}
	}
	return s.U_S
}

func (s StructLit) TypingWithEnv(env *Env, delta Delta, gamma Gamma,
	allowStupid bool) Type {
	s.U_S.Ok(env.fggDecls, delta)
	fs := fields(env.fggDecls, s.U_S)
	if len(s.Elems) != len(fs) {
		var b strings.Builder
		b.WriteString("Arity mismatch: args=[")
		writeExprs(&b, s.Elems)
		b.WriteString("], fields=[")
		writeFieldDecls(&b, fs)
		b.WriteString("]\n\t")
		b.WriteString(s.String())
		panic(b.String())
	}
	for i := 0; i < len(s.Elems); i++ {
		u := s.Elems[i].TypingWithEnv(env, delta, gamma, allowStupid)
		r := fs[i].u
		if !u.ImplsDeltaWithEnv(env, delta, r) {
			panic("Arg expr must implement field type: arg=" + u.String() +
				", field=" + r.String() + "\n\t" + s.String())
		}
	}
	return s.U_S
}

// From base.Expr
func (s StructLit) IsValue() bool {
	for _, v := range s.Elems {
		if !v.IsValue() {
			return false
		}
	}
	return true
}

func (s StructLit) CanEval(ds []Decl) bool {
	for _, v := range s.Elems {
		if v.CanEval(ds) {
			return true
		} else if !v.IsValue() {
			return false
		}
	}
	return false
}

func (s StructLit) String() string {
	var b strings.Builder
	b.WriteString(s.U_S.String())
	b.WriteString("{")
	writeExprs(&b, s.Elems)
	b.WriteString("}")
	return b.String()
}

func (s StructLit) ToGoString(ds []Decl) string {
	var b strings.Builder
	b.WriteString(s.U_S.ToGoString(ds))
	b.WriteString("{")
	td := GetTDecl(ds, s.U_S.TName).(STypeLit)
	if len(s.Elems) > 0 {
		b.WriteString(td.fDecls[0].field)
		b.WriteString(":")
		b.WriteString(s.Elems[0].ToGoString(ds))
		for i, v := range s.Elems[1:] {
			b.WriteString(", ")
			b.WriteString(td.fDecls[i+1].field)
			b.WriteString(":")
			b.WriteString(v.ToGoString(ds))
		}
	}
	b.WriteString("}")
	return b.String()
}

/* Select */

type Select struct {
	E_S   FGGExpr
	field Name
}

var _ FGGExpr = Select{}

func (s Select) GetExpr() FGGExpr { return s.E_S }
func (s Select) GetField() Name   { return s.field }

func (s Select) Subs(subs map[Variable]FGGExpr) FGGExpr {
	return Select{s.E_S.Subs(subs), s.field}
}

func (s Select) TSubs(subs map[TParam]Type) FGGExpr {
	return Select{s.E_S.TSubs(subs), s.field}
}

func (s Select) Eval(ds []Decl) (FGGExpr, string) {
	if !s.E_S.IsValue() {
		e, rule := s.E_S.Eval(ds)
		return Select{e, s.field}, rule
	}
	v := s.E_S.(StructLit)
	fds := fields(ds, v.U_S)
	for i := 0; i < len(fds); i++ {
		if fds[i].field == s.field {
			return v.Elems[i], "Select"
		}
	}
	panic("Field not found: " + s.field + "\n\t" + s.String())
}

func (s Select) Typing(ds []Decl, delta Delta, gamma Gamma,
	allowStupid bool) Type {
	u := s.E_S.Typing(ds, delta, gamma, allowStupid)
	if !IsStructType(ds, u) {
		panic("Illegal select on expr of non-struct type: " + u.String() +
			"\n\t" + s.String())
	}
	fds := fields(ds, u.(TNamed))
	for _, v := range fds {
		if v.field == s.field {
			return v.u
		}
	}
	panic("Field " + s.field + " not found in type: " + u.String() +
		"\n\t" + s.String())
}

func (s Select) TypingWithEnv(env *Env, delta Delta, gamma Gamma,
	allowStupid bool) Type {
	u := s.E_S.TypingWithEnv(env, delta, gamma, allowStupid)
	if !env.IsStructType(u) {
		panic("Illegal select on expr of non-struct type: " + u.String() +
			"\n\t" + s.String())
	}
	fds := fields(env.fggDecls, u.(TNamed))
	for _, v := range fds {
		if v.field == s.field {
			return v.u
		}
	}
	panic("Field " + s.field + " not found in type: " + u.String() +
		"\n\t" + s.String())
}

// From base.Expr
func (s Select) IsValue() bool {
	return false
}

func (s Select) CanEval(ds []Decl) bool {
	if s.E_S.CanEval(ds) {
		return true
	} else if !s.E_S.IsValue() {
		return false
	}
	for _, v := range fields(ds, s.E_S.(StructLit).U_S) { // N.B. "purely operational", no typing aspect
		if v.field == s.field {
			return true
		}
	}
	return false
}

func (s Select) String() string {
	return s.E_S.String() + "." + s.field
}

func (s Select) ToGoString(ds []Decl) string {
	return s.E_S.ToGoString(ds) + "." + s.field
}

/* Call */

type Call struct {
	E_recv FGGExpr
	Meth   Name
	t_args []Type // Rename UArgs?
	Args   []FGGExpr
}

var _ FGGExpr = Call{}

func (c Call) GetRecv() FGGExpr   { return c.E_recv } // Called GetReceiver in fg
func (c Call) GetMethod() Name    { return c.Meth }
func (c Call) GetTArgs() []Type   { return c.t_args }
func (c Call) GetArgs() []FGGExpr { return c.Args }

func (c Call) Subs(subs map[Variable]FGGExpr) FGGExpr {
	e := c.E_recv.Subs(subs)
	args := make([]FGGExpr, len(c.Args))
	for i := 0; i < len(c.Args); i++ {
		args[i] = c.Args[i].Subs(subs)
	}
	return Call{e, c.Meth, c.t_args, args}
}

func (c Call) TSubs(subs map[TParam]Type) FGGExpr {
	targs := make([]Type, len(c.t_args))
	for i := 0; i < len(c.t_args); i++ {
		targs[i] = c.t_args[i].TSubs(subs)
	}
	args := make([]FGGExpr, len(c.Args))
	for i := 0; i < len(c.Args); i++ {
		args[i] = c.Args[i].TSubs(subs)
	}
	return Call{c.E_recv.TSubs(subs), c.Meth, targs, args}
}

func (c Call) Eval(ds []Decl) (FGGExpr, string) {
	if !c.E_recv.IsValue() {
		e, rule := c.E_recv.Eval(ds)
		return Call{e, c.Meth, c.t_args, c.Args}, rule
	}
	args := make([]FGGExpr, len(c.Args))
	done := false
	var rule string
	for i := 0; i < len(c.Args); i++ {
		e := c.Args[i]
		if !done && !e.IsValue() {
			e, rule = e.Eval(ds)
			done = true
		}
		args[i] = e
	}
	if done {
		return Call{c.E_recv, c.Meth, c.t_args, args}, rule
	}
	// c.e and c.args all values
	s := c.E_recv.(StructLit)
	x0, xs, stmts := body(ds, s.U_S, c.Meth, c.t_args) // panics if method not found
	subs := make(map[Variable]FGGExpr)
	subs[Variable{x0.Name}] = c.E_recv
	for i := 0; i < len(xs); i++ {
		subs[Variable{xs[i].Name}] = c.Args[i]
	}
	return stmts.Subs(subs), "Call" // N.B. single combined substitution map slightly different to R-Call
}

func (c Call) Typing(ds []Decl, delta Delta, gamma Gamma, allowStupid bool) Type {
	u0 := c.E_recv.Typing(ds, delta, gamma, allowStupid)
	var g Sig
	t1, _ := MethodsDelta(ds, delta, bounds(delta, u0))
	if tmp, ok := t1[c.Meth]; !ok { // !!! submission version had "methods(m)"
		panic("Method not found: " + c.Meth + " in " + u0.String())
	} else {
		g = tmp
	}
	if len(c.t_args) != len(g.Psi.TFormals) {
		var b strings.Builder
		b.WriteString("Arity mismatch: type actuals=[")
		writeTypes(&b, c.t_args)
		b.WriteString("], formals=[")
		b.WriteString(g.Psi.String())
		b.WriteString("]\n\t")
		b.WriteString(c.String())
		panic(b.String())
	}
	if len(c.Args) != len(g.pDecls) {
		var b strings.Builder
		b.WriteString("Arity mismatch: args=[")
		writeExprs(&b, c.Args)
		b.WriteString("], params=[")
		writeParamDecls(&b, g.pDecls)
		b.WriteString("]\n\t")
		b.WriteString(c.String())
		panic(b.String())
	}
	subs := make(map[TParam]Type) // CHECKME: applying this subs vs. adding to a new delta?  // Cf. MakeEta
	for i := 0; i < len(c.t_args); i++ {
		subs[g.Psi.TFormals[i].Name] = c.t_args[i]
	}
	for i := 0; i < len(c.t_args); i++ {
		u := g.Psi.TFormals[i].U_I.TSubs(subs)
		if !c.t_args[i].ImplsDelta(ds, delta, u) {
			panic("Type actual must implement type formal: actual=" +
				c.t_args[i].String() + ", param=" + u.String() + "\n\t" + c.String())
		}
	}
	for i := 0; i < len(c.Args); i++ {
		// CHECKME: submission version's notation, (~\tau :> ~\rho)[subs], slightly unclear
		u_a := c.Args[i].Typing(ds, delta, gamma, allowStupid)
		//.TSubs(subs)  // !!! submission version, subs also applied to ~tau, ..
		// ..falsely captures "repeat" var occurrences in recursive calls, ..
		// ..e.g., bad monomorph (Box) example.
		// The ~beta morally do not occur in ~tau, they only bind ~rho
		u_p := g.pDecls[i].U.TSubs(subs)
		if !u_a.ImplsDelta(ds, delta, u_p) {
			panic("Arg expr type must implement param type: arg=" + u_a.String() +
				", param=" + u_p.String() + "\n\t" + c.String())
		}
	}
	return g.u_ret.TSubs(subs) // subs necessary, c.psi info (i.e., bounds) will be "lost" after leaving this context
}

func (c Call) TypingWithEnv(env *Env, delta Delta, gamma Gamma, allowStupid bool) Type {
	u0 := c.E_recv.TypingWithEnv(env, delta, gamma, allowStupid)
	var g *Sig
	t1, _ := env.MethodsDelta(delta, bounds(delta, u0))
	if tmp, ok := t1[c.Meth]; !ok { // !!! submission version had "methods(m)"
		panic("Method not found: " + c.Meth + " in " + u0.String())
	} else {
		g = tmp
	}
	if len(c.t_args) != len(g.Psi.TFormals) {
		var b strings.Builder
		b.WriteString("Arity mismatch: type actuals=[")
		writeTypes(&b, c.t_args)
		b.WriteString("], formals=[")
		b.WriteString(g.Psi.String())
		b.WriteString("]\n\t")
		b.WriteString(c.String())
		panic(b.String())
	}
	if len(c.Args) != len(g.pDecls) {
		var b strings.Builder
		b.WriteString("Arity mismatch: args=[")
		writeExprs(&b, c.Args)
		b.WriteString("], params=[")
		writeParamDecls(&b, g.pDecls)
		b.WriteString("]\n\t")
		b.WriteString(c.String())
		panic(b.String())
	}
	subs := make(map[TParam]Type) // CHECKME: applying this subs vs. adding to a new delta?  // Cf. MakeEta
	for i := 0; i < len(c.t_args); i++ {
		subs[g.Psi.TFormals[i].Name] = c.t_args[i]
	}
	for i := 0; i < len(c.t_args); i++ {
		u := g.Psi.TFormals[i].U_I.TSubs(subs)
		if !c.t_args[i].ImplsDeltaWithEnv(env, delta, u) {
			panic("Type actual must implement type formal: actual=" +
				c.t_args[i].String() + ", param=" + u.String() + "\n\t" + c.String())
		}
	}
	for i := 0; i < len(c.Args); i++ {
		// CHECKME: submission version's notation, (~\tau :> ~\rho)[subs], slightly unclear
		u_a := c.Args[i].TypingWithEnv(env, delta, gamma, allowStupid)
		//.TSubs(subs)  // !!! submission version, subs also applied to ~tau, ..
		// ..falsely captures "repeat" var occurrences in recursive calls, ..
		// ..e.g., bad monomorph (Box) example.
		// The ~beta morally do not occur in ~tau, they only bind ~rho
		u_p := g.pDecls[i].U.TSubs(subs)
		if !u_a.ImplsDeltaWithEnv(env, delta, u_p) {
			panic("Arg expr type must implement param type: arg=" + u_a.String() +
				", param=" + u_p.String() + "\n\t" + c.String())
		}
	}
	return g.u_ret.TSubs(subs) // subs necessary, c.psi info (i.e., bounds) will be "lost" after leaving this context
}

// From base.Expr
func (c Call) IsValue() bool {
	return false
}

func (c Call) CanEval(ds []Decl) bool {
	if c.E_recv.CanEval(ds) {
		return true
	} else if !c.E_recv.IsValue() {
		return false
	}
	for _, v := range c.Args {
		if v.CanEval(ds) {
			return true
		} else if !v.IsValue() {
			return false
		}
	}
	u_S := c.E_recv.(StructLit).U_S
	for _, d := range ds { // TODO: factor out GetMethDecl
		if md, ok := d.(MethDecl); ok &&
			md.TRecv == u_S.TName && md.name == c.Meth {
			// Disregard type bounds -- cf. actual typing, methods aux
			return len(md.Psi_recv.TFormals) == len(u_S.UArgs) && // Needed, or also disregard?
				len(md.Psi_meth.TFormals) == len(c.t_args) &&
				len(md.pDecls) == len(c.Args)
		}
	}
	return false
}

func (c Call) String() string {
	var b strings.Builder
	b.WriteString(c.E_recv.String())
	b.WriteString(".")
	b.WriteString(c.Meth)
	b.WriteString("[")
	writeTypes(&b, c.t_args)
	b.WriteString("](")
	writeExprs(&b, c.Args)
	b.WriteString(")")
	return b.String()
}

func (c Call) ToGoString(ds []Decl) string {
	var b strings.Builder
	b.WriteString(c.E_recv.ToGoString(ds))
	b.WriteString(".")
	b.WriteString(c.Meth)
	b.WriteString("(")
	writeToGoTypes(ds, &b, c.t_args)
	b.WriteString(")(")
	writeToGoExprs(ds, &b, c.Args)
	b.WriteString(")")
	return b.String()
}

/* Assert */

type Assert struct {
	E_I    FGGExpr
	U_cast Type
}

var _ FGGExpr = Assert{}

func (a Assert) GetExpr() FGGExpr { return a.E_I }
func (a Assert) GetType() Type    { return a.U_cast }

func (a Assert) Subs(subs map[Variable]FGGExpr) FGGExpr {
	return Assert{a.E_I.Subs(subs), a.U_cast}
}

func (a Assert) TSubs(subs map[TParam]Type) FGGExpr {
	return Assert{a.E_I.TSubs(subs), a.U_cast.TSubs(subs)}
}

func (a Assert) Eval(ds []Decl) (FGGExpr, string) {
	if !a.E_I.IsValue() {
		e, rule := a.E_I.Eval(ds)
		return Assert{e, a.U_cast}, rule
	}
	u_S := a.E_I.(StructLit).U_S
	if !IsStructType(ds, u_S) {
		panic("Non struct type found in struct lit: " + u_S.String())
	}
	if u_S.ImplsDelta(ds, make(map[TParam]Type), a.U_cast) { // Empty Delta -- not super clear in submission version
		return a.E_I, "Assert"
	}
	panic("Cannot reduce: " + a.String())
}

func (a Assert) Typing(ds []Decl, delta Delta, gamma Gamma, allowStupid bool) Type {
	u := a.E_I.Typing(ds, delta, gamma, allowStupid)
	a.U_cast.Ok(ds, delta)
	if IsStructType(ds, u) {
		if allowStupid {
			return a.U_cast
		} else {
			panic("Expr must be an interface type (in a non-stupid context): found " +
				u.String() + " for\n\t" + a.String())
		}
	}
	// u is a TParam or an interface type TName
	if _, ok := a.U_cast.(TParam); ok || IsNamedIfaceType(ds, a.U_cast) {
		return a.U_cast // No further checks -- N.B., Robert said they are looking to refine this
	}
	//fmt.Printf("Assert.Typing: \nu_cast=%#v\ntype of a.e_I=%#v\n", a.U_cast, u)
	// a.u is a struct type TName
	if a.U_cast.ImplsDelta(ds, delta, u) {
		return a.U_cast
	}
	panic("Struct type assertion must implement expr type: asserted=" +
		a.U_cast.String() + ", expr=" + u.String())
}

func (a Assert) TypingWithEnv(env *Env, delta Delta, gamma Gamma, allowStupid bool) Type {
	u := a.E_I.TypingWithEnv(env, delta, gamma, allowStupid)
	a.U_cast.Ok(env.fggDecls, delta)
	if env.IsStructType(u) {
		if allowStupid {
			return a.U_cast
		} else {
			panic("Expr must be an interface type (in a non-stupid context): found " +
				u.String() + " for\n\t" + a.String())
		}
	}
	// u is a TParam or an interface type TName
	if _, ok := a.U_cast.(TParam); ok || env.IsInterfaceType(a.U_cast) {
		return a.U_cast // No further checks -- N.B., Robert said they are looking to refine this
	}
	//fmt.Printf("Assert.TypingWithEnv: \nu_cast=%#v\ntype of a.e_I=%#v\n", a.U_cast, u)
	// a.u is a struct type TName
	if a.U_cast.ImplsDeltaWithEnv(env, delta, u) {
		return a.U_cast
	}
	panic("Struct type assertion must implement expr type: asserted=" +
		a.U_cast.String() + ", expr=" + u.String())
}

// CHECKME: make isStuck alternative? (i.e., bad cast)
// From base.fgg
func (a Assert) IsValue() bool {
	return false
}

func (a Assert) CanEval(ds []Decl) bool {
	if a.E_I.CanEval(ds) {
		return true
	} else if !a.E_I.IsValue() {
		return false
	}
	return a.E_I.(StructLit).U_S.Impls(ds, a.U_cast)
}

func (a Assert) String() string {
	var b strings.Builder
	b.WriteString(a.E_I.String())
	b.WriteString(".(")
	b.WriteString(a.U_cast.String())
	b.WriteString(")")
	return b.String()
}

func (a Assert) ToGoString(ds []Decl) string {
	var b strings.Builder
	b.WriteString(a.E_I.ToGoString(ds))
	b.WriteString(".(")
	b.WriteString(a.U_cast.ToGoString(ds))
	b.WriteString(")")
	return b.String()
}

/* StringLit, fmt.Sprintf */

type StringLit struct {
	val string
}

var _ FGGExpr = StringLit{}

func (s StringLit) GetValue() string { return s.val }

func (s StringLit) Subs(subs map[Variable]FGGExpr) FGGExpr {
	return s
}

func (s StringLit) TSubs(subs map[TParam]Type) FGGExpr {
	return s
}

func (s StringLit) Eval(ds []Decl) (FGGExpr, string) {
	panic("Cannot reduce: " + s.String())
}

func (s StringLit) Typing(ds []Decl, delta Delta, gamma Gamma, allowStupid bool) Type {
	return STRING_TYPE
}

func (s StringLit) TypingWithEnv(env *Env, delta Delta, gamma Gamma, allowStupid bool) Type {
	return STRING_TYPE
}

// From base.Expr
func (s StringLit) IsValue() bool {
	return true
}

func (s StringLit) CanEval(ds []Decl) bool {
	return false
}

func (s StringLit) String() string {
	return "\"" + s.val + "\""
}

func (s StringLit) ToGoString(ds []Decl) string {
	return "\"" + s.val + "\""
}

type Sprintf struct {
	format string // Includes surrounding quotes
	Args   []FGGExpr
}

var _ FGGExpr = Sprintf{}

func (s Sprintf) GetFormat() string  { return s.format }
func (s Sprintf) GetArgs() []FGGExpr { return s.Args }

func (s Sprintf) Subs(subs map[Variable]FGGExpr) FGGExpr {
	args := make([]FGGExpr, len(s.Args))
	for i := 0; i < len(args); i++ {
		args[i] = s.Args[i].Subs(subs)
	}
	return Sprintf{s.format, args}
}

func (s Sprintf) TSubs(subs map[TParam]Type) FGGExpr {
	return s
}

func (s Sprintf) Eval(ds []Decl) (FGGExpr, string) {
	args := make([]FGGExpr, len(s.Args))
	done := false
	var rule string
	for i := 0; i < len(s.Args); i++ {
		v := s.Args[i]
		if !done && !v.IsValue() {
			v, rule = v.Eval(ds)
			done = true
		}
		args[i] = v
	}
	if done {
		return Sprintf{s.format, args}, rule
	} else {
		cast := make([]interface{}, len(args))
		for i := range args {
			cast[i] = args[i] // N.B. inside fgg this is, e.g., a StructLit (not the struct itself, as in native Go)
		}
		template := s.format[1 : len(s.format)-1] // Remove surrounding quote chars
		str := fmt.Sprintf(template, cast...)
		str = strings.ReplaceAll(str, "\"", "") // HACK because StringLit.String() includes quotes
		// FIXME: currently user remplates cannot include xplicit quote chars
		return StringLit{str}, "Sprintf"
	}
}

// TODO: [Warning] not "fully" type checked, cf. MISSING/EXTRA
func (s Sprintf) Typing(ds []Decl, delta Delta, gamma Gamma, allowStupid bool) Type {
	for i := 0; i < len(s.Args); i++ {
		s.Args[i].Typing(ds, delta, gamma, allowStupid)
	}
	return STRING_TYPE
}

func (s Sprintf) TypingWithEnv(env *Env, delta Delta, gamma Gamma, allowStupid bool) Type {
	for i := 0; i < len(s.Args); i++ {
		s.Args[i].TypingWithEnv(env, delta, gamma, allowStupid)
	}
	return STRING_TYPE
}

// From base.Expr
func (s Sprintf) IsValue() bool {
	return false
}

func (s Sprintf) CanEval(ds []Decl) bool {
	return true
}

func (s Sprintf) String() string {
	var b strings.Builder
	b.WriteString("fmt.Sprintf(")
	b.WriteString(s.format)
	if len(s.Args) > 0 {
		b.WriteString(", ")
		writeExprs(&b, s.Args)
	}
	b.WriteString(")")
	return b.String()
}

func (s Sprintf) ToGoString(ds []Decl) string {
	var b strings.Builder
	b.WriteString("fmt.Sprintf(")
	b.WriteString(s.format)
	if len(s.Args) > 0 {
		b.WriteString(", ")
		writeToGoExprs(ds, &b, s.Args)
	}
	b.WriteString(")")
	return b.String()
}

/* Aux, helpers */

func writeExprs(b *strings.Builder, es []FGGExpr) {
	if len(es) > 0 {
		b.WriteString(es[0].String())
		for _, v := range es[1:] {
			b.WriteString(", " + v.String())
		}
	}
}

func writeToGoExprs(ds []Decl, b *strings.Builder, es []FGGExpr) {
	if len(es) > 0 {
		b.WriteString(es[0].ToGoString(ds))
		for _, v := range es[1:] {
			b.WriteString(", " + v.ToGoString(ds))
		}
	}
}
