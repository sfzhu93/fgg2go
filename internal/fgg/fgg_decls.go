/*
 * This file contains defs for "concrete" syntax w.r.t. programs and decls.
 * Base ("abstract") types, interfaces, etc. are in fgg.go.
 */

package fgg

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/sfzhu93/fgg2go/internal/base"
)

var _ = fmt.Errorf
var _ = reflect.Append

/* Public constructors */

// TODO: rename NewFGGProgram
func NewProgram(ds []Decl, e FGGExpr, printf bool) FGGProgram {
	return FGGProgram{ds, e, printf}
}

func NewSTypeLit(t Name, Psi BigPsi, fds []FieldDecl) STypeLit { return STypeLit{t, Psi, fds} }

func NewITypeLit(t_I Name, Psi BigPsi, specs []Spec) ITypeLit {
	return ITypeLit{t_I, Psi, specs}
}

func NewMethDecl(
	x_recv Name,
	t_recv Name,
	Psi_recv BigPsi,
	name Name,
	Psi_meth BigPsi,
	paramDecls []ParamDecl,
	u_ret Type,
	bodyStmt MethBody) MethDecl {
	return MethDecl{x_recv, t_recv, Psi_recv, name, Psi_meth, paramDecls, u_ret, nil, bodyStmt}
}

// TODO: rename NewMethDecl
func NewMDecl(x_recv Name,
	t_recv Name,
	Psi_recv BigPsi,
	name Name,
	Psi_meth BigPsi,
	pDecls []ParamDecl,
	u_ret Type,
	stmt MethBody) MethDecl {
	return MethDecl{x_recv, t_recv, Psi_recv, name, Psi_meth, pDecls, u_ret, nil, stmt}
}
func NewFieldDecl(f Name, t Type) FieldDecl                  { return FieldDecl{f, t} }
func NewParamDecl(x Name, t Type) ParamDecl                  { return ParamDecl{x, t} }     // For fgg_monom.MakeWMap
func NewSig(m Name, Psi BigPsi, pds []ParamDecl, t Type) Sig { return Sig{m, Psi, pds, t} } // For fgg_monom.MakeWMap

/* Program */

type FGGProgram struct {
	decls  []Decl
	e_main FGGExpr
	printf bool // false = "original" `_ = e_main` syntax; true = import-fmt/printf syntax
	// N.B. coincidentally "behaves" like an actual printf because interpreter prints out final eval result
}

var _ base.Program = FGGProgram{}
var _ FGGNode = FGGProgram{}

func (p FGGProgram) GetDecls() []Decl   { return p.decls } // Return a copy?
func (p FGGProgram) GetMain() base.Expr { return p.e_main }
func (p FGGProgram) IsPrintf() bool     { return p.printf } // HACK

func (p FGGProgram) Ok(allowStupid bool) base.Type {
	tds := make(map[string]TypeDecl) // Type name
	mds := make(map[string]MethDecl) // Hack, string = md.recv.t + "." + md.name
	for _, v := range p.decls {
		switch d := v.(type) {
		case TypeDecl:
			d.Ok(p.decls)
			t := d.GetName()
			if _, ok := tds[t]; ok {
				panic("Multiple declarations of type name: " + t + "\n\t" +
					d.String())
			}
			tds[t] = d
		case MethDecl:
			d.Ok(p.decls)
			hash := string(d.TRecv) + "." + d.name
			if _, ok := mds[hash]; ok {
				panic("Multiple declarations for receiver " + string(d.TRecv) +
					" of the method name: " + d.name + "\n\t" + d.String())
			}
			mds[hash] = d
		default:
			panic("Unknown decl: " + reflect.TypeOf(v).String() + "\n\t" +
				v.String())
		}
	}
	// Empty envs for main
	var delta Delta
	var gamma Gamma
	return p.e_main.Typing(p.decls, delta, gamma, allowStupid)
}

func (p FGGProgram) Eval() (base.Program, string) {
	e, rule := p.e_main.Eval(p.decls)
	return FGGProgram{p.decls, e.(FGGExpr), p.printf}, rule
}

func (p FGGProgram) String() string {
	var b strings.Builder
	b.WriteString("package main;\n")
	if p.printf {
		b.WriteString("import \"fmt\";\n")
	}
	for _, v := range p.decls {
		b.WriteString(v.String())
		b.WriteString(";\n")
	}
	b.WriteString("func main() { ")
	if p.printf {
		b.WriteString("fmt.Printf(\"%#v\", ")
		b.WriteString(p.e_main.String())
		b.WriteString(")")
	} else {
		b.WriteString("_ = ")
		b.WriteString(p.e_main.String())
	}
	b.WriteString(" }")
	return b.String()
}

/* STypeLit, FieldDecl */

type STypeLit struct {
	T_name Name
	Psi    BigPsi
	fDecls []FieldDecl
}

var _ TypeDecl = STypeLit{}

func (s STypeLit) GetName() Name              { return s.T_name }
func (s STypeLit) GetBigPsi() BigPsi          { return s.Psi }
func (s STypeLit) GetFieldDecls() []FieldDecl { return s.fDecls }

func (s STypeLit) Ok(ds []Decl) {
	s.Psi.Ok(ds, PRIMITIVE_PSI)
	seen := make(map[Name]FieldDecl)
	//delta := s.Psi.ToDelta()
	root := makeRootPsi(s.Psi)
	delta := root.ToDelta()
	for _, v := range s.fDecls {
		if _, ok := seen[v.field]; ok {
			panic("Duplicate field name: " + v.field + "\n\t" + s.String())
		}
		seen[v.field] = v
		v.u.Ok(ds, delta)
	}
	u_S := TNamed{s.T_name, s.Psi.Hat()}
	if isRecursiveFieldType(ds, make(map[string]TNamed), u_S) {
		panic("Invalid recursive struct type:\n\t" + s.String())
	}
}

func (s STypeLit) String() string {
	var b strings.Builder
	b.WriteString("type ")
	b.WriteString(string(s.T_name))
	b.WriteString(s.Psi.String())
	b.WriteString(" struct {")
	if len(s.fDecls) > 0 {
		b.WriteString(" ")
		writeFieldDecls(&b, s.fDecls)
		b.WriteString(" ")
	}
	b.WriteString("}")
	return b.String()
}

type FieldDecl struct {
	field Name
	u     Type // u=tau
}

var _ FGGNode = FieldDecl{}

func (fd FieldDecl) GetName() Name { return fd.field }
func (fd FieldDecl) GetType() Type { return fd.u }

func (fd FieldDecl) Subs(subs map[TParam]Type) FieldDecl {
	return FieldDecl{fd.field, fd.u.TSubs(subs)}
}

func (fd FieldDecl) String() string {
	return fd.field + " " + fd.u.String()
}

/* MethDecl, ParamDecl */

type MethDecl struct {
	XRecv    Name // CHECKME: better to be Variable?  (etc. for other such Names)
	TRecv    Name // N.B. t_S
	Psi_recv BigPsi
	// N.B. receiver elements "decomposed" because Psi (not TNamed, cf. fg.MDecl uses ParamDecl)
	name     Name // Refactor to embed Sig?
	Psi_meth BigPsi
	pDecls   []ParamDecl
	U_ret    Type // Return
	E_body   FGGExpr
	BodyStmt MethBody //[]FGGStmt
}

var _ Decl = MethDecl{}

func (md MethDecl) GetRecvName() Name          { return md.XRecv }
func (md MethDecl) GetRecvTypeName() Name      { return md.TRecv }
func (md MethDecl) GetRecvPsi() BigPsi         { return md.Psi_recv }
func (md MethDecl) GetName() Name              { return md.name }
func (md MethDecl) GetMDeclPsi() BigPsi        { return md.Psi_meth } // MDecl in name to prevent false capture by TDecl interface
func (md MethDecl) GetParamDecls() []ParamDecl { return md.pDecls }
func (md MethDecl) GetReturn() Type            { return md.U_ret }
func (md MethDecl) GetBody() FGGExpr           { return md.E_body }
func (md MethDecl) GetMethBody() MethBody      { return md.BodyStmt }

//T-func?
func (md MethDecl) Ok(ds []Decl) {
	if md.GetName() == "Next" {
		println("11")
	}
	if !isStructName(ds, md.TRecv) {
		panic("Receiver must be a struct type: not " + md.TRecv +
			"\n\t" + md.String())
	}

	//Φ; Ψ ok Δ
	md.Psi_recv.Ok(ds, BigPsi{}) // !!! premise ok missing
	md.Psi_meth.Ok(ds, md.Psi_recv)
	delta := md.Psi_recv.ToDelta()
	for _, v := range md.Psi_meth.TFormals {
		delta[v.Name] = v.U_I
	}

	td := GetTDecl(ds, md.TRecv)
	tfs_td := td.GetBigPsi().TFormals
	if len(tfs_td) != len(md.Psi_recv.TFormals) {
		panic("Receiver type parameter arity mismatch:\n\tmdecl=" + md.TRecv +
			md.Psi_recv.String() + ", tdecl=" + td.GetName() +
			"\n\t" + td.GetBigPsi().String())
	}
	for i := 0; i < len(tfs_td); i++ {
		subs_md := MakeParamIndexSubs(md.Psi_recv)
		subs_td := MakeParamIndexSubs(td.GetBigPsi())
		if !md.Psi_recv.TFormals[i].U_I.TSubs(subs_md). // Canonicalised
								Impls(ds, tfs_td[i].U_I.TSubs(subs_td)) {
			panic("Receiver parameter upperbound not a subtype of type decl upperbound:" +
				"\n\tmdecl=" + md.Psi_recv.TFormals[i].String() + ", tdecl=" +
				tfs_td[i].String())
		}
	}

	as := md.Psi_recv.Hat()                        // !!! submission version, x:t_S(a) => x:t_S(~a)
	gamma := Gamma{md.XRecv: TNamed{md.TRecv, as}} // CHECKME: can we give the bounds directly here instead of 'as'?
	seen := make(map[Name]Name)
	seen[md.XRecv] = md.XRecv
	for _, v := range md.pDecls {
		if _, ok := seen[v.Name]; ok {
			panic("Duplicate receiver/param name: " + v.Name + "\n\t" + md.String())
		}
		seen[v.Name] = v.Name
		v.U.Ok(ds, delta)
		gamma[v.Name] = v.U
	}
	md.U_ret.Ok(ds, delta)
	allowStupid := false

	u := md.BodyStmt.Typing(ds, delta, gamma, md.U_ret, allowStupid)
	//u := md.E_body.Typing(ds, delta, gamma, allowStupid)

	/*fmt.Println("a:", u)
	fmt.Println("b:", md.u_ret)
	fmt.Println("c:", u.ImplsDelta(ds, delta, md.u_ret))*/

	if !u.ImplsDelta(ds, delta, md.U_ret) {
		panic("Method body type must implement declared return type: found=" +
			u.String() + ", expected=" + md.U_ret.String() + "\n\t" + md.String())
	}
}

func (md MethDecl) ToSig() Sig {
	return Sig{md.name, md.Psi_meth, md.pDecls, md.U_ret}
}

func (md MethDecl) String() string {
	var b strings.Builder
	b.WriteString("func (")
	//b.WriteString(md.recv.String())
	b.WriteString(md.XRecv)
	b.WriteString(" ")
	b.WriteString(md.TRecv)
	b.WriteString(md.Psi_recv.String())
	b.WriteString(") ")
	b.WriteString(md.name)
	b.WriteString(md.Psi_meth.String())
	b.WriteString("(")
	writeParamDecls(&b, md.pDecls)
	b.WriteString(") ")
	b.WriteString(md.U_ret.String())
	b.WriteString(" { return ")
	b.WriteString(md.BodyStmt.String())
	b.WriteString(" }")
	return b.String()
}

// Cf. FieldDecl
type ParamDecl struct {
	Name Name // CHECKME: Variable?
	U    Type
}

var _ FGGNode = ParamDecl{}

func (pd ParamDecl) GetName() Name { return pd.Name }
func (pd ParamDecl) GetType() Type { return pd.U }

func (pd ParamDecl) String() string {
	return pd.Name + " " + pd.U.String()
}

/* ITypeLit, Sig */

type ITypeLit struct {
	t_I   Name
	Psi   BigPsi
	specs []Spec
}

var _ TypeDecl = ITypeLit{}

func (c ITypeLit) GetName() Name     { return c.t_I }
func (c ITypeLit) GetBigPsi() BigPsi { return c.Psi }
func (c ITypeLit) GetSpecs() []Spec  { return c.specs }

func (c ITypeLit) Ok(ds []Decl) {
	c.Psi.Ok(ds, PRIMITIVE_PSI)
	root := makeRootPsi(c.Psi)
	delta := root.ToDelta()
	seen_g := make(map[Name]Sig)    // !!! unique(~S) more flexible
	seen_u := make(map[string]Type) // key is u.String()
	for _, v := range c.specs {
		switch s := v.(type) {
		case Sig:
			if _, ok := seen_g[s.meth]; ok {
				panic("Multiple sigs with name: " + s.meth + "\n\t" + c.String())
			}
			seen_g[s.meth] = s
			s.Ok(ds, root)
		case TNamed:
			k := s.String()
			if _, ok := seen_u[k]; ok {
				panic("Repeat embedding of type: " + k + "\n\t" + c.String())
			}
			seen_u[k] = s
			if !IsNamedIfaceType(ds, s) { // CHECKME: allow embed type param?
				panic("Embedded type must be a named interface, not: " + k + "\n\t" + c.String())
			}
			s.Ok(ds, delta)
			if isRecursiveInterfaceEmbedding(ds, make(map[string]TNamed), s) {
				panic("Invalid recursive interface embedding type:\n\t" + c.String())
			}
		default:
			panic("Unknown Spec kind: " + reflect.TypeOf(v).String() + "\n\t" +
				c.String())
		}
	}
}

func (c ITypeLit) String() string {
	var b strings.Builder
	b.WriteString("type ")
	b.WriteString(string(c.t_I))
	b.WriteString(c.Psi.String())
	b.WriteString(" interface {")
	if len(c.specs) > 0 {
		b.WriteString(" ")
		b.WriteString(c.specs[0].String())
		for _, v := range c.specs[1:] {
			b.WriteString("; ")
			b.WriteString(v.String())
		}
		b.WriteString(" ")
	}
	b.WriteString("}")
	return b.String()
}

type Sig struct {
	meth   Name
	Psi    BigPsi // Add-meth-tparams
	pDecls []ParamDecl
	u_ret  Type
}

var _ Spec = Sig{}

func (g Sig) GetMethod() Name            { return g.meth }
func (g Sig) GetPsi() BigPsi             { return g.Psi }
func (g Sig) GetParamDecls() []ParamDecl { return g.pDecls }
func (g Sig) GetReturn() Type            { return g.u_ret }

func (g Sig) TSubs(subs map[TParam]Type) Sig {
	tfs := make([]TFormal, len(g.Psi.TFormals))
	for i := 0; i < len(g.Psi.TFormals); i++ {
		tf := g.Psi.TFormals[i]
		tfs[i] = TFormal{tf.Name.TSubs(subs).(TParam), tf.U_I.TSubs(subs)}
	}
	ps := make([]ParamDecl, len(g.pDecls))
	for i := 0; i < len(ps); i++ {
		pd := g.pDecls[i]
		ps[i] = ParamDecl{pd.Name, pd.U.TSubs(subs)}
	}
	u := g.u_ret.TSubs(subs)
	return Sig{g.meth, BigPsi{tfs}, ps, u}
}

func (g Sig) Ok(ds []Decl, env BigPsi) {
	env.Ok(ds, BigPsi{})
	g.Psi.Ok(ds, env)
	delta := env.ToDelta()
	for _, v := range g.Psi.TFormals {
		delta[v.Name] = v.U_I
	}
	seen := make(map[Name]ParamDecl)
	for _, v := range g.pDecls {
		if _, ok := seen[v.Name]; ok {
			panic("Duplicate variable name " + v.Name + ":\n\t" + g.String())
		}
		seen[v.Name] = v
		v.U.Ok(ds, delta)
	}
	g.u_ret.Ok(ds, delta)
}

func (g Sig) GetSigs(_ []Decl) []Sig {
	return []Sig{g}
}

func (g Sig) String() string {
	var b strings.Builder
	b.WriteString(g.meth)
	b.WriteString(g.Psi.String())
	b.WriteString("(")
	writeParamDecls(&b, g.pDecls)
	b.WriteString(") ")
	b.WriteString(g.u_ret.String())
	return b.String()
}

/* Aux, helpers */

/*func BigPsiOk(ds []Decl, env BigPsi, Psi BigPsi) {
	Psi.Ok(ds)
	delta := Psi.ToDelta()
	for _, v := range Psi.tFormals {
		u_I, _ := v.u_I.(TNamed) // \tau_I, already checked by psi.Ok
		u_I.Ok(ds, delta)        // !!! Submission version T-Type, t_i => t_I
	}
}*/

func makeRootPsi(Psi BigPsi) BigPsi {
	tFormals := make([]TFormal, len(PRIMITIVE_PSI.TFormals)+len(Psi.TFormals))
	for i, v := range PRIMITIVE_PSI.TFormals {
		tFormals[i] = v
	}
	for i, v := range Psi.TFormals {
		tFormals[i+len(PRIMITIVE_PSI.TFormals)] = v
	}
	return BigPsi{tFormals}
}

// Pre: isStruct(ds, u_S)
func isRecursiveFieldType(ds []Decl, seen map[string]TNamed,
	u_S TNamed) bool {
	k := u_S.String()
	if _, ok := seen[k]; ok {
		return true
	}
	for _, v := range fields(ds, u_S) {
		if !isStructType(ds, v.u) { // v.u is TNamed -- params OK, type args don't allow writing recursive struct
			continue
		}
		seen1 := make(map[string]TNamed)
		for k1, v1 := range seen {
			seen1[k1] = v1
		}
		seen1[k] = u_S
		if isRecursiveFieldType(ds, seen1, v.u.(TNamed)) {
			return true
		}
	}
	return false
}

// Pre: isNamedIfaceType(ds, t_I), t_I OK already checked
func isRecursiveInterfaceEmbedding(ds []Decl, seen map[string]TNamed,
	u_I TNamed) bool {
	k := u_I.String()
	if _, ok := seen[k]; ok {
		return true
	}
	td := GetTDecl(ds, u_I.TName).(ITypeLit)
	for _, v := range td.specs {
		emb, ok := v.(TNamed)
		if !ok {
			continue
		}
		seen1 := make(map[string]TNamed)
		for k1, v1 := range seen {
			seen1[k1] = v1
		}
		seen1[k] = u_I
		if isRecursiveInterfaceEmbedding(ds, seen1, emb) {
			return true
		}
	}
	return false
}

// Doesn't include "(...)" -- slightly more convenient for debug messages
func writeFieldDecls(b *strings.Builder, fds []FieldDecl) {
	if len(fds) > 0 {
		b.WriteString(fds[0].String())
		for _, v := range fds[1:] {
			b.WriteString("; " + v.String())
		}
	}
}

func writeParamDecls(b *strings.Builder, pds []ParamDecl) {
	if len(pds) > 0 {
		b.WriteString(pds[0].String())
		for _, v := range pds[1:] {
			b.WriteString(", " + v.String())
		}
	}
}
