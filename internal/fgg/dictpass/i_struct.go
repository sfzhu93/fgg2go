package dictpass

import (
	"fmt"
	"strings"

	"github.com/sfzhu93/fgg2go/internal/fg"
	"github.com/sfzhu93/fgg2go/internal/fgg"
)

//I-STRUCT
func (this DictPassTranslator) TranslateSTypeLit(sdecl fgg.STypeLit) []fg.Decl {
	fields := []fg.FieldOrParamDecl{}
	n := len(sdecl.Psi.TFormals)
	this.env.N_bar[n] = struct{}{}
	for _, fldDecl := range sdecl.GetFieldDecls() {
		fields = append(fields, fg.NewFieldDecl(AuxBracket(fldDecl.GetName()), fg.Type("Any")))
	}
	fields = append(fields, AuxAsParam(sdecl.Psi)...)
	ret := []fg.Decl{fg.NewSTypeLit(fg.Type(AuxBracket(sdecl.GetName())), ToFieldDecl(fields))}
	if !this.env.No_assertion {

		metadata_struct_name := AuxMetadataName(sdecl.GetName()) //spec_metadata_name + strconv.Itoa(n)
		params := []string{}
		for i := 0; i < len(sdecl.Psi.TFormals); i++ {
			params = append(params, fmt.Sprintf("_type_%d _type_metadata", i))
		}
		metadata_struct := fg.OtherDecls{Def: fmt.Sprintf("type %s struct { %s }", metadata_struct_name, strings.Join(params, ";"))}
		ret = append(ret, metadata_struct)
		//make trycast
		trycast_b := []fg.FGStmt{}
		trycast_b = append(trycast_b, fg.OtherStmt{Body: "_ = x.(" + AuxBracket(sdecl.GetName()) + ")"})
		for i := 0; i < n; i++ {
			if this.env.OptAssertionForZeroParam {
				if !this.env.CanOptimizedOutTypeMetadata(sdecl.Psi.TFormals[i].U_I) {
					trycast_b = append(trycast_b, fg.OtherStmt{
						Body: fmt.Sprintf("this._type_%[1]d.assertEq(x.(%s).dict_%[1]d._type)", i, AuxBracket(sdecl.GetName())),
					})
				}
			}
		}
		trycast_b = append(trycast_b, fg.OtherStmt{"return x"})
		func_tryCast := fg.NewMDecl(
			fg.NewParamDecl("this", fg.Type(metadata_struct_name)),
			"tryCast",
			[]fg.ParamDecl{fg.NewParamDecl(fg.Name("x"), fg.Type("Any"))},
			fg.Type("Any"),
			nil,
			trycast_b,
		)
		ret = append(ret, func_tryCast)

		//make asserteq
		assertEq_b := []fg.FGStmt{}
		for i := 0; i < n; i++ {
			assertEq_b = append(assertEq_b, fg.OtherStmt{
				Body: fmt.Sprintf("this._type_%[1]d.assertEq(x.(%s)._type_%[1]d)", i, metadata_struct_name),
			})
		}
		// TODO: we need a x.(t_S) assertion here, according to d-struct in the oopsla paper.
		// However, this expression is incorrect in standard Go.
		assertEq_b = append(assertEq_b, fg.OtherStmt{"return this"})
		func_assertEq := fg.NewMDecl(
			fg.NewParamDecl("this", fg.Type(metadata_struct_name)),
			"assertEq",
			[]fg.ParamDecl{fg.NewParamDecl(fg.Name("x"), fg.Type("_type_metadata"))},
			fg.Type("Any"),
			nil,
			assertEq_b,
		)
		ret = append(ret, func_assertEq)
	}
	return ret
}
