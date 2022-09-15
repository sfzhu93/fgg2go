package fgg

import (
	"fmt"
	"reflect"
	"strings"
)

var _ = fmt.Errorf
var _ = reflect.Append
var _ = strings.Compare

// Return true if *not* nomono
func IsMonomOK(p FGGProgram) (bool, string) {
	ds := p.GetDecls()
	for _, v := range ds {
		if md, ok := v.(MethDecl); ok {
			omega := Nomega{make(map[string]Type), make(map[string]MethInstanOpen)}
			delta := md.Psi_recv.ToDelta()
			for _, v := range md.Psi_meth.TFormals {
				delta[v.Name] = v.U_I
			}
			gamma := make(Gamma)
			psi_recv := make(SmallPsi, len(md.Psi_recv.TFormals))
			for i, v := range md.Psi_recv.TFormals {
				psi_recv[i] = v.Name
			}
			//psi_recv = md.Psi_recv.Hat()
			u_recv := TNamed{md.TRecv, psi_recv}
			gamma[md.XRecv] = u_recv
			omega.us[ToKey_Wt(u_recv)] = u_recv
			for _, v := range md.pDecls { // TODO: factor out
				gamma[v.Name] = v.U
			}
			collectExprForStmtSeq(ds, delta, gamma, md.BodyStmt.StmtBody, omega)
			//collectExprOpen(ds, delta, gamma, md.E_body, omega)
			if ok, msg := nomonoOmega(ds, delta, md, omega); ok {
				return false, msg
			}
		}
	}
	return true, ""
}

// Return true if nomono
func nomonoOmega(ds []Decl, delta Delta, md MethDecl, omega Nomega) (bool, string) {
	for auxGOpen(ds, delta, omega) {
		for _, v := range omega.ms {
			if !isStructType(ds, v.u_recv) {
				continue
			}
			u_S := v.u_recv.(TNamed)
			if u_S.TName == md.TRecv && v.meth == md.name {
				if occurs(md.Psi_recv, u_S.UArgs) {
					return true, md.TRecv + md.Psi_recv.String() + " ->* " + md.TRecv +
						"(" + SmallPsi(u_S.UArgs).String() + ")"
				}
				if occurs(md.Psi_meth, v.psi) {
					return true, md.TRecv + md.Psi_recv.String() + "." + md.name +
						md.Psi_meth.String() + " ->* " + md.name + "(" + v.psi.String() + ")"
				}
			}
		}
	}
	return false, ""
}

// Pre: len(Psi) == len(psi)
func occurs(Psi BigPsi, psi SmallPsi) bool {
	for i, v := range Psi.TFormals {
		if cast, ok := psi[i].(TNamed); ok { // !!! simplified
			for _, x := range fv(cast) {
				if x.Equals(v.Name) {
					return true
				}
			}
		}
	}
	return false
}

func fv(u Type) []TParam {
	if cast, ok := u.(TParam); ok {
		return []TParam{cast}
	}
	res := []TParam{}
	cast := u.(TNamed)
	for _, v := range cast.UArgs {
		res = append(res, fv(v)...)
	}
	return res
}

/* !!! Duplication of Omega (etc.) for non-ground types -- if only Go had generics! */

type Nomega struct {
	us map[string]Type
	ms map[string]MethInstanOpen
}

func (w Nomega) clone() Nomega {
	us := make(map[string]Type)
	ms := make(map[string]MethInstanOpen)
	for k, v := range w.us {
		us[k] = v
	}
	for k, v := range w.ms {
		ms[k] = v
	}
	return Nomega{us, ms}
}

func (w Nomega) Println() {
	fmt.Println("=== Type instances:")
	for _, v := range w.us {
		fmt.Println(v)
	}
	fmt.Println("--- Method instances:")
	for _, v := range w.ms {
		fmt.Println(v.u_recv, v.meth, v.psi)
	}
	fmt.Println("===")
}

// Factor out with MethInstan
type MethInstanOpen struct {
	u_recv Type
	meth   Name
	psi    SmallPsi
}

func tokeyWtOpen(u Type) string {
	return u.String()
}

func tokeyWmOpen(x MethInstanOpen) string {
	return x.u_recv.String() + "_" + x.meth + "_" + x.psi.String()
}

func collectExprForStmtSeq(ds []Decl, delta Delta, gamma Gamma, body []FGGStmt, omega Nomega) bool {
	new_gamma := make(Gamma)
	for k, v := range gamma {
		new_gamma[k] = v
	}
	ret := false
	for _, x := range body {
		switch x_ := x.(type) {
		case ReturnStmt:
			ret = ret || collectExprOpen(ds, delta, new_gamma, x_.ExprBody, omega)
		case FGGCaseSelectStmt:
			for _, casebody := range x_.Cases {
				new_stmt_seq := []FGGStmt{casebody.CaseGuard.ToFGGStmt()}
				ret = ret || collectExprForStmtSeq(ds, delta, new_gamma, append(new_stmt_seq, casebody.Body...), omega)
			}
			if x_.DefaultCase != nil && len(x_.DefaultCase) > 0 {
				ret = ret || collectExprForStmtSeq(ds, delta, new_gamma, x_.DefaultCase, omega)
			}
		case FGGChCloseStmt:
			ret = ret || collectExprOpen(ds, delta, new_gamma, x_.Body, omega)
		case FGGChDispatchStmt:
			ret = ret || collectExprOpen(ds, delta, new_gamma, x_.LeftExpr, omega)
			ret = ret || collectExprOpen(ds, delta, new_gamma, x_.RightExpr, omega)
		case FGGAssignmentStmt:
			ret = ret || collectExprOpen(ds, delta, new_gamma, x_.Body, omega)
			new_gamma[x_.VarName] = x_.Body.Typing(ds, delta, new_gamma, false)
			//TODO: add to omega set later on? no
		case GoroutineStmt:
			ret = ret || collectExprOpen(ds, delta, new_gamma, x_.CallExpr, omega)
		default:
			panic("unknown statement type.")
		}
	}
	return ret
}

func collectExprOpen(ds []Decl, delta Delta, gamma Gamma, e FGGExpr, omega Nomega) bool {
	res := false
	switch e1 := e.(type) {
	case Variable:
		return res
	case StructLit:
		for _, elem := range e1.Elems {
			res = collectExprOpen(ds, delta, gamma, elem, omega) || res
		}
		k := tokeyWtOpen(e1.U_S)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = e1.U_S
			res = true
		}
	case Select:
		return collectExprOpen(ds, delta, gamma, e1.E_S, omega)
	case Call:
		res = collectExprOpen(ds, delta, gamma, e1.E_recv, omega) || res
		for _, e_arg := range e1.Args {
			res = collectExprOpen(ds, delta, gamma, e_arg, omega) || res
		}
		gamma1 := make(Gamma)
		for k, v := range gamma {
			gamma1[k] = v
		}
		u_recv := e1.E_recv.Typing(ds, delta, gamma1, false)
		k_t := tokeyWtOpen(u_recv)
		if _, ok := omega.us[k_t]; !ok {
			omega.us[k_t] = u_recv
			res = true
		}
		m := MethInstanOpen{u_recv, e1.Meth, e1.GetTArgs()} // CHECKME: why add u_recv separately?
		k_m := tokeyWmOpen(m)
		if _, ok := omega.ms[k_m]; !ok {
			omega.ms[k_m] = m
			res = true
		}
	case Assert:
		res = collectExprOpen(ds, delta, gamma, e1.E_I, omega) || res
		u := e1.U_cast
		k := tokeyWtOpen(u)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = u
			res = true
		}
	case StringLit: // CHECKME
		k := tokeyWtOpen(STRING_TYPE)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = STRING_TYPE
			res = true // CHECKME
		}
	case Sprintf:
		k := tokeyWtOpen(STRING_TYPE)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = STRING_TYPE
			res = true
		}
		for _, arg := range e1.Args {
			res = collectExprOpen(ds, delta, gamma, arg, omega) || res
		}
	case MakeChan:
		k := e1.ChType.String()
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = STRING_TYPE
			res = true
		}
	case ChanRecv:
		res = collectExprOpen(ds, delta, gamma, e1.Body, omega) || res
	default:
		panic("Unknown Expr kind: " + reflect.TypeOf(e).String() + "\n\t" +
			e.String())
	}
	return res
}

/* Aux */

// Return true if omega has changed
// N.B. no closure over types occurring in bounds, or *interface decl* method sigs
func auxGOpen(ds []Decl, delta Delta, omega Nomega) bool {
	res := false
	res = auxFOpen(ds, omega) || res
	res = auxI2(ds, delta, omega) || res
	res = auxMOpen(ds, delta, omega) || res
	res = auxSOpen(ds, delta, omega) || res
	// I/face embeddings
	res = auxE1Open(ds, omega) || res
	res = auxE2Open(ds, omega) || res
	return res
}

func auxFOpen(ds []Decl, omega Nomega) bool {
	res := false
	tmp := make(map[string]Type)
	for _, u := range omega.us {
		if !isStructType(ds, u) { //|| u.Equals(STRING_TYPE) { // CHECKME
			continue
		}
		for _, u_f := range Fields(ds, u.(TNamed)) {
			cast := u_f.u
			tmp[tokeyWtOpen(cast)] = cast
		}
	}
	for k, v := range tmp {
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = v
			res = true
		}
	}
	return res
}

func auxI2(ds []Decl, delta Delta, omega Nomega) bool {
	res := false
	tmp := make(map[string]MethInstanOpen)
	for _, m := range omega.ms {
		if !IsNamedIfaceType(ds, m.u_recv) {
			continue
		}
		for _, m1 := range omega.ms {
			if !IsNamedIfaceType(ds, m1.u_recv) {
				continue
			}
			if m1.u_recv.ImplsDelta(ds, delta, m.u_recv) {
				mm := MethInstanOpen{m1.u_recv, m.meth, m.psi}
				tmp[tokeyWmOpen(mm)] = mm
			}
		}
	}
	for k, v := range tmp {
		if _, ok := omega.ms[k]; !ok {
			omega.ms[k] = v
			res = true
		}
	}
	return res
}

func auxMOpen(ds []Decl, delta Delta, omega Nomega) bool {
	res := false
	tmp := make(map[string]Type)
	for _, m := range omega.ms {
		gs, _ := MethodsDelta(ds, delta, m.u_recv)
		for _, g := range gs { // Should be only g s.t. g.meth == m.meth
			if g.meth != m.meth {
				continue
			}
			eta := MakeEtaOpen(g.Psi, m.psi)
			for _, pd := range g.pDecls {
				u_pd := pd.U.SubsEtaOpen(eta) // HERE: need receiver subs also? cf. map.fgg "type b Eq(b)" -- methods should be ok?
				tmp[tokeyWtOpen(u_pd)] = u_pd
			}
			u_ret := g.u_ret.SubsEtaOpen(eta)
			tmp[tokeyWtOpen(u_ret)] = u_ret
		}
	}
	for k, v := range tmp {
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = v
			res = true
		}
	}
	return res
}

func auxSOpen(ds []Decl, delta Delta, omega Nomega) bool {
	res := false
	tmp := make(map[string]MethInstanOpen)
	clone := omega.clone()
	for _, m := range clone.ms {
		for _, u := range clone.us {
			u_recv := bounds(delta, m.u_recv) // !!! cf. plain type param
			if !isStructType(ds, u) || !u.ImplsDelta(ds, delta, u_recv) {
				continue
			}
			u_S := u.(TNamed)
			x0, xs, stmts := body(ds, u_S, m.meth, m.psi)
			gamma := make(Gamma)
			gamma[x0.Name] = x0.U.(TNamed)
			for _, pd := range xs {
				gamma[pd.Name] = pd.U
			}
			m1 := MethInstanOpen{u_S, m.meth, m.psi}
			k := tokeyWmOpen(m1)
			//if _, ok := omega.ms[k]; !ok { // No: initial collectExpr already adds to omega.ms
			tmp[k] = m1
			res = collectExprForStmtSeq(ds, delta, gamma, stmts.StmtBody, omega) || res //collectExprOpen(ds, delta, gamma, e, omega) || res
			//}
		}
	}
	for k, v := range tmp {
		if _, ok := omega.ms[k]; !ok {
			omega.ms[k] = v
			res = true
		}
	}
	return res
}

// Add embedded types
func auxE1Open(ds []Decl, omega Nomega) bool {
	res := false
	tmp := make(map[string]TNamed)
	for _, u := range omega.us {
		if !isNamedIfaceType(ds, u) { // CHECKME: type param
			continue
		}
		u_I := u.(TNamed)
		td_I := GetTDecl(ds, u_I.TName).(ITypeLit)
		eta := MakeEtaOpen(td_I.Psi, u_I.UArgs)
		for _, s := range td_I.specs {
			if u_emb, ok := s.(TNamed); ok {
				u_sub := u_emb.SubsEtaOpen(eta).(TNamed)
				tmp[tokeyWtOpen(u_sub)] = u_sub
			}
		}
	}
	for k, v := range tmp {
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = v
			res = true
		}
	}
	return res
}

// Propagate method instances up to embedded supertypes
func auxE2Open(ds []Decl, omega Nomega) bool {
	res := false
	tmp := make(map[string]MethInstanOpen)
	for _, m := range omega.ms {
		if !isNamedIfaceType(ds, m.u_recv) { // CHECKME: type param
			continue
		}
		u_I := m.u_recv.(TNamed)
		td_I := GetTDecl(ds, u_I.TName).(ITypeLit)
		eta := MakeEtaOpen(td_I.Psi, u_I.UArgs)
		for _, s := range td_I.specs {
			if u_emb, ok := s.(TNamed); ok {
				u_sub := u_emb.SubsEtaOpen(eta).(TNamed)
				gs := methods(ds, u_sub)
				for _, g := range gs {
					if m.meth == g.meth {
						m_emb := MethInstanOpen{u_sub, m.meth, m.psi}
						tmp[tokeyWmOpen(m_emb)] = m_emb
					}
				}
			}
		}
	}
	for k, v := range tmp {
		if _, ok := omega.ms[k]; !ok {
			omega.ms[k] = v
			res = true
		}
	}
	return res
}

/*

















































 */

/*
 * Deprecated: old CFG-based test
 */

type RecvMethPair struct {
	t_recv Name // Pre: t_S
	m      Name // TODO rename
}

func (x0 RecvMethPair) equals(x RecvMethPair) bool {
	return x0.t_recv == x.t_recv && x0.m == x.m
}

type cTypeArgs struct {
	psi_recv SmallPsi
	psi_meth SmallPsi
}

func (x0 cTypeArgs) equals(x cTypeArgs) bool {
	return x0.psi_recv.Equals(x.psi_recv) && x0.psi_meth.Equals(x.psi_meth)
}

// Static call graph, agnostic of specific type args (cf. MethInstan)
// N.B. nodes are for struct types
type cgraph struct {
	edges map[RecvMethPair]map[RecvMethPair]([]cTypeArgs)
}

func (x0 cgraph) String() string {
	var b strings.Builder
	for k, v := range x0.edges {
		b.WriteString(k.t_recv)
		b.WriteString(".")
		b.WriteString(k.m)
		b.WriteString(": ")
		b.WriteString(fmt.Sprintf("%v", v))
		b.WriteString("\n")
	}
	return b.String()
}

// CFG-based occurs check -- needs method set unification ("open" impls)
// CHECKME: generally, covariant receiver bounds specialisation
func IsMonomOK_CFG(p FGGProgram) bool {
	ds := p.GetDecls()
	graph := cgraph{make(map[RecvMethPair]map[RecvMethPair]([]cTypeArgs))}
	for _, v := range ds {
		if md, ok := v.(MethDecl); ok {
			buildGraph(ds, md, graph)
		}
	}
	//buildGraphExpr(ds, make(Delta), make(Gamma), ...)  // visit main unnecessary -- CHECKME: all type instans seen?
	//fmt.Println("111:\n", graph.String(), "---")
	cycles := make(map[string]cycle)
	findCycles(graph, cycles)
	/*for _, v := range cycles {
		fmt.Println("aaa:", v)
	}*/
	for _, v := range cycles {
		//fmt.Println("bbb:", v)
		if isNomonoCycle(ds, graph, v) {
			return false
		}
		return true
	}
	return true
}

// Occurs check -- N.B. conservative w.r.t. whether type params actually used
func isNomonoCycle(ds []Decl, graph cgraph, c cycle) bool {
	for _, tArgs := range graph.edges[c[0]][c[1]] {
		if isNomonoTypeArgs(tArgs) || isNomonoCycleAux(ds, graph, c, tArgs, 1) {
			return true
		}
	}
	return false
}

func isNomonoTypeArgs(tArgs cTypeArgs) bool {
	for _, v := range tArgs.psi_recv {
		if containsNestedTParam(v) {
			return true
		}
	}
	for _, v := range tArgs.psi_meth {
		if containsNestedTParam(v) {
			return true
		}
	}
	return false
}

func isNomonoCycleAux(ds []Decl, graph cgraph, c cycle, tArgs cTypeArgs, i int) bool {
	if i >= (len(c) - 1) {
		return false
	}
	next := c[i]
	md := getMDecl(ds, next.t_recv, next.m)
	subs := make(Delta)
	for i, v := range tArgs.psi_recv {
		subs[md.Psi_recv.TFormals[i].Name] = v
	}
	for i, v := range tArgs.psi_meth {
		subs[md.Psi_meth.TFormals[i].Name] = v
	}

	for _, v := range graph.edges[c[i]][c[i+1]] {
		tArgs1 := cTypeArgs{v.psi_recv.TSubs(subs), v.psi_meth.TSubs(subs)}
		if isNomonoTypeArgs(tArgs1) {
			return true
		}
		isNomonoCycleAux(ds, graph, c, tArgs1, i+1)
	}
	return false
}

func getMDecl(ds []Decl, t_recv Name, meth Name) MethDecl {
	for _, v := range ds {
		if md, ok := v.(MethDecl); ok && md.TRecv == t_recv && md.name == meth {
			return md
		}
	}
	panic("MethDecl not found: " + t_recv + "." + meth)
}

func containsNestedTParam(u Type) bool {
	if cast, ok := u.(TNamed); ok {
		for _, v := range cast.UArgs {
			if isOrContainsTParam(v) {
				return true
			}
		}
	}
	return false
}

type cycle []RecvMethPair

func (x0 cycle) toHash() string {
	return fmt.Sprintf("%v", x0)
}

func findCycles(graph cgraph, cycles map[string]cycle) {
	for k, _ := range graph.edges {
		stack := []RecvMethPair{k}
		findCyclesAux(graph, stack, cycles)
	}
}

// DFS -- TODO: start from main more efficient? -- CHECKME: maybe more "correct", w.r.t. omega method discarding
func findCyclesAux(graph cgraph, stack []RecvMethPair, cycles map[string]cycle) {
	targets := graph.edges[stack[len(stack)-1]]
	if targets == nil {
		panic("Shouldn't get in here:")
	}
lab:
	for next, _ := range targets {
		stack1 := append(stack, next)
		if stack1[0].equals(next) {
			cycles[cycle(stack1).toHash()] = stack1
			continue
		}
		for _, prev := range stack[1:] {
			if prev.equals(next) {
				continue lab
			}
		}
		findCyclesAux(graph, stack1, cycles)
	}
}

// "Flat" graph building -- calls not visited (i.e., `body` not used)
// Output: mutates cgraph
func buildGraph(ds []Decl, md MethDecl, graph cgraph) {
	n := RecvMethPair{md.TRecv, md.name}
	graph.edges[n] = make(map[RecvMethPair]([]cTypeArgs))
	delta := md.Psi_meth.ToDelta() // recv params added below
	gamma := make(Gamma)
	psi := make(SmallPsi, len(md.Psi_recv.TFormals))
	for i, v := range md.Psi_recv.TFormals {
		delta[v.Name] = v.U_I
		psi[i] = v.Name
	}
	gamma[md.XRecv] = TNamed{md.TRecv, psi}
	for _, v := range md.pDecls { // TODO: factor out
		gamma[v.Name] = v.U
	}
	buildGraphExpr(ds, delta, gamma, n, md.E_body, graph)
}

// "Flat" graph building -- calls not visited (i.e., `body` not used)
func buildGraphExpr(ds []Decl, delta Delta, gamma Gamma, curr RecvMethPair, e1 FGGExpr, graph cgraph) {
	switch e := e1.(type) {
	case Variable:
	case StructLit:
		for _, elem := range e.Elems {
			buildGraphExpr(ds, delta, gamma, curr, elem, graph)
		}
	case Select:
		buildGraphExpr(ds, delta, gamma, curr, e.E_S, graph)
	case Call:
		buildGraphExpr(ds, delta, gamma, curr, e.E_recv, graph)
		for _, arg := range e.Args {
			buildGraphExpr(ds, delta, gamma, curr, arg, graph)
		}
		u_recv := e.E_recv.Typing(ds, delta, gamma, true)

		if isStructType(ds, u_recv) { // u_recv is a TNamed struct
			u_S := u_recv.(TNamed)
			putTArgs(graph, curr, u_S, e.Meth, e.t_args)

		} else { // TNamed interface or TParam
			u_I := u_recv // Or type param
			if _, ok := u_I.(TParam); ok {
				u_I = u_I.TSubs(delta) // CHECKME
			}
			for _, v := range ds {
				if d, ok := v.(STypeLit); ok {

					// method set unification instead of basic impls? -- or using bounds (hat) sufficient?
					u_S := TNamed{d.T_name, d.Psi.Hat()} // !!!
					if u_S.ImplsDelta(ds, delta, u_I) {
						putTArgs(graph, curr, u_S, e.Meth, e.t_args)
					}

				}
			}
		}
	case Assert:
		buildGraphExpr(ds, delta, gamma, curr, e.E_I, graph)
	default:
		panic("Unknown Expr kind: " + reflect.TypeOf(e1).String() + "\n\t" +
			e1.String())
	}
}

// u_recv is target u_S
func putTArgs(graph cgraph, curr RecvMethPair, u_recv TNamed, meth Name, psi_meth SmallPsi) {
	edges := graph.edges[curr]
	/*if edges == nil {
		edges = make(map[node]([]cTypeArgs))
		graph.edges[curr] = edges
	}*/
	target := RecvMethPair{u_recv.TName, meth}
	tArgs := edges[target]
	if tArgs == nil {
		tArgs = []cTypeArgs{}
	}
	tArgs = append(tArgs, cTypeArgs{u_recv.UArgs, psi_meth})
	edges[target] = tArgs
}
