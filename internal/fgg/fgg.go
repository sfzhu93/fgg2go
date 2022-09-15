package fgg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/sfzhu93/fgg2go/internal/base"
)

var _ = fmt.Errorf
var _ = reflect.Append
var _ = strconv.AppendBool

/* Export */

func NewTName(t Name, us []Type) TNamed        { return TNamed{t, us} }
func IsStructType(ds []Decl, u Type) bool      { return isStructType(ds, u) }
func IsNamedIfaceType(ds []Decl, u Type) bool  { return isNamedIfaceType(ds, u) }
func NewTFormal(name TParam, u_I Type) TFormal { return TFormal{name, u_I} }
func NewBigPsi(tFormals []TFormal) BigPsi      { return BigPsi{tFormals} }

/* Constants */

// Hacks
var STRING_TYPE = TParam("string")
var PRIMITIVE_TYPES = make(map[TParam]TParam)
var PRIMITIVE_PSI BigPsi // Because prim types parsed as TParams, need to check OK

func init() {
	PRIMITIVE_TYPES[STRING_TYPE] = STRING_TYPE
	tfs := []TFormal{}
	for k, v := range PRIMITIVE_TYPES {
		tfs = append(tfs, TFormal{k, v})
	}
	PRIMITIVE_PSI = BigPsi{tfs}
}

/* Aliases from base */

type Name = base.Name
type FGGNode = base.AstNode
type Decl = base.Decl

/* Name, Type, Type param, Type name -- !!! submission version, "Type name" overloaded */

// Name: see Aliases (at top)

type Type interface {
	base.Type
	ImplsDelta(ds []Decl, delta Delta, u Type) bool
	ImplsDeltaWithEnv(env *Env, delta Delta, u Type) bool
	TSubs(subs map[TParam]Type) Type // N.B. map is Delta -- factor out a Subs type?
	SubsEta(eta Eta) Type
	SubsEtaOpen(eta EtaOpen) Type
	Ok(ds []Decl, delta Delta)
	ToGoString(ds []Decl) string
}

type TParam Name

var _ Type = TParam("")

func (a TParam) TSubs(subs map[TParam]Type) Type {
	res, ok := subs[a]
	if !ok {
		//panic("Unknown param: " + a.String())
		return a // CHECKME: ok? -- see TSubs in methods aux, w.r.t. meth-tparams that aren't in the subs map
		// Cf. Variable.Subs?
	}
	return res
}

func (a TParam) SubsEta(eta Eta) Type {
	if _, ok := PRIMITIVE_TYPES[a]; ok {
		return STRING_TYPE_MONOM // HACK TODO: refactor prims map as TParam->TNamed (map to monom rep)
	}
	res, ok := eta[a]
	if !ok {
		panic("Shouldn't get here: " + a)
	}
	return res
}

func (a TParam) SubsEtaOpen(eta EtaOpen) Type {
	res, ok := eta[a]
	if !ok {
		return a
	}
	return res
}

// u0 <: u
func (a TParam) ImplsDelta(ds []Decl, delta Delta, u Type) bool {
	if a1, ok := u.(TParam); ok {
		return a == a1
	} else {
		//return bounds(delta, a).ImplsDelta(ds, delta, u) // !!! more efficient?
		gs0, _ := MethodsDelta(ds, delta, a)
		gs, _ := MethodsDelta(ds, delta, u)
		for k, g := range gs {
			g0, ok := gs0[k]
			//if !ok || !sigAlphaEquals(g0, g) {
			if !ok || g0.String() != g.String() {
				return false
			}
		}
		return true
	}
}

func (a TParam) ImplsDeltaWithEnv(env *Env, delta Delta, u Type) bool {
	if a1, ok := u.(TParam); ok {
		return a == a1
	} else {
		//return bounds(delta, a).ImplsDelta(ds, delta, u) // !!! more efficient?
		gs0, _ := env.MethodsDelta(delta, a)
		gs, _ := env.MethodsDelta(delta, u)
		for k, g := range gs {
			g0, ok := gs0[k]
			//if !ok || !sigAlphaEquals(g0, g) {
			if !ok || g0.String() != g.String() {
				return false
			}
		}
		return true
	}
}

// Cf. base.Type
func (a TParam) Impls(ds []Decl, u base.Type) bool {
	if _, ok := u.(Type); !ok {
		panic("Expected FGG type, not " + reflect.TypeOf(u).String() +
			":\n\t" + u.String())
	}
	return a.ImplsDelta(ds, make(Delta), u.(Type))
}

func (a TParam) Ok(ds []Decl, delta Delta) {
	if _, ok := PRIMITIVE_TYPES[a]; ok {
		return
	}
	if _, ok := delta[a]; !ok {
		panic("Type param " + a.String() + " unknown in context: " + delta.String())
	}
}

func (a TParam) Equals(u base.Type) bool {
	if _, ok := u.(Type); !ok {
		panic("Expected FGG type, not " + reflect.TypeOf(u).String() +
			":\n\t" + u.String())
	}
	if b, ok := u.(TParam); ok {
		return a == b // Handles primitives
	}
	return false
}

func (a TParam) String() string {
	return string(a)
}

func (a TParam) ToGoString(ds []Decl) string {
	return string(a)
}

// Convention: t=type name (t), u=FGG type (tau)
type TNamed struct {
	TName Name
	UArgs []Type // SmallPsi
}

var _ Type = TNamed{}
var _ Spec = TNamed{}

func (u0 TNamed) GetName() Name    { return u0.TName }
func (u0 TNamed) GetTArgs() []Type { return u0.UArgs } // SmallPsi

func (u0 TNamed) TSubs(subs map[TParam]Type) Type {
	us := make([]Type, len(u0.UArgs))
	for i := 0; i < len(us); i++ {
		us[i] = u0.UArgs[i].TSubs(subs)
	}
	return TNamed{u0.TName, us}
}

func (u0 TNamed) SubsEta(eta Eta) Type {
	//fmt.Println("555:", u0, eta)
	us := make([]Type, len(u0.UArgs))
	for i := 0; i < len(us); i++ {
		us[i] = u0.UArgs[i].SubsEta(eta)
	}
	return TNamed{u0.TName, us}
}

func (u0 TNamed) SubsEtaOpen(eta EtaOpen) Type {
	//fmt.Println("555:", u0, eta)
	us := make([]Type, len(u0.UArgs))
	for i := 0; i < len(us); i++ {
		us[i] = u0.UArgs[i].SubsEtaOpen(eta)
	}
	return TNamed{u0.TName, us}
}

// u0 <: u
// delta unused here (cf. TParam.ImplsDelta)
func (u0 TNamed) ImplsDelta(ds []Decl, delta Delta, u Type) bool {
	if isStructType(ds, u) {
		return isStructType(ds, u0) && u0.Equals(u) // Asks equality of nested TParam
	}
	if _, ok := u.(TParam); ok { // e.g., fgg_test.go, Test014
		panic("Type name does not implement open type param: found=" +
			u0.String() + ", expected=" + u.String())
	}

	gs, _ := MethodsDelta(ds, delta, u)   // u is a t_I
	gs0, _ := MethodsDelta(ds, delta, u0) // t0 may be any
	for k, g := range gs {
		g0, ok := gs0[k]
		//fmt.Printf("TNamed.ImplsDelta: \nu.method=%#v\nu0.method=%#v\n", g, g0)
		if !ok || !sigAlphaEquals(&g0, &g) {
			return false
		}
	}
	return true
}

func (u0 TNamed) ImplsDeltaWithEnv(env *Env, delta Delta, u Type) bool {
	if env.IsStructType(u) {
		return env.IsStructType(u0) && u0.Equals(u)
	}
	if _, ok := u.(TParam); ok { // e.g., fgg_test.go, Test014
		panic("Type name does not implement open type param: found=" +
			u0.String() + ", expected=" + u.String())
	}

	gs, _ := env.MethodsDelta(delta, u)   // u is a t_I
	gs0, _ := env.MethodsDelta(delta, u0) // t0 may be any
	for k, g := range gs {
		g0, ok := gs0[k]
		//fmt.Printf("TNamed.ImplsDeltaWithEnv: \nu.method=%#v\nu0.method=%#v\n", g, g0)
		if !ok || !sigAlphaEquals(g0, g) {
			return false
		}
	}
	return true
}

// !!! Sig in FGG includes ~a and ~x, which naively breaks "impls"
func sigAlphaEquals(g0 *Sig, g *Sig) bool {
	//fmt.Printf("sigAlphaEquals: \nu.method=%#v\nu0.method=%#v\n", g, g0)

	if len(g0.Psi.TFormals) != len(g.Psi.TFormals) || len(g0.pDecls) != len(g.pDecls) {
		return false
	}
	subs0 := MakeParamIndexSubs(g0.Psi)
	subs := MakeParamIndexSubs(g.Psi)
	for i := 0; i < len(g0.Psi.TFormals); i++ {
		if !g0.Psi.TFormals[i].U_I.TSubs(subs0).
			Equals(g.Psi.TFormals[i].U_I.TSubs(subs)) {
			//fmt.Println("z:")
			return false
		}
	}
	for i := 0; i < len(g0.pDecls); i++ {
		if !g0.pDecls[i].U.TSubs(subs0).Equals(g.pDecls[i].U.TSubs(subs)) {
			/*fmt.Println("w1: ", g0.pDecls[i].u, g0.pDecls[i].u.TSubs(subs0))
			fmt.Println("w2: ", g.pDecls[i].u, g.pDecls[i].u.TSubs(subs))
			fmt.Println("y:")*/
			return false
		}
	}
	/*fmt.Println("1:", g0)
	fmt.Println("2:", g)
	fmt.Println("3:", g0.meth == g.meth, g0.u_ret.Equals(g.u_ret))
	fmt.Println("4:", g0.u_ret.TSubs(subs0).Equals(g.u_ret.TSubs(subs)))*/
	return g0.meth == g.meth && g0.u_ret.TSubs(subs0).Equals(g.u_ret.TSubs(subs))
}

// CHECKME: Used by sigAlphaEquals, and MDecl.OK (for covariant receiver bounds)
func MakeParamIndexSubs(Psi BigPsi) Delta {
	subs := make(Delta)
	for j := 0; j < len(Psi.TFormals); j++ {
		//subs[Psi.tFormals[j].name] = Psi.tFormals[j].name
		subs[Psi.TFormals[j].Name] = TParam("α" + strconv.Itoa(j+1))
	}
	return subs
}

// Cf. base.Type
func (u0 TNamed) Impls(ds []Decl, u base.Type) bool {
	if _, ok := u.(Type); !ok {
		panic("Expected FGG type, not " + reflect.TypeOf(u).String() +
			":\n\t" + u.String())
	}
	return u0.ImplsDelta(ds, make(Delta), u.(Type))
}

func (u0 TNamed) Ok(ds []Decl, delta Delta) {
	//if _, ok
	td := GetTDecl(ds, u0.TName) // Panics if type not found
	psi := td.GetBigPsi()
	if len(psi.TFormals) != len(u0.UArgs) {
		var b strings.Builder
		b.WriteString("Arity mismatch between type formals and actuals: formals=")
		b.WriteString(psi.String())
		b.WriteString(" actuals=")
		writeTypes(&b, u0.UArgs)
		b.WriteString("\n\t")
		b.WriteString(u0.String())
		panic(b.String())
	}
	subs := make(map[TParam]Type)
	for i := 0; i < len(psi.TFormals); i++ {
		subs[psi.TFormals[i].Name] = u0.UArgs[i]
	}
	for i := 0; i < len(psi.TFormals); i++ {
		actual := psi.TFormals[i].Name.TSubs(subs)
		// CHECKME: submission T-Named, subs applied to Delta? -- already applied, Delta is coming from the subs context
		formal := psi.TFormals[i].U_I.TSubs(subs)
		if !actual.ImplsDelta(ds, delta, formal) { // tfs[i].u is a \tau_I, checked by TDecl.Ok
			panic("Type actual must implement type formal: actual=" +
				actual.String() + " formal=" + formal.String())
		}
	}
	for _, v := range u0.UArgs {
		v.Ok(ds, delta)
	}
}

// \tau_I is a Spec, but not \tau_S -- this aspect is currently "dynamically typed"
// From Spec
func (u TNamed) GetSigs(ds []Decl) []Sig {
	if !isNamedIfaceType(ds, u) { // isStructType would be more efficient
		panic("Cannot use non-interface type as a Spec: " + u.String() +
			" is a " + reflect.TypeOf(u).String())
	}
	td := GetTDecl(ds, u.TName).(ITypeLit)
	var res []Sig
	for _, s := range td.specs {
		res = append(res, s.GetSigs(ds)...)
	}
	return res
}

func (u0 TNamed) Equals(u base.Type) bool {
	if _, ok := u.(Type); !ok {
		panic("Expected FGG type, not " + reflect.TypeOf(u).String() +
			":\n\t" + u.String())
	}
	if _, ok := u.(TNamed); !ok {
		return false
	}
	u1 := u.(TNamed)
	if u0.TName != u1.TName || len(u0.UArgs) != len(u1.UArgs) {
		return false
	}
	for i := 0; i < len(u0.UArgs); i++ {
		if !u0.UArgs[i].Equals(u1.UArgs[i]) { // Asks equality of nested TParam
			return false
		}
	}
	return true
}

func (u TNamed) String() string {
	var b strings.Builder
	b.WriteString(string(u.TName))
	b.WriteString("(")
	writeTypes(&b, u.UArgs)
	b.WriteString(")")
	return b.String()
}

func (u TNamed) ToGoString(ds []Decl) string {
	var b strings.Builder
	b.WriteString("main.")
	b.WriteString(string(u.TName))
	b.WriteString("(")
	writeToGoTypes(ds, &b, u.UArgs)
	b.WriteString(")")
	return b.String()
}

/* Type formals and actuals */

// Pre: len(as) == len(us)
// Wrapper for []TFormal (cf. e.g., FieldDecl), only because of "(type ...)" syntax
// Also ranged over by big phi
type BigPsi struct {
	TFormals []TFormal
}

func (Psi BigPsi) GetTFormals() []TFormal { return Psi.TFormals }

func (Psi BigPsi) Ok(ds []Decl, env BigPsi) {
	delta := env.ToDelta()
	for _, v := range Psi.TFormals {
		if _, ok := delta[v.Name]; ok {
			panic("Duplicate param name " + string(v.Name) + " under context: " +
				env.String() + "\n\t" + Psi.String())
		}
		delta[v.Name] = v.U_I
	} // Delta built
	for _, v := range Psi.TFormals {
		u_I, ok := v.U_I.(TNamed)
		if !ok {
			if _, foo := PRIMITIVE_TYPES[v.U_I.(TParam)]; !foo { // Only because PRIMITIVE_PSI hacks the upperbound like this
				panic("Upper bound must be a named interface type: not " + v.U_I.String() +
					"\n\t" + Psi.String())
			}
		} else {
			if !isNamedIfaceType(ds, u_I) {
				panic("Upper bound must be a named interface type: not " + v.U_I.String() +
					"\n\t" + Psi.String())
			}
			u_I.Ok(ds, delta) // Checks params bound under delta -- N.B. can forward ref (not restricted left-to-right)
		}
	}
}

func (Psi BigPsi) ToDelta() Delta {
	delta := make(map[TParam]Type)
	for _, v := range Psi.TFormals {
		delta[v.Name] = v.U_I
	}
	return delta
}

// The ordered value set of ToDelta
func (Psi BigPsi) Hat() SmallPsi {
	res := make(SmallPsi, len(Psi.TFormals))
	for i, v := range Psi.TFormals {
		res[i] = v.Name
	}
	return res
}

func (Psi BigPsi) String() string {
	var b strings.Builder
	b.WriteString("(type ") // Includes "(...)" -- cf. e.g., writeFieldDecls
	if len(Psi.TFormals) > 0 {
		b.WriteString(Psi.TFormals[0].String())
		for _, v := range Psi.TFormals[1:] {
			b.WriteString(", ")
			b.WriteString(v.String())
		}
	}
	b.WriteString(")")
	return b.String()
}

type TFormal struct {
	Name TParam
	U_I  Type
	// CHECKME: submission version, upper bound \tau_I is only "of the form t_I(~\tau)"? -- i.e., not \alpha?
	// ^If so, then can refine to TNamed
}

func (tf TFormal) GetTParam() TParam   { return tf.Name }
func (tf TFormal) GetUpperBound() Type { return tf.U_I }

func (tf TFormal) String() string {
	return string(tf.Name) + " " + tf.U_I.String()
}

// Type actuals
// Also ranged over by small phi
type SmallPsi []Type // CHECKME: Currently only used in omega/monom, maybe deprecate?

func (x0 SmallPsi) TSubs(subs map[TParam]Type) SmallPsi {
	res := make(SmallPsi, len(x0))
	for i, v := range x0 {
		res[i] = v.TSubs(subs)
	}
	return res
}

func (x0 SmallPsi) String() string {
	var b strings.Builder
	for _, v := range x0 {
		b.WriteString(v.String())
	}
	return b.String()
}

func (x0 SmallPsi) Equals(x SmallPsi) bool {
	if len(x0) != len(x) {
		return false
	}
	for i := 0; i < len(x0); i++ {
		if !x0[i].Equals(x[i]) {
			return false
		}
	}
	return true
}

/* Context, Type context, Substitutions */

//type Gamma map[Variable]Type
type Gamma map[Name]Type
type Delta map[TParam]Type // Type intended to be an upper bound
type Eta map[TParam]Type   // TNamed intended to be a ground

type EtaOpen map[TParam]Type // cf. Delta

func (delta Delta) String() string {
	res := "["
	first := true
	for k, v := range delta {
		if first {
			first = false
		} else {
			res = res + ", "
		}
		res = k.String() + ":" + v.String()
	}
	return res + "]"
}

// Pre: len(psi) == len(Psi.GetTFormals()); psi all ground
func MakeEta(Psi BigPsi, psi SmallPsi) Eta {
	eta := make(Eta)
	tfs := Psi.TFormals
	for i := 0; i < len(tfs); i++ {
		eta[tfs[i].Name] = psi[i].(TNamed)
	}
	return eta
}

func MakeEtaDelta(ds []Decl, delta Delta, Psi BigPsi, psi SmallPsi) (bool, EtaOpen) {
	eta := MakeEtaOpen(Psi, psi)
	for _, v := range Psi.TFormals {
		a := v.Name.SubsEtaOpen(eta)
		u_I := v.U_I.SubsEtaOpen(eta)
		//fmt.Printf("a=%#v\nu_I=%#v\n", a, u_I)
		if !a.ImplsDelta(ds, delta, u_I) {
			return false, eta
		}
	}
	return true, eta
}

func MakeEtaOpen(Psi BigPsi, psi SmallPsi) EtaOpen {
	eta := make(EtaOpen)
	tfs := Psi.TFormals
	for i := 0; i < len(tfs); i++ {
		eta[tfs[i].Name] = psi[i]
	}
	return eta
}

/* AST base intefaces: FGGNode, Decl, TypeDecl, Spec, Expr */

// FGGNode, Decl: see Aliases (at top)

type TypeDecl interface {
	Decl
	GetBigPsi() BigPsi // TODO: rename? potential clash with, e.g., MDecl, can cause "false" interface satisfaction
}

type Spec interface {
	FGGNode
	GetSigs(ds []Decl) []Sig
}

type FGGExpr interface {
	base.Expr
	Subs(subs map[Variable]FGGExpr) FGGExpr
	TSubs(subs map[TParam]Type) FGGExpr
	// gamma and delta should be treated immutably
	Typing(ds []Decl, delta Delta, gamma Gamma, allowStupid bool) Type
	TypingWithEnv(env *Env, delta Delta, gamma Gamma, allowStupid bool) Type
	Eval(ds []Decl) (FGGExpr, string)
}

/* Helpers */

// Based on FG version -- but currently no FGG equiv of isInterfaceType
// Helpful for MDecl.TRecv
func isStructName(ds []Decl, t Name) bool {
	for _, v := range ds {
		d, ok := v.(STypeLit)
		if ok && d.T_name == t {
			return true
		}
	}
	return false
}

// Check if u is a \tau_S -- implicitly must be a TNamed
func isStructType(ds []Decl, u Type) bool {
	if u1, ok := u.(TNamed); ok {
		for _, v := range ds {
			d, ok := v.(STypeLit)
			if ok && d.T_name == u1.TName {
				return true
			}
		}
	}
	return false
}

// Check if u is a \tau_I -- N.B. looks for a *TNamed*, i.e., not a TParam
func isNamedIfaceType(ds []Decl, u Type) bool {
	if u1, ok := u.(TNamed); ok {
		for _, v := range ds {
			d, ok := v.(ITypeLit)
			if ok && d.t_I == u1.TName {
				return true
			}
		}
	}
	return false
}

func writeTypes(b *strings.Builder, us []Type) {
	if len(us) > 0 {
		b.WriteString(us[0].String())
		for _, v := range us[1:] {
			b.WriteString(", " + v.String())
		}
	}
}

func writeToGoTypes(ds []Decl, b *strings.Builder, us []Type) {
	if len(us) > 0 {
		b.WriteString(us[0].ToGoString(ds))
		for _, v := range us[1:] {
			b.WriteString(", " + v.ToGoString(ds))
		}
	}
}
