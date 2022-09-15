package fg

import "reflect"

import "github.com/sfzhu93/fgg2go/internal/base"

/* Aliases from base */

type Name = base.Name
type FGNode = base.AstNode
type Decl = base.Decl
type FieldOrParamDecl interface {
	base.AstNode
	GetName() Name
	GetType() TypeBase
}

/* Constants */

var STRING_TYPE = Type("string")
var PRIMITIVE_TYPES = make(map[Type]Type)

func init() {
	PRIMITIVE_TYPES[STRING_TYPE] = STRING_TYPE
}

/* Name, Context, Type */

// Name: see Aliases (at top)

type Gamma map[Name]TypeBase // Variable? though is an Expr

type Type Name // Type definition (cf. alias)

var _ base.Type = Type("")
var _ Spec = Type("")

// Pre: t0, t are known types
// t0 <: t
func (t0 Type) Impls(ds []Decl, t base.Type) bool {
	if _, ok := PRIMITIVE_TYPES[t0]; ok {
		if _, ok := t.(Type); ok {
			return t0.Equals(t)
		}
	}

	if _, ok := t.(Type); !ok {
		panic("Expected FGR type, not " + reflect.TypeOf(t).String() +
			":\n\t" + t.String())
	}
	t_fg := t.(Type)
	if IsStructType(ds, t_fg) {
		return IsStructType(ds, t0) && t0 == t_fg
	}

	gs := methods(ds, t_fg) // t is a t_I
	gs0 := methods(ds, t0)  // t0 may be any
	for k, g := range gs {
		g0, ok := gs0[k]
		if !ok || !g.EqExceptVars(g0) {
			return false
		}
	}
	return true
}

// t_I is a Spec, but not t_S -- this aspect is currently "dynamically typed"
// From Spec
func (t Type) GetSigs(ds []Decl) []Sig {
	if !isInterfaceType(ds, t) { // IsStructType would be more efficient
		panic("Cannot use non-interface type as a Spec: " + t.String())
	}
	td := getTDecl(ds, t).(ITypeLit)
	var res []Sig
	for _, s := range td.specs {
		res = append(res, s.GetSigs(ds)...)
	}
	return res
}

func (t0 Type) Equals(t base.Type) bool {
	if _, ok := t.(Type); !ok {
		panic("Expected FGR type, not " + reflect.TypeOf(t).String() +
			":\n\t" + t.String())
	}
	return t0 == t.(Type)
}

func (t Type) String() string {
	return string(t)
}

/* AST base intefaces: FGNode, Decl, TDecl, Spec, Expr */

// FGNode, Decl: see Aliases (at top)

type TDecl interface { // Rename TypeDecl
	Decl
	GetType() Type // In FG, GetType() == Type(GetName())
}

// A Sig or a Type (specifically a t_I -- bad t_S usage raises a run-time error, cf. Type.GetSigs)
type Spec interface {
	FGNode
	GetSigs(ds []Decl) []Sig
}

type FGExpr interface {
	base.Expr
	Subs(subs map[Variable]FGExpr) FGExpr

	// N.B. gamma should be treated immutably (and ds, of course)
	// (No typing rule modifies gamma, except the T-Func bootstrap)
	Typing(ds []Decl, gamma Gamma, allowStupid bool) TypeBase

	// string is the type name of the "actually evaluated" expr (within the eval context)
	// CHECKME: resulting Exprs are not "parsed" from source, OK?
	Eval(ds []Decl) (FGExpr, string)

	//IsPanic() bool  // TODO "explicit" FG panic -- cf. underlying runtime panic
}

/* Helpers */

func IsStructType(ds []Decl, t TypeBase) bool {
	for _, v := range ds {
		if d, ok := v.(STypeLit); ok && d.t_S == t {
			return true
		}
	}
	return false
}

func isInterfaceType(ds []Decl, t TypeBase) bool {
	for _, v := range ds {
		d, ok := v.(ITypeLit)
		if ok && d.t_I == t {
			return true
		}
	}
	return false
}
