package dictpass

import (
	"github.com/sfzhu93/fgg2go/internal/fg"
	"github.com/sfzhu93/fgg2go/internal/fgg"
)

func AuxRechan(zeta map[fgg.Name]fg.FGExpr, e_ddagger fg.FGExpr, isStructLit bool, tau fgg.Type) fg.FGExpr {
	switch tau.(type) {
	case fgg.ChannelType:
		if isStructLit {
			return fg.NewStructLit(fg.Type("chan_wrapper"), []fg.FGExpr{
				fg.NewSelect(e_ddagger, fg.Name("ch")),
				AuxTypeMeta(zeta, tau),
			})
		} else {
			return fg.NewStructLit(fg.Type("chan_wrapper"), []fg.FGExpr{
				fg.NewSelect(fg.NewAssert(e_ddagger, fg.Type(chan_wrapper_name)), fg.Name("ch")),
				AuxTypeMeta(zeta, tau),
			})
		}
	default:
		return e_ddagger
	}
}

func composeEtaToZetaDummy(eta map[fgg.TParam]fg.FGExpr) map[fgg.Name]fg.FGExpr {
	zeta := make(map[fgg.Name]fg.FGExpr)
	for k, _ := range eta {
		zeta[k.String()] = fg.NewVariable("nil")
	}
	return zeta
}

func composeEtaToZeta(eta map[fgg.TParam]fg.FGExpr) map[fgg.Name]fg.FGExpr {
	zeta := make(map[fgg.Name]fg.FGExpr)
	for k, v := range eta {
		zeta[k.String()] = fg.NewSelect(v, "_type")
	}
	return zeta
}

func (this DictPassTranslator) TranslateExpr(e fgg.FGGExpr, ds []fgg.Decl, delta fgg.Delta, eta map[fgg.TParam]fg.FGExpr, gamma fgg.Gamma) (expr fg.FGExpr, isStructLit bool) {
	switch e_ := e.(type) {
	case fgg.Variable: //I-VAR
		ret := fg.NewVariable(AuxBracket(e_.GetName()))
		//this.env.fgExprIsStructLit[&ret] = true
		_, ok := this.env.StructVariables[e_.GetName()]
		return ret, ok //fgg.IsStructType(ds, gamma[e_.GetName()]) || (fgg.IsChannelType(gamma[e_.GetName()]))
	case fgg.StructLit: //I-VALUE
		zeta := composeEtaToZeta(eta)
		new_es := []fg.FGExpr{}
		fields_ret := fgg.Fields(ds, e_.U_S)
		for i, old_e := range e_.Elems {
			_expr, isStructLit := this.TranslateExpr(old_e, ds, delta, eta, gamma)
			new_es = append(new_es, AuxRechan(zeta, _expr, isStructLit, fields_ret[i].GetType()))
		}
		tdecl := fgg.GetTDecl(ds, e_.U_S.TName)
		new_es = append(new_es, this.MakeDictForLists(ds, eta, delta, e_.U_S.UArgs, tdecl.GetBigPsi())...)
		ret := fg.NewStructLit(fg.Type(AuxBracket(e_.U_S.GetName())), new_es)
		//this.env.fgExprIsStructLit[&ret] = true
		return ret, true
	case fgg.Select: //I-field
		e_ty := e_.GetExpr().TypingWithEnv(&this.env, delta, gamma, false) //TODO: add a compile error report here

		new_e, isStructLit := this.TranslateExpr(e_.GetExpr(), ds, delta, eta, gamma)
		var ret fg.FGExpr
		if !this.env.Stupid && isStructLit {
			ret = fg.NewSelect(new_e, AuxBracket(e_.GetField()))
		} else {
			ret = fg.NewSelect(fg.NewAssert(new_e, fg.Type(AuxBracket(e_ty.(fgg.TNamed).TName))), AuxBracket(e_.GetField()))
		}
		//e_prt_ty := e_.Typing(ds, delta, gamma, false)
		return ret, false //fgg.IsStructType(ds, e_prt_ty) //|| fgg.IsChannelType(e_prt_ty)
	case fgg.Assert: //I-ASSERT
		new_e, _ := this.TranslateExpr(e_.E_I, ds, delta, eta, gamma)
		zeta := composeEtaToZeta(eta)
		ret := fg.NewCall(AuxTypeMeta(zeta, e_.U_cast), "tryCast", []fg.FGExpr{new_e})
		return ret, false
	case fgg.Call:
		callee_e_ty := e_.E_recv.TypingWithEnv(&this.env, delta, gamma, false)
		new_e_bar := []fg.FGExpr{}
		var m_sig *fgg.Sig
		md, _ := this.env.MethodsDelta(delta, callee_e_ty) //fgg.MethodsDelta(ds, delta, callee_e_ty)
		if tmp, ok := md[e_.Meth]; !ok {
			panic("unknown type argument")
		} else {
			m_sig = tmp
		}
		zeta := composeEtaToZeta(eta)
		for i, old_e := range e_.Args {
			_expr, isStructLit := this.TranslateExpr(old_e, ds, delta, eta, gamma)
			new_e_bar = append(new_e_bar, AuxRechan(zeta, _expr, isStructLit, m_sig.GetParamDecls()[i].U.TSubs(m_sig.GetPsi().ToDelta())))
		}

		var small_psi_dagger []fg.FGExpr
		small_psi_dagger = this.MakeDictForLists(ds, eta, delta, e_.GetTArgs(), m_sig.GetPsi())
		e_dagger, isStructLit := this.TranslateExpr(e_.E_recv, ds, delta, eta, gamma)
		switch callee_e_ty_ := callee_e_ty.(type) {
		case fgg.TParam: //I-DICTCALL
			new_arg_list := []fg.FGExpr{}
			new_arg_list = append(new_arg_list, e_dagger)
			new_arg_list = append(new_arg_list, small_psi_dagger...)
			new_arg_list = append(new_arg_list, new_e_bar...)
			var ret fg.Call
			ret = fg.NewCall(eta[callee_e_ty_], AuxBracket(m_sig.GetMethod()), new_arg_list)
			return ret, false
		case fgg.TNamed: //I-CALL
			new_arg_list := []fg.FGExpr{}
			new_arg_list = append(new_arg_list, small_psi_dagger...)
			new_arg_list = append(new_arg_list, new_e_bar...)
			var ret fg.Call
			if !this.env.Stupid && isStructLit {
				ret = fg.NewCall(e_dagger, AuxBracket(m_sig.GetMethod()), new_arg_list)
			} else {
				e_assert := fg.NewAssert(e_dagger, fg.Type(AuxBracket(callee_e_ty_.TName)))
				ret = fg.NewCall(e_assert, AuxBracket(m_sig.GetMethod()), new_arg_list)
			}
			return ret, false //fgg.IsStructType(ds, m_sig.GetReturn()) || fgg.IsChannelType(m_sig.GetReturn())
		default:
			panic("not implemented yet!")
		}
	case fgg.MakeChan: //i-make
		exprs := []fg.FGExpr{}
		exprs = append(exprs, fg.MakeChan{fg.ChannelType{
			ElementType: fg.Type("Any"),
			ChType:      e_.ChType.ChType,
		}})
		zeta := composeEtaToZeta(eta)
		exprs = append(exprs, AuxTypeMeta(zeta, e_.ChType))
		chan_wrapper_struct := fg.NewStructLit(fg.Type(chan_wrapper_name), exprs)
		//TODO: add type_meta_zeta here
		return chan_wrapper_struct, true
	case fgg.ChanRecv: //i-receive
		new_sub_expr, isSLit := this.TranslateExpr(e_.Body, ds, delta, eta, gamma)
		new_assert := fg.NewAssert(new_sub_expr, fg.Type(chan_wrapper_name))
		var new_expr fg.FGExpr
		if !this.env.Stupid && isSLit {
			new_expr = fg.NewSelect(new_sub_expr, fg.Name("ch"))
		} else {
			new_expr = fg.NewSelect(new_assert, fg.Name("ch"))
		}
		new_recv := fg.ChanRecv{new_expr}
		elem_ty := e_.Typing(ds, delta, gamma, false)
		newExprIsSLit := fgg.IsChannelType(elem_ty) || fgg.IsStructType(ds, elem_ty)
		return new_recv, newExprIsSLit
	default:
		panic("not implemented yet!")
	}
}
