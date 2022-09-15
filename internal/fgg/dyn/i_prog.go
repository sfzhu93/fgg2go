package dyn

import (
	"github.com/sfzhu93/fgg2go/internal/fg"
	"github.com/sfzhu93/fgg2go/internal/fgg"
	"strconv"
)

const type_metadata_name = "_type_metadata"
const bidirectional_channel_meta_name = "__BidirectionalChannel"
const receive_only_channel_meta_name = "__receiveOnlyChannel"
const send_only_channel_meta_name = "__sendOnlyChannel"
const spec_metadata_name = "spec_metadata_"

type Env struct {
	n_bar           map[int]struct{} //we update the arity set while translating the program, and call makeMetas in the end
	param_index_set map[int]struct{}
	//fgExprIsStructLit map[fg.FGExpr]bool
	StructVariables map[string]struct{}
	stupid          bool
	NoUnusedVar     bool
	defVar          map[string]struct{}
	usedVar         map[string]struct{}
}

type DictPassTranslator struct {
	env     Env
	fggProg fgg.FGGProgram
	fgProg  fg.FGProgram
}

func (this DictPassTranslator) MakeTypeMeta() fg.Decl {
	return fg.NewITypeLit(fg.Type(type_metadata_name), []fg.Spec{
		fg.NewSig("tryCast",
			[]fg.ParamDecl{
				fg.NewParamDecl(fg.Name("x"), fg.Type("Any")),
			},
			fg.Type("Any"),
		),
		fg.NewSig("assertEq",
			[]fg.ParamDecl{
				fg.NewParamDecl(fg.Name("x"), fg.Type(type_metadata_name)),
			},
			fg.Type("Any"),
		),
	})
}

func (this DictPassTranslator) MakeChannelDirectionDefs() []fg.Decl {
	ret := []fg.Decl{}
	for _, chDirName := range []string{bidirectional_channel_meta_name, send_only_channel_meta_name, receive_only_channel_meta_name} {
		ret = append(ret, fg.NewSTypeLit(fg.Type(chDirName), []fg.FieldDecl{}))
		ret = append(ret, fg.NewMDecl(fg.NewParamDecl(fg.Name("this"), fg.Type(chDirName)),
			fg.Name("assertEq"),
			[]fg.ParamDecl{
				fg.NewParamDecl(fg.Name("x"), fg.Type("Any")),
			},
			fg.Type("Any"),
			nil,
			[]fg.FGStmt{
				fg.FGReturnStmt{fg.NewAssert(fg.NewVariable(fg.Name("x")), fg.Type(chDirName))},
			},
		))
	}
	return ret
}

func (this DictPassTranslator) MakeChannelDefs() []fg.Decl {
	ret := []fg.Decl{}
	ret = append(ret, fg.NewSTypeLit(fg.Type(chan_wrapper_name), []fg.FieldDecl{
		fg.NewFieldDecl(fg.Name("ch"), fg.ChannelType{
			ElementType: fg.Type("Any"),
			ChType:      fg.ChBidirection,
		}),
		//fg.NewFieldDecl(fg.Name("_type"), fg.Type(type_metadata_name)),
	}))
	ret = append(ret, fg.NewITypeLit(fg.Type("chan_direction"), []fg.Spec{
		fg.NewSig("assertEq",
			[]fg.ParamDecl{
				fg.NewParamDecl(fg.Name("x"), fg.Type("Any")),
			},
			fg.Type("Any"),
		),
	}))
	ret = append(ret, fg.NewSTypeLit(fg.Type("chan_metadata"), []fg.FieldDecl{
		fg.NewFieldDecl(fg.Name("dir"), fg.Type("chan_direction")),
		//fg.NewFieldDecl(fg.Name("_type"), fg.Type(type_metadata_name)),
	}))
	ret = append(ret, fg.OtherDecls{Def: `func (this chan_metadata) tryCast(x Any) Any {
//_ = x.(chan_wrapper)._type.assertEq(this);
return x
}`})
	ret = append(ret, fg.OtherDecls{Def: `func (this chan_metadata) assertEq(x _type_metadata) Any {
//_ = x.(chan_metadata)._type.assertEq(this._type); 
return x.(chan_metadata).dir.assertEq(this.dir)
}`})
	return ret
}

/*func (this DictPassTranslator) MakeArities() {
	for _, d := range this.fggProg.GetDecls() {
		if d_, ok := d.(fgg.ITypeLit); ok {
			arity := len(d_.Psi.TFormals) + len(d_.GetSpecs())
			this.env.n_bar
		}


	}
}*/

func MakeSpecMetadata(n int) fg.Decl {
	fields := []fg.FieldDecl{}
	for i := 0; i <= n; i++ {
		field := fg.NewFieldDecl(fg.Name("_type_"+strconv.Itoa(i)), fg.Type(type_metadata_name))
		fields = append(fields, field)
	}
	return fg.NewSTypeLit(fg.Type(spec_metadata_name+strconv.Itoa(n)), fields)
}

func (this DictPassTranslator) MakeSpecMetadataStructs() []fg.Decl {
	ret := []fg.Decl{}
	for k, _ := range this.env.n_bar {
		ret = append(ret, MakeSpecMetadata(k))
	}
	return ret
}

func TranslateToDynFG(p fgg.FGGProgram) fg.FGProgram {
	//omega := GetOmega(p.GetDecls(), p.GetMain().(fgg.FGGExpr))
	dictCache = getDict(p.GetDecls())
	dpt := DictPassTranslator{
		env:     Env{},
		fggProg: p,
		fgProg:  fg.FGProgram{},
	}
	dpt.env.n_bar = make(map[int]struct{})
	//dpt.env.fgExprIsStructLit = make(map[fg.FGExpr]bool)
	dpt.env.stupid = false
	dpt.env.NoUnusedVar = true
	//fmt.Printf("%+v\n", omega)
	decls := dpt.TranslateDecls(p.GetDecls())
	decls = append(decls, dpt.MakeSpecMetadataStructs()...)
	decls = append(decls, dpt.MakeTypeMeta())
	decls = append(decls, dpt.MakeChannelDirectionDefs()...)
	decls = append(decls, dpt.MakeChannelDefs()...)
	//decls := eraseDecls(p)
	e_dagger, _ := dpt.TranslateExpr(p.GetMain().(fgg.FGGExpr), p.GetDecls(), fgg.Delta{}, make(map[fgg.TParam]fg.FGExpr), fgg.Gamma{})
	return fg.NewFGProgram(decls, e_dagger, false)
}
