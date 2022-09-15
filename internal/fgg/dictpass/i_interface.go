package dictpass

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sfzhu93/fgg2go/internal/fg"
	"github.com/sfzhu93/fgg2go/internal/fgg"
)

var no_assertion_global bool

//I-SPEC
func TranslateSig(spec fgg.Sig) fg.Sig {
	newPDecls := ToParamDecl(AuxAsParam(spec.Psi))
	pdecls := spec.GetParamDecls()
	for _, pdecl := range pdecls {
		newpd := fg.NewParamDecl(AuxBracket(pdecl.Name), fg.Type("Any"))
		newPDecls = append(newPDecls, newpd)
	}
	return fg.NewSig(AuxBracket(spec.GetMethod()), newPDecls, fg.Type("Any"))
}

func AuxSpecName(name fgg.Name) fg.Name {
	return "spec_" + name
}
func (this DictPassTranslator) AuxSpecMetadata(spec fgg.Sig) fg.Sig {
	metadataIndex := len(spec.Psi.TFormals) + len(spec.GetParamDecls())
	this.env.N_bar[metadataIndex] = struct{}{}
	return fg.NewSig(AuxSpecName(spec.GetMethod()), []fg.ParamDecl{}, fg.Type("spec_metadata_"+strconv.Itoa(metadataIndex)))
}

//I-INTERFACE
func (this DictPassTranslator) TranslateITypeLit(idecl fgg.ITypeLit) []fg.Decl {
	ret := []fg.Decl{}
	ispecs := []fg.Spec{}
	ispec_metadata := []fg.Spec{}
	dictfields := []fg.FieldDecl{}
	for _, spec := range idecl.GetSpecs() {
		switch spec_ := spec.(type) {
		case fgg.Sig:
			//if InDictionaries(idecl.GetName()) {
			dictfields = append(dictfields, TranslateSigDict(&spec_))
			//}
			ispecs = append(ispecs, TranslateSig(spec_))
			if !this.env.No_assertion {
				ispec_metadata = append(ispec_metadata, this.AuxSpecMetadata(spec_))
			}
		case fgg.TNamed:
			panic("Not a standard FGG syntax!")
		}
	}

	ispecs = append(ispecs, ispec_metadata...)
	ret = append(ret, fg.NewITypeLit(fg.Type(AuxBracket(idecl.GetName())), ispecs))

	//if InDictionaries(idecl.GetName()) {
	if !this.env.No_assertion && (!this.env.OptAssertionForZeroParam || this.env.OptAssertionForZeroParam && len(idecl.Psi.TFormals) > 0) {
		dictfields = append(dictfields, fg.NewFieldDecl(fg.Name("_type"), fg.Type("_type_metadata")))
	}
	ret = append(ret, fg.NewSTypeLit(fg.Type(AuxTypeDict(fg.Type(AuxBracket(idecl.GetName())))), dictfields))
	if !this.env.No_assertion {
		ret = append(ret, this.MakeOtherDefinitionsForInterface(idecl)...)
	}

	//}

	return ret
}

func AuxTypeMeta(zeta map[fgg.Name]fg.FGExpr, ty_ fgg.Type) fg.FGExpr {
	switch ty := ty_.(type) {
	case fgg.TNamed:
		fields := []fg.FGExpr{}
		for _, field := range ty.UArgs {
			fields = append(fields, AuxTypeMeta(zeta, field))
		}
		return fg.NewStructLit(fg.Type(AuxMetadataName(ty.TName)), fields)
	case fgg.TParam:
		return zeta[ty.String()]
	case fgg.ChannelType:
		chan_meta_name := ""
		switch ty.ChType {
		case fgg.ChBidirection:
			chan_meta_name = bidirectional_channel_meta_name
		case fgg.ChRecvOnly:
			chan_meta_name = receive_only_channel_meta_name
		case fgg.ChSendOnly:
			chan_meta_name = send_only_channel_meta_name
		}
		ty_meta := AuxTypeMeta(zeta, ty.ElementType)
		if ty_meta == nil {
			println("debug point")
		}
		return fg.NewStructLit(fg.Type("chan_metadata"), []fg.FGExpr{
			fg.NewStructLit(fg.Type(chan_meta_name), []fg.FGExpr{}),
			ty_meta,
		})
	default:
		panic("unknown type kind.")
	}
}
func AuxMetadataName(idecl fgg.Name) fg.Name {
	return AuxBracket(idecl) + "_meta"
}

func (this DictPassTranslator) MakeTryCastBodyForInterface(idecl fgg.ITypeLit, zeta map[fgg.Name]fg.FGExpr) []fg.FGStmt {
	ret := []fg.FGStmt{}
	ret = append(ret, fg.OtherStmt{"x_ := x.(" + AuxBracket(idecl.GetName()) + ")"})
	for _, spec := range idecl.GetSpecs() {
		if sig, ok := spec.(fgg.Sig); ok {
			new_zeta := make(map[fgg.Name]fg.FGExpr)
			for k, v := range zeta {
				new_zeta[k] = v
			}
			for i, tformal := range sig.Psi.TFormals {
				this.env.Param_index_set[i] = struct{}{}
				new_zeta[tformal.Name.String()] = fg.NewStructLit(fg.Type("param_index_"+strconv.Itoa(i)), []fg.FGExpr{})
			}
			ret = append(ret, fg.OtherStmt{
				fmt.Sprintf("%s_actual := x_.%s()\n", sig.GetMethod(), AuxSpecName(sig.GetMethod())),
			})
			len_beta := len(sig.Psi.TFormals)
			for i, tformal := range sig.Psi.TFormals {
				ret = append(ret, fg.OtherStmt{
					Body: fmt.Sprintf("%s_actual._type_%d.assertEq(%s)", sig.GetMethod(), i, AuxTypeMeta(new_zeta, tformal.U_I)),
				})
			}
			for i, pd := range sig.GetParamDecls() {
				ret = append(ret, fg.OtherStmt{
					fmt.Sprintf("%s_actual._type_%d.assertEq(%s)", sig.GetMethod(), i+len_beta, AuxTypeMeta(new_zeta, pd.U)),
				})
			}
			ret = append(ret, fg.OtherStmt{
				fmt.Sprintf("%s_actual._type_%d.assertEq(%s)", sig.GetMethod(), len(sig.GetParamDecls())+len_beta, AuxTypeMeta(new_zeta, sig.GetReturn())),
			})
		} else {
			panic(spec.String() + "is not a sig.")
		}
	}
	ret = append(ret, fg.OtherStmt{"return x_"})
	return ret
}

func (this DictPassTranslator) MakeAssertEqBodyForInterface(idecl fgg.ITypeLit) []fg.FGStmt {
	ret := []fg.FGStmt{}
	ret = append(ret, fg.OtherStmt{"x_ := x.(" + AuxMetadataName(idecl.GetName()) + ")"})
	for i := 0; i < len(idecl.Psi.TFormals); i++ {
		stmt := fg.OtherStmt{
			"this._type_" + strconv.Itoa(i) + ".assertEq(x_._type_" + strconv.Itoa(i) + ")",
		}
		ret = append(ret, stmt)
	}
	ret = append(ret, fg.FGReturnStmt{fg.NewVariable("x_")})
	return ret
}

func (this DictPassTranslator) MakeOtherDefinitionsForInterface(idecl fgg.ITypeLit) []fg.Decl {
	ret := []fg.Decl{}
	meta_name := AuxMetadataName(idecl.GetName())
	params := []string{}
	for i := 0; i < len(idecl.Psi.TFormals); i++ {
		params = append(params, fmt.Sprintf("_type_%d _type_metadata", i))
	}
	metadata_struct := fg.OtherDecls{Def: fmt.Sprintf("type %s struct { %s }", meta_name, strings.Join(params, ";"))}
	ret = append(ret, metadata_struct)
	zeta := make(map[fgg.Name]fg.FGExpr)
	for i, tfml := range idecl.Psi.TFormals {
		zeta[tfml.Name.String()] = fg.NewSelect(fg.NewVariable("this"), fg.Name("_type_"+strconv.Itoa(i)))
		//fg.NewSelect(fg.NewSelect(fg.NewVariable("this"), fg.Name("dict_" + strconv.Itoa(i))), fg.Name("_type"))
	}
	trycast_b := this.MakeTryCastBodyForInterface(idecl, zeta)
	func_tryCast := fg.NewMDecl(
		fg.NewParamDecl("this", fg.Type(meta_name)),
		"tryCast",
		[]fg.ParamDecl{fg.NewParamDecl(fg.Name("x"), fg.Type("Any"))},
		fg.Type("Any"),
		nil,
		trycast_b,
	)
	ret = append(ret, func_tryCast)
	func_assertEq := fg.NewMDecl(
		fg.NewParamDecl("this", fg.Type(meta_name)),
		"assertEq",
		[]fg.ParamDecl{fg.NewParamDecl(fg.Name("x"), fg.Type("_type_metadata"))},
		fg.Type("Any"),
		nil,
		this.MakeAssertEqBodyForInterface(idecl),
	)
	ret = append(ret, func_assertEq)
	return ret
}
