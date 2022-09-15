package fgg

import (
	"fmt"
	"reflect"
)

var _ = fmt.Errorf

/* Constants */

// Hack
var STRING_TYPE_MONOM = TNamed{string(STRING_TYPE), SmallPsi{}} // Because TNamed required

/* GroundEnv */

// Basically a Gamma for only ground TNamed
type GroundGamma map[Name]Type // Pre: forall TName, isGround

/**
 * Build Omega -- (morally) a map from ground FGG types to Sigs of (potential)
 * calls on that receiver.  N.B., calls are recorded only as seen for each
 * specific receiver type -- i.e., omega does not attempt to "respect"
 * subtyping (cf. "zigzagging" in fgg_monom).
 */

// Pre: IsMonomOK
func GetOmega(ds []Decl, e_main FGGExpr) Omega {
	omega := Omega{make(map[string]Type), make(map[string]MethInstan)}
	collectExpr(ds, make(GroundGamma), e_main, omega)
	fixomega(ds, omega)
	//omega.Println()
	return omega
}

/* Omega, MethInstan */

type Omega struct {
	// Keys given by toKey_Wt, toKey_Wm
	us map[string]Type // Pre: all TNamed are isGround FCGG needs to support channels. changed it to Type
	ms map[string]MethInstan
}

func (w Omega) clone() Omega {
	us := make(map[string]Type)
	ms := make(map[string]MethInstan)
	for k, v := range w.us {
		us[k] = v
	}
	for k, v := range w.ms {
		ms[k] = v
	}
	return Omega{us, ms}
}

func (w Omega) Println() {
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

type MethInstan struct {
	u_recv TNamed // Pre: isGround
	meth   Name
	psi    SmallPsi // Pre: all isGround
}

// Pre: isGround(u_ground)
func ToKey_Wt(u_ground Type) string {
	return u_ground.String()
}

// Pre: isGround(x.u_ground)
func ToKey_Wm(x MethInstan) string {
	return x.u_recv.String() + "_" + x.meth + "_" + x.psi.String()
}

/* fixOmega */

func fixomega(ds []Decl, omega Omega) {
	/*fmt.Println("......initial.........", len(omega.us), len(omega.ms))
	omega.Println()
	fmt.Println(".............", len(omega.us), len(omega.ms))*/
	for auxG(ds, omega) {
		//omega.Println()
		//fmt.Println(".............", len(omega.us), len(omega.ms))
	}
}

/* Expressions */

func collectStmtSeqForMonom(ds []Decl, gamma GroundGamma, stmts []FGGStmt, omega Omega) bool {
	new_gamma := make(GroundGamma)
	for k, v := range gamma {
		new_gamma[k] = v
	}
	ret := false
	for _, x := range stmts {
		switch x_ := x.(type) {
		case ReturnStmt:
			ret = ret || collectExpr(ds, new_gamma, x_.ExprBody, omega)
		case FGGCaseSelectStmt:
			for _, casebody := range x_.Cases {
				new_stmt_seq := []FGGStmt{casebody.CaseGuard.ToFGGStmt()}
				ret = ret || collectStmtSeqForMonom(ds, new_gamma, append(new_stmt_seq, casebody.Body...), omega)
			}
			if x_.DefaultCase != nil && len(x_.DefaultCase) > 0 {
				ret = ret || collectStmtSeqForMonom(ds, new_gamma, x_.DefaultCase, omega)
			}
		case FGGChCloseStmt:
			ret = ret || collectExpr(ds, new_gamma, x_.Body, omega)
		case FGGChDispatchStmt:
			ret = ret || collectExpr(ds, new_gamma, x_.LeftExpr, omega)
			ret = ret || collectExpr(ds, new_gamma, x_.RightExpr, omega)
		case FGGAssignmentStmt:
			ret = ret || collectExpr(ds, new_gamma, x_.Body, omega)
			gamma1 := make(Gamma)
			for k, v := range new_gamma {
				gamma1[k] = v
			}
			new_gamma[x_.VarName] = x_.Body.Typing(ds, make(Delta), gamma1, false)
			//TODO: add to omega set later on? no
		case GoroutineStmt:
			ret = ret || collectExpr(ds, new_gamma, x_.CallExpr, omega)
		default:
			panic("unknown statement type.")
		}
	}
	return ret
}

// gamma used to type Call receiver
func collectExpr(ds []Decl, gamma GroundGamma, e FGGExpr, omega Omega) bool {
	res := false
	switch e1 := e.(type) {
	case Variable:
		return res
	case StructLit:
		for _, elem := range e1.Elems {
			res = collectExpr(ds, gamma, elem, omega) || res
		}
		k := ToKey_Wt(e1.U_S)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = e1.U_S
			res = true
		}
	case Select:
		return collectExpr(ds, gamma, e1.E_S, omega)
	case Call:
		res = collectExpr(ds, gamma, e1.E_recv, omega) || res
		for _, e_arg := range e1.Args {
			res = collectExpr(ds, gamma, e_arg, omega) || res
		}
		gamma1 := make(Gamma)
		for k, v := range gamma {
			gamma1[k] = v
		}
		u_recv := e1.E_recv.Typing(ds, make(Delta), gamma1, false).(TNamed)
		k_t := ToKey_Wt(u_recv)
		if _, ok := omega.us[k_t]; !ok {
			omega.us[k_t] = u_recv
			res = true
		}
		m := MethInstan{u_recv, e1.Meth, e1.GetTArgs()} // N.B. type/method instans recorded separately
		k_m := ToKey_Wm(m)
		if _, ok := omega.ms[k_m]; !ok {
			omega.ms[k_m] = m
			res = true
		}
	case Assert:
		res = collectExpr(ds, gamma, e1.E_I, omega) || res
		u := e1.U_cast.(TNamed)
		k := ToKey_Wt(u)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = u
			res = true
		}
	case StringLit: // CHECKME
		//k := toKey_Wt(STRING_TYPE)
		k := string(STRING_TYPE)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = STRING_TYPE_MONOM
			res = true // CHECKME
		}
	case Sprintf:
		//k := toKey_Wt(STRING_TYPE)
		k := string(STRING_TYPE)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = STRING_TYPE_MONOM
			res = true
		}
		for _, arg := range e1.Args {
			res = collectExpr(ds, gamma, arg, omega) || res
		}
	case ChanRecv:
		return collectExpr(ds, gamma, e1.Body, omega)
	case MakeChan:
		u := e1.ChType
		k := ToKey_Wt(u)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = u
			res = true
		}
	default:
		panic("Unknown Expr kind: " + reflect.TypeOf(e).String() + "\n\t" +
			e.String())
	}
	return res
}

/* Aux */

// Return true if omega has changed
// N.B. mutating omega in each sub-step -- can produce many levels of nesting within one G step
// ^also non-deterministic progress, because mutating maps while ranging; also side-effect results may depend on iteration order over maps
// N.B. no closure over types occurring in bounds, or *interface decl* method sigs
func auxG(ds []Decl, omega Omega) bool {
	res := false
	res = auxF(ds, omega) || res
	res = auxI(ds, omega) || res
	res = auxM(ds, omega) || res
	res = auxS(ds, make(Delta), omega) || res
	res = auxChan(ds, omega) || res
	// I/face embeddings
	res = auxE1(ds, omega) || res
	res = auxE2(ds, omega) || res
	//res = auxP(ds, omega) || res
	return res
}

func auxChan(ds []Decl, omega Omega) bool {
	res := false
	tmp := make(map[string]Type)
	for _, u := range omega.us {
		if u_, ok := u.(ChannelType); ok {
			tmp[ToKey_Wt(u_.ElementType)] = u_.ElementType
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

func auxF(ds []Decl, omega Omega) bool {
	res := false
	tmp := make(map[string]Type)
	for _, u := range omega.us {
		if !isStructType(ds, u) { //|| u.Equals(STRING_TYPE) { // CHECKME
			continue
		}
		if u_tnamed, ok := u.(TNamed); ok {
			for _, u_f := range Fields(ds, u_tnamed) {
				//cast := u_f.u.(TNamed)
				tmp[ToKey_Wt(u_f.u)] = u_f.u
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

func auxI(ds []Decl, omega Omega) bool {
	res := false
	tmp := make(map[string]MethInstan)
	for _, m := range omega.ms {
		if !IsNamedIfaceType(ds, m.u_recv) {
			continue
		}
		for _, m1 := range omega.ms {
			if !IsNamedIfaceType(ds, m1.u_recv) {
				continue
			}
			if m1.u_recv.Impls(ds, m.u_recv) {
				mm := MethInstan{m1.u_recv, m.meth, m.psi}
				tmp[ToKey_Wm(mm)] = mm
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

func auxM(ds []Decl, omega Omega) bool {
	res := false
	tmp := make(map[string]Type)
	for _, m := range omega.ms {
		gs := methods(ds, m.u_recv)
		for _, g := range gs { // Should be only g s.t. g.meth == m.meth
			if g.meth != m.meth {
				continue
			}
			eta := MakeEta(g.Psi, m.psi)
			for _, pd := range g.pDecls {
				u_pd := pd.U.SubsEta(eta) // HERE: need receiver subs also? cf. map.fgg "type b Eq(b)" -- methods should be ok?
				tmp[ToKey_Wt(u_pd)] = u_pd
			}
			u_ret := g.u_ret.SubsEta(eta)
			tmp[ToKey_Wt(u_ret)] = u_ret
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

func auxS(ds []Decl, delta Delta, omega Omega) bool {
	res := false
	tmp := make(map[string]MethInstan)
	clone := omega.clone()
	for _, m := range clone.ms {
		for _, u := range clone.us {
			if !isStructType(ds, u) || !u.ImplsDelta(ds, delta, m.u_recv) {
				continue
			}
			u_ := u.(TNamed)
			x0, xs, stmts := body(ds, u_, m.meth, m.psi)
			gamma := make(GroundGamma)
			gamma[x0.Name] = x0.U
			for _, pd := range xs {
				gamma[pd.Name] = pd.U
			}
			m1 := MethInstan{u_, m.meth, m.psi}
			k := ToKey_Wm(m1)
			//if _, ok := omega.ms[k]; !ok { // No: initial collectExpr already adds to omega.ms
			tmp[k] = m1
			res = collectStmtSeqForMonom(ds, gamma, stmts.StmtBody, omega) || res //collectExpr(ds, gamma, e, omega) || res
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
func auxE1(ds []Decl, omega Omega) bool {
	res := false
	tmp := make(map[string]Type)
	for _, u := range omega.us {
		if !isNamedIfaceType(ds, u) {
			continue
		}
		u_ := u.(TNamed)
		td_I := GetTDecl(ds, u_.TName).(ITypeLit)
		eta := MakeEta(td_I.Psi, u_.UArgs)
		for _, s := range td_I.specs {
			if u_emb, ok := s.(TNamed); ok {
				u_sub := u_emb.SubsEta(eta)
				tmp[ToKey_Wt(u_sub)] = u_sub
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
func auxE2(ds []Decl, omega Omega) bool {
	res := false
	tmp := make(map[string]MethInstan)
	for _, m := range omega.ms {
		if !isNamedIfaceType(ds, m.u_recv) {
			continue
		}
		td_I := GetTDecl(ds, m.u_recv.TName).(ITypeLit)
		eta := MakeEta(td_I.Psi, m.u_recv.UArgs)
		for _, s := range td_I.specs {
			if u_emb, ok := s.(TNamed); ok {
				u_sub := u_emb.SubsEta(eta).(TNamed)
				gs := methods(ds, u_sub)
				for _, g := range gs {
					if m.meth == g.meth {
						m_emb := MethInstan{u_sub, m.meth, m.psi}
						tmp[ToKey_Wm(m_emb)] = m_emb
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

/* Helpers */

func isGround(u TNamed) bool {
	for _, v := range u.UArgs {
		if u1, ok := v.(TNamed); !ok {
			return false
		} else if !isGround(u1) {
			return false
		}
	}
	return true
}
