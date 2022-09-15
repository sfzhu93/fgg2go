package dyn

import (
	"github.com/sfzhu93/fgg2go/internal/fg"
	"github.com/sfzhu93/fgg2go/internal/fgg"
)

//I-SPEC
func TranslateSig(spec fgg.Sig) fg.Sig {
	newPDecls := []fg.ParamDecl{} //ToParamDecl(AuxAsParam(spec.Psi))
	pdecls := spec.GetParamDecls()
	for _, pdecl := range pdecls {
		newpd := fg.NewParamDecl(AuxBracket(pdecl.Name), fg.Type("Any"))
		newPDecls = append(newPDecls, newpd)
	}
	return fg.NewSig(AuxBracket(spec.GetMethod()), newPDecls, fg.Type("Any"))
}

//I-INTERFACE
func (this DictPassTranslator) TranslateITypeLit(idecl fgg.ITypeLit) []fg.Decl {
	ret := []fg.Decl{}
	ispecs := []fg.Spec{}
	//ispec_metadata := []fg.Spec{}
	//dictfields := []fg.FieldDecl{}
	for _, spec := range idecl.GetSpecs() {
		switch spec_ := spec.(type) {
		case fgg.Sig:
			//if InDictionaries(idecl.GetName()) {
			//dictfields = append(dictfields, TranslateSigDict(spec_))
			//}
			ispecs = append(ispecs, TranslateSig(spec_))
			//ispec_metadata = append(ispec_metadata, this.AuxSpecMetadata(spec_))
		case fgg.TNamed:
			panic("!!")
		}
	}
	//ispecs = append(ispecs, ispec_metadata...)
	ret = append(ret, fg.NewITypeLit(fg.Type(AuxBracket(idecl.GetName())), ispecs))
	//if InDictionaries(idecl.GetName()) {
	//dictfields = append(dictfields, fg.NewFieldDecl(fg.Name("_type"), fg.Type("_type_metadata")))
	//ret = append(ret, fg.NewSTypeLit(fg.Type(AuxTypeDict(fg.Type(AuxBracket(idecl.GetName())))), dictfields))
	//}
	//ret = append(ret, this.MakeOtherDefinitionsForInterface(idecl)...)
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
	//case fgg.ChannelType:
	//	chan_meta_name := ""
	//	switch ty.ChType {
	//	case fgg.ChBidirection:
	//		chan_meta_name = bidirectional_channel_meta_name
	//	case fgg.ChRecvOnly:
	//		chan_meta_name = receive_only_channel_meta_name
	//	case fgg.ChSendOnly:
	//		chan_meta_name = send_only_channel_meta_name
	//	}
	//	ty_meta := AuxTypeMeta(zeta, ty.ElementType)
	//	if ty_meta == nil {
	//		println("debug point")
	//	}
	//	return fg.NewStructLit(fg.Type("chan_metadata"), []fg.FGExpr{
	//		fg.NewStructLit(fg.Type(chan_meta_name), []fg.FGExpr{}),
	//		ty_meta,
	//	})
	default:
		panic("unknown type kind.")
	}
}
func AuxMetadataName(idecl fgg.Name) fg.Name {
	return AuxBracket(idecl) + "_meta"
}
