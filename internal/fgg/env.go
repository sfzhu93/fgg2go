package fgg

import (
	"reflect"
)

type MethodDeltaResult struct {
	Name2Sig    map[Name]*Sig
	OrderedName []Name
}

type Env struct {
	No_assertion             bool
	OptAssertionForZeroParam bool
	N_bar                    map[int]struct{} //we update the arity set while translating the program, and call makeMetas in the end
	Param_index_set          map[int]struct{}
	//fgExprIsStructLit map[fg.FGExpr]bool
	StructVariables         map[string]struct{}
	Stupid                  bool
	NoUnusedVar             bool
	defVar                  map[string]struct{}
	usedVar                 map[string]struct{}
	fggDecls                []Decl
	SName2Lit               map[Name]*STypeLit //TODO: optimize with pointers to interned objects
	IName2Lit               map[Name]*ITypeLit
	name2Methods            map[Name][]*MethDecl
	IType2MethDelta         map[string]MethodDeltaResult
	UseStructMethDeltaCache bool
	SType2MethDelta         map[string]MethodDeltaResult
	ProgramHasAssertion     bool
}

func NewEnv(ds []Decl) Env {
	ret := Env{
		No_assertion:    false,
		N_bar:           nil,
		Param_index_set: nil,
		StructVariables: nil,
		Stupid:          false,
		NoUnusedVar:     false,
		defVar:          nil,
		usedVar:         nil,
		fggDecls:        nil,
		SName2Lit:       nil,
		IName2Lit:       nil,
		name2Methods:    nil,
	}
	ret.fggDecls = ds
	ret.IType2MethDelta = make(map[string]MethodDeltaResult)
	ret.InitNameSets(ds)
	return ret
}

func (env *Env) EnterMethod() {
	env.SType2MethDelta = make(map[string]MethodDeltaResult)
}

func (env *Env) ExitMethod() {
	env.SType2MethDelta = nil
}

func (env *Env) InitNameSets(ds []Decl) {
	env.SName2Lit = make(map[Name]*STypeLit)
	env.IName2Lit = make(map[Name]*ITypeLit)
	env.name2Methods = make(map[Name][]*MethDecl)
	for _, decl := range ds {
		switch decl_ := decl.(type) {
		case STypeLit:
			env.SName2Lit[decl_.T_name] = &decl_
		case ITypeLit:
			env.IName2Lit[decl_.GetName()] = &decl_
		case MethDecl:
			if _, ok := env.name2Methods[decl_.TRecv]; !ok {
				env.name2Methods[decl_.TRecv] = []*MethDecl{}
			}
			env.name2Methods[decl_.TRecv] = append(env.name2Methods[decl_.TRecv], &decl_)
		}
	}
}

// Check if u is a \tau_S -- implicitly must be a TNamed
func (env *Env) IsStructType(u Type) bool {
	if u1, ok := u.(TNamed); ok {
		if _, ok := env.SName2Lit[u1.TName]; ok {
			return true
		}
	}
	return false
}

func (env *Env) IsInterfaceType(u Type) bool {
	if u1, ok := u.(TNamed); ok {
		return env.IsInterfaceName(u1.TName)
	}
	return false
}

func (env *Env) IsInterfaceName(u Name) bool {
	_, ok := env.IName2Lit[u]
	return ok
}

func (env *Env) IsStructName(u Name) bool {
	_, ok := env.SName2Lit[u]
	return ok
}

func (env *Env) MakeEtaDelta(delta Delta, Psi BigPsi, psi SmallPsi) (bool, EtaOpen) {
	eta := MakeEtaOpen(Psi, psi)
	for _, v := range Psi.TFormals {
		a := v.Name.SubsEtaOpen(eta)
		u_I := v.U_I.SubsEtaOpen(eta)
		if !a.ImplsDeltaWithEnv(env, delta, u_I) {
			return false, eta
		}
	}
	return true, eta
}

func (env *Env) CanOptimizedOutTypeMetadata(u Type) bool {
	switch u_ := u.(type) {
	case TNamed:
		if ty, ok := env.IName2Lit[u_.TName]; ok {
			println("can optimize ?", u_.TName)
			println("ty.psi.tformals.length = ", len(ty.Psi.TFormals))
			return len(ty.Psi.TFormals) == 0
		}
	default:
		return false
	}
	return false
}

func (env *Env) MethodsDelta(delta Delta, u Type) (map[Name]*Sig, []Name) {
	//fmt.Printf("MethodsDelta:\ndelta = %#v\nu = %#v\n", delta, u)
	res := make(map[Name]*Sig)
	order := []Name{}
	switch u_ := u.(type) {
	case TNamed:
		if _, ok := env.SName2Lit[u_.TName]; ok {
			//memorize here
			ok := false
			var cachedMethodsDelta MethodDeltaResult
			if env.SType2MethDelta != nil {
				cachedMethodsDelta, ok = env.SType2MethDelta[u_.String()]
			}
			if env.UseStructMethDeltaCache && ok {
				res = cachedMethodsDelta.Name2Sig
				order = cachedMethodsDelta.OrderedName
			} else {
				//can we remove this one, because Phi = Phi'?
				//no, we have to init psi also
				for _, methdecl := range env.name2Methods[u_.TName] {
					if ok, eta := env.MakeEtaDelta(delta, methdecl.Psi_recv, u_.UArgs); ok {
						tmp := methdecl.ToSig().TSubs(eta)
						res[methdecl.GetName()] = &tmp
						order = append(order, methdecl.GetName())
					}
				}
			}
		} else if td, ok := env.IName2Lit[u_.TName]; ok { // N.B. u is a TName, \tau_I (not a TParam)
			if cachedMethodsDelta, ok := env.IType2MethDelta[u_.String()]; ok {
				res = cachedMethodsDelta.Name2Sig
				order = cachedMethodsDelta.OrderedName
			} else {
				subs := make(map[TParam]Type) // Cf. MakeEta
				for i := 0; i < len(td.Psi.TFormals); i++ {
					subs[td.Psi.TFormals[i].Name] = u_.UArgs[i]
				}
				for _, s := range td.GetSpecs() {
					switch s1 := s.(type) {
					case Sig:
						tmp := s1.TSubs(subs)
						res[s1.GetMethod()] = &tmp
						order = append(order, s1.GetMethod())
					case TNamed: // Embedded u_I
						for k, v := range Methods(env.fggDecls, s1.TSubs(subs)) { // cycles? (cf. submission version)
							res[k] = &v
						}
					default:
						panic("Unknown Spec kind: " + reflect.TypeOf(s).String())
					}
				}
				env.IType2MethDelta[u_.String()] = MethodDeltaResult{Name2Sig: res, OrderedName: order}
			}
		}
	case TParam:
		upper, ok := delta[u_]
		if !ok {
			panic("Unknown type: " + u.String())
		}
		//return methodsDelta(ds, delta, bounds(delta, cast)) // !!! delegate to bounds
		return env.MethodsDelta(delta, upper) //fgg.MethodsDelta(env.fggDecls, delta, upper)
	default:
		panic("Unknown type: " + u.String()) // Perhaps redundant if all TDecl OK checked first
	}
	return res, order
}
